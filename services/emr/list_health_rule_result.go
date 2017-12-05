package emr

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

func (client *Client) ListHealthRuleResult(request *ListHealthRuleResultRequest) (response *ListHealthRuleResultResponse, err error) {
	response = CreateListHealthRuleResultResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListHealthRuleResultWithChan(request *ListHealthRuleResultRequest) (<-chan *ListHealthRuleResultResponse, <-chan error) {
	responseChan := make(chan *ListHealthRuleResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListHealthRuleResult(request)
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

func (client *Client) ListHealthRuleResultWithCallback(request *ListHealthRuleResultRequest, callback func(response *ListHealthRuleResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListHealthRuleResultResponse
		var err error
		defer close(result)
		response, err = client.ListHealthRuleResult(request)
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

type ListHealthRuleResultRequest struct {
	*requests.RpcRequest
	PageSize        string `position:"Query" name:"PageSize"`
	PageNumber      string `position:"Query" name:"PageNumber"`
	Component       string `position:"Query" name:"Component"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	Service         string `position:"Query" name:"Service"`
	HostName        string `position:"Query" name:"HostName"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
	Pass            string `position:"Query" name:"Pass"`
}

type ListHealthRuleResultResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	Total                int    `json:"Total" xml:"Total"`
	PageNumber           int    `json:"PageNumber" xml:"PageNumber"`
	PageSize             int    `json:"PageSize" xml:"PageSize"`
	HealthRuleResultList []struct {
		Id              int64  `json:"Id" xml:"Id"`
		ClusterId       int64  `json:"ClusterId" xml:"ClusterId"`
		RuleId          int64  `json:"RuleId" xml:"RuleId"`
		RuleName        string `json:"RuleName" xml:"RuleName"`
		RuleTitle       string `json:"RuleTitle" xml:"RuleTitle"`
		RuleStatus      string `json:"RuleStatus" xml:"RuleStatus"`
		RuleDescription string `json:"RuleDescription" xml:"RuleDescription"`
		Service         string `json:"Service" xml:"Service"`
		Component       string `json:"Component" xml:"Component"`
		Pass            string `json:"Pass" xml:"Pass"`
		HostNames       string `json:"HostNames" xml:"HostNames"`
	} `json:"HealthRuleResultList" xml:"HealthRuleResultList"`
}

func CreateListHealthRuleResultRequest() (request *ListHealthRuleResultRequest) {
	request = &ListHealthRuleResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListHealthRuleResult", "", "")
	return
}

func CreateListHealthRuleResultResponse() (response *ListHealthRuleResultResponse) {
	response = &ListHealthRuleResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}