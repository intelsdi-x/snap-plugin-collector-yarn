/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Simple HTTP mocking mechanism inspired by https://github.com/jarcoal/httpmock

package httpmock

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

//Mock is a boolean value indicates that Mock is turned off or turned on
var Mock bool = false

type responder struct {
	reqType  string
	reqURL   string
	response string
	status   int
}

var responders = []responder{}

//RegisterResponder registering http responder
func RegisterResponder(reqType, reqURL, response string, status int) {
	newResponder := responder{reqType, reqURL, response, status}
	responders = append(responders, newResponder)
}

//ResetResponders resetting http responder
func ResetResponders() {
	responders = nil
}

//Get returns Mock or real http object
func Get(url string) (resp *http.Response, err error) {
	if Mock {
		return createResponse(url, "GET")
	}
	return http.Get(url)

}

//PostForm returns Mock or real http post response
func PostForm(url string, data url.Values) (resp *http.Response, err error) {
	if Mock {
		return createResponse(url, "POST")
	}
	return http.PostForm(url, data)

}

func createResponse(url, reqType string) (resp *http.Response, err error) {
	for _, r := range responders {
		if r.reqURL == url && strings.ToUpper(r.reqType) == reqType {
			return &http.Response{
				StatusCode: r.status,
				Body:       responseBodyFromString(r.response),
			}, nil
		}
	}
	return nil, fmt.Errorf("URL %s not registered as responder for %s", url, reqType)
}

func responseBodyFromString(body string) io.ReadCloser {
	return &dummyReadCloser{strings.NewReader(body)}
}

type dummyReadCloser struct {
	body io.ReadSeeker
}

func (d *dummyReadCloser) Read(p []byte) (n int, err error) {
	n, err = d.body.Read(p)
	if err == io.EOF {
		d.body.Seek(0, 0)
	}
	return n, err
}

func (d *dummyReadCloser) Close() error {
	return nil
}
