// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Original source: github.com/micro/go-micro/v3/util/kubernetes/api/response.go

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Errors ...
var (
	ErrNotFound = errors.New("kubernetes: resource not found")
	ErrDecode   = errors.New("kubernetes: error decoding")
	ErrUnknown  = errors.New("kubernetes: unknown error")
)

// Status is an object that is returned when a request
// failed or delete succeeded.
type Status struct {
	Kind    string `json:"kind"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
	Code    int    `json:"code"`
}

// Response ...
type Response struct {
	res *http.Response
	err error

	body []byte
}

// Error returns an error
func (r *Response) Error() error {
	return r.err
}

// StatusCode returns status code for response
func (r *Response) StatusCode() int {
	return r.res.StatusCode
}

// Into decode body into `data`
func (r *Response) Into(data interface{}) error {
	if r.err != nil {
		return r.err
	}

	defer r.res.Body.Close()
	decoder := json.NewDecoder(r.res.Body)
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("%v: %v", ErrDecode, err)
	}

	return r.err
}

func (r *Response) Close() error {
	return r.res.Body.Close()
}

func newResponse(res *http.Response, err error) *Response {
	r := &Response{
		res: res,
		err: err,
	}

	if err != nil {
		return r
	}

	if r.res.StatusCode == http.StatusOK ||
		r.res.StatusCode == http.StatusCreated ||
		r.res.StatusCode == http.StatusNoContent {
		// Non error status code
		return r
	}

	if r.res.StatusCode == http.StatusNotFound {
		r.err = ErrNotFound
		return r
	}

	b, err := ioutil.ReadAll(r.res.Body)
	if err == nil {
		r.err = errors.New(string(b))
		return r
	}

	r.err = ErrUnknown

	return r
}
