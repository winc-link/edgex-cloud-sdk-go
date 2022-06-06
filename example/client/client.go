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
package main

import (
    "context"
    "edgex-cloud-sdk-go/proto/device"
    "fmt"
    "google.golang.org/grpc"
    "log"
    "time"
)

func main() {
    // 连接grpc服务器
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    // 延迟关闭连接
    defer conn.Close()
    
    // 初始化Greeter服务客户端
    c := device.NewRpcDeviceClient(conn)
    
    // 初始化上下文，设置请求超时时间为1秒
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    // 延迟关闭请求会话
    defer cancel()
    
    dev := device.DeviceAddInfo{
        Id: "1",
    }
    // 调用SayHello接口，发送一条消息
    addDeviceInfo := device.AddDeviceRequest{
        Device: &dev,
    }
    r, err := c.AddDevice(ctx,&addDeviceInfo)
    if err != nil {
        fmt.Println(err)
    }
    
    // 打印服务的返回的消息
    fmt.Println("Greeting: %s",r)
}
