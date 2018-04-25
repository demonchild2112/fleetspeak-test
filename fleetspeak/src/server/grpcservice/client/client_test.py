# Copyright 2017 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Tests for grpcservice.client.client."""

import threading
import unittest

import grpc

from fleetspeak.src.server.grpcservice.client import client
from fleetspeak.src.common.proto.fleetspeak import common_pb2


# More of a mock than a fake.
# TODO: migrate to use a mock library.
class FakeStub(object):

  def __init__(self):
    self.event = threading.Event()
    self.insert_done = False
    self.insert_errors = 1

  def KeepAlive(self, unused, timeout=None):
    del unused
    del timeout
    self.event.set()

  def InsertMessage(self, message, timeout=None):
    del timeout
    if self.insert_errors:
      self.insert_errors -= 1
      raise grpc.RpcError("insert_errors is positive, try again")
    self.insert_done = True
    self.message = message


class ClientTest(unittest.TestCase):

  def testKeepAlive(self):
    t = FakeStub()
    s = client.OutgoingConnection(None, 'test', t)
    self.assertTrue(t.event.wait(10))
    s.Shutdown()

  def testInsertMessage(self):
    t = FakeStub()
    s = client.OutgoingConnection(None, 'test', t)
    s.InsertMessage(common_pb2.Message())
    self.assertTrue(t.insert_done)
    self.assertFalse(t.insert_errors)
    self.assertEqual(t.message.source.service_name, 'test')

if __name__ == '__main__':
  unittest.main()
