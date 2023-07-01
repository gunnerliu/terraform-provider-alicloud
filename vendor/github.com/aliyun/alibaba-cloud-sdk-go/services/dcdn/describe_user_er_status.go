package dcdn

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

// DescribeUserErStatus invokes the dcdn.DescribeUserErStatus API synchronously
func (client *Client) DescribeUserErStatus(request *DescribeUserErStatusRequest) (response *DescribeUserErStatusResponse, err error) {
	response = CreateDescribeUserErStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserErStatusWithChan invokes the dcdn.DescribeUserErStatus API asynchronously
func (client *Client) DescribeUserErStatusWithChan(request *DescribeUserErStatusRequest) (<-chan *DescribeUserErStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeUserErStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserErStatus(request)
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

// DescribeUserErStatusWithCallback invokes the dcdn.DescribeUserErStatus API asynchronously
func (client *Client) DescribeUserErStatusWithCallback(request *DescribeUserErStatusRequest, callback func(response *DescribeUserErStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserErStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserErStatus(request)
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

// DescribeUserErStatusRequest is the request struct for api DescribeUserErStatus
type DescribeUserErStatusRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeUserErStatusResponse is the response struct for api DescribeUserErStatus
type DescribeUserErStatusResponse struct {
	*responses.BaseResponse
	InDebt        bool   `json:"InDebt" xml:"InDebt"`
	OnService     bool   `json:"OnService" xml:"OnService"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	InDebtOverdue bool   `json:"InDebtOverdue" xml:"InDebtOverdue"`
	Enabled       bool   `json:"Enabled" xml:"Enabled"`
}

// CreateDescribeUserErStatusRequest creates a request to invoke DescribeUserErStatus API
func CreateDescribeUserErStatusRequest() (request *DescribeUserErStatusRequest) {
	request = &DescribeUserErStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeUserErStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeUserErStatusResponse creates a response to parse from DescribeUserErStatus response
func CreateDescribeUserErStatusResponse() (response *DescribeUserErStatusResponse) {
	response = &DescribeUserErStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
