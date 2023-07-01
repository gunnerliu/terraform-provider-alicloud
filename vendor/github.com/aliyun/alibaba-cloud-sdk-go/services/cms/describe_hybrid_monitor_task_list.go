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

// DescribeHybridMonitorTaskList invokes the cms.DescribeHybridMonitorTaskList API synchronously
func (client *Client) DescribeHybridMonitorTaskList(request *DescribeHybridMonitorTaskListRequest) (response *DescribeHybridMonitorTaskListResponse, err error) {
	response = CreateDescribeHybridMonitorTaskListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHybridMonitorTaskListWithChan invokes the cms.DescribeHybridMonitorTaskList API asynchronously
func (client *Client) DescribeHybridMonitorTaskListWithChan(request *DescribeHybridMonitorTaskListRequest) (<-chan *DescribeHybridMonitorTaskListResponse, <-chan error) {
	responseChan := make(chan *DescribeHybridMonitorTaskListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHybridMonitorTaskList(request)
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

// DescribeHybridMonitorTaskListWithCallback invokes the cms.DescribeHybridMonitorTaskList API asynchronously
func (client *Client) DescribeHybridMonitorTaskListWithCallback(request *DescribeHybridMonitorTaskListRequest, callback func(response *DescribeHybridMonitorTaskListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHybridMonitorTaskListResponse
		var err error
		defer close(result)
		response, err = client.DescribeHybridMonitorTaskList(request)
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

// DescribeHybridMonitorTaskListRequest is the request struct for api DescribeHybridMonitorTaskList
type DescribeHybridMonitorTaskListRequest struct {
	*requests.RpcRequest
	IncludeInstance   requests.Boolean `position:"Query" name:"IncludeInstance"`
	PageNumber        requests.Integer `position:"Query" name:"PageNumber"`
	ExtraInfo         string           `position:"Query" name:"ExtraInfo"`
	TargetUserId      requests.Integer `position:"Query" name:"TargetUserId"`
	CollectTargetType string           `position:"Query" name:"CollectTargetType"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	Keyword           string           `position:"Query" name:"Keyword"`
	TaskId            string           `position:"Query" name:"TaskId"`
	TaskType          string           `position:"Query" name:"TaskType"`
	GroupId           string           `position:"Query" name:"GroupId"`
	IncludeAliyunTask requests.Boolean `position:"Query" name:"IncludeAliyunTask"`
	Namespace         string           `position:"Query" name:"Namespace"`
}

// DescribeHybridMonitorTaskListResponse is the response struct for api DescribeHybridMonitorTaskList
type DescribeHybridMonitorTaskListResponse struct {
	*responses.BaseResponse
	RequestId  string         `json:"RequestId" xml:"RequestId"`
	Success    string         `json:"Success" xml:"Success"`
	Code       string         `json:"Code" xml:"Code"`
	Message    string         `json:"Message" xml:"Message"`
	PageSize   int            `json:"PageSize" xml:"PageSize"`
	PageNumber int            `json:"PageNumber" xml:"PageNumber"`
	Total      int            `json:"Total" xml:"Total"`
	TaskList   []TaskListItem `json:"TaskList" xml:"TaskList"`
}

// CreateDescribeHybridMonitorTaskListRequest creates a request to invoke DescribeHybridMonitorTaskList API
func CreateDescribeHybridMonitorTaskListRequest() (request *DescribeHybridMonitorTaskListRequest) {
	request = &DescribeHybridMonitorTaskListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeHybridMonitorTaskList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeHybridMonitorTaskListResponse creates a response to parse from DescribeHybridMonitorTaskList response
func CreateDescribeHybridMonitorTaskListResponse() (response *DescribeHybridMonitorTaskListResponse) {
	response = &DescribeHybridMonitorTaskListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
