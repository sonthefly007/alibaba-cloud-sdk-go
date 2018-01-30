package ecs

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

func (client *Client) CreateVpc(request *CreateVpcRequest) (response *CreateVpcResponse, err error) {
	response = CreateCreateVpcResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateVpcWithChan(request *CreateVpcRequest) (<-chan *CreateVpcResponse, <-chan error) {
	responseChan := make(chan *CreateVpcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVpc(request)
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

func (client *Client) CreateVpcWithCallback(request *CreateVpcRequest, callback func(response *CreateVpcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVpcResponse
		var err error
		defer close(result)
		response, err = client.CreateVpc(request)
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

type CreateVpcRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	VpcName              string           `position:"Query" name:"VpcName"`
	UserCidr             string           `position:"Query" name:"UserCidr"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	CidrBlock            string           `position:"Query" name:"CidrBlock"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type CreateVpcResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	VpcId        string `json:"VpcId" xml:"VpcId"`
	VRouterId    string `json:"VRouterId" xml:"VRouterId"`
	RouteTableId string `json:"RouteTableId" xml:"RouteTableId"`
}

func CreateCreateVpcRequest() (request *CreateVpcRequest) {
	request = &CreateVpcRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateVpc", "ecs", "openAPI")
	return
}

func CreateCreateVpcResponse() (response *CreateVpcResponse) {
	response = &CreateVpcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
