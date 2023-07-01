package polardb

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

// ModifyBackupPolicy invokes the polardb.ModifyBackupPolicy API synchronously
func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (response *ModifyBackupPolicyResponse, err error) {
	response = CreateModifyBackupPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBackupPolicyWithChan invokes the polardb.ModifyBackupPolicy API asynchronously
func (client *Client) ModifyBackupPolicyWithChan(request *ModifyBackupPolicyRequest) (<-chan *ModifyBackupPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyBackupPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBackupPolicy(request)
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

// ModifyBackupPolicyWithCallback invokes the polardb.ModifyBackupPolicy API asynchronously
func (client *Client) ModifyBackupPolicyWithCallback(request *ModifyBackupPolicyRequest, callback func(response *ModifyBackupPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBackupPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyBackupPolicy(request)
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

// ModifyBackupPolicyRequest is the request struct for api ModifyBackupPolicy
type ModifyBackupPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DataLevel2BackupRetentionPeriod              string           `position:"Query" name:"DataLevel2BackupRetentionPeriod"`
	DataLevel1BackupPeriod                       string           `position:"Query" name:"DataLevel1BackupPeriod"`
	DataLevel2BackupPeriod                       string           `position:"Query" name:"DataLevel2BackupPeriod"`
	PreferredBackupPeriod                        string           `position:"Query" name:"PreferredBackupPeriod"`
	DataLevel1BackupRetentionPeriod              string           `position:"Query" name:"DataLevel1BackupRetentionPeriod"`
	BackupRetentionPolicyOnClusterDeletion       string           `position:"Query" name:"BackupRetentionPolicyOnClusterDeletion"`
	ResourceOwnerAccount                         string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId                                  string           `position:"Query" name:"DBClusterId"`
	OwnerAccount                                 string           `position:"Query" name:"OwnerAccount"`
	DataLevel2BackupAnotherRegionRetentionPeriod string           `position:"Query" name:"DataLevel2BackupAnotherRegionRetentionPeriod"`
	OwnerId                                      requests.Integer `position:"Query" name:"OwnerId"`
	PreferredBackupTime                          string           `position:"Query" name:"PreferredBackupTime"`
	BackupRetentionPeriod                        string           `position:"Query" name:"BackupRetentionPeriod"`
	BackupFrequency                              string           `position:"Query" name:"BackupFrequency"`
	DataLevel1BackupFrequency                    string           `position:"Query" name:"DataLevel1BackupFrequency"`
	DataLevel2BackupAnotherRegionRegion          string           `position:"Query" name:"DataLevel2BackupAnotherRegionRegion"`
	DataLevel1BackupTime                         string           `position:"Query" name:"DataLevel1BackupTime"`
}

// ModifyBackupPolicyResponse is the response struct for api ModifyBackupPolicy
type ModifyBackupPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyBackupPolicyRequest creates a request to invoke ModifyBackupPolicy API
func CreateModifyBackupPolicyRequest() (request *ModifyBackupPolicyRequest) {
	request = &ModifyBackupPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyBackupPolicy", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyBackupPolicyResponse creates a response to parse from ModifyBackupPolicy response
func CreateModifyBackupPolicyResponse() (response *ModifyBackupPolicyResponse) {
	response = &ModifyBackupPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
