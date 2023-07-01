package alikafka

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

// ModifyTopicRemark invokes the alikafka.ModifyTopicRemark API synchronously
func (client *Client) ModifyTopicRemark(request *ModifyTopicRemarkRequest) (response *ModifyTopicRemarkResponse, err error) {
	response = CreateModifyTopicRemarkResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyTopicRemarkWithChan invokes the alikafka.ModifyTopicRemark API asynchronously
func (client *Client) ModifyTopicRemarkWithChan(request *ModifyTopicRemarkRequest) (<-chan *ModifyTopicRemarkResponse, <-chan error) {
	responseChan := make(chan *ModifyTopicRemarkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyTopicRemark(request)
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

// ModifyTopicRemarkWithCallback invokes the alikafka.ModifyTopicRemark API asynchronously
func (client *Client) ModifyTopicRemarkWithCallback(request *ModifyTopicRemarkRequest, callback func(response *ModifyTopicRemarkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyTopicRemarkResponse
		var err error
		defer close(result)
		response, err = client.ModifyTopicRemark(request)
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

// ModifyTopicRemarkRequest is the request struct for api ModifyTopicRemark
type ModifyTopicRemarkRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Topic      string `position:"Query" name:"Topic"`
	Remark     string `position:"Query" name:"Remark"`
}

// ModifyTopicRemarkResponse is the response struct for api ModifyTopicRemark
type ModifyTopicRemarkResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateModifyTopicRemarkRequest creates a request to invoke ModifyTopicRemark API
func CreateModifyTopicRemarkRequest() (request *ModifyTopicRemarkRequest) {
	request = &ModifyTopicRemarkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "ModifyTopicRemark", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyTopicRemarkResponse creates a response to parse from ModifyTopicRemark response
func CreateModifyTopicRemarkResponse() (response *ModifyTopicRemarkResponse) {
	response = &ModifyTopicRemarkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
