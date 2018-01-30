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

func (client *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
	response = CreateCreateImageResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateImageWithChan(request *CreateImageRequest) (<-chan *CreateImageResponse, <-chan error) {
	responseChan := make(chan *CreateImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateImage(request)
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

func (client *Client) CreateImageWithCallback(request *CreateImageRequest, callback func(response *CreateImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateImageResponse
		var err error
		defer close(result)
		response, err = client.CreateImage(request)
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

type CreateImageRequest struct {
	*requests.RpcRequest
	Architecture         string                          `position:"Query" name:"Architecture"`
	Tag5Value            string                          `position:"Query" name:"Tag.5.Value"`
	Tag3Key              string                          `position:"Query" name:"Tag.3.Key"`
	ResourceOwnerAccount string                          `position:"Query" name:"ResourceOwnerAccount"`
	SnapshotId           string                          `position:"Query" name:"SnapshotId"`
	Description          string                          `position:"Query" name:"Description"`
	DiskDeviceMapping    *[]CreateImageDiskDeviceMapping `position:"Query" name:"DiskDeviceMapping"  type:"Repeated"`
	Tag1Key              string                          `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string                          `position:"Query" name:"Tag.1.Value"`
	ResourceOwnerId      requests.Integer                `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string                          `position:"Query" name:"OwnerAccount"`
	Tag4Value            string                          `position:"Query" name:"Tag.4.Value"`
	Platform             string                          `position:"Query" name:"Platform"`
	ClientToken          string                          `position:"Query" name:"ClientToken"`
	ImageVersion         string                          `position:"Query" name:"ImageVersion"`
	OwnerId              requests.Integer                `position:"Query" name:"OwnerId"`
	Tag5Key              string                          `position:"Query" name:"Tag.5.Key"`
	ImageName            string                          `position:"Query" name:"ImageName"`
	Tag2Key              string                          `position:"Query" name:"Tag.2.Key"`
	InstanceId           string                          `position:"Query" name:"InstanceId"`
	Tag3Value            string                          `position:"Query" name:"Tag.3.Value"`
	Tag4Key              string                          `position:"Query" name:"Tag.4.Key"`
	Tag2Value            string                          `position:"Query" name:"Tag.2.Value"`
}

type CreateImageDiskDeviceMapping struct {
	Size       string `name:"Size"`
	SnapshotId string `name:"SnapshotId"`
	Device     string `name:"Device"`
	DiskType   string `name:"DiskType"`
}

type CreateImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ImageId   string `json:"ImageId" xml:"ImageId"`
}

func CreateCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateImage", "ecs", "openAPI")
	return
}

func CreateCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
