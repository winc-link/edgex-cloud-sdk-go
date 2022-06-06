/*******************************************************************************
 * Copyright 2017 Dell Inc.
 * Copyright (c) 2019 Intel Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package handlers

import (
    "context"
    "edgex-cloud-sdk-go/proto/common"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

type CommonServer struct {
    common.UnimplementedCommonServer
}

func (cs *CommonServer) Ping(ctx context.Context, empty *emptypb.Empty) (*common.Pong, error) {
    panic("implement me")
}

func (cs *CommonServer) Version(ctx context.Context, empty *emptypb.Empty) (*common.VersionResponse, error) {
    panic("implement me")
}

func NewCommonServer() *CommonServer {
    return &CommonServer{
    }
}

func (cs *CommonServer) RegisterCommonServer(s *grpc.Server) {
    common.RegisterCommonServer(s, cs)
}