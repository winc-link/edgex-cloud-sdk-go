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
    "fmt"
    
    "github.com/winc-link/edgex-cloud-sdk-go/providers"
    "github.com/winc-link/edgex-cloud-sdk-go/service"
)

var _ providers.EdgeXCloud = (*st)(nil)

type st struct {

}


func (s st) CreateProduct(productName string,nodeType int ,description string,protocolType string ,extra map[string]interface{}) {
    panic("implement me")
}

func (s st) AddDevice(id string) {
    fmt.Println("add device",id)
}

func Newst()  *st{
    return &st{}
}

func main()  {
    fmt.Println(1)
    s := service.NewEdgeXCloudServer(Newst())
    err := s.Start()
    if err != nil{
        panic(err)
    }
}
