package ess

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyEciScalingConfiguration invokes the ess.ModifyEciScalingConfiguration API synchronously
func (client *Client) ModifyEciScalingConfiguration(request *ModifyEciScalingConfigurationRequest) (response *ModifyEciScalingConfigurationResponse, err error) {
	response = CreateModifyEciScalingConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyEciScalingConfigurationWithChan invokes the ess.ModifyEciScalingConfiguration API asynchronously
func (client *Client) ModifyEciScalingConfigurationWithChan(request *ModifyEciScalingConfigurationRequest) (<-chan *ModifyEciScalingConfigurationResponse, <-chan error) {
	responseChan := make(chan *ModifyEciScalingConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyEciScalingConfiguration(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyEciScalingConfigurationWithCallback invokes the ess.ModifyEciScalingConfiguration API asynchronously
func (client *Client) ModifyEciScalingConfigurationWithCallback(request *ModifyEciScalingConfigurationRequest, callback func(response *ModifyEciScalingConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyEciScalingConfigurationResponse
		var err error
		defer close(result)
		response, err = client.ModifyEciScalingConfiguration(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyEciScalingConfigurationRequest is the request struct for api ModifyEciScalingConfiguration
type ModifyEciScalingConfigurationRequest struct {
	*requests.RpcRequest
	Container                     *[]ModifyEciScalingConfigurationContainer               `position:"Query" name:"Container"  type:"Repeated"`
	NtpServer                     *[]string                                               `position:"Query" name:"NtpServer"  type:"Repeated"`
	SpotPriceLimit                requests.Float                                          `position:"Query" name:"SpotPriceLimit"`
	EnableSls                     requests.Boolean                                        `position:"Query" name:"EnableSls"`
	ResourceGroupId               string                                                  `position:"Query" name:"ResourceGroupId"`
	HostName                      string                                                  `position:"Query" name:"HostName"`
	ActiveDeadlineSeconds         requests.Integer                                        `position:"Query" name:"ActiveDeadlineSeconds"`
	EgressBandwidth               requests.Integer                                        `position:"Query" name:"EgressBandwidth"`
	DnsConfigSearch               *[]string                                               `position:"Query" name:"DnsConfigSearch"  type:"Repeated"`
	HostAliase                    *[]ModifyEciScalingConfigurationHostAliase              `position:"Query" name:"HostAliase"  type:"Repeated"`
	ImageSnapshotId               string                                                  `position:"Query" name:"ImageSnapshotId"`
	Tag                           *[]ModifyEciScalingConfigurationTag                     `position:"Query" name:"Tag"  type:"Repeated"`
	CpuOptionsThreadsPerCore      requests.Integer                                        `position:"Query" name:"CpuOptionsThreadsPerCore"`
	Ipv6AddressCount              requests.Integer                                        `position:"Query" name:"Ipv6AddressCount"`
	Cpu                           requests.Float                                          `position:"Query" name:"Cpu"`
	OwnerId                       requests.Integer                                        `position:"Query" name:"OwnerId"`
	ScalingConfigurationName      string                                                  `position:"Query" name:"ScalingConfigurationName"`
	ScalingConfigurationId        string                                                  `position:"Query" name:"ScalingConfigurationId"`
	SpotStrategy                  string                                                  `position:"Query" name:"SpotStrategy"`
	Volume                        *[]ModifyEciScalingConfigurationVolume                  `position:"Query" name:"Volume"  type:"Repeated"`
	InstanceFamilyLevel           string                                                  `position:"Query" name:"InstanceFamilyLevel"`
	DnsConfigOption               *[]ModifyEciScalingConfigurationDnsConfigOption         `position:"Query" name:"DnsConfigOption"  type:"Repeated"`
	ContainersUpdateType          string                                                  `position:"Query" name:"ContainersUpdateType"`
	EphemeralStorage              requests.Integer                                        `position:"Query" name:"EphemeralStorage"`
	EipBandwidth                  requests.Integer                                        `position:"Query" name:"EipBandwidth"`
	CostOptimization              requests.Boolean                                        `position:"Query" name:"CostOptimization"`
	Memory                        requests.Float                                          `position:"Query" name:"Memory"`
	SecurityGroupId               string                                                  `position:"Query" name:"SecurityGroupId"`
	Description                   string                                                  `position:"Query" name:"Description"`
	IngressBandwidth              requests.Integer                                        `position:"Query" name:"IngressBandwidth"`
	DnsPolicy                     string                                                  `position:"Query" name:"DnsPolicy"`
	SecurityContextSysctl         *[]ModifyEciScalingConfigurationSecurityContextSysctl   `position:"Query" name:"SecurityContextSysctl"  type:"Repeated"`
	DnsConfigNameServer           *[]string                                               `position:"Query" name:"DnsConfigNameServer"  type:"Repeated"`
	InitContainer                 *[]ModifyEciScalingConfigurationInitContainer           `position:"Query" name:"InitContainer"  type:"Repeated"`
	TerminationGracePeriodSeconds requests.Integer                                        `position:"Query" name:"TerminationGracePeriodSeconds"`
	ImageRegistryCredential       *[]ModifyEciScalingConfigurationImageRegistryCredential `position:"Query" name:"ImageRegistryCredential"  type:"Repeated"`
	ResourceOwnerAccount          string                                                  `position:"Query" name:"ResourceOwnerAccount"`
	RestartPolicy                 string                                                  `position:"Query" name:"RestartPolicy"`
	CpuOptionsCore                requests.Integer                                        `position:"Query" name:"CpuOptionsCore"`
	RamRoleName                   string                                                  `position:"Query" name:"RamRoleName"`
	AcrRegistryInfo               *[]ModifyEciScalingConfigurationAcrRegistryInfo         `position:"Query" name:"AcrRegistryInfo"  type:"Repeated"`
	AutoMatchImageCache           requests.Boolean                                        `position:"Query" name:"AutoMatchImageCache"`
	LoadBalancerWeight            requests.Integer                                        `position:"Query" name:"LoadBalancerWeight"`
	ContainerGroupName            string                                                  `position:"Query" name:"ContainerGroupName"`
	AutoCreateEip                 requests.Boolean                                        `position:"Query" name:"AutoCreateEip"`
}

// ModifyEciScalingConfigurationContainer is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationContainer struct {
	Stdin                                 string                                                  `name:"Stdin"`
	Memory                                string                                                  `name:"Memory"`
	LivenessProbeExecCommand              *[]string                                               `name:"LivenessProbe.Exec.Command" type:"Repeated"`
	WorkingDir                            string                                                  `name:"WorkingDir"`
	ReadinessProbeHttpGetPort             string                                                  `name:"ReadinessProbe.HttpGet.Port"`
	ReadinessProbeFailureThreshold        string                                                  `name:"ReadinessProbe.FailureThreshold"`
	LivenessProbeHttpGetPort              string                                                  `name:"LivenessProbe.HttpGet.Port"`
	Arg                                   *[]string                                               `name:"Arg" type:"Repeated"`
	ReadinessProbeSuccessThreshold        string                                                  `name:"ReadinessProbe.SuccessThreshold"`
	VolumeMount                           *[]ModifyEciScalingConfigurationContainerVolumeMount    `name:"VolumeMount" type:"Repeated"`
	Image                                 string                                                  `name:"Image"`
	SecurityContextCapabilityAdd          *[]string                                               `name:"SecurityContext.Capability.Add" type:"Repeated"`
	ReadinessProbeInitialDelaySeconds     string                                                  `name:"ReadinessProbe.InitialDelaySeconds"`
	Cpu                                   string                                                  `name:"Cpu"`
	ReadinessProbeExecCommand             *[]string                                               `name:"ReadinessProbe.Exec.Command" type:"Repeated"`
	ReadinessProbeHttpGetScheme           string                                                  `name:"ReadinessProbe.HttpGet.Scheme"`
	ReadinessProbeHttpGetPath             string                                                  `name:"ReadinessProbe.HttpGet.Path"`
	Gpu                                   string                                                  `name:"Gpu"`
	StdinOnce                             string                                                  `name:"StdinOnce"`
	ImagePullPolicy                       string                                                  `name:"ImagePullPolicy"`
	Command                               *[]string                                               `name:"Command" type:"Repeated"`
	LivenessProbeSuccessThreshold         string                                                  `name:"LivenessProbe.SuccessThreshold"`
	SecurityContextRunAsUser              string                                                  `name:"SecurityContext.RunAsUser"`
	LivenessProbeHttpGetPath              string                                                  `name:"LivenessProbe.HttpGet.Path"`
	LivenessProbePeriodSeconds            string                                                  `name:"LivenessProbe.PeriodSeconds"`
	LivenessProbeInitialDelaySeconds      string                                                  `name:"LivenessProbe.InitialDelaySeconds"`
	LivenessProbeTimeoutSeconds           string                                                  `name:"LivenessProbe.TimeoutSeconds"`
	LivenessProbeTcpSocketPort            string                                                  `name:"LivenessProbe.TcpSocket.Port"`
	Port                                  *[]ModifyEciScalingConfigurationContainerPort           `name:"Port" type:"Repeated"`
	ReadinessProbePeriodSeconds           string                                                  `name:"ReadinessProbe.PeriodSeconds"`
	EnvironmentVar                        *[]ModifyEciScalingConfigurationContainerEnvironmentVar `name:"EnvironmentVar" type:"Repeated"`
	Tty                                   string                                                  `name:"Tty"`
	Name                                  string                                                  `name:"Name"`
	SecurityContextReadOnlyRootFilesystem string                                                  `name:"SecurityContext.ReadOnlyRootFilesystem"`
	LivenessProbeFailureThreshold         string                                                  `name:"LivenessProbe.FailureThreshold"`
	ReadinessProbeTimeoutSeconds          string                                                  `name:"ReadinessProbe.TimeoutSeconds"`
	ReadinessProbeTcpSocketPort           string                                                  `name:"ReadinessProbe.TcpSocket.Port"`
	LivenessProbeHttpGetScheme            string                                                  `name:"LivenessProbe.HttpGet.Scheme"`
}

// ModifyEciScalingConfigurationHostAliase is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationHostAliase struct {
	Hostname *[]string `name:"Hostname" type:"Repeated"`
	Ip       string    `name:"Ip"`
}

// ModifyEciScalingConfigurationTag is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ModifyEciScalingConfigurationVolume is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationVolume struct {
	DiskVolumeFsType                 string                                                                 `name:"DiskVolume.FsType"`
	NFSVolumePath                    string                                                                 `name:"NFSVolume.Path"`
	DiskVolumeDiskId                 string                                                                 `name:"DiskVolume.DiskId"`
	FlexVolumeFsType                 string                                                                 `name:"FlexVolume.FsType"`
	Type                             string                                                                 `name:"Type"`
	FlexVolumeDriver                 string                                                                 `name:"FlexVolume.Driver"`
	FlexVolumeOptions                string                                                                 `name:"FlexVolume.Options"`
	NFSVolumeServer                  string                                                                 `name:"NFSVolume.Server"`
	EmptyDirVolumeMedium             string                                                                 `name:"EmptyDirVolume.Medium"`
	HostPathVolumePath               string                                                                 `name:"HostPathVolume.Path"`
	Name                             string                                                                 `name:"Name"`
	ConfigFileVolumeConfigFileToPath *[]ModifyEciScalingConfigurationVolumeConfigFileVolumeConfigFileToPath `name:"ConfigFileVolumeConfigFileToPath" type:"Repeated"`
	DiskVolumeDiskSize               string                                                                 `name:"DiskVolume.DiskSize"`
	ConfigFileVolumeDefaultMode      string                                                                 `name:"ConfigFileVolumeDefaultMode"`
	HostPathVolumeType               string                                                                 `name:"HostPathVolume.Type"`
	NFSVolumeReadOnly                string                                                                 `name:"NFSVolume.ReadOnly"`
}

// ModifyEciScalingConfigurationDnsConfigOption is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationDnsConfigOption struct {
	Name  string `name:"Name"`
	Value string `name:"Value"`
}

// ModifyEciScalingConfigurationSecurityContextSysctl is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationSecurityContextSysctl struct {
	Name  string `name:"Name"`
	Value string `name:"Value"`
}

// ModifyEciScalingConfigurationInitContainer is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationInitContainer struct {
	Image                                 string                                                                   `name:"Image"`
	InitContainerEnvironmentVar           *[]ModifyEciScalingConfigurationInitContainerInitContainerEnvironmentVar `name:"InitContainerEnvironmentVar" type:"Repeated"`
	SecurityContextCapabilityAdd          *[]string                                                                `name:"SecurityContext.Capability.Add" type:"Repeated"`
	Memory                                string                                                                   `name:"Memory"`
	WorkingDir                            string                                                                   `name:"WorkingDir"`
	Cpu                                   string                                                                   `name:"Cpu"`
	Gpu                                   string                                                                   `name:"Gpu"`
	ImagePullPolicy                       string                                                                   `name:"ImagePullPolicy"`
	Command                               *[]string                                                                `name:"Command" type:"Repeated"`
	SecurityContextRunAsUser              string                                                                   `name:"SecurityContext.RunAsUser"`
	InitContainerPort                     *[]ModifyEciScalingConfigurationInitContainerInitContainerPort           `name:"InitContainerPort" type:"Repeated"`
	Arg                                   *[]string                                                                `name:"Arg" type:"Repeated"`
	Name                                  string                                                                   `name:"Name"`
	InitContainerVolumeMount              *[]ModifyEciScalingConfigurationInitContainerInitContainerVolumeMount    `name:"InitContainerVolumeMount" type:"Repeated"`
	SecurityContextReadOnlyRootFilesystem string                                                                   `name:"SecurityContext.ReadOnlyRootFilesystem"`
}

// ModifyEciScalingConfigurationImageRegistryCredential is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationImageRegistryCredential struct {
	Server   string `name:"Server"`
	Password string `name:"Password"`
	UserName string `name:"UserName"`
}

// ModifyEciScalingConfigurationAcrRegistryInfo is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationAcrRegistryInfo struct {
	InstanceName string    `name:"InstanceName"`
	InstanceId   string    `name:"InstanceId"`
	RegionId     string    `name:"RegionId"`
	Domain       *[]string `name:"Domain" type:"Repeated"`
}

// ModifyEciScalingConfigurationContainerVolumeMount is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationContainerVolumeMount struct {
	MountPath        string `name:"MountPath"`
	ReadOnly         string `name:"ReadOnly"`
	MountPropagation string `name:"MountPropagation"`
	Name             string `name:"Name"`
	SubPath          string `name:"SubPath"`
}

// ModifyEciScalingConfigurationContainerPort is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationContainerPort struct {
	Protocol string `name:"Protocol"`
	Port     string `name:"Port"`
}

// ModifyEciScalingConfigurationContainerEnvironmentVar is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationContainerEnvironmentVar struct {
	FieldRefFieldPath string `name:"FieldRef.FieldPath"`
	Value             string `name:"Value"`
	Key               string `name:"Key"`
}

// ModifyEciScalingConfigurationVolumeConfigFileVolumeConfigFileToPath is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationVolumeConfigFileVolumeConfigFileToPath struct {
	Mode    string `name:"Mode"`
	Path    string `name:"Path"`
	Content string `name:"Content"`
}

// ModifyEciScalingConfigurationInitContainerInitContainerEnvironmentVar is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationInitContainerInitContainerEnvironmentVar struct {
	FieldRefFieldPath string `name:"FieldRef.FieldPath"`
	Value             string `name:"Value"`
	Key               string `name:"Key"`
}

// ModifyEciScalingConfigurationInitContainerInitContainerPort is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationInitContainerInitContainerPort struct {
	Protocol string `name:"Protocol"`
	Port     string `name:"Port"`
}

// ModifyEciScalingConfigurationInitContainerInitContainerVolumeMount is a repeated param struct in ModifyEciScalingConfigurationRequest
type ModifyEciScalingConfigurationInitContainerInitContainerVolumeMount struct {
	MountPath        string `name:"MountPath"`
	ReadOnly         string `name:"ReadOnly"`
	MountPropagation string `name:"MountPropagation"`
	Name             string `name:"Name"`
	SubPath          string `name:"SubPath"`
}

// ModifyEciScalingConfigurationResponse is the response struct for api ModifyEciScalingConfiguration
type ModifyEciScalingConfigurationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyEciScalingConfigurationRequest creates a request to invoke ModifyEciScalingConfiguration API
func CreateModifyEciScalingConfigurationRequest() (request *ModifyEciScalingConfigurationRequest) {
	request = &ModifyEciScalingConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ModifyEciScalingConfiguration", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyEciScalingConfigurationResponse creates a response to parse from ModifyEciScalingConfiguration response
func CreateModifyEciScalingConfigurationResponse() (response *ModifyEciScalingConfigurationResponse) {
	response = &ModifyEciScalingConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
