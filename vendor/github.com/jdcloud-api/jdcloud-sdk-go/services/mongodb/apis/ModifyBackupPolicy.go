// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type ModifyBackupPolicyRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* 备份时间，格式：HH:mmZ- HH:mmZ，只允许间隔时间为1小时的整点.  */
    PreferredBackupTime string `json:""`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param preferredBackupTime: 备份时间，格式：HH:mmZ- HH:mmZ，只允许间隔时间为1小时的整点. (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyBackupPolicyRequest(
    regionId string,
    instanceId string,
    preferredBackupTime string,
) *ModifyBackupPolicyRequest {

	return &ModifyBackupPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/backupPolicy",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        PreferredBackupTime: preferredBackupTime,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param preferredBackupTime: 备份时间，格式：HH:mmZ- HH:mmZ，只允许间隔时间为1小时的整点. (Required)
 */
func NewModifyBackupPolicyRequestWithAllParams(
    regionId string,
    instanceId string,
    preferredBackupTime string,
) *ModifyBackupPolicyRequest {

    return &ModifyBackupPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/backupPolicy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        PreferredBackupTime: preferredBackupTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyBackupPolicyRequestWithoutParam() *ModifyBackupPolicyRequest {

    return &ModifyBackupPolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/backupPolicy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyBackupPolicyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: Instance ID(Required) */
func (r *ModifyBackupPolicyRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param preferredBackupTime: 备份时间，格式：HH:mmZ- HH:mmZ，只允许间隔时间为1小时的整点.(Required) */
func (r *ModifyBackupPolicyRequest) SetPreferredBackupTime(preferredBackupTime string) {
    r.PreferredBackupTime = preferredBackupTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyBackupPolicyRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyBackupPolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyBackupPolicyResult `json:"result"`
}

type ModifyBackupPolicyResult struct {
    PreferredBackupPeriod string `json:"preferredBackupPeriod"`
    PreferredBackupWindow string `json:"preferredBackupWindow"`
    BackupRetentionPeriod string `json:"backupRetentionPeriod"`
}