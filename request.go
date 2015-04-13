/*
 * Copyright 2015 Xuyuan Pang
 * Author: Xuyuan Pang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package httputil

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// ParseJSON parses request body into v as json format.
func ParseJSON(req *http.Request, v interface{}) error {
	return json.NewDecoder(req.Body).Decode(v)
}

// ParseXML parses request body into v as xml format.
func ParseXML(req *http.Request, v interface{}) error {
	return xml.NewDecoder(req.Body).Decode(v)
}

// PostJSON sends post request to the url with v in json format
func PostJSON(url string, v interface{}) (*http.Response, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buffer).Encode(v); err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json; charset=utf-8", buffer)
	return resp, err
}

// PostXML sends post request to the url with v in xml format
func PostXML(url string, v interface{}) (*http.Response, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := xml.NewEncoder(buffer).Encode(v); err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json; charset=utf-8", buffer)
	return resp, err
}
