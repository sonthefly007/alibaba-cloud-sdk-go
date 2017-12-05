package cloudapi

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

func (client *Client) DescribeVpcAccesses(request *DescribeVpcAccessesRequest) (response *DescribeVpcAccessesResponse, err error) {
	response = CreateDescribeVpcAccessesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeVpcAccessesWithChan(request *DescribeVpcAccessesRequest) (<-chan *DescribeVpcAccessesResponse, <-chan error) {
	responseChan := make(chan *DescribeVpcAccessesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpcAccesses(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeVpcAccessesWithCallback(request *DescribeVpcAccessesRequest, callback func(response *DescribeVpcAccessesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpcAccessesResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpcAccesses(request)
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

type DescribeVpcAccessesRequest struct {
	*requests.RpcRequest
	PageSize   string `position:"Query" name:"PageSize"`
	PageNumber string `position:"Query" name:"PageNumber"`
}

type DescribeVpcAccessesResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	TotalCount          int    `json:"TotalCount" xml:"TotalCount"`
	PageSize            int    `json:"PageSize" xml:"PageSize"`
	PageNumber          int    `json:"PageNumber" xml:"PageNumber"`
	VpcAccessAttributes []struct {
		VpcId       string `json:"VpcId" xml:"VpcId"`
		InstanceId  string `json:"InstanceId" xml:"InstanceId"`
		CreatedTime string `json:"CreatedTime" xml:"CreatedTime"`
		Port        int    `json:"Port" xml:"Port"`
		RegionId    string `json:"RegionId" xml:"RegionId"`
		Name        string `json:"Name" xml:"Name"`
	} `json:"VpcAccessAttributes" xml:"VpcAccessAttributes"`
}

func CreateDescribeVpcAccessesRequest() (request *DescribeVpcAccessesRequest) {
	request = &DescribeVpcAccessesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeVpcAccesses", "", "")
	return
}

func CreateDescribeVpcAccessesResponse() (response *DescribeVpcAccessesResponse) {
	response = &DescribeVpcAccessesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}