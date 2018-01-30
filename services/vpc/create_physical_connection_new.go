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

func (client *Client) CreatePhysicalConnectionNew(request *CreatePhysicalConnectionNewRequest) (response *CreatePhysicalConnectionNewResponse, err error) {
	response = CreateCreatePhysicalConnectionNewResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreatePhysicalConnectionNewWithChan(request *CreatePhysicalConnectionNewRequest) (<-chan *CreatePhysicalConnectionNewResponse, <-chan error) {
	responseChan := make(chan *CreatePhysicalConnectionNewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePhysicalConnectionNew(request)
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

func (client *Client) CreatePhysicalConnectionNewWithCallback(request *CreatePhysicalConnectionNewRequest, callback func(response *CreatePhysicalConnectionNewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePhysicalConnectionNewResponse
		var err error
		defer close(result)
		response, err = client.CreatePhysicalConnectionNew(request)
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

type CreatePhysicalConnectionNewRequest struct {
	*requests.RpcRequest
	PeerLocation                  string           `position:"Query" name:"PeerLocation"`
	ClientToken                   string           `position:"Query" name:"ClientToken"`
	InterfaceName                 string           `position:"Query" name:"InterfaceName"`
	AccessPointId                 string           `position:"Query" name:"AccessPointId"`
	RedundantPhysicalConnectionId string           `position:"Query" name:"RedundantPhysicalConnectionId"`
	OwnerId                       requests.Integer `position:"Query" name:"OwnerId"`
	Type                          string           `position:"Query" name:"Type"`
	Bandwidth                     requests.Integer `position:"Query" name:"bandwidth"`
	LineOperator                  string           `position:"Query" name:"LineOperator"`
	CircuitCode                   string           `position:"Query" name:"CircuitCode"`
	ResourceOwnerAccount          string           `position:"Query" name:"ResourceOwnerAccount"`
	Description                   string           `position:"Query" name:"Description"`
	PortType                      string           `position:"Query" name:"PortType"`
	Name                          string           `position:"Query" name:"Name"`
	DeviceName                    string           `position:"Query" name:"DeviceName"`
	ResourceOwnerId               requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount                  string           `position:"Query" name:"OwnerAccount"`
}

type CreatePhysicalConnectionNewResponse struct {
	*responses.BaseResponse
	RequestId            string `json:"RequestId" xml:"RequestId"`
	PhysicalConnectionId string `json:"PhysicalConnectionId" xml:"PhysicalConnectionId"`
}

func CreateCreatePhysicalConnectionNewRequest() (request *CreatePhysicalConnectionNewRequest) {
	request = &CreatePhysicalConnectionNewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreatePhysicalConnectionNew", "vpc", "openAPI")
	return
}

func CreateCreatePhysicalConnectionNewResponse() (response *CreatePhysicalConnectionNewResponse) {
	response = &CreatePhysicalConnectionNewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
