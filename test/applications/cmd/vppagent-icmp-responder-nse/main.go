// Copyright 2018-2019 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/sirupsen/logrus"

	"github.com/networkservicemesh/networkservicemesh/pkg/tools"
	"github.com/networkservicemesh/networkservicemesh/sdk/common"
	"github.com/networkservicemesh/networkservicemesh/sdk/endpoint"
	"github.com/networkservicemesh/networkservicemesh/sdk/vppagent"
)

var version string

func main() {
	logrus.Info("Starting vppagent-icmp-responder-nse...")
	logrus.Infof("Version: %v", version)
	// Capture signals to cleanup before exiting
	c := tools.NewOSSignalChannel()

	configuration := &common.NSConfiguration{
		MechanismType: "mem",
	}

	composite := endpoint.NewCompositeEndpoint(
		endpoint.NewMonitorEndpoint(configuration),
		vppagent.NewFlush(configuration, "localhost:9112"),
		vppagent.NewMemifConnect(configuration),
		endpoint.NewIpamEndpoint(nil),
		endpoint.NewConnectionEndpoint(configuration))

	nsmEndpoint, err := endpoint.NewNSMEndpoint(nil, configuration, composite)
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	nsmEndpoint.Start()
	defer nsmEndpoint.Delete()

	<-c
}
