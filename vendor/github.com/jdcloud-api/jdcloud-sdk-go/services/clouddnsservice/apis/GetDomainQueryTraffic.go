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

type GetDomainQueryTrafficRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 域名ID  */
    DomainId string `json:"domainId"`

    /* 域名  */
    DomainName string `json:"domainName"`

    /* 起始时间, UTC时间例如2017-11-10T23:00:00Z  */
    Start string `json:"start"`

    /* 终止时间, UTC时间例如2017-11-10T23:00:00Z  */
    End string `json:"end"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID (Required)
 * param domainName: 域名 (Required)
 * param start: 起始时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 * param end: 终止时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetDomainQueryTrafficRequest(
    regionId string,
    domainId string,
    domainName string,
    start string,
    end string,
) *GetDomainQueryTrafficRequest {

	return &GetDomainQueryTrafficRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/queryTraffic",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DomainId: domainId,
        DomainName: domainName,
        Start: start,
        End: end,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param domainId: 域名ID (Required)
 * param domainName: 域名 (Required)
 * param start: 起始时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 * param end: 终止时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 */
func NewGetDomainQueryTrafficRequestWithAllParams(
    regionId string,
    domainId string,
    domainName string,
    start string,
    end string,
) *GetDomainQueryTrafficRequest {

    return &GetDomainQueryTrafficRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/queryTraffic",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DomainId: domainId,
        DomainName: domainName,
        Start: start,
        End: end,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetDomainQueryTrafficRequestWithoutParam() *GetDomainQueryTrafficRequest {

    return &GetDomainQueryTrafficRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/queryTraffic",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *GetDomainQueryTrafficRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID(Required) */
func (r *GetDomainQueryTrafficRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param domainName: 域名(Required) */
func (r *GetDomainQueryTrafficRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

/* param start: 起始时间, UTC时间例如2017-11-10T23:00:00Z(Required) */
func (r *GetDomainQueryTrafficRequest) SetStart(start string) {
    r.Start = start
}

/* param end: 终止时间, UTC时间例如2017-11-10T23:00:00Z(Required) */
func (r *GetDomainQueryTrafficRequest) SetEnd(end string) {
    r.End = end
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetDomainQueryTrafficRequest) GetRegionId() string {
    return r.RegionId
}

type GetDomainQueryTrafficResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetDomainQueryTrafficResult `json:"result"`
}

type GetDomainQueryTrafficResult struct {
    Time []int `json:"time"`
    Unit string `json:"unit"`
    Traffic []float64 `json:"traffic"`
}