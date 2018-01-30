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

func (client *Client) DetachPolicyFromRole(request *DetachPolicyFromRoleRequest) (response *DetachPolicyFromRoleResponse, err error) {
	response = CreateDetachPolicyFromRoleResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DetachPolicyFromRoleWithChan(request *DetachPolicyFromRoleRequest) (<-chan *DetachPolicyFromRoleResponse, <-chan error) {
	responseChan := make(chan *DetachPolicyFromRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachPolicyFromRole(request)
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

func (client *Client) DetachPolicyFromRoleWithCallback(request *DetachPolicyFromRoleRequest, callback func(response *DetachPolicyFromRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachPolicyFromRoleResponse
		var err error
		defer close(result)
		response, err = client.DetachPolicyFromRole(request)
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

type DetachPolicyFromRoleRequest struct {
	*requests.RpcRequest
	RoleName   string `position:"Query" name:"RoleName"`
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

type DetachPolicyFromRoleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDetachPolicyFromRoleRequest() (request *DetachPolicyFromRoleRequest) {
	request = &DetachPolicyFromRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "DetachPolicyFromRole", "", "")
	return
}

func CreateDetachPolicyFromRoleResponse() (response *DetachPolicyFromRoleResponse) {
	response = &DetachPolicyFromRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
