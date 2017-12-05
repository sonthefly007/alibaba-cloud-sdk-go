package cloudwf

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

func (client *Client) GetRadioRunHistoryTimeSer(request *GetRadioRunHistoryTimeSerRequest) (response *GetRadioRunHistoryTimeSerResponse, err error) {
	response = CreateGetRadioRunHistoryTimeSerResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetRadioRunHistoryTimeSerWithChan(request *GetRadioRunHistoryTimeSerRequest) (<-chan *GetRadioRunHistoryTimeSerResponse, <-chan error) {
	responseChan := make(chan *GetRadioRunHistoryTimeSerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRadioRunHistoryTimeSer(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) GetRadioRunHistoryTimeSerWithCallback(request *GetRadioRunHistoryTimeSerRequest, callback func(response *GetRadioRunHistoryTimeSerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRadioRunHistoryTimeSerResponse
		var err error
		defer close(result)
		response, err = client.GetRadioRunHistoryTimeSer(request)
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

type GetRadioRunHistoryTimeSerRequest struct {
	*requests.RpcRequest
	Id string `position:"Query" name:"Id"`
}

type GetRadioRunHistoryTimeSerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

func CreateGetRadioRunHistoryTimeSerRequest() (request *GetRadioRunHistoryTimeSerRequest) {
	request = &GetRadioRunHistoryTimeSerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetRadioRunHistoryTimeSer", "", "")
	return
}

func CreateGetRadioRunHistoryTimeSerResponse() (response *GetRadioRunHistoryTimeSerResponse) {
	response = &GetRadioRunHistoryTimeSerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}