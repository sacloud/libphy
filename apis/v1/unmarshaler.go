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

package v1

import "encoding/json"

// UnmarshalJSON RaidLevelに文字列 or 数値を受け入れるための実装
//
// 定義上は文字列になっているが数値が入るケースがあるため暫定的にここで違いを吸収し文字列としてUnmarshalする
// see:https://github.com/sacloud/phy-api-go/issues/92
func (v *RaidLogicalVolume) UnmarshalJSON(data []byte) error {
	type alias struct {
		PhysicalDeviceIds []string                `json:"physical_device_ids"`
		RaidLevel         json.Number             `json:"raid_level"`
		Status            RaidLogicalVolumeStatus `json:"status"`
		VolumeId          string                  `json:"volume_id"`
	}

	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = RaidLogicalVolume{
		PhysicalDeviceIds: tmp.PhysicalDeviceIds,
		RaidLevel:         tmp.RaidLevel.String(),
		Status:            tmp.Status,
		VolumeId:          tmp.VolumeId,
	}
	return nil
}
