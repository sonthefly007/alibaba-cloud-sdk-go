package vod

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

// CDNMetric is a nested struct in vod response
type CDNMetric struct {
	StatTime              string `json:"StatTime" xml:"StatTime"`
	FlowDataDomesticValue int    `json:"FlowDataDomesticValue" xml:"FlowDataDomesticValue"`
	FlowDataOverseasValue int    `json:"FlowDataOverseasValue" xml:"FlowDataOverseasValue"`
	FlowDataValue         int    `json:"FlowDataValue" xml:"FlowDataValue"`
	BpsDataDomesticValue  int    `json:"BpsDataDomesticValue" xml:"BpsDataDomesticValue"`
	BpsDataOverseasValue  int    `json:"BpsDataOverseasValue" xml:"BpsDataOverseasValue"`
	BpsDataValue          int    `json:"BpsDataValue" xml:"BpsDataValue"`
}
