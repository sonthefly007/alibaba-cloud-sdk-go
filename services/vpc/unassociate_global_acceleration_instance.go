package vpc

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

func (client *Client) UnassociateGlobalAccelerationInstance(request *UnassociateGlobalAccelerationInstanceRequest) (response *UnassociateGlobalAccelerationInstanceResponse, err error) {
	response = CreateUnassociateGlobalAccelerationInstanceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) UnassociateGlobalAccelerationInstanceWithChan(request *UnassociateGlobalAccelerationInstanceRequest) (<-chan *UnassociateGlobalAccelerationInstanceResponse, <-chan error) {
	responseChan := make(chan *UnassociateGlobalAccelerationInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnassociateGlobalAccelerationInstance(request)
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

func (client *Client) UnassociateGlobalAccelerationInstanceWithCallback(request *UnassociateGlobalAccelerationInstanceRequest, callback func(response *UnassociateGlobalAccelerationInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnassociateGlobalAccelerationInstanceResponse
		var err error
		defer close(result)
		response, err = client.UnassociateGlobalAccelerationInstance(request)
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

type UnassociateGlobalAccelerationInstanceRequest struct {
	*requests.RpcRequest
	GlobalAccelerationInstanceId string           `position:"Query" name:"GlobalAccelerationInstanceId"`
	ResourceOwnerAccount         string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount                 string           `position:"Query" name:"OwnerAccount"`
	OwnerId                      requests.Integer `position:"Query" name:"OwnerId"`
	InstanceType                 string           `position:"Query" name:"InstanceType"`
}

type UnassociateGlobalAccelerationInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateUnassociateGlobalAccelerationInstanceRequest() (request *UnassociateGlobalAccelerationInstanceRequest) {
	request = &UnassociateGlobalAccelerationInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "UnassociateGlobalAccelerationInstance", "vpc", "openAPI")
	return
}

func CreateUnassociateGlobalAccelerationInstanceResponse() (response *UnassociateGlobalAccelerationInstanceResponse) {
	response = &UnassociateGlobalAccelerationInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
