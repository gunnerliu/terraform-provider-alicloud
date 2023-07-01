package cms

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

// CreateCmsSmspackageOrder invokes the cms.CreateCmsSmspackageOrder API synchronously
func (client *Client) CreateCmsSmspackageOrder(request *CreateCmsSmspackageOrderRequest) (response *CreateCmsSmspackageOrderResponse, err error) {
	response = CreateCreateCmsSmspackageOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCmsSmspackageOrderWithChan invokes the cms.CreateCmsSmspackageOrder API asynchronously
func (client *Client) CreateCmsSmspackageOrderWithChan(request *CreateCmsSmspackageOrderRequest) (<-chan *CreateCmsSmspackageOrderResponse, <-chan error) {
	responseChan := make(chan *CreateCmsSmspackageOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCmsSmspackageOrder(request)
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

// CreateCmsSmspackageOrderWithCallback invokes the cms.CreateCmsSmspackageOrder API asynchronously
func (client *Client) CreateCmsSmspackageOrderWithCallback(request *CreateCmsSmspackageOrderRequest, callback func(response *CreateCmsSmspackageOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCmsSmspackageOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateCmsSmspackageOrder(request)
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

// CreateCmsSmspackageOrderRequest is the request struct for api CreateCmsSmspackageOrder
type CreateCmsSmspackageOrderRequest struct {
	*requests.RpcRequest
	AutoRenewPeriod requests.Integer `position:"Query" name:"AutoRenewPeriod"`
	Period          requests.Integer `position:"Query" name:"Period"`
	AutoPay         requests.Boolean `position:"Query" name:"AutoPay"`
	SmsCount        string           `position:"Query" name:"SmsCount"`
	AutoUseCoupon   requests.Boolean `position:"Query" name:"AutoUseCoupon"`
	PeriodUnit      string           `position:"Query" name:"PeriodUnit"`
}

// CreateCmsSmspackageOrderResponse is the response struct for api CreateCmsSmspackageOrder
type CreateCmsSmspackageOrderResponse struct {
	*responses.BaseResponse
	OrderId   string `json:"OrderId" xml:"OrderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateCmsSmspackageOrderRequest creates a request to invoke CreateCmsSmspackageOrder API
func CreateCreateCmsSmspackageOrderRequest() (request *CreateCmsSmspackageOrderRequest) {
	request = &CreateCmsSmspackageOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "CreateCmsSmspackageOrder", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCmsSmspackageOrderResponse creates a response to parse from CreateCmsSmspackageOrder response
func CreateCreateCmsSmspackageOrderResponse() (response *CreateCmsSmspackageOrderResponse) {
	response = &CreateCmsSmspackageOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
