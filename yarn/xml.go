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

//Root yarn type
type Root struct {
	Scheduler *Scheduler `json:"scheduler"`
}

//Scheduler yarn type
type Scheduler struct {
	SchedulerInfo *SchedulerInfo `json:"schedulerInfo"`
}

//SchedulerInfo yarn type
type SchedulerInfo struct {
	Capacity     int     `json:"capacity"`
	MaxCapacity  int     `json:"maxCapacity"`
	QueueName    string  `json:"queueName"`
	Queues       *Queues `json:"queues,omitempty"`
	Type         string  `json:"type"`
	UsedCapacity int     `json:"usedCapacity"`
}

//ResourcesUsed yarn type
type ResourcesUsed struct {
	Memory int `json:"memory"`
	VCores int `json:"vCores"`
}

//Queue yarn type
type Queue struct {
	AbsoluteCapacity             float64        `json:"absoluteCapacity"`
	AbsoluteMaxCapacity          int            `json:"absoluteMaxCapacity"`
	AbsoluteUsedCapacity         int            `json:"absoluteUsedCapacity"`
	Capacity                     int            `json:"capacity"`
	MaxActiveApplications        int            `json:"maxActiveApplications"`
	MaxActiveApplicationsPerUser int            `json:"maxActiveApplicationsPerUser"`
	MaxApplications              int            `json:"maxApplications"`
	MaxApplicationsPerUser       int            `json:"maxApplicationsPerUser"`
	MaxCapacity                  int            `json:"maxCapacity"`
	NumActiveApplications        int            `json:"numActiveApplications"`
	NumApplications              int            `json:"numApplications"`
	NumContainers                int            `json:"numContainers"`
	NumPendingApplications       int            `json:"numPendingApplications"`
	QueueName                    string         `json:"queueName"`
	ResourcesUsed                *ResourcesUsed `json:"resourcesUsed"`
	State                        string         `json:"state"`
	Type                         string         `json:"type"`
	UsedCapacity                 int            `json:"usedCapacity"`
	UsedResources                string         `json:"usedResources"`
	UserLimit                    int            `json:"userLimit"`
	UserLimitFactor              int            `json:"userLimitFactor"`
	Users                        interface{}    `json:"users"`
	Queues                       *Queues        `json:"queues"`
}

//Queues yarn type
type Queues struct {
	Queue        []Queue `json:"queue"`
	Type         string  `json:"type"`
	UsedCapacity int     `json:"usedCapacity"`
}
