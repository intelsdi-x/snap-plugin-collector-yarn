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

package yarn

import (
	"encoding/json"
	"errors"
)

// Yarn interface
type Yarn interface {
	HadoopRequest(string, string) ([]byte, error)
	GetSchedulerInfo([]byte) (SchedulerInfo, error)
}

//HadoopRequest calls hadoop
func HadoopRequest(hostname string, port int64, path string) ([]byte, error) {

	response, err := hadoopRestGet(hostname, port, path)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//GetSchedulerInfo returns Yarn Scheduling Information
func GetSchedulerInfo(response []byte) (*Scheduler, error) {
	yarnRoot := Root{}

	json.Unmarshal(response, &yarnRoot)
	if yarnRoot.Scheduler == nil {
		return &Scheduler{}, errors.New("Error parsing Yarn Scheduler Info Response")
	}
	return yarnRoot.Scheduler, nil
}
