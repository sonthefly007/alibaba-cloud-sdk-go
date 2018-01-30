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

func (client *Client) RevokeSecurityGroup(request *RevokeSecurityGroupRequest) (response *RevokeSecurityGroupResponse, err error) {
	response = CreateRevokeSecurityGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RevokeSecurityGroupWithChan(request *RevokeSecurityGroupRequest) (<-chan *RevokeSecurityGroupResponse, <-chan error) {
	responseChan := make(chan *RevokeSecurityGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokeSecurityGroup(request)
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

func (client *Client) RevokeSecurityGroupWithCallback(request *RevokeSecurityGroupRequest, callback func(response *RevokeSecurityGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokeSecurityGroupResponse
		var err error
		defer close(result)
		response, err = client.RevokeSecurityGroup(request)
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

type RevokeSecurityGroupRequest struct {
	*requests.RpcRequest
	SourceGroupOwnerAccount string           `position:"Query" name:"SourceGroupOwnerAccount"`
	PortRange               string           `position:"Query" name:"PortRange"`
	DestCidrIp              string           `position:"Query" name:"DestCidrIp"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	Description             string           `position:"Query" name:"Description"`
	Priority                string           `position:"Query" name:"Priority"`
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	SourceGroupOwnerId      requests.Integer `position:"Query" name:"SourceGroupOwnerId"`
	ClientToken             string           `position:"Query" name:"ClientToken"`
	SecurityGroupId         string           `position:"Query" name:"SecurityGroupId"`
	SourcePortRange         string           `position:"Query" name:"SourcePortRange"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	IpProtocol              string           `position:"Query" name:"IpProtocol"`
	SourceGroupId           string           `position:"Query" name:"SourceGroupId"`
	NicType                 string           `position:"Query" name:"NicType"`
	Policy                  string           `position:"Query" name:"Policy"`
	SourceCidrIp            string           `position:"Query" name:"SourceCidrIp"`
}

type RevokeSecurityGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateRevokeSecurityGroupRequest() (request *RevokeSecurityGroupRequest) {
	request = &RevokeSecurityGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "RevokeSecurityGroup", "ecs", "openAPI")
	return
}

func CreateRevokeSecurityGroupResponse() (response *RevokeSecurityGroupResponse) {
	response = &RevokeSecurityGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
