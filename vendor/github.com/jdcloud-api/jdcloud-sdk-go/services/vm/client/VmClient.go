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

package client

import (
	"encoding/json"
	"errors"
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/apis"
)

type VmClient struct {
	core.JDCloudClient
}

func NewVmClient(credential *core.Credential) *VmClient {
	if credential == nil {
		return nil
	}

	config := core.NewConfig()
	config.SetEndpoint("vm.jdcloud-api.com")

	return &VmClient{
		core.JDCloudClient{
			Credential:  *credential,
			Config:      *config,
			ServiceName: "vm",
			Revision:    "1.0.8",
			Logger:      core.NewDefaultLogger(core.LogInfo),
		}}
}

func (c *VmClient) SetConfig(config *core.Config) {
	c.Config = *config
}

func (c *VmClient) SetLogger(logger core.Logger) {
	c.Logger = logger
}

/* 为一台云主机挂载一块数据盘(云硬盘)，云主机和云硬盘没有正在进行中的的任务时才可挂载。<br>
云主机状态必须是<b>running</b>或<b>stopped</b>状态。<br>
本地盘(local类型)做系统盘的云主机可挂载8块数据盘，云硬盘(cloud类型)做系统盘的云主机可挂载7块数据盘。
*/
func (c *VmClient) AttachDisk(request *vm.AttachDiskRequest) (*vm.AttachDiskResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.AttachDiskResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 查询镜像共享帐户列表，只允许操作您的个人私有镜像。
 */
func (c *VmClient) DescribeImageMembers(request *vm.DescribeImageMembersRequest) (*vm.DescribeImageMembersResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeImageMembersResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 查询镜像详情。
 */
func (c *VmClient) DescribeImage(request *vm.DescribeImageRequest) (*vm.DescribeImageResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeImageResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 云主机使用指定镜像重置云主机系统<br>
云主机的状态必须为<b>stopped</b>状态。<br>
若当前云主机的系统盘类型为local类型，那么更换的镜像必须为localDisk类型的镜像；同理若当前云主机的系统盘为cloud类型，那么更换的镜像必须为cloudDisk类型的镜像。可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeimages">DescribeImages</a>接口获得指定地域的镜像信息。<br>
若不指定镜像ID，默认使用当前主机的原镜像重置系统。<br>
指定的镜像必须能够支持当前主机的实例规格(instanceType)，否则会返回错误。可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeimageconstraints">DescribeImageConstraints</a>接口获得指定镜像支持的系统盘类型信息。
*/
func (c *VmClient) RebuildInstance(request *vm.RebuildInstanceRequest) (*vm.RebuildInstanceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.RebuildInstanceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 创建ssh密钥对。公钥部分存储在京东云，并返回未加密的 PEM 编码的 PKCS#8 格式私钥，您只有一次机会保存您的私钥。请妥善保管。<br>
若传入已存在的密钥名称，会返回错误。
*/
func (c *VmClient) CreateKeypair(request *vm.CreateKeypairRequest) (*vm.CreateKeypairResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.CreateKeypairResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 重启单个云主机，只能重启<b>running</b>状态的云主机，云主机没有正在进行中的任务才可重启。
 */
func (c *VmClient) RebootInstance(request *vm.RebootInstanceRequest) (*vm.RebootInstanceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.RebootInstanceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 共享镜像，只允许操作您的个人私有镜像，单个镜像最多可共享给20个京东云帐户。<br>
整机镜像目前不支持共享。
*/
func (c *VmClient) ShareImage(request *vm.ShareImageRequest) (*vm.ShareImageResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.ShareImageResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 批量查询云主机内网IP地址，查询的是主网卡内网主IP地址。 */
func (c *VmClient) DescribeInstancePrivateIpAddress(request *vm.DescribeInstancePrivateIpAddressRequest) (*vm.DescribeInstancePrivateIpAddressResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeInstancePrivateIpAddressResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 查询镜像信息列表。<br>
通过此接口可以查询到京东云官方镜像、第三方镜像、私有镜像、或其他用户共享给您的镜像。<br>
此接口支持分页查询，默认每页20条。
*/
func (c *VmClient) DescribeImages(request *vm.DescribeImagesRequest) (*vm.DescribeImagesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeImagesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 云主机挂载一块弹性网卡。<br>
云主机状态必须为<b>running</b>或<b>stopped</b>状态，并且没有正在进行中的任务才可操作。<br>
弹性网卡上如果绑定了公网IP，那么公网IP所在az需要与云主机的az保持一致，或者公网IP属于全可用区，才可挂载。<br>
云主机挂载弹性网卡的数量，不能超过实例规格的限制。可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeinstancetypes">DescribeInstanceTypes</a>接口获得指定规格可挂载弹性网卡的数量上限。<br>
弹性网卡与云主机必须在相同vpc下。
*/
func (c *VmClient) AttachNetworkInterface(request *vm.AttachNetworkInterfaceRequest) (*vm.AttachNetworkInterfaceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.AttachNetworkInterfaceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 导入由其他工具生成的密钥对的公钥部分。<br>
若传入已存在的密钥名称，会返回错误。
*/
func (c *VmClient) ImportKeypair(request *vm.ImportKeypairRequest) (*vm.ImportKeypairResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.ImportKeypairResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 停止单个云主机，只能停止<b>running</b>状态的云主机，云主机没有正在进行中的任务才可停止
 */
func (c *VmClient) StopInstance(request *vm.StopInstanceRequest) (*vm.StopInstanceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.StopInstanceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 修改镜像信息，包括名称、描述；只允许操作您的个人私有镜像。
 */
func (c *VmClient) ModifyImageAttribute(request *vm.ModifyImageAttributeRequest) (*vm.ModifyImageAttributeResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.ModifyImageAttributeResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 批量查询镜像的实例规格限制。<br>
通过此接口可以查看镜像不支持的实例规格。只有官方镜像、第三方镜像有实例规格的限制，个人的私有镜像没有此限制。
*/
func (c *VmClient) DescribeImageConstraintsBatch(request *vm.DescribeImageConstraintsBatchRequest) (*vm.DescribeImageConstraintsBatchResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeImageConstraintsBatchResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 批量查询云主机状态 */
func (c *VmClient) DescribeInstanceStatus(request *vm.DescribeInstanceStatusRequest) (*vm.DescribeInstanceStatusResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeInstanceStatusResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 获取云主机vnc，用于连接管理云主机。<br>
vnc地址的有效期为1个小时，调用接口获取vnc地址后如果1个小时内没有使用，vnc地址自动失效，再次使用需要重新获取。
*/
func (c *VmClient) DescribeInstanceVncUrl(request *vm.DescribeInstanceVncUrlRequest) (*vm.DescribeInstanceVncUrlResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeInstanceVncUrlResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 删除一个私有镜像，只允许操作您的个人私有镜像。<br>
若镜像已共享给其他用户，需先取消共享才可删除。
*/
func (c *VmClient) DeleteImage(request *vm.DeleteImageRequest) (*vm.DeleteImageResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DeleteImageResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 修改虚机弹性网卡属性，包括是否随云主机一起删除。<br>
不能修改主网卡。
*/
func (c *VmClient) ModifyInstanceNetworkAttribute(request *vm.ModifyInstanceNetworkAttributeRequest) (*vm.ModifyInstanceNetworkAttributeResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.ModifyInstanceNetworkAttributeResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 查询一台云主机的详细信息
 */
func (c *VmClient) DescribeInstance(request *vm.DescribeInstanceRequest) (*vm.DescribeInstanceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeInstanceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 云主机变更实例规格<br>
云主机的状态必须为<b>stopped</b>状态。<br>
16年创建的云硬盘做系统盘的主机，一代与二代实例规格不允许相互调整。<br>
本地盘(local类型)做系统盘的主机，一代与二代实例规格不允许相互调整。<br>
使用高可用组(Ag)创建的主机，一代与二代实例规格不允许相互调整。<br>
云硬盘(cloud类型)做系统盘的主机，一代与二代实例规格允许相互调整。<br>
如果当前主机中的弹性网卡数量，大于新实例规格允许的弹性网卡数量，会返回错误。可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeinstancetypes">DescribeInstanceTypes</a>接口获得指定地域及可用区下的实例规格信息。<br>
当前主机所使用的镜像，需要支持要变更的目标实例规格，否则返回错误。可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeimageconstraints">DescribeImageConstraints</a>接口获得指定镜像的实例规格限制信息。<br>
云主机欠费或到期时，无法更改实例规格。
*/
func (c *VmClient) ResizeInstance(request *vm.ResizeInstanceRequest) (*vm.ResizeInstanceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.ResizeInstanceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 修改云主机挂载的数据盘属性，包括是否随主机删除。
 */
func (c *VmClient) ModifyInstanceDiskAttribute(request *vm.ModifyInstanceDiskAttributeRequest) (*vm.ModifyInstanceDiskAttributeResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.ModifyInstanceDiskAttributeResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 查询配额，支持：云主机、镜像、密钥、模板、镜像共享
 */
func (c *VmClient) DescribeQuotas(request *vm.DescribeQuotasRequest) (*vm.DescribeQuotasResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeQuotasResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 删除ssh密钥对。
 */
func (c *VmClient) DeleteKeypair(request *vm.DeleteKeypairRequest) (*vm.DeleteKeypairResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DeleteKeypairResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 为云主机创建私有镜像。云主机状态必须为<b>stopped</b>。<br>
云主机没有正在进行中的任务才可制作镜像。<br>
制作镜像以备份系统盘为基础，在此之上可选择全部或部分挂载数据盘制作整机镜像（如不做任何更改将默认制作整机镜像），制作镜像过程会为所挂载云硬盘创建快照并与镜像关联。<br>
调用接口后，需要等待镜像状态变为<b>ready</b>后，才能正常使用镜像。
*/
func (c *VmClient) CreateImage(request *vm.CreateImageRequest) (*vm.CreateImageResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.CreateImageResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 批量查询云主机的详细信息<br>
此接口支持分页查询，默认每页20条。
*/
func (c *VmClient) DescribeInstances(request *vm.DescribeInstancesRequest) (*vm.DescribeInstancesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeInstancesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 创建一台或多台指定配置的云主机，创建模式分为三种：1.普通方式、2.使用高可用组、3.使用启动模板。三种方式创建云主机时参数的必传与非必传是不同的，具体请参考<a href="http://docs.jdcloud.com/virtual-machines/api/create_vm_sample">参数详细说明</a><br>
- 创建云主机需要通过实名认证
- 实例规格
    - 可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeinstancetypes">DescribeInstanceTypes</a>接口获得指定地域或可用区的规格信息。
    - 不能使用已下线、或已售馨的规格ID
- 镜像
    - Windows Server 2012 R2标准版 64位 中文版 SQL Server 2014 标准版 SP2内存需大于1GB；
    - Windows Server所有镜像CPU不可选超过64核CPU。
    - 可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeimages">DescribeImages</a>接口获得指定地域的镜像信息。
    - 选择的镜像必须支持选择的实例规格。可查询<a href="http://docs.jdcloud.com/virtual-machines/api/describeimageconstraints">DescribeImageConstraints</a>接口获得指定镜像的实例规格限制信息。<br>
- 网络配置
    - 指定主网卡配置信息
        - 必须指定subnetId
        - 可以指定elasticIp规格来约束创建的弹性IP，带宽取值范围[1-100]Mbps，步进1Mbps
        - 可以指定主网卡的内网主IP(primaryIpAddress)，此时maxCount只能为1
        - 安全组securityGroup需与子网Subnet在同一个私有网络VPC内
        - 一台云主机创建时必须指定一个安全组，至多指定5个安全组，如果没有指定安全组，默认使用默认安全组
        - 主网卡deviceIndex设置为1
- 存储
    - 系统盘
        - 磁盘分类，系统盘支持local或cloud
        - 磁盘大小
            - local：不能指定大小，默认为40GB
            - cloud：取值范围: 40-500GB，并且不能小于镜像的最小系统盘大小，如果没有指定，默认以镜像中的系统盘大小为准
        - 自动删除
            - 如果是local，默认自动删除，不能修改此属性
            - 如果是cloud类型的按配置计费的云硬盘，可以指定为True
    - 数据盘
        - 磁盘分类，数据盘仅支持cloud
        - 云硬盘类型可以选择ssd、premium-hdd
        - 磁盘大小
            - premium-hdd：范围[20,3000]GB，步长为10G
            - ssd：范围[20,1000]GB，步长为10G
        - 自动删除
            - 默认自动删除，如果是包年包月的数据盘或共享型数据盘，此参数不生效
            - 可以指定SnapshotId创建云硬盘
        - 可以从快照创建磁盘
    - local类型系统的云主机可以挂载8块云硬盘
    - cloud类型系统的云主机可以挂载7块云硬盘
- 计费
    - 弹性IP的计费模式，如果选择按用量类型可以单独设置，其它计费模式都以主机为准
    - 云硬盘的计费模式以主机为准
- 其他
    - 创建完成后，主机状态为running
    - 仅Linux系统云主机可以指定密钥
    - maxCount为最大努力，不保证一定能达到maxCount
    - 虚机的az会覆盖磁盘的az属性
- 密码
    - <a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>
*/
func (c *VmClient) CreateInstances(request *vm.CreateInstancesRequest) (*vm.CreateInstancesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.CreateInstancesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 修改云主机密码，主机没有正在进行中的任务时才可操作。<br>
修改密码后，需要重启云主机后生效。
*/
func (c *VmClient) ModifyInstancePassword(request *vm.ModifyInstancePasswordRequest) (*vm.ModifyInstancePasswordResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.ModifyInstancePasswordResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 取消共享镜像，只允许操作您的个人私有镜像。
 */
func (c *VmClient) UnShareImage(request *vm.UnShareImageRequest) (*vm.UnShareImageResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.UnShareImageResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 修改云主机部分信息，包括名称、描述。
 */
func (c *VmClient) ModifyInstanceAttribute(request *vm.ModifyInstanceAttributeRequest) (*vm.ModifyInstanceAttributeResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.ModifyInstanceAttributeResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 云主机解绑弹性公网IP，解绑的是主网卡、内网主IP对应的弹性公网IP。
 */
func (c *VmClient) DisassociateElasticIp(request *vm.DisassociateElasticIpRequest) (*vm.DisassociateElasticIpResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DisassociateElasticIpResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 镜像跨区复制，将私有镜像复制到其它地域下，只允许操作您的个人私有镜像。<br>
只支持rootDeviceType为cloudDisk的云硬盘系统盘镜像操作。
*/
func (c *VmClient) CopyImages(request *vm.CopyImagesRequest) (*vm.CopyImagesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.CopyImagesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 云主机缷载一块弹性网卡。<br>
云主机状态必须为<b>running</b>或<b>stopped</b>状态，并且没有正在进行中的任务才可操作。<br>
不能缷载主网卡。
*/
func (c *VmClient) DetachNetworkInterface(request *vm.DetachNetworkInterfaceRequest) (*vm.DetachNetworkInterfaceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DetachNetworkInterfaceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 查询镜像的实例规格限制。<br>
通过此接口可以查看镜像不支持的实例规格。只有官方镜像、第三方镜像有实例规格的限制，个人的私有镜像没有此限制。
*/
func (c *VmClient) DescribeImageConstraints(request *vm.DescribeImageConstraintsRequest) (*vm.DescribeImageConstraintsResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeImageConstraintsResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 启动单个云主机，只能启动<b>stopped</b>状态的云主机，云主机没有正在进行中的任务才可启动。<br>
只能启动正常计费状态的云主机。
*/
func (c *VmClient) StartInstance(request *vm.StartInstanceRequest) (*vm.StartInstanceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.StartInstanceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 删除按配置计费、或包年包月已到期的单个云主机。不能删除没有计费信息的云主机。<br>
云主机状态必须为运行<b>running</b>、停止<b>stopped</b>、错误<b>error</b>，同时云主机没有正在进行中的任务才可删除。<br>
包年包月未到期的云主机不能删除。<br>
如果主机中挂载的数据盘为按配置计费的云硬盘，并且不是共享型云硬盘，并且AutoDelete属性为true，那么数据盘会随主机一起删除。
 [MFA enabled] */
func (c *VmClient) DeleteInstance(request *vm.DeleteInstanceRequest) (*vm.DeleteInstanceResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DeleteInstanceResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 为云主机主网卡下的主内网IP绑定弹性公网IP。<br>
一台云主机只能绑定一个弹性公网IP(主网卡)，若主网卡已存在弹性公网IP，会返回错误。<br>
*/
func (c *VmClient) AssociateElasticIp(request *vm.AssociateElasticIpRequest) (*vm.AssociateElasticIpResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.AssociateElasticIpResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 批量查询密钥对。<br>
此接口支持分页查询，默认每页20条。
*/
func (c *VmClient) DescribeKeypairs(request *vm.DescribeKeypairsRequest) (*vm.DescribeKeypairsResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeKeypairsResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 云主机缷载数据盘，云主机和云硬盘没有正在进行中的任务时才可缷载。<br>
 */
func (c *VmClient) DetachDisk(request *vm.DetachDiskRequest) (*vm.DetachDiskResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DetachDiskResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 查询实例规格信息列表
 */
func (c *VmClient) DescribeInstanceTypes(request *vm.DescribeInstanceTypesRequest) (*vm.DescribeInstanceTypesResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DescribeInstanceTypesResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		return nil, err
	}

	return jdResp, err
}

/* 创建一个指定参数的启动模板，启动模板中包含创建云主机时的大部分配置参数，避免每次创建云主机时的重复性工作。<br>
如果是使用启动模板创建云主机，如果指定了某些参数与模板中的参数相冲突，那么新指定的参数会替换模板中的参数。<br>
如果是使用启动模板创建云主机，如果指定了镜像ID与模板中的镜像ID不一致，那么模板中的dataDisks参数会失效。<br>
如果使用高可用组(Ag)创建云主机，那么Ag所关联的模板中的参数都不可以被调整，只能以模板为准。
*/
func (c *VmClient) CreateInstanceTemplate(request *vm.CreateInstanceTemplateRequest) (*vm.CreateInstanceTemplateResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.CreateInstanceTemplateResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}

/* 删除一个启动模板
 */
func (c *VmClient) DeleteInstanceTemplate(request *vm.DeleteInstanceTemplateRequest) (*vm.DeleteInstanceTemplateResponse, error) {
	if request == nil {
		return nil, errors.New("Request object is nil. ")
	}
	resp, err := c.Send(request, c.ServiceName)
	if err != nil {
		return nil, err
	}

	jdResp := &vm.DeleteInstanceTemplateResponse{}
	err = json.Unmarshal(resp, jdResp)
	if err != nil {
		c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
		return nil, err
	}

	return jdResp, err
}
