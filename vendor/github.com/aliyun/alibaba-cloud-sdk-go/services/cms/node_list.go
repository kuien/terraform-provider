package cms

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

// NodeList invokes the cms.NodeList API synchronously
// api document: https://help.aliyun.com/api/cms/nodelist.html
func (client *Client) NodeList(request *NodeListRequest) (response *NodeListResponse, err error) {
	response = CreateNodeListResponse()
	err = client.DoAction(request, response)
	return
}

// NodeListWithChan invokes the cms.NodeList API asynchronously
// api document: https://help.aliyun.com/api/cms/nodelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) NodeListWithChan(request *NodeListRequest) (<-chan *NodeListResponse, <-chan error) {
	responseChan := make(chan *NodeListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.NodeList(request)
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

// NodeListWithCallback invokes the cms.NodeList API asynchronously
// api document: https://help.aliyun.com/api/cms/nodelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) NodeListWithCallback(request *NodeListRequest, callback func(response *NodeListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *NodeListResponse
		var err error
		defer close(result)
		response, err = client.NodeList(request)
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

// NodeListRequest is the request struct for api NodeList
type NodeListRequest struct {
	*requests.RpcRequest
	PageNumber       requests.Integer `position:"Query" name:"PageNumber"`
	UserId           requests.Integer `position:"Query" name:"UserId"`
	HostName         string           `position:"Query" name:"HostName"`
	InstanceIds      string           `position:"Query" name:"InstanceIds"`
	SerialNumbers    string           `position:"Query" name:"SerialNumbers"`
	KeyWord          string           `position:"Query" name:"KeyWord"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	Status           string           `position:"Query" name:"Status"`
	InstanceRegionId string           `position:"Query" name:"InstanceRegionId"`
}

// NodeListResponse is the response struct for api NodeList
type NodeListResponse struct {
	*responses.BaseResponse
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	PageNumber   int    `json:"PageNumber" xml:"PageNumber"`
	PageSize     int    `json:"PageSize" xml:"PageSize"`
	PageTotal    int    `json:"PageTotal" xml:"PageTotal"`
	Total        int    `json:"Total" xml:"Total"`
	Nodes        Nodes  `json:"Nodes" xml:"Nodes"`
}

// CreateNodeListRequest creates a request to invoke NodeList API
func CreateNodeListRequest() (request *NodeListRequest) {
	request = &NodeListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "NodeList", "cms", "openAPI")
	return
}

// CreateNodeListResponse creates a response to parse from NodeList response
func CreateNodeListResponse() (response *NodeListResponse) {
	response = &NodeListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
