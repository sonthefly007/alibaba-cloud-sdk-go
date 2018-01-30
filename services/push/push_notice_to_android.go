package push

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

func (client *Client) PushNoticeToAndroid(request *PushNoticeToAndroidRequest) (response *PushNoticeToAndroidResponse, err error) {
	response = CreatePushNoticeToAndroidResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) PushNoticeToAndroidWithChan(request *PushNoticeToAndroidRequest) (<-chan *PushNoticeToAndroidResponse, <-chan error) {
	responseChan := make(chan *PushNoticeToAndroidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushNoticeToAndroid(request)
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

func (client *Client) PushNoticeToAndroidWithCallback(request *PushNoticeToAndroidRequest, callback func(response *PushNoticeToAndroidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushNoticeToAndroidResponse
		var err error
		defer close(result)
		response, err = client.PushNoticeToAndroid(request)
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

type PushNoticeToAndroidRequest struct {
	*requests.RpcRequest
	Body          string           `position:"Query" name:"Body"`
	Title         string           `position:"Query" name:"Title"`
	ExtParameters string           `position:"Query" name:"ExtParameters"`
	TargetValue   string           `position:"Query" name:"TargetValue"`
	JobKey        string           `position:"Query" name:"JobKey"`
	Target        string           `position:"Query" name:"Target"`
	AppKey        requests.Integer `position:"Query" name:"AppKey"`
}

type PushNoticeToAndroidResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MessageId string `json:"MessageId" xml:"MessageId"`
}

func CreatePushNoticeToAndroidRequest() (request *PushNoticeToAndroidRequest) {
	request = &PushNoticeToAndroidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "PushNoticeToAndroid", "", "")
	return
}

func CreatePushNoticeToAndroidResponse() (response *PushNoticeToAndroidResponse) {
	response = &PushNoticeToAndroidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
