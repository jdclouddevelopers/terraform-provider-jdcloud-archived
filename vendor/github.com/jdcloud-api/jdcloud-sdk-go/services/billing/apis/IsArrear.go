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

type IsArrearRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    Pin *string `json:"pin"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewIsArrearRequest(
    regionId string,
) *IsArrearRequest {

	return &IsArrearRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/isArrear",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param pin:  (Optional)
 */
func NewIsArrearRequestWithAllParams(
    regionId string,
    pin *string,
) *IsArrearRequest {

    return &IsArrearRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/isArrear",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewIsArrearRequestWithoutParam() *IsArrearRequest {

    return &IsArrearRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/isArrear",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *IsArrearRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: (Optional) */
func (r *IsArrearRequest) SetPin(pin string) {
    r.Pin = &pin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r IsArrearRequest) GetRegionId() string {
    return r.RegionId
}

type IsArrearResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result IsArrearResult `json:"result"`
}

type IsArrearResult struct {
    Result bool `json:"result"`
}