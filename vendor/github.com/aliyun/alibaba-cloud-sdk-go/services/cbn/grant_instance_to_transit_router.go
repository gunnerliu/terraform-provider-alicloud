package cbn

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

// GrantInstanceToTransitRouter invokes the cbn.GrantInstanceToTransitRouter API synchronously
func (client *Client) GrantInstanceToTransitRouter(request *GrantInstanceToTransitRouterRequest) (response *GrantInstanceToTransitRouterResponse, err error) {
	response = CreateGrantInstanceToTransitRouterResponse()
	err = client.DoAction(request, response)
	return
}

// GrantInstanceToTransitRouterWithChan invokes the cbn.GrantInstanceToTransitRouter API asynchronously
func (client *Client) GrantInstanceToTransitRouterWithChan(request *GrantInstanceToTransitRouterRequest) (<-chan *GrantInstanceToTransitRouterResponse, <-chan error) {
	responseChan := make(chan *GrantInstanceToTransitRouterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GrantInstanceToTransitRouter(request)
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

// GrantInstanceToTransitRouterWithCallback invokes the cbn.GrantInstanceToTransitRouter API asynchronously
func (client *Client) GrantInstanceToTransitRouterWithCallback(request *GrantInstanceToTransitRouterRequest, callback func(response *GrantInstanceToTransitRouterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GrantInstanceToTransitRouterResponse
		var err error
		defer close(result)
		response, err = client.GrantInstanceToTransitRouter(request)
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

// GrantInstanceToTransitRouterRequest is the request struct for api GrantInstanceToTransitRouter
type GrantInstanceToTransitRouterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                string           `position:"Query" name:"CenId"`
	CenOwnerId           requests.Integer `position:"Query" name:"CenOwnerId"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Version              string           `position:"Query" name:"Version"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	OrderType            string           `position:"Query" name:"OrderType"`
}

// GrantInstanceToTransitRouterResponse is the response struct for api GrantInstanceToTransitRouter
type GrantInstanceToTransitRouterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGrantInstanceToTransitRouterRequest creates a request to invoke GrantInstanceToTransitRouter API
func CreateGrantInstanceToTransitRouterRequest() (request *GrantInstanceToTransitRouterRequest) {
	request = &GrantInstanceToTransitRouterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "GrantInstanceToTransitRouter", "", "")
	request.Method = requests.POST
	return
}

// CreateGrantInstanceToTransitRouterResponse creates a response to parse from GrantInstanceToTransitRouter response
func CreateGrantInstanceToTransitRouterResponse() (response *GrantInstanceToTransitRouterResponse) {
	response = &GrantInstanceToTransitRouterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
