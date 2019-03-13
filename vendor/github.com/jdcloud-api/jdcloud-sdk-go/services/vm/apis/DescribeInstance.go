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
    vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
)

type DescribeInstanceRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云主机ID  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceRequest(
    regionId string,
    instanceId string,
) *DescribeInstanceRequest {

	return &DescribeInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 */
func NewDescribeInstanceRequestWithAllParams(
    regionId string,
    instanceId string,
) *DescribeInstanceRequest {

    return &DescribeInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstanceRequestWithoutParam() *DescribeInstanceRequest {

    return &DescribeInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 云主机ID(Required) */
func (r *DescribeInstanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceResult `json:"result"`
}

type DescribeInstanceResult struct {
    Instance vm.Instance `json:"instance"`
}
