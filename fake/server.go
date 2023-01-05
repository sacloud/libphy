// Copyright 2021-2023 The phy-api-go authors
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
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

type Server struct {
	Server       *v1.Server
	RaidStatus   *v1.RaidStatus
	OSImages     []*v1.OsImage
	PowerStatus  *v1.ServerPowerStatus
	TrafficGraph *v1.TrafficGraph
}

func (s *Server) Id() string {
	return s.Server.ServerId
}

func (s *Server) getPortChannelById(portChannelId v1.PortChannelId) (*v1.PortChannel, error) {
	for i := range s.Server.PortChannels {
		portChannel := s.Server.PortChannels[i]
		if portChannel.PortChannelId == portChannelId {
			var channel v1.PortChannel
			if err := deepcopy.Copy(&channel, &portChannel); err != nil {
				return nil, err
			}
			return &channel, nil
		}
	}
	return nil, NewError(ErrorTypeNotFound, "port-channel", portChannelId, "server[%s]", s.Id())
}

func (s *Server) getPortById(portId v1.PortId) (*v1.InterfacePort, error) {
	for i := range s.Server.Ports {
		port := s.Server.Ports[i]
		if port.PortId == portId {
			var pt v1.InterfacePort
			if err := deepcopy.Copy(&pt, &port); err != nil {
				return nil, err
			}
			return &pt, nil
		}
	}
	return nil, NewError(ErrorTypeNotFound, "port", portId, "server[%s]", s.Id())
}

func (s *Server) updatePortChannel(portChannel *v1.PortChannel) {
	for i, pc := range s.Server.PortChannels {
		if pc.PortChannelId == portChannel.PortChannelId {
			s.Server.PortChannels[i] = *portChannel
			break
		}
	}
}

func (s *Server) updatePort(port *v1.InterfacePort) {
	for i, p := range s.Server.Ports {
		if p.PortId == port.PortId {
			s.Server.Ports[i] = *port
			break
		}
	}
}
