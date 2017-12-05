package ram

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

func (client *Client) ListPolicies(request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
	response = CreateListPoliciesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListPoliciesWithChan(request *ListPoliciesRequest) (<-chan *ListPoliciesResponse, <-chan error) {
	responseChan := make(chan *ListPoliciesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPolicies(request)
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

func (client *Client) ListPoliciesWithCallback(request *ListPoliciesRequest, callback func(response *ListPoliciesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPoliciesResponse
		var err error
		defer close(result)
		response, err = client.ListPolicies(request)
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

type ListPoliciesRequest struct {
	*requests.RpcRequest
	Marker     string `position:"Query" name:"Marker"`
	PolicyType string `position:"Query" name:"PolicyType"`
	MaxItems   string `position:"Query" name:"MaxItems"`
}

type ListPoliciesResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	IsTruncated bool   `json:"IsTruncated" xml:"IsTruncated"`
	Marker      string `json:"Marker" xml:"Marker"`
	Policies    []struct {
		PolicyName      string `json:"PolicyName" xml:"PolicyName"`
		PolicyType      string `json:"PolicyType" xml:"PolicyType"`
		Description     string `json:"Description" xml:"Description"`
		DefaultVersion  string `json:"DefaultVersion" xml:"DefaultVersion"`
		CreateDate      string `json:"CreateDate" xml:"CreateDate"`
		UpdateDate      string `json:"UpdateDate" xml:"UpdateDate"`
		AttachmentCount int    `json:"AttachmentCount" xml:"AttachmentCount"`
	} `json:"Policies" xml:"Policies"`
}

func CreateListPoliciesRequest() (request *ListPoliciesRequest) {
	request = &ListPoliciesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "ListPolicies", "", "")
	return
}

func CreateListPoliciesResponse() (response *ListPoliciesResponse) {
	response = &ListPoliciesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}