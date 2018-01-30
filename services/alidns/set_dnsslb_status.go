package alidns

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

func (client *Client) SetDNSSLBStatus(request *SetDNSSLBStatusRequest) (response *SetDNSSLBStatusResponse, err error) {
	response = CreateSetDNSSLBStatusResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetDNSSLBStatusWithChan(request *SetDNSSLBStatusRequest) (<-chan *SetDNSSLBStatusResponse, <-chan error) {
	responseChan := make(chan *SetDNSSLBStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDNSSLBStatus(request)
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

func (client *Client) SetDNSSLBStatusWithCallback(request *SetDNSSLBStatusRequest, callback func(response *SetDNSSLBStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDNSSLBStatusResponse
		var err error
		defer close(result)
		response, err = client.SetDNSSLBStatus(request)
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

type SetDNSSLBStatusRequest struct {
	*requests.RpcRequest
	Open         requests.Boolean `position:"Query" name:"Open"`
	SubDomain    string           `position:"Query" name:"SubDomain"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
}

type SetDNSSLBStatusResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	RecordCount int    `json:"RecordCount" xml:"RecordCount"`
	Open        bool   `json:"Open" xml:"Open"`
}

func CreateSetDNSSLBStatusRequest() (request *SetDNSSLBStatusRequest) {
	request = &SetDNSSLBStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "SetDNSSLBStatus", "", "")
	return
}

func CreateSetDNSSLBStatusResponse() (response *SetDNSSLBStatusResponse) {
	response = &SetDNSSLBStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
