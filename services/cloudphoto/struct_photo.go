package cloudphoto

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

type Photo struct {
	FileId          string `json:"FileId" xml:"FileId"`
	Width           int    `json:"Width" xml:"Width"`
	State           string `json:"State" xml:"State"`
	Md5             string `json:"Md5" xml:"Md5"`
	Title           string `json:"Title" xml:"Title"`
	Height          int    `json:"Height" xml:"Height"`
	Mtime           int    `json:"Mtime" xml:"Mtime"`
	TakenAt         int    `json:"TakenAt" xml:"TakenAt"`
	Ctime           int    `json:"Ctime" xml:"Ctime"`
	Location        string `json:"Location" xml:"Location"`
	IsVideo         bool   `json:"IsVideo" xml:"IsVideo"`
	Like            int    `json:"Like" xml:"Like"`
	Id              int    `json:"Id" xml:"Id"`
	Size            int    `json:"Size" xml:"Size"`
	Remark          string `json:"Remark" xml:"Remark"`
	InactiveTime    int    `json:"InactiveTime" xml:"InactiveTime"`
	ShareExpireTime int    `json:"ShareExpireTime" xml:"ShareExpireTime"`
}