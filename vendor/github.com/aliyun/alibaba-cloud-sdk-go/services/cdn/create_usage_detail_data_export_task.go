package cdn

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

// CreateUsageDetailDataExportTask invokes the cdn.CreateUsageDetailDataExportTask API synchronously
func (client *Client) CreateUsageDetailDataExportTask(request *CreateUsageDetailDataExportTaskRequest) (response *CreateUsageDetailDataExportTaskResponse, err error) {
	response = CreateCreateUsageDetailDataExportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUsageDetailDataExportTaskWithChan invokes the cdn.CreateUsageDetailDataExportTask API asynchronously
func (client *Client) CreateUsageDetailDataExportTaskWithChan(request *CreateUsageDetailDataExportTaskRequest) (<-chan *CreateUsageDetailDataExportTaskResponse, <-chan error) {
	responseChan := make(chan *CreateUsageDetailDataExportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUsageDetailDataExportTask(request)
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

// CreateUsageDetailDataExportTaskWithCallback invokes the cdn.CreateUsageDetailDataExportTask API asynchronously
func (client *Client) CreateUsageDetailDataExportTaskWithCallback(request *CreateUsageDetailDataExportTaskRequest, callback func(response *CreateUsageDetailDataExportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUsageDetailDataExportTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateUsageDetailDataExportTask(request)
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

// CreateUsageDetailDataExportTaskRequest is the request struct for api CreateUsageDetailDataExportTask
type CreateUsageDetailDataExportTaskRequest struct {
	*requests.RpcRequest
	DomainNames string `position:"Query" name:"DomainNames"`
	EndTime     string `position:"Query" name:"EndTime"`
	TaskName    string `position:"Query" name:"TaskName"`
	Language    string `position:"Query" name:"Language"`
	StartTime   string `position:"Query" name:"StartTime"`
	Type        string `position:"Query" name:"Type"`
	Group       string `position:"Query" name:"Group"`
}

// CreateUsageDetailDataExportTaskResponse is the response struct for api CreateUsageDetailDataExportTask
type CreateUsageDetailDataExportTaskResponse struct {
	*responses.BaseResponse
	EndTime   string `json:"EndTime" xml:"EndTime"`
	StartTime string `json:"StartTime" xml:"StartTime"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateCreateUsageDetailDataExportTaskRequest creates a request to invoke CreateUsageDetailDataExportTask API
func CreateCreateUsageDetailDataExportTaskRequest() (request *CreateUsageDetailDataExportTaskRequest) {
	request = &CreateUsageDetailDataExportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "CreateUsageDetailDataExportTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateUsageDetailDataExportTaskResponse creates a response to parse from CreateUsageDetailDataExportTask response
func CreateCreateUsageDetailDataExportTaskResponse() (response *CreateUsageDetailDataExportTaskResponse) {
	response = &CreateUsageDetailDataExportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
