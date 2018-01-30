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

func (client *Client) DescribeDomainsUsageByDay(request *DescribeDomainsUsageByDayRequest) (response *DescribeDomainsUsageByDayResponse, err error) {
	response = CreateDescribeDomainsUsageByDayResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainsUsageByDayWithChan(request *DescribeDomainsUsageByDayRequest) (<-chan *DescribeDomainsUsageByDayResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainsUsageByDayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainsUsageByDay(request)
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

func (client *Client) DescribeDomainsUsageByDayWithCallback(request *DescribeDomainsUsageByDayRequest, callback func(response *DescribeDomainsUsageByDayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainsUsageByDayResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainsUsageByDay(request)
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

type DescribeDomainsUsageByDayRequest struct {
	*requests.RpcRequest
	EndTime       string           `position:"Query" name:"EndTime"`
	StartTime     string           `position:"Query" name:"StartTime"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeDomainsUsageByDayResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DomainName   string `json:"DomainName" xml:"DomainName"`
	DataInterval string `json:"DataInterval" xml:"DataInterval"`
	StartTime    string `json:"StartTime" xml:"StartTime"`
	EndTime      string `json:"EndTime" xml:"EndTime"`
	UsageTotal   struct {
		BytesHitRate   string `json:"BytesHitRate" xml:"BytesHitRate"`
		RequestHitRate string `json:"RequestHitRate" xml:"RequestHitRate"`
		MaxBps         string `json:"MaxBps" xml:"MaxBps"`
		MaxBpsTime     string `json:"MaxBpsTime" xml:"MaxBpsTime"`
		MaxSrcBps      string `json:"MaxSrcBps" xml:"MaxSrcBps"`
		MaxSrcBpsTime  string `json:"MaxSrcBpsTime" xml:"MaxSrcBpsTime"`
		TotalAccess    string `json:"TotalAccess" xml:"TotalAccess"`
		TotalTraffic   string `json:"TotalTraffic" xml:"TotalTraffic"`
	} `json:"UsageTotal" xml:"UsageTotal"`
	UsageByDays struct {
		UsageByDay []struct {
			TimeStamp      string `json:"TimeStamp" xml:"TimeStamp"`
			Qps            string `json:"Qps" xml:"Qps"`
			BytesHitRate   string `json:"BytesHitRate" xml:"BytesHitRate"`
			RequestHitRate string `json:"RequestHitRate" xml:"RequestHitRate"`
			MaxBps         string `json:"MaxBps" xml:"MaxBps"`
			MaxBpsTime     string `json:"MaxBpsTime" xml:"MaxBpsTime"`
			MaxSrcBps      string `json:"MaxSrcBps" xml:"MaxSrcBps"`
			MaxSrcBpsTime  string `json:"MaxSrcBpsTime" xml:"MaxSrcBpsTime"`
			TotalAccess    string `json:"TotalAccess" xml:"TotalAccess"`
			TotalTraffic   string `json:"TotalTraffic" xml:"TotalTraffic"`
		} `json:"UsageByDay" xml:"UsageByDay"`
	} `json:"UsageByDays" xml:"UsageByDays"`
}

func CreateDescribeDomainsUsageByDayRequest() (request *DescribeDomainsUsageByDayRequest) {
	request = &DescribeDomainsUsageByDayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainsUsageByDay", "", "")
	return
}

func CreateDescribeDomainsUsageByDayResponse() (response *DescribeDomainsUsageByDayResponse) {
	response = &DescribeDomainsUsageByDayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
