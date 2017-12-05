package ons

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

func (client *Client) OnsTrendTopicInputTps(request *OnsTrendTopicInputTpsRequest) (response *OnsTrendTopicInputTpsResponse, err error) {
	response = CreateOnsTrendTopicInputTpsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) OnsTrendTopicInputTpsWithChan(request *OnsTrendTopicInputTpsRequest) (<-chan *OnsTrendTopicInputTpsResponse, <-chan error) {
	responseChan := make(chan *OnsTrendTopicInputTpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsTrendTopicInputTps(request)
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

func (client *Client) OnsTrendTopicInputTpsWithCallback(request *OnsTrendTopicInputTpsRequest, callback func(response *OnsTrendTopicInputTpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsTrendTopicInputTpsResponse
		var err error
		defer close(result)
		response, err = client.OnsTrendTopicInputTps(request)
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

type OnsTrendTopicInputTpsRequest struct {
	*requests.RpcRequest
	EndTime      string `position:"Query" name:"EndTime"`
	Topic        string `position:"Query" name:"Topic"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	BeginTime    string `position:"Query" name:"BeginTime"`
	PreventCache string `position:"Query" name:"PreventCache"`
	Type         string `position:"Query" name:"Type"`
	Period       string `position:"Query" name:"Period"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
}

type OnsTrendTopicInputTpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      struct {
		Title   string `json:"Title" xml:"Title"`
		XUnit   string `json:"XUnit" xml:"XUnit"`
		YUnit   string `json:"YUnit" xml:"YUnit"`
		Records []struct {
			X int64   `json:"X" xml:"X"`
			Y float64 `json:"Y" xml:"Y"`
		} `json:"Records" xml:"Records"`
	} `json:"Data" xml:"Data"`
}

func CreateOnsTrendTopicInputTpsRequest() (request *OnsTrendTopicInputTpsRequest) {
	request = &OnsTrendTopicInputTpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2017-09-18", "OnsTrendTopicInputTps", "", "")
	return
}

func CreateOnsTrendTopicInputTpsResponse() (response *OnsTrendTopicInputTpsResponse) {
	response = &OnsTrendTopicInputTpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}