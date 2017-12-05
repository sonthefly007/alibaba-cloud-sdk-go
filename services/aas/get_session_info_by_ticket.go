package aas

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

func (client *Client) GetSessionInfoByTicket(request *GetSessionInfoByTicketRequest) (response *GetSessionInfoByTicketResponse, err error) {
	response = CreateGetSessionInfoByTicketResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetSessionInfoByTicketWithChan(request *GetSessionInfoByTicketRequest) (<-chan *GetSessionInfoByTicketResponse, <-chan error) {
	responseChan := make(chan *GetSessionInfoByTicketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSessionInfoByTicket(request)
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

func (client *Client) GetSessionInfoByTicketWithCallback(request *GetSessionInfoByTicketRequest, callback func(response *GetSessionInfoByTicketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSessionInfoByTicketResponse
		var err error
		defer close(result)
		response, err = client.GetSessionInfoByTicket(request)
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

type GetSessionInfoByTicketRequest struct {
	*requests.RpcRequest
	Ticket string `position:"Query" name:"Ticket"`
}

type GetSessionInfoByTicketResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	SessionInfo string `json:"SessionInfo" xml:"SessionInfo"`
}

func CreateGetSessionInfoByTicketRequest() (request *GetSessionInfoByTicketRequest) {
	request = &GetSessionInfoByTicketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Aas", "2015-07-01", "GetSessionInfoByTicket", "", "")
	return
}

func CreateGetSessionInfoByTicketResponse() (response *GetSessionInfoByTicketResponse) {
	response = &GetSessionInfoByTicketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}