// Copyright 2021 The phy-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fake

import (
	"github.com/getlantern/deepcopy"
	"github.com/sacloud/phy-go/openapi"
)

type Server struct {
	Server       *openapi.Server
	RaidStatus   *openapi.RaidStatus
	OSImages     []*openapi.OsImage
	PowerStatus  *openapi.ServerPowerStatus
	TrafficGraph *openapi.TrafficGraph
}

func (s *Server) Id() string {
	return s.Server.ServerId
}

func (s *Server) getPortChannelById(portChannelId openapi.PortChannelId) (*openapi.PortChannel, error) {
	for _, portChannel := range s.Server.PortChannels {
		if portChannel.PortChannelId == int(portChannelId) {
			var channel openapi.PortChannel
			if err := deepcopy.Copy(&channel, &portChannel); err != nil {
				return nil, err
			}
			return &channel, nil
		}
	}
	return nil, NewError(ErrorTypeNotFound, "port-channel", portChannelId, "server[%s]", s.Id())
}

func (s *Server) getPortById(portId openapi.PortId) (*openapi.InterfacePort, error) {
	for _, port := range s.Server.Ports {
		if port.PortId == int(portId) {
			var pt openapi.InterfacePort
			if err := deepcopy.Copy(&pt, &port); err != nil {
				return nil, err
			}
			return &pt, nil
		}
	}
	return nil, NewError(ErrorTypeNotFound, "port", portId, "server[%s]", s.Id())
}

func (s *Server) updatePortChannel(portChannel *openapi.PortChannel) {
	for i, pc := range s.Server.PortChannels {
		if pc.PortChannelId == portChannel.PortChannelId {
			s.Server.PortChannels[i] = *portChannel
			break
		}
	}
}

func (s *Server) updatePort(port *openapi.InterfacePort) {
	for i, p := range s.Server.Ports {
		if p.PortId == port.PortId {
			s.Server.Ports[i] = *port
			break
		}
	}
}
