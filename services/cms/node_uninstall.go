package cms

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

func (client *Client) NodeUninstall(request *NodeUninstallRequest) (response *NodeUninstallResponse, err error) {
	response = CreateNodeUninstallResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) NodeUninstallWithChan(request *NodeUninstallRequest) (<-chan *NodeUninstallResponse, <-chan error) {
	responseChan := make(chan *NodeUninstallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.NodeUninstall(request)
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

func (client *Client) NodeUninstallWithCallback(request *NodeUninstallRequest, callback func(response *NodeUninstallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *NodeUninstallResponse
		var err error
		defer close(result)
		response, err = client.NodeUninstall(request)
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

type NodeUninstallRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

type NodeUninstallResponse struct {
	*responses.BaseResponse
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

func CreateNodeUninstallRequest() (request *NodeUninstallRequest) {
	request = &NodeUninstallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "NodeUninstall", "cms", "openAPI")
	return
}

func CreateNodeUninstallResponse() (response *NodeUninstallResponse) {
	response = &NodeUninstallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
