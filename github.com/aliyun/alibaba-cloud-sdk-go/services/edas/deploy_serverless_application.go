package edas

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

// DeployServerlessApplication invokes the edas.DeployServerlessApplication API synchronously
// api document: https://help.aliyun.com/api/edas/deployserverlessapplication.html
func (client *Client) DeployServerlessApplication(request *DeployServerlessApplicationRequest) (response *DeployServerlessApplicationResponse, err error) {
	response = CreateDeployServerlessApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// DeployServerlessApplicationWithChan invokes the edas.DeployServerlessApplication API asynchronously
// api document: https://help.aliyun.com/api/edas/deployserverlessapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeployServerlessApplicationWithChan(request *DeployServerlessApplicationRequest) (<-chan *DeployServerlessApplicationResponse, <-chan error) {
	responseChan := make(chan *DeployServerlessApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeployServerlessApplication(request)
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

// DeployServerlessApplicationWithCallback invokes the edas.DeployServerlessApplication API asynchronously
// api document: https://help.aliyun.com/api/edas/deployserverlessapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeployServerlessApplicationWithCallback(request *DeployServerlessApplicationRequest, callback func(response *DeployServerlessApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeployServerlessApplicationResponse
		var err error
		defer close(result)
		response, err = client.DeployServerlessApplication(request)
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

// DeployServerlessApplicationRequest is the request struct for api DeployServerlessApplication
type DeployServerlessApplicationRequest struct {
	*requests.RoaRequest
	WebContainer      string           `position:"Query" name:"WebContainer"`
	JarStartArgs      string           `position:"Query" name:"JarStartArgs"`
	CommandArgs       string           `position:"Query" name:"CommandArgs"`
	Readiness         string           `position:"Query" name:"Readiness"`
	BatchWaitTime     requests.Integer `position:"Query" name:"BatchWaitTime"`
	Liveness          string           `position:"Query" name:"Liveness"`
	Envs              string           `position:"Query" name:"Envs"`
	PackageVersion    string           `position:"Query" name:"PackageVersion"`
	Command           string           `position:"Query" name:"Command"`
	CustomHostAlias   string           `position:"Query" name:"CustomHostAlias"`
	Jdk               string           `position:"Query" name:"Jdk"`
	JarStartOptions   string           `position:"Query" name:"JarStartOptions"`
	MinReadyInstances requests.Integer `position:"Query" name:"MinReadyInstances"`
	PackageUrl        string           `position:"Query" name:"PackageUrl"`
	AppId             string           `position:"Query" name:"AppId"`
	ImageUrl          string           `position:"Query" name:"ImageUrl"`
}

// DeployServerlessApplicationResponse is the response struct for api DeployServerlessApplication
type DeployServerlessApplicationResponse struct {
	*responses.BaseResponse
	Code    int    `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
	Data    Data   `json:"Data" xml:"Data"`
}

// CreateDeployServerlessApplicationRequest creates a request to invoke DeployServerlessApplication API
func CreateDeployServerlessApplicationRequest() (request *DeployServerlessApplicationRequest) {
	request = &DeployServerlessApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "DeployServerlessApplication", "/pop/v5/k8s/pop/pop_serverless_app_deploy", "", "")
	request.Method = requests.POST
	return
}

// CreateDeployServerlessApplicationResponse creates a response to parse from DeployServerlessApplication response
func CreateDeployServerlessApplicationResponse() (response *DeployServerlessApplicationResponse) {
	response = &DeployServerlessApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
