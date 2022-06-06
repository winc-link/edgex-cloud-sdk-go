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
    "github.com/winc-link/edgex-cloud-sdk-go/proto/device"
    "github.com/winc-link/edgex-cloud-sdk-go/providers"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

var _ device.RpcDeviceServer = (*DeviceServer)(nil)

type DeviceServer struct {
    provider providers.EdgeXCloud
    device.UnsafeRpcDeviceServer
}

func (ds *DeviceServer) AddDevice(ctx context.Context, request *device.AddDeviceRequest) (*emptypb.Empty, error) {
    fmt.Println("....",request.GetDevice())
    ds.provider.AddDevice(request.GetDevice().Id)
    return nil,nil
    //panic("implement me")
}

func (ds *DeviceServer) ActivateDevice(ctx context.Context, active *device.DeviceActive) (*device.ActiveDeviceResponse, error) {
    panic("implement me")
}

func (ds *DeviceServer) UpdateDevice(ctx context.Context, request *device.UpdateDeviceRequest) (*emptypb.Empty, error) {
    panic("implement me")
}

func (ds *DeviceServer) DeviceById(ctx context.Context, request *device.DeviceByIdRequest) (*device.DeviceInfo, error) {
    panic("implement me")
}

func (ds *DeviceServer) DeleteDeviceById(ctx context.Context, request *device.DeleteDeviceByIdRequest) (*emptypb.Empty, error) {
    panic("implement me")
}

func (ds *DeviceServer) DevicesSearch(ctx context.Context, request *device.DeviceSearchQueryRequest) (*device.MultiDeviceResponse, error) {
    panic("implement me")
}

func (ds *DeviceServer) ReportDevicesOnlineAndOffline(ctx context.Context, list *device.DeviceOnlineAndOfflineList) (*emptypb.Empty, error) {
    panic("implement me")
}

//func (ds *DeviceServer) mustEmbedUnimplementedRpcDeviceServer() {
//    panic("implement me")
//}

func NewDeviceServer(provider providers.EdgeXCloud) *DeviceServer {
    return &DeviceServer{
        provider: provider,
    }
}

func (ds *DeviceServer) RegisterServer(s *grpc.Server) {
    device.RegisterRpcDeviceServer(s, ds)
}




