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

type UpdateAlarmRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 规则id  */
    AlarmId string `json:"alarmId"`

    /* 统计方法：平均值=avg、最大值=max、最小值=min、总和=sum  */
    Calculation string `json:"calculation"`

    /* 通知的联系组，如 [“联系组1”,”联系组2”] (Optional) */
    ContactGroups []string `json:"contactGroups"`

    /* 通知的联系人，如 [“联系人1”,”联系人2”] (Optional) */
    ContactPersons []string `json:"contactPersons"`

    /* 取样频次 (Optional) */
    DownSample *string `json:"downSample"`

    /* 根据产品线查询可用监控项列表 接口 返回的Metric字段  */
    Metric string `json:"metric"`

    /* 通知周期 单位：小时 (Optional) */
    NoticePeriod *int `json:"noticePeriod"`

    /* >=、>、<、<=、==、!=  */
    Operation string `json:"operation"`

    /* 统计周期（单位：分钟），可选值：2,5,15,30,60  */
    Period int `json:"period"`

    /* 产品名称  */
    ServiceCode string `json:"serviceCode"`

    /* 阈值  */
    Threshold float64 `json:"threshold"`

    /* 连续多少次后报警，可选值:1,2,3,5  */
    Times int `json:"times"`
}

/*
 * param regionId: 地域 Id (Required)
 * param alarmId: 规则id (Required)
 * param calculation: 统计方法：平均值=avg、最大值=max、最小值=min、总和=sum (Required)
 * param metric: 根据产品线查询可用监控项列表 接口 返回的Metric字段 (Required)
 * param operation: >=、>、<、<=、==、!= (Required)
 * param period: 统计周期（单位：分钟），可选值：2,5,15,30,60 (Required)
 * param serviceCode: 产品名称 (Required)
 * param threshold: 阈值 (Required)
 * param times: 连续多少次后报警，可选值:1,2,3,5 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateAlarmRequest(
    regionId string,
    alarmId string,
    calculation string,
    metric string,
    operation string,
    period int,
    serviceCode string,
    threshold float64,
    times int,
) *UpdateAlarmRequest {

	return &UpdateAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/alarms/{alarmId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AlarmId: alarmId,
        Calculation: calculation,
        Metric: metric,
        Operation: operation,
        Period: period,
        ServiceCode: serviceCode,
        Threshold: threshold,
        Times: times,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param alarmId: 规则id (Required)
 * param calculation: 统计方法：平均值=avg、最大值=max、最小值=min、总和=sum (Required)
 * param contactGroups: 通知的联系组，如 [“联系组1”,”联系组2”] (Optional)
 * param contactPersons: 通知的联系人，如 [“联系人1”,”联系人2”] (Optional)
 * param downSample: 取样频次 (Optional)
 * param metric: 根据产品线查询可用监控项列表 接口 返回的Metric字段 (Required)
 * param noticePeriod: 通知周期 单位：小时 (Optional)
 * param operation: >=、>、<、<=、==、!= (Required)
 * param period: 统计周期（单位：分钟），可选值：2,5,15,30,60 (Required)
 * param serviceCode: 产品名称 (Required)
 * param threshold: 阈值 (Required)
 * param times: 连续多少次后报警，可选值:1,2,3,5 (Required)
 */
func NewUpdateAlarmRequestWithAllParams(
    regionId string,
    alarmId string,
    calculation string,
    contactGroups []string,
    contactPersons []string,
    downSample *string,
    metric string,
    noticePeriod *int,
    operation string,
    period int,
    serviceCode string,
    threshold float64,
    times int,
) *UpdateAlarmRequest {

    return &UpdateAlarmRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarms/{alarmId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AlarmId: alarmId,
        Calculation: calculation,
        ContactGroups: contactGroups,
        ContactPersons: contactPersons,
        DownSample: downSample,
        Metric: metric,
        NoticePeriod: noticePeriod,
        Operation: operation,
        Period: period,
        ServiceCode: serviceCode,
        Threshold: threshold,
        Times: times,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateAlarmRequestWithoutParam() *UpdateAlarmRequest {

    return &UpdateAlarmRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarms/{alarmId}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *UpdateAlarmRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param alarmId: 规则id(Required) */
func (r *UpdateAlarmRequest) SetAlarmId(alarmId string) {
    r.AlarmId = alarmId
}

/* param calculation: 统计方法：平均值=avg、最大值=max、最小值=min、总和=sum(Required) */
func (r *UpdateAlarmRequest) SetCalculation(calculation string) {
    r.Calculation = calculation
}

/* param contactGroups: 通知的联系组，如 [“联系组1”,”联系组2”](Optional) */
func (r *UpdateAlarmRequest) SetContactGroups(contactGroups []string) {
    r.ContactGroups = contactGroups
}

/* param contactPersons: 通知的联系人，如 [“联系人1”,”联系人2”](Optional) */
func (r *UpdateAlarmRequest) SetContactPersons(contactPersons []string) {
    r.ContactPersons = contactPersons
}

/* param downSample: 取样频次(Optional) */
func (r *UpdateAlarmRequest) SetDownSample(downSample string) {
    r.DownSample = &downSample
}

/* param metric: 根据产品线查询可用监控项列表 接口 返回的Metric字段(Required) */
func (r *UpdateAlarmRequest) SetMetric(metric string) {
    r.Metric = metric
}

/* param noticePeriod: 通知周期 单位：小时(Optional) */
func (r *UpdateAlarmRequest) SetNoticePeriod(noticePeriod int) {
    r.NoticePeriod = &noticePeriod
}

/* param operation: >=、>、<、<=、==、!=(Required) */
func (r *UpdateAlarmRequest) SetOperation(operation string) {
    r.Operation = operation
}

/* param period: 统计周期（单位：分钟），可选值：2,5,15,30,60(Required) */
func (r *UpdateAlarmRequest) SetPeriod(period int) {
    r.Period = period
}

/* param serviceCode: 产品名称(Required) */
func (r *UpdateAlarmRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param threshold: 阈值(Required) */
func (r *UpdateAlarmRequest) SetThreshold(threshold float64) {
    r.Threshold = threshold
}

/* param times: 连续多少次后报警，可选值:1,2,3,5(Required) */
func (r *UpdateAlarmRequest) SetTimes(times int) {
    r.Times = times
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateAlarmRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateAlarmResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateAlarmResult `json:"result"`
}

type UpdateAlarmResult struct {
    AlarmId string `json:"alarmId"`
}