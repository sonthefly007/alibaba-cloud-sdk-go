package slb

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

func (client *Client) CreateLoadBalancerUDPListener(request *CreateLoadBalancerUDPListenerRequest) (response *CreateLoadBalancerUDPListenerResponse, err error) {
	response = CreateCreateLoadBalancerUDPListenerResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateLoadBalancerUDPListenerWithChan(request *CreateLoadBalancerUDPListenerRequest) (<-chan *CreateLoadBalancerUDPListenerResponse, <-chan error) {
	responseChan := make(chan *CreateLoadBalancerUDPListenerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLoadBalancerUDPListener(request)
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

func (client *Client) CreateLoadBalancerUDPListenerWithCallback(request *CreateLoadBalancerUDPListenerRequest, callback func(response *CreateLoadBalancerUDPListenerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLoadBalancerUDPListenerResponse
		var err error
		defer close(result)
		response, err = client.CreateLoadBalancerUDPListener(request)
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

type CreateLoadBalancerUDPListenerRequest struct {
	*requests.RpcRequest
	Tags                      string           `position:"Query" name:"Tags"`
	BackendServerPort         requests.Integer `position:"Query" name:"BackendServerPort"`
	Scheduler                 string           `position:"Query" name:"Scheduler"`
	UnhealthyThreshold        requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	Bandwidth                 requests.Integer `position:"Query" name:"Bandwidth"`
	VServerGroupId            string           `position:"Query" name:"VServerGroupId"`
	HealthCheckInterval       requests.Integer `position:"Query" name:"healthCheckInterval"`
	HealthCheckExp            string           `position:"Query" name:"healthCheckExp"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	ListenerPort              requests.Integer `position:"Query" name:"ListenerPort"`
	AccessKeyId               string           `position:"Query" name:"access_key_id"`
	MaxConnection             requests.Integer `position:"Query" name:"MaxConnection"`
	HealthCheckReq            string           `position:"Query" name:"healthCheckReq"`
	HealthCheckConnectPort    requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	LoadBalancerId            string           `position:"Query" name:"LoadBalancerId"`
	HealthyThreshold          requests.Integer `position:"Query" name:"HealthyThreshold"`
	MasterSlaveServerGroupId  string           `position:"Query" name:"MasterSlaveServerGroupId"`
	HealthCheckConnectTimeout requests.Integer `position:"Query" name:"HealthCheckConnectTimeout"`
	PersistenceTimeout        requests.Integer `position:"Query" name:"PersistenceTimeout"`
}

type CreateLoadBalancerUDPListenerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateCreateLoadBalancerUDPListenerRequest() (request *CreateLoadBalancerUDPListenerRequest) {
	request = &CreateLoadBalancerUDPListenerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "CreateLoadBalancerUDPListener", "slb", "openAPI")
	return
}

func CreateCreateLoadBalancerUDPListenerResponse() (response *CreateLoadBalancerUDPListenerResponse) {
	response = &CreateLoadBalancerUDPListenerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
