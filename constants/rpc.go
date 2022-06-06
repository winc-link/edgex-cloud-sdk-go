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
package constants

import (
    "google.golang.org/grpc/keepalive"
    "time"
)

var Kaep = keepalive.EnforcementPolicy{
    MinTime:             5 * time.Second,
    PermitWithoutStream: true,
}

var Kasp = keepalive.ServerParameters{
    MaxConnectionIdle:     30 * time.Second,
    MaxConnectionAge:      30 * time.Second,
    MaxConnectionAgeGrace: 5 * time.Second,
    Time:                  5 * time.Second,
    Timeout:               3 * time.Second,
}
