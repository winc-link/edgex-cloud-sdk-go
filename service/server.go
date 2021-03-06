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
package service

import (
	"github.com/winc-link/edgex-cloud-sdk-go/constants"
	"github.com/winc-link/edgex-cloud-sdk-go/internal/pkg/handlers"
	"github.com/winc-link/edgex-cloud-sdk-go/providers"
	"google.golang.org/grpc"
	"net"
)

type Service struct {
	provider providers.EdgeXCloud
}

func NewEdgeXCloudServer(p providers.EdgeXCloud) *Service {
	return &Service{
		provider: p,
	}
}



func (b *Service) Start() error {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		return err
	}
	s := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(constants.Kaep), grpc.KeepaliveParams(constants.Kasp))
	handlers.NewCommonServer().RegisterCommonServer(s)
	handlers.NewDeviceServer(b.provider).RegisterServer(s)

	err = s.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
