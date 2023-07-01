package adb

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

// DescribeTableStatistics invokes the adb.DescribeTableStatistics API synchronously
func (client *Client) DescribeTableStatistics(request *DescribeTableStatisticsRequest) (response *DescribeTableStatisticsResponse, err error) {
	response = CreateDescribeTableStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTableStatisticsWithChan invokes the adb.DescribeTableStatistics API asynchronously
func (client *Client) DescribeTableStatisticsWithChan(request *DescribeTableStatisticsRequest) (<-chan *DescribeTableStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeTableStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTableStatistics(request)
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

// DescribeTableStatisticsWithCallback invokes the adb.DescribeTableStatistics API asynchronously
func (client *Client) DescribeTableStatisticsWithCallback(request *DescribeTableStatisticsRequest, callback func(response *DescribeTableStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTableStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeTableStatistics(request)
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

// DescribeTableStatisticsRequest is the request struct for api DescribeTableStatistics
type DescribeTableStatisticsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	Order                string           `position:"Query" name:"Order"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeTableStatisticsResponse is the response struct for api DescribeTableStatistics
type DescribeTableStatisticsResponse struct {
	*responses.BaseResponse
	TotalCount  string                         `json:"TotalCount" xml:"TotalCount"`
	PageSize    string                         `json:"PageSize" xml:"PageSize"`
	RequestId   string                         `json:"RequestId" xml:"RequestId"`
	PageNumber  string                         `json:"PageNumber" xml:"PageNumber"`
	DBClusterId string                         `json:"DBClusterId" xml:"DBClusterId"`
	Items       ItemsInDescribeTableStatistics `json:"Items" xml:"Items"`
}

// CreateDescribeTableStatisticsRequest creates a request to invoke DescribeTableStatistics API
func CreateDescribeTableStatisticsRequest() (request *DescribeTableStatisticsRequest) {
	request = &DescribeTableStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "DescribeTableStatistics", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTableStatisticsResponse creates a response to parse from DescribeTableStatistics response
func CreateDescribeTableStatisticsResponse() (response *DescribeTableStatisticsResponse) {
	response = &DescribeTableStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
