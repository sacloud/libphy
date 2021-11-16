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

package phy

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	v1 "github.com/sacloud/phy-go/apis/v1"
	"github.com/stretchr/testify/require"
)

const (
	EnvNameServerID                  = "PHY_TEST_SERVER_ID"
	EnvNameDisableServiceAPI         = "PHY_TEST_DISABLE_SERVICE_API"
	EnvNameDisableDedicatedSubnetAPI = "PHY_TEST_DISABLE_DEDICATED_SUBNET_API"
	EnvNameDisablePrivateNetworkAPI  = "PHY_TEST_DISABLE_PRIVATE_NETWORK_API"
	EnvNameDisableServerAPI          = "PHY_TEST_DISABLE_SERVER_API"
)

var serverPowerControlWaitDuration = time.Duration(1 * time.Minute)

// TestAcc_ServiceAPI サービスAPIのテスト
//
// このテストの実行にはサービスが少なくとも1つ必要です。
func TestAcc_ServiceAPI(t *testing.T) {
	onlyAcceptanceTest(t)
	skipIfEnvIsSet(t, EnvNameDisableServiceAPI)

	client := testClient(t)
	serviceOp := NewServiceOp(client)

	ctx := context.Background()

	// List
	limit := v1.Limit(1)
	found, err := serviceOp.List(ctx, &v1.ListServicesParams{
		Limit: &limit,
	})
	require.NoError(t, err)
	// require.Equal(t, 1, found.Meta.Count) 取得した件数ではなく総件数
	require.Equal(t, 1, len(found.Services))

	// Read
	serviceId := v1.ServiceId(found.Services[0].ServiceId)
	read, err := serviceOp.Read(ctx, serviceId)
	require.NoError(t, err)
	require.NotNil(t, read)

	// Update
	originalDesc := ""
	if read.Description != nil {
		originalDesc = *read.Description
	}
	updDesc := originalDesc + "-upd"

	upd, err := serviceOp.Update(ctx, serviceId, v1.UpdateServiceParameter{
		Description: &updDesc,
		Nickname:    read.Nickname,
	})
	require.NoError(t, err)
	require.EqualValues(t, updDesc, *upd.Description)

	// 元に戻しておく
	upd2, err := serviceOp.Update(ctx, serviceId, v1.UpdateServiceParameter{
		Description: &originalDesc,
		Nickname:    read.Nickname,
	})
	require.NoError(t, err)
	require.EqualValues(t, originalDesc, *upd2.Description)
}

// TestAcc_PrivateNetworkAPI ローカルネットワークAPIのテスト
//
// このテストの実行にはローカルネットワークが少なくとも1つ必要です。
func TestAcc_PrivateNetworkAPI(t *testing.T) {
	onlyAcceptanceTest(t)
	skipIfEnvIsSet(t, EnvNameDisablePrivateNetworkAPI)

	client := testClient(t)
	privateNetworkOp := NewPrivateNetworkOp(client)

	ctx := context.Background()

	// List
	limit := v1.Limit(1)
	found, err := privateNetworkOp.List(ctx, &v1.ListPrivateNetworksParams{
		Limit: &limit,
	})
	require.NoError(t, err)
	require.Len(t, found.PrivateNetworks, 1)

	privateNetworkId := v1.PrivateNetworkId(found.PrivateNetworks[0].PrivateNetworkId)

	// Read
	read, err := privateNetworkOp.Read(ctx, privateNetworkId)
	require.NoError(t, err)
	require.NotNil(t, read)
}

// TestAcc_DedicatedSubnetAPI グローバルネットワークのテスト
//
// このテストの実行にはグローバルネットワークが少なくとも1つ必要です。
func TestAcc_DedicatedSubnetAPI(t *testing.T) {
	onlyAcceptanceTest(t)
	skipIfEnvIsSet(t, EnvNameDisableDedicatedSubnetAPI)

	client := testClient(t)
	dedicatedSubnetOp := NewDedicatedSubnetOp(client)

	ctx := context.Background()

	// List
	limit := v1.Limit(1)
	found, err := dedicatedSubnetOp.List(ctx, &v1.ListDedicatedSubnetsParams{
		Limit: &limit,
	})
	require.NoError(t, err)
	require.Len(t, found.DedicatedSubnets, 1)

	id := v1.DedicatedSubnetId(found.DedicatedSubnets[0].DedicatedSubnetId)

	// Read
	read, err := dedicatedSubnetOp.Read(ctx, id, false)
	require.NoError(t, err)
	require.NotNil(t, read)
}

// TestAcc_ServerAPI サーバのテスト
//
// このテストの実行には"PHY_TEST_SERVER_ID"環境変数の指定が必要です。
//
// また、以下の更新系APIはテスト対象外です。
//   - OSInstall
//   - ConfigureBonding
//   - AssignNetwork
func TestAcc_ServerAPI(t *testing.T) {
	onlyAcceptanceTest(t)
	skipIfEnvIsSet(t, EnvNameDisableServerAPI)

	envId := strings.TrimSpace(os.Getenv(EnvNameServerID))
	if envId == "" {
		t.Skip("environment variable 'PHY_TEST_SERVER_ID' is not set. skip")
	}
	serverId := v1.ServerId(envId)

	client := testClient(t)
	serverOp := NewServerOp(client)
	ctx := context.Background()

	var portChannelId v1.PortChannelId
	var portId v1.PortId
	var port *v1.InterfacePort
	var currentPowerStatus v1.ServerPowerStatusStatus

	t.Run("List", func(t *testing.T) {
		limit := v1.Limit(1)
		found, err := serverOp.List(ctx, &v1.ListServersParams{
			Limit: &limit,
		})
		require.NoError(t, err)
		require.Len(t, found.Servers, 1)
	})

	t.Run("Read", func(t *testing.T) {
		read, err := serverOp.Read(ctx, serverId)
		require.NoError(t, err)
		require.NotNil(t, read)

		// 後のテストのためにIDを保持しておく
		portChannelId = v1.PortChannelId(read.PortChannels[0].PortChannelId)
		portId = v1.PortId(read.Ports[0].PortId)
	})

	t.Run("ListOSImages", func(t *testing.T) {
		osImages, err := serverOp.ListOSImages(ctx, serverId)
		require.NoError(t, err)
		require.NotEmpty(t, osImages)
	})

	t.Run("ReadPortChannel", func(t *testing.T) {
		portChannel, err := serverOp.ReadPortChannel(ctx, serverId, portChannelId)
		require.NoError(t, err)
		require.NotNil(t, portChannel)
	})

	t.Run("ReadPort", func(t *testing.T) {
		read, err := serverOp.ReadPort(ctx, serverId, portId)
		require.NoError(t, err)
		require.NotNil(t, read)

		// 後のテストのために保持しておく
		port = read
	})

	t.Run("UpdatePort", func(t *testing.T) {
		originalPortName := port.Nickname
		updatedPort, err := serverOp.UpdatePort(ctx, serverId, portId, v1.UpdateServerPortParameter{
			Nickname: originalPortName + "-upd",
		})
		require.NoError(t, err)
		require.Equal(t, originalPortName+"-upd", updatedPort.Nickname)

		// 元に戻す
		_, err = serverOp.UpdatePort(ctx, serverId, portId, v1.UpdateServerPortParameter{
			Nickname: originalPortName,
		})
		require.NoError(t, err)
	})

	// EnablePort
	t.Run("EnablePort", func(t *testing.T) {
		current := port.Enabled

		updated, err := serverOp.EnablePort(ctx, serverId, portId, !current)
		require.NoError(t, err)
		require.NotNil(t, updated)
		require.Equal(t, updated.Enabled, !current)

		// 元に戻す
		_, err = serverOp.EnablePort(ctx, serverId, portId, !current)
		require.NoError(t, err)
	})

	t.Run("ReadTrafficByPort", func(t *testing.T) {
		trafficData, err := serverOp.ReadTrafficByPort(ctx, serverId, portId, v1.ReadServerTrafficByPortParams{})
		require.NoError(t, err)
		require.NotNil(t, trafficData)
		require.NotEmpty(t, trafficData.Transmit)
		require.NotEmpty(t, trafficData.Receive)
	})

	t.Run("ReadRAIDStatus", func(t *testing.T) {
		// 電源ONの場合のみ
		if currentPowerStatus != v1.ServerPowerStatusStatusOn {
			t.Skipf("Server must be running, skip")
		}
		raidStatus, err := serverOp.ReadRAIDStatus(ctx, serverId, true)
		require.NoError(t, err)
		require.NotNil(t, raidStatus)
	})

	// PowerControl/ReadPowerStatus
	t.Run("PowerControl/ReadPowerStatus", func(t *testing.T) {
		powerStatus, err := serverOp.ReadPowerStatus(ctx, serverId)
		require.NoError(t, err)
		require.NotNil(t, powerStatus)

		currentPowerStatus = powerStatus.Status

		var operation v1.ServerPowerOperations
		var reverse v1.ServerPowerStatusStatus
		switch currentPowerStatus {
		case v1.ServerPowerStatusStatusOn:
			operation = v1.ServerPowerOperationsSoft
			reverse = v1.ServerPowerStatusStatusOff
		case v1.ServerPowerStatusStatusOff:
			operation = v1.ServerPowerOperationsOn
			reverse = v1.ServerPowerStatusStatusOn
		}

		require.NoError(t, serverOp.PowerControl(ctx, serverId, operation))
		// 電源状態が反映されるまで少し待つ
		time.Sleep(serverPowerControlWaitDuration)

		powerStatus, err = serverOp.ReadPowerStatus(ctx, serverId)
		require.NoError(t, err)
		require.Equal(t, reverse, powerStatus.Status)

		// 元に戻しておく
		switch reverse {
		case v1.ServerPowerStatusStatusOn:
			operation = v1.ServerPowerOperationsSoft
		case v1.ServerPowerStatusStatusOff:
			operation = v1.ServerPowerOperationsOn
		}
		require.NoError(t, serverOp.PowerControl(ctx, serverId, operation))
	})

}
