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

// DeleteTrafficControl invokes the cloudapi.DeleteTrafficControl API synchronously
func (client *Client) DeleteTrafficControl(request *DeleteTrafficControlRequest) (response *DeleteTrafficControlResponse, err error) {
	response = CreateDeleteTrafficControlResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTrafficControlWithChan invokes the cloudapi.DeleteTrafficControl API asynchronously
func (client *Client) DeleteTrafficControlWithChan(request *DeleteTrafficControlRequest) (<-chan *DeleteTrafficControlResponse, <-chan error) {
	responseChan := make(chan *DeleteTrafficControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTrafficControl(request)
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

// DeleteTrafficControlWithCallback invokes the cloudapi.DeleteTrafficControl API asynchronously
func (client *Client) DeleteTrafficControlWithCallback(request *DeleteTrafficControlRequest, callback func(response *DeleteTrafficControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTrafficControlResponse
		var err error
		defer close(result)
		response, err = client.DeleteTrafficControl(request)
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

// DeleteTrafficControlRequest is the request struct for api DeleteTrafficControl
type DeleteTrafficControlRequest struct {
	*requests.RpcRequest
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
}

// DeleteTrafficControlResponse is the response struct for api DeleteTrafficControl
type DeleteTrafficControlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTrafficControlRequest creates a request to invoke DeleteTrafficControl API
func CreateDeleteTrafficControlRequest() (request *DeleteTrafficControlRequest) {
	request = &DeleteTrafficControlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteTrafficControl", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteTrafficControlResponse creates a response to parse from DeleteTrafficControl response
func CreateDeleteTrafficControlResponse() (response *DeleteTrafficControlResponse) {
	response = &DeleteTrafficControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
