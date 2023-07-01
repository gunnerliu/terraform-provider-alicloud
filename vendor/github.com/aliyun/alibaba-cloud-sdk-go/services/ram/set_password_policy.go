package ram

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

// SetPasswordPolicy invokes the ram.SetPasswordPolicy API synchronously
func (client *Client) SetPasswordPolicy(request *SetPasswordPolicyRequest) (response *SetPasswordPolicyResponse, err error) {
	response = CreateSetPasswordPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// SetPasswordPolicyWithChan invokes the ram.SetPasswordPolicy API asynchronously
func (client *Client) SetPasswordPolicyWithChan(request *SetPasswordPolicyRequest) (<-chan *SetPasswordPolicyResponse, <-chan error) {
	responseChan := make(chan *SetPasswordPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetPasswordPolicy(request)
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

// SetPasswordPolicyWithCallback invokes the ram.SetPasswordPolicy API asynchronously
func (client *Client) SetPasswordPolicyWithCallback(request *SetPasswordPolicyRequest, callback func(response *SetPasswordPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetPasswordPolicyResponse
		var err error
		defer close(result)
		response, err = client.SetPasswordPolicy(request)
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

// SetPasswordPolicyRequest is the request struct for api SetPasswordPolicy
type SetPasswordPolicyRequest struct {
	*requests.RpcRequest
	PasswordReusePrevention    requests.Integer `position:"Query" name:"PasswordReusePrevention"`
	RequireUppercaseCharacters requests.Boolean `position:"Query" name:"RequireUppercaseCharacters"`
	MinimumPasswordLength      requests.Integer `position:"Query" name:"MinimumPasswordLength"`
	RequireNumbers             requests.Boolean `position:"Query" name:"RequireNumbers"`
	RequireLowercaseCharacters requests.Boolean `position:"Query" name:"RequireLowercaseCharacters"`
	MaxPasswordAge             requests.Integer `position:"Query" name:"MaxPasswordAge"`
	MaxLoginAttemps            requests.Integer `position:"Query" name:"MaxLoginAttemps"`
	HardExpiry                 requests.Boolean `position:"Query" name:"HardExpiry"`
	RequireSymbols             requests.Boolean `position:"Query" name:"RequireSymbols"`
}

// SetPasswordPolicyResponse is the response struct for api SetPasswordPolicy
type SetPasswordPolicyResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	PasswordPolicy PasswordPolicy `json:"PasswordPolicy" xml:"PasswordPolicy"`
}

// CreateSetPasswordPolicyRequest creates a request to invoke SetPasswordPolicy API
func CreateSetPasswordPolicyRequest() (request *SetPasswordPolicyRequest) {
	request = &SetPasswordPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "SetPasswordPolicy", "Ram", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetPasswordPolicyResponse creates a response to parse from SetPasswordPolicy response
func CreateSetPasswordPolicyResponse() (response *SetPasswordPolicyResponse) {
	response = &SetPasswordPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
