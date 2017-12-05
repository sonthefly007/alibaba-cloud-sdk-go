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

func (client *Client) ListApAssetCanBeAdded(request *ListApAssetCanBeAddedRequest) (response *ListApAssetCanBeAddedResponse, err error) {
	response = CreateListApAssetCanBeAddedResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListApAssetCanBeAddedWithChan(request *ListApAssetCanBeAddedRequest) (<-chan *ListApAssetCanBeAddedResponse, <-chan error) {
	responseChan := make(chan *ListApAssetCanBeAddedResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListApAssetCanBeAdded(request)
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

func (client *Client) ListApAssetCanBeAddedWithCallback(request *ListApAssetCanBeAddedRequest, callback func(response *ListApAssetCanBeAddedResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListApAssetCanBeAddedResponse
		var err error
		defer close(result)
		response, err = client.ListApAssetCanBeAdded(request)
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

type ListApAssetCanBeAddedRequest struct {
	*requests.RpcRequest
	SearchModel string `position:"Query" name:"SearchModel"`
	SearchName  string `position:"Query" name:"SearchName"`
	Length      string `position:"Query" name:"Length"`
	PageIndex   string `position:"Query" name:"PageIndex"`
	ApgroupId   string `position:"Query" name:"ApgroupId"`
	SearchMac   string `position:"Query" name:"SearchMac"`
}

type ListApAssetCanBeAddedResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

func CreateListApAssetCanBeAddedRequest() (request *ListApAssetCanBeAddedRequest) {
	request = &ListApAssetCanBeAddedRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ListApAssetCanBeAdded", "", "")
	return
}

func CreateListApAssetCanBeAddedResponse() (response *ListApAssetCanBeAddedResponse) {
	response = &ListApAssetCanBeAddedResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}