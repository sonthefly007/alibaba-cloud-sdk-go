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

func (client *Client) ModifyBandwidthPackageSpec(request *ModifyBandwidthPackageSpecRequest) (response *ModifyBandwidthPackageSpecResponse, err error) {
	response = CreateModifyBandwidthPackageSpecResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyBandwidthPackageSpecWithChan(request *ModifyBandwidthPackageSpecRequest) (<-chan *ModifyBandwidthPackageSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyBandwidthPackageSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBandwidthPackageSpec(request)
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

func (client *Client) ModifyBandwidthPackageSpecWithCallback(request *ModifyBandwidthPackageSpecRequest, callback func(response *ModifyBandwidthPackageSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBandwidthPackageSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyBandwidthPackageSpec(request)
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

type ModifyBandwidthPackageSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Bandwidth            string           `position:"Query" name:"Bandwidth"`
	BandwidthPackageId   string           `position:"Query" name:"BandwidthPackageId"`
}

type ModifyBandwidthPackageSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyBandwidthPackageSpecRequest() (request *ModifyBandwidthPackageSpecRequest) {
	request = &ModifyBandwidthPackageSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyBandwidthPackageSpec", "ecs", "openAPI")
	return
}

func CreateModifyBandwidthPackageSpecResponse() (response *ModifyBandwidthPackageSpecResponse) {
	response = &ModifyBandwidthPackageSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
