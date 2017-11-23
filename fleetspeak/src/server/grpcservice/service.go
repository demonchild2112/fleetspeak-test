// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package grpcservice defines a service.Service which passes all received messages to
// a destination host using grpc.
package grpcservice

import (
	"fmt"
	"time"

	"context"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/google/fleetspeak/fleetspeak/src/server/service"

	fspb "github.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeak"
	gpb "github.com/google/fleetspeak/fleetspeak/src/server/grpcservice/proto/fleetspeak_grpcservice"
	ggrpc "github.com/google/fleetspeak/fleetspeak/src/server/grpcservice/proto/fleetspeak_grpcservice"
	spb "github.com/google/fleetspeak/fleetspeak/src/server/proto/fleetspeak_server"
)

// GRPCService is a service.Service which forwards all received
// messages to an implementation of ggrpc.Processor.
type GRPCService struct {
	sctx   service.Context
	conn   grpc.ClientConn
	client ggrpc.ProcessorClient
}

// NewGRPCService returns a service.Service which forwards received
// messages to c. Implementations which wish to implement transport
// security or otherwise control the connection used should define a
// service.Factory based on this.
func NewGRPCService(c *grpc.ClientConn) *GRPCService {
	return &GRPCService{
		client: ggrpc.NewProcessorClient(c),
	}
}

// Start implements service.Service.
func (s *GRPCService) Start(sctx service.Context) error {
	s.sctx = sctx
	return nil
}

// Stop implements service.Service.
func (s *GRPCService) Stop() error {
	s.conn.Close()
	return nil
}

// ProcessMessage implements service.Service.
func (s *GRPCService) ProcessMessage(ctx context.Context, m *fspb.Message) error {
	// TODO: Remove retry logic when possible.
	var err error
	d := time.Second
L:
	for {
		_, err = s.client.Process(ctx, m)
		if err == nil {
			return nil
		}
		t := time.NewTimer(d)
		select {
		case <-ctx.Done():
			t.Stop()
			break L
		case <-t.C:
			d = d * 2
		}
	}
	// Tell Fleetspeak to retry (minutes later) in the most obviously retryable
	// cases. Assume permission issues are configuration errors which will be
	// fixed eventually.
	if s, ok := status.FromError(err); ok {
		c := s.Code()
		if c == codes.DeadlineExceeded ||
			c == codes.Unavailable ||
			c == codes.Aborted ||
			c == codes.Canceled ||
			c == codes.Unauthenticated ||
			c == codes.PermissionDenied {
			err = service.TemporaryError{err}
		}
	}
	return err
}

// Factory is a server.ServiceFactory that creates a GRPCService.
//
// cfg must contain a fleetspeak.grpcservice.Config message describing
// how to dial the target grpc server.
func Factory(cfg *spb.ServiceConfig) (service.Service, error) {
	var conf gpb.Config
	if err := ptypes.UnmarshalAny(cfg.Config, &conf); err != nil {
		return nil, err
	}

	switch {
	case conf.Insecure:
		con, err := grpc.DialContext(context.Background(), conf.Target, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		return NewGRPCService(con), nil
	case conf.CertFile != "":
		cred, err := credentials.NewClientTLSFromFile(conf.CertFile, "")
		if err != nil {
			return nil, err
		}
		con, err := grpc.DialContext(context.Background(), conf.Target, grpc.WithTransportCredentials(cred))
		if err != nil {
			return nil, err
		}
		return NewGRPCService(con), nil
	default:
		return nil, fmt.Errorf("GRPCService requires either insecure or cert_file to be set, got: %+v", conf)
	}
}
