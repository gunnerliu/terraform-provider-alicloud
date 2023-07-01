package cloudapi

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

// DeleteBackendModel invokes the cloudapi.DeleteBackendModel API synchronously
func (client *Client) DeleteBackendModel(request *DeleteBackendModelRequest) (response *DeleteBackendModelResponse, err error) {
	response = CreateDeleteBackendModelResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBackendModelWithChan invokes the cloudapi.DeleteBackendModel API asynchronously
func (client *Client) DeleteBackendModelWithChan(request *DeleteBackendModelRequest) (<-chan *DeleteBackendModelResponse, <-chan error) {
	responseChan := make(chan *DeleteBackendModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBackendModel(request)
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

// DeleteBackendModelWithCallback invokes the cloudapi.DeleteBackendModel API asynchronously
func (client *Client) DeleteBackendModelWithCallback(request *DeleteBackendModelRequest, callback func(response *DeleteBackendModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBackendModelResponse
		var err error
		defer close(result)
		response, err = client.DeleteBackendModel(request)
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

// DeleteBackendModelRequest is the request struct for api DeleteBackendModel
type DeleteBackendModelRequest struct {
	*requests.RpcRequest
	StageName      string `position:"Query" name:"StageName"`
	BackendId      string `position:"Query" name:"BackendId"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	BackendModelId string `position:"Query" name:"BackendModelId"`
}

// DeleteBackendModelResponse is the response struct for api DeleteBackendModel
type DeleteBackendModelResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	OperationId string `json:"OperationId" xml:"OperationId"`
}

// CreateDeleteBackendModelRequest creates a request to invoke DeleteBackendModel API
func CreateDeleteBackendModelRequest() (request *DeleteBackendModelRequest) {
	request = &DeleteBackendModelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteBackendModel", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteBackendModelResponse creates a response to parse from DeleteBackendModel response
func CreateDeleteBackendModelResponse() (response *DeleteBackendModelResponse) {
	response = &DeleteBackendModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
