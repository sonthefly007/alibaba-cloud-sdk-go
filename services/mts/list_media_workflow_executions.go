package mts

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

func (client *Client) ListMediaWorkflowExecutions(request *ListMediaWorkflowExecutionsRequest) (response *ListMediaWorkflowExecutionsResponse, err error) {
	response = CreateListMediaWorkflowExecutionsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListMediaWorkflowExecutionsWithChan(request *ListMediaWorkflowExecutionsRequest) (<-chan *ListMediaWorkflowExecutionsResponse, <-chan error) {
	responseChan := make(chan *ListMediaWorkflowExecutionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMediaWorkflowExecutions(request)
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

func (client *Client) ListMediaWorkflowExecutionsWithCallback(request *ListMediaWorkflowExecutionsRequest, callback func(response *ListMediaWorkflowExecutionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMediaWorkflowExecutionsResponse
		var err error
		defer close(result)
		response, err = client.ListMediaWorkflowExecutions(request)
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

type ListMediaWorkflowExecutionsRequest struct {
	*requests.RpcRequest
	InputFileURL         string           `position:"Query" name:"InputFileURL"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	MediaWorkflowName    string           `position:"Query" name:"MediaWorkflowName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	MaximumPageSize      requests.Integer `position:"Query" name:"MaximumPageSize"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NextPageToken        string           `position:"Query" name:"NextPageToken"`
	MediaWorkflowId      string           `position:"Query" name:"MediaWorkflowId"`
}

type ListMediaWorkflowExecutionsResponse struct {
	*responses.BaseResponse
	RequestId                  string `json:"RequestId" xml:"RequestId"`
	NextPageToken              string `json:"NextPageToken" xml:"NextPageToken"`
	MediaWorkflowExecutionList struct {
		MediaWorkflowExecution []struct {
			RunId           string `json:"RunId" xml:"RunId"`
			MediaWorkflowId string `json:"MediaWorkflowId" xml:"MediaWorkflowId"`
			Name            string `json:"Name" xml:"Name"`
			State           string `json:"State" xml:"State"`
			MediaId         string `json:"MediaId" xml:"MediaId"`
			CreationTime    string `json:"CreationTime" xml:"CreationTime"`
			Input           struct {
				UserData  string `json:"UserData" xml:"UserData"`
				InputFile struct {
					Bucket   string `json:"Bucket" xml:"Bucket"`
					Location string `json:"Location" xml:"Location"`
					Object   string `json:"Object" xml:"Object"`
				} `json:"InputFile" xml:"InputFile"`
			} `json:"Input" xml:"Input"`
			ActivityList struct {
				Activity []struct {
					Name             string `json:"Name" xml:"Name"`
					Type             string `json:"Type" xml:"Type"`
					JobId            string `json:"JobId" xml:"JobId"`
					State            string `json:"State" xml:"State"`
					Code             string `json:"Code" xml:"Code"`
					Message          string `json:"Message" xml:"Message"`
					StartTime        string `json:"StartTime" xml:"StartTime"`
					EndTime          string `json:"EndTime" xml:"EndTime"`
					MNSMessageResult struct {
						MessageId    string `json:"MessageId" xml:"MessageId"`
						ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
						ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
					} `json:"MNSMessageResult" xml:"MNSMessageResult"`
				} `json:"Activity" xml:"Activity"`
			} `json:"ActivityList" xml:"ActivityList"`
		} `json:"MediaWorkflowExecution" xml:"MediaWorkflowExecution"`
	} `json:"MediaWorkflowExecutionList" xml:"MediaWorkflowExecutionList"`
}

func CreateListMediaWorkflowExecutionsRequest() (request *ListMediaWorkflowExecutionsRequest) {
	request = &ListMediaWorkflowExecutionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListMediaWorkflowExecutions", "mts", "openAPI")
	return
}

func CreateListMediaWorkflowExecutionsResponse() (response *ListMediaWorkflowExecutionsResponse) {
	response = &ListMediaWorkflowExecutionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
