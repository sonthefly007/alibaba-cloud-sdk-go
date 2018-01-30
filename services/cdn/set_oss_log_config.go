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

func (client *Client) SetOssLogConfig(request *SetOssLogConfigRequest) (response *SetOssLogConfigResponse, err error) {
	response = CreateSetOssLogConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetOssLogConfigWithChan(request *SetOssLogConfigRequest) (<-chan *SetOssLogConfigResponse, <-chan error) {
	responseChan := make(chan *SetOssLogConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetOssLogConfig(request)
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

func (client *Client) SetOssLogConfigWithCallback(request *SetOssLogConfigRequest, callback func(response *SetOssLogConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetOssLogConfigResponse
		var err error
		defer close(result)
		response, err = client.SetOssLogConfig(request)
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

type SetOssLogConfigRequest struct {
	*requests.RpcRequest
	Prefix        string           `position:"Query" name:"Prefix"`
	Bucket        string           `position:"Query" name:"Bucket"`
	Enable        string           `position:"Query" name:"Enable"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type SetOssLogConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetOssLogConfigRequest() (request *SetOssLogConfigRequest) {
	request = &SetOssLogConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetOssLogConfig", "", "")
	return
}

func CreateSetOssLogConfigResponse() (response *SetOssLogConfigResponse) {
	response = &SetOssLogConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
