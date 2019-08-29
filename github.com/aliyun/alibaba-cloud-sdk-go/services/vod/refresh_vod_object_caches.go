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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// RefreshVodObjectCaches invokes the vod.RefreshVodObjectCaches API synchronously
// api document: https://help.aliyun.com/api/vod/refreshvodobjectcaches.html
func (client *Client) RefreshVodObjectCaches(request *RefreshVodObjectCachesRequest) (response *RefreshVodObjectCachesResponse, err error) {
	response = CreateRefreshVodObjectCachesResponse()
	err = client.DoAction(request, response)
	return
}

// RefreshVodObjectCachesWithChan invokes the vod.RefreshVodObjectCaches API asynchronously
// api document: https://help.aliyun.com/api/vod/refreshvodobjectcaches.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshVodObjectCachesWithChan(request *RefreshVodObjectCachesRequest) (<-chan *RefreshVodObjectCachesResponse, <-chan error) {
	responseChan := make(chan *RefreshVodObjectCachesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefreshVodObjectCaches(request)
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

// RefreshVodObjectCachesWithCallback invokes the vod.RefreshVodObjectCaches API asynchronously
// api document: https://help.aliyun.com/api/vod/refreshvodobjectcaches.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshVodObjectCachesWithCallback(request *RefreshVodObjectCachesRequest, callback func(response *RefreshVodObjectCachesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefreshVodObjectCachesResponse
		var err error
		defer close(result)
		response, err = client.RefreshVodObjectCaches(request)
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

// RefreshVodObjectCachesRequest is the request struct for api RefreshVodObjectCaches
type RefreshVodObjectCachesRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	ObjectPath    string           `position:"Query" name:"ObjectPath"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	ObjectType    string           `position:"Query" name:"ObjectType"`
}

// RefreshVodObjectCachesResponse is the response struct for api RefreshVodObjectCaches
type RefreshVodObjectCachesResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	RefreshTaskId string `json:"RefreshTaskId" xml:"RefreshTaskId"`
}

// CreateRefreshVodObjectCachesRequest creates a request to invoke RefreshVodObjectCaches API
func CreateRefreshVodObjectCachesRequest() (request *RefreshVodObjectCachesRequest) {
	request = &RefreshVodObjectCachesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "RefreshVodObjectCaches", "vod", "openAPI")
	return
}

// CreateRefreshVodObjectCachesResponse creates a response to parse from RefreshVodObjectCaches response
func CreateRefreshVodObjectCachesResponse() (response *RefreshVodObjectCachesResponse) {
	response = &RefreshVodObjectCachesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}