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

func (client *Client) DescribeUserConfigs(request *DescribeUserConfigsRequest) (response *DescribeUserConfigsResponse, err error) {
	response = CreateDescribeUserConfigsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeUserConfigsWithChan(request *DescribeUserConfigsRequest) (<-chan *DescribeUserConfigsResponse, <-chan error) {
	responseChan := make(chan *DescribeUserConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserConfigs(request)
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

func (client *Client) DescribeUserConfigsWithCallback(request *DescribeUserConfigsRequest, callback func(response *DescribeUserConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserConfigsResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserConfigs(request)
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

type DescribeUserConfigsRequest struct {
	*requests.RpcRequest
	Config        string           `position:"Query" name:"Config"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeUserConfigsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Configs   struct {
		OssLogConfig struct {
			Enable string `json:"Enable" xml:"Enable"`
			Bucket string `json:"Bucket" xml:"Bucket"`
			Prefix string `json:"Prefix" xml:"Prefix"`
		} `json:"OssLogConfig" xml:"OssLogConfig"`
		GreenManagerConfig struct {
			Quota string `json:"Quota" xml:"Quota"`
			Ratio string `json:"Ratio" xml:"Ratio"`
		} `json:"GreenManagerConfig" xml:"GreenManagerConfig"`
	} `json:"Configs" xml:"Configs"`
}

func CreateDescribeUserConfigsRequest() (request *DescribeUserConfigsRequest) {
	request = &DescribeUserConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeUserConfigs", "", "")
	return
}

func CreateDescribeUserConfigsResponse() (response *DescribeUserConfigsResponse) {
	response = &DescribeUserConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
