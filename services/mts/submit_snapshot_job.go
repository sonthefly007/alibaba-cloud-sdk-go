package mts

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

func (client *Client) SubmitSnapshotJob(request *SubmitSnapshotJobRequest) (response *SubmitSnapshotJobResponse, err error) {
	response = CreateSubmitSnapshotJobResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SubmitSnapshotJobWithChan(request *SubmitSnapshotJobRequest) (<-chan *SubmitSnapshotJobResponse, <-chan error) {
	responseChan := make(chan *SubmitSnapshotJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitSnapshotJob(request)
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

func (client *Client) SubmitSnapshotJobWithCallback(request *SubmitSnapshotJobRequest, callback func(response *SubmitSnapshotJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitSnapshotJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitSnapshotJob(request)
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

type SubmitSnapshotJobRequest struct {
	*requests.RpcRequest
	UserData             string           `position:"Query" name:"UserData"`
	Input                string           `position:"Query" name:"Input"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	SnapshotConfig       string           `position:"Query" name:"SnapshotConfig"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type SubmitSnapshotJobResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	SnapshotJob struct {
		Id           string `json:"Id" xml:"Id"`
		UserData     string `json:"UserData" xml:"UserData"`
		PipelineId   string `json:"PipelineId" xml:"PipelineId"`
		State        string `json:"State" xml:"State"`
		Code         string `json:"Code" xml:"Code"`
		Count        string `json:"Count" xml:"Count"`
		TileCount    string `json:"TileCount" xml:"TileCount"`
		Message      string `json:"Message" xml:"Message"`
		CreationTime string `json:"CreationTime" xml:"CreationTime"`
		Input        struct {
			Bucket   string `json:"Bucket" xml:"Bucket"`
			Location string `json:"Location" xml:"Location"`
			Object   string `json:"Object" xml:"Object"`
			RoleArn  string `json:"RoleArn" xml:"RoleArn"`
		} `json:"Input" xml:"Input"`
		SnapshotConfig struct {
			Time       string `json:"Time" xml:"Time"`
			Interval   string `json:"Interval" xml:"Interval"`
			Num        string `json:"Num" xml:"Num"`
			Width      string `json:"Width" xml:"Width"`
			Height     string `json:"Height" xml:"Height"`
			FrameType  string `json:"FrameType" xml:"FrameType"`
			OutputFile struct {
				Bucket   string `json:"Bucket" xml:"Bucket"`
				Location string `json:"Location" xml:"Location"`
				Object   string `json:"Object" xml:"Object"`
				RoleArn  string `json:"RoleArn" xml:"RoleArn"`
			} `json:"OutputFile" xml:"OutputFile"`
			TileOutputFile struct {
				Bucket   string `json:"Bucket" xml:"Bucket"`
				Location string `json:"Location" xml:"Location"`
				Object   string `json:"Object" xml:"Object"`
				RoleArn  string `json:"RoleArn" xml:"RoleArn"`
			} `json:"TileOutputFile" xml:"TileOutputFile"`
			TileOut struct {
				Lines         string `json:"Lines" xml:"Lines"`
				Columns       string `json:"Columns" xml:"Columns"`
				CellWidth     string `json:"CellWidth" xml:"CellWidth"`
				CellHeight    string `json:"CellHeight" xml:"CellHeight"`
				Margin        string `json:"Margin" xml:"Margin"`
				Padding       string `json:"Padding" xml:"Padding"`
				Color         string `json:"Color" xml:"Color"`
				IsKeepCellPic string `json:"IsKeepCellPic" xml:"IsKeepCellPic"`
				CellSelStep   string `json:"CellSelStep" xml:"CellSelStep"`
			} `json:"TileOut" xml:"TileOut"`
		} `json:"SnapshotConfig" xml:"SnapshotConfig"`
		MNSMessageResult struct {
			MessageId    string `json:"MessageId" xml:"MessageId"`
			ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
			ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
		} `json:"MNSMessageResult" xml:"MNSMessageResult"`
	} `json:"SnapshotJob" xml:"SnapshotJob"`
}

func CreateSubmitSnapshotJobRequest() (request *SubmitSnapshotJobRequest) {
	request = &SubmitSnapshotJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitSnapshotJob", "mts", "openAPI")
	return
}

func CreateSubmitSnapshotJobResponse() (response *SubmitSnapshotJobResponse) {
	response = &SubmitSnapshotJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
