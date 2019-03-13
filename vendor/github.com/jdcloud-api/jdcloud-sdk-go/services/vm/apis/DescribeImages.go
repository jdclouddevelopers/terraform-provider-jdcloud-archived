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

type DescribeImagesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 镜像来源，如果没有指定ids参数，此参数必传；取值范围：public、shared、thirdparty、private (Optional) */
    ImageSource *string `json:"imageSource"`

    /* 操作系统平台，取值范围：Windows Server、CentOS、Ubuntu (Optional) */
    Platform *string `json:"platform"`

    /* 镜像ID列表，如果指定了此参数，其它参数可为空 (Optional) */
    Ids []string `json:"ids"`

    /* 镜像支持的系统盘类型，[localDisk,cloudDisk] (Optional) */
    RootDeviceType *string `json:"rootDeviceType"`

    /* <a href="http://docs.jdcloud.com/virtual-machines/api/image_status">参考镜像状态</a> (Optional) */
    Status *string `json:"status"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeImagesRequest(
    regionId string,
) *DescribeImagesRequest {

	return &DescribeImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/images",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param imageSource: 镜像来源，如果没有指定ids参数，此参数必传；取值范围：public、shared、thirdparty、private (Optional)
 * param platform: 操作系统平台，取值范围：Windows Server、CentOS、Ubuntu (Optional)
 * param ids: 镜像ID列表，如果指定了此参数，其它参数可为空 (Optional)
 * param rootDeviceType: 镜像支持的系统盘类型，[localDisk,cloudDisk] (Optional)
 * param status: <a href="http://docs.jdcloud.com/virtual-machines/api/image_status">参考镜像状态</a> (Optional)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 */
func NewDescribeImagesRequestWithAllParams(
    regionId string,
    imageSource *string,
    platform *string,
    ids []string,
    rootDeviceType *string,
    status *string,
    pageNumber *int,
    pageSize *int,
) *DescribeImagesRequest {

    return &DescribeImagesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ImageSource: imageSource,
        Platform: platform,
        Ids: ids,
        RootDeviceType: rootDeviceType,
        Status: status,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeImagesRequestWithoutParam() *DescribeImagesRequest {

    return &DescribeImagesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/images",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeImagesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param imageSource: 镜像来源，如果没有指定ids参数，此参数必传；取值范围：public、shared、thirdparty、private(Optional) */
func (r *DescribeImagesRequest) SetImageSource(imageSource string) {
    r.ImageSource = &imageSource
}

/* param platform: 操作系统平台，取值范围：Windows Server、CentOS、Ubuntu(Optional) */
func (r *DescribeImagesRequest) SetPlatform(platform string) {
    r.Platform = &platform
}

/* param ids: 镜像ID列表，如果指定了此参数，其它参数可为空(Optional) */
func (r *DescribeImagesRequest) SetIds(ids []string) {
    r.Ids = ids
}

/* param rootDeviceType: 镜像支持的系统盘类型，[localDisk,cloudDisk](Optional) */
func (r *DescribeImagesRequest) SetRootDeviceType(rootDeviceType string) {
    r.RootDeviceType = &rootDeviceType
}

/* param status: <a href="http://docs.jdcloud.com/virtual-machines/api/image_status">参考镜像状态</a>(Optional) */
func (r *DescribeImagesRequest) SetStatus(status string) {
    r.Status = &status
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeImagesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[10, 100](Optional) */
func (r *DescribeImagesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeImagesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeImagesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeImagesResult `json:"result"`
}

type DescribeImagesResult struct {
    Images []vm.Image `json:"images"`
    TotalCount int `json:"totalCount"`
}
