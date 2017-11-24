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

# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: fleetspeak/src/server/proto/fleetspeak_server/server.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from fleetspeak.src.server.proto.fleetspeak_server import services_pb2 as fleetspeak_dot_src_dot_server_dot_proto_dot_fleetspeak__server_dot_services__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='fleetspeak/src/server/proto/fleetspeak_server/server.proto',
  package='fleetspeak.server',
  syntax='proto3',
  serialized_pb=_b('\n:fleetspeak/src/server/proto/fleetspeak_server/server.proto\x12\x11\x66leetspeak.server\x1a\x1egoogle/protobuf/duration.proto\x1a<fleetspeak/src/server/proto/fleetspeak_server/services.proto\"z\n\x0cServerConfig\x12\x32\n\x08services\x18\x01 \x03(\x0b\x32 .fleetspeak.server.ServiceConfig\x12\x36\n\x13\x62roadcast_poll_time\x18\x02 \x01(\x0b\x32\x19.google.protobuf.Durationb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,fleetspeak_dot_src_dot_server_dot_proto_dot_fleetspeak__server_dot_services__pb2.DESCRIPTOR,])




_SERVERCONFIG = _descriptor.Descriptor(
  name='ServerConfig',
  full_name='fleetspeak.server.ServerConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='services', full_name='fleetspeak.server.ServerConfig.services', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='broadcast_poll_time', full_name='fleetspeak.server.ServerConfig.broadcast_poll_time', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=175,
  serialized_end=297,
)

_SERVERCONFIG.fields_by_name['services'].message_type = fleetspeak_dot_src_dot_server_dot_proto_dot_fleetspeak__server_dot_services__pb2._SERVICECONFIG
_SERVERCONFIG.fields_by_name['broadcast_poll_time'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
DESCRIPTOR.message_types_by_name['ServerConfig'] = _SERVERCONFIG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ServerConfig = _reflection.GeneratedProtocolMessageType('ServerConfig', (_message.Message,), dict(
  DESCRIPTOR = _SERVERCONFIG,
  __module__ = 'fleetspeak.src.server.proto.fleetspeak_server.server_pb2'
  # @@protoc_insertion_point(class_scope:fleetspeak.server.ServerConfig)
  ))
_sym_db.RegisterMessage(ServerConfig)


try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)