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

// DescribeDomainRealTimeSrcBpsData invokes the cdn.DescribeDomainRealTimeSrcBpsData API synchronously
func (client *Client) DescribeDomainRealTimeSrcBpsData(request *DescribeDomainRealTimeSrcBpsDataRequest) (response *DescribeDomainRealTimeSrcBpsDataResponse, err error) {
	response = CreateDescribeDomainRealTimeSrcBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainRealTimeSrcBpsDataWithChan invokes the cdn.DescribeDomainRealTimeSrcBpsData API asynchronously
func (client *Client) DescribeDomainRealTimeSrcBpsDataWithChan(request *DescribeDomainRealTimeSrcBpsDataRequest) (<-chan *DescribeDomainRealTimeSrcBpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainRealTimeSrcBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainRealTimeSrcBpsData(request)
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

// DescribeDomainRealTimeSrcBpsDataWithCallback invokes the cdn.DescribeDomainRealTimeSrcBpsData API asynchronously
func (client *Client) DescribeDomainRealTimeSrcBpsDataWithCallback(request *DescribeDomainRealTimeSrcBpsDataRequest, callback func(response *DescribeDomainRealTimeSrcBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainRealTimeSrcBpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainRealTimeSrcBpsData(request)
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

// DescribeDomainRealTimeSrcBpsDataRequest is the request struct for api DescribeDomainRealTimeSrcBpsData
type DescribeDomainRealTimeSrcBpsDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
}

// DescribeDomainRealTimeSrcBpsDataResponse is the response struct for api DescribeDomainRealTimeSrcBpsData
type DescribeDomainRealTimeSrcBpsDataResponse struct {
	*responses.BaseResponse
	EndTime                       string                        `json:"EndTime" xml:"EndTime"`
	StartTime                     string                        `json:"StartTime" xml:"StartTime"`
	RequestId                     string                        `json:"RequestId" xml:"RequestId"`
	DomainName                    string                        `json:"DomainName" xml:"DomainName"`
	DataInterval                  string                        `json:"DataInterval" xml:"DataInterval"`
	RealTimeSrcBpsDataPerInterval RealTimeSrcBpsDataPerInterval `json:"RealTimeSrcBpsDataPerInterval" xml:"RealTimeSrcBpsDataPerInterval"`
}

// CreateDescribeDomainRealTimeSrcBpsDataRequest creates a request to invoke DescribeDomainRealTimeSrcBpsData API
func CreateDescribeDomainRealTimeSrcBpsDataRequest() (request *DescribeDomainRealTimeSrcBpsDataRequest) {
	request = &DescribeDomainRealTimeSrcBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainRealTimeSrcBpsData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainRealTimeSrcBpsDataResponse creates a response to parse from DescribeDomainRealTimeSrcBpsData response
func CreateDescribeDomainRealTimeSrcBpsDataResponse() (response *DescribeDomainRealTimeSrcBpsDataResponse) {
	response = &DescribeDomainRealTimeSrcBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
