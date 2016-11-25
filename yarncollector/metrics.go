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

package yarncollector

import (
	"github.com/intelsdi-x/snap-plugin-collector-yarn/yarn"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

var nsTypes = struct {
	queue         []string
	schedulerInfo []string
}{
	queue: []string{"absolutecapacity", "absolutemaxcapacity", "absoluteusedcapacity", "capacity", "maxactiveapplications", "maxactiveapplicationsperuser", "maxapplications",
		"maxapplicationsperuser", "maxcapacity", "numactiveapplications", "numapplications", "numcontainers", "numpendingapplications", "usedcapacity", "usedresources", "userlimit", "userlimitfactor",
		"resources_memory", "resources_vcores"},
	schedulerInfo: []string{"capacity", "maxcapacity"},
}

func getQueueInfo(info *yarn.SchedulerInfo, mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}

	for _, y := range getQueues(info.Queues.Queue) {
		metrics = append(metrics, getQueueMetric(y, mts)...)

		for _, x := range getQueues(y.Queues.Queue) {
			metrics = append(metrics, getQueueMetric(x, mts)...)
		}
	}
	return metrics, nil

}

//getQueues return Queues and check if Queue is not empty
func getQueues(queue []yarn.Queue) []yarn.Queue {
	if queue != nil {
		return queue
	}
	return []yarn.Queue{}
}

func getQueueMetric(queue yarn.Queue, mts []plugin.Metric) []plugin.Metric {
	metrics := []plugin.Metric{}

	for _, mt := range mts {
		ns := copyNamespace(mt)
		ns[nsSubMetric].Value = queue.QueueName
		switch ns[nsSubSubMetric].Value {
		case "absolutecapacity":
			metrics = append(metrics, createMeasurement(mt, queue.AbsoluteCapacity, ns))
		case "absoluteusedcapacity":
			metrics = append(metrics, createMeasurement(mt, queue.AbsoluteUsedCapacity, ns))
		case "capacity":
			metrics = append(metrics, createMeasurement(mt, queue.Capacity, ns))
		case "maxactiveapplications":
			metrics = append(metrics, createMeasurement(mt, queue.MaxActiveApplications, ns))
		case "maxactiveapplicationsperuser":
			metrics = append(metrics, createMeasurement(mt, queue.MaxActiveApplicationsPerUser, ns))
		case "maxapplications":
			metrics = append(metrics, createMeasurement(mt, queue.MaxApplications, ns))
		case "maxapplicationsperuser":
			metrics = append(metrics, createMeasurement(mt, queue.MaxApplicationsPerUser, ns))
		case "numactiveapplications":
			metrics = append(metrics, createMeasurement(mt, queue.NumActiveApplications, ns))
		case "numapplications":
			metrics = append(metrics, createMeasurement(mt, queue.NumApplications, ns))
		case "numcontainers":
			metrics = append(metrics, createMeasurement(mt, queue.NumContainers, ns))
		case "numpendingapplications":
			metrics = append(metrics, createMeasurement(mt, queue.NumPendingApplications, ns))
		case "usedcapacity":
			metrics = append(metrics, createMeasurement(mt, queue.UsedCapacity, ns))
		case "usedresources":
			metrics = append(metrics, createMeasurement(mt, queue.UsedResources, ns))
		case "userlimit":
			metrics = append(metrics, createMeasurement(mt, queue.UserLimit, ns))
		case "userlimitfactor":
			metrics = append(metrics, createMeasurement(mt, queue.UserLimitFactor, ns))
		case "resources_vcore":
			if queue.ResourcesUsed != nil {
				metrics = append(metrics, createMeasurement(mt, queue.ResourcesUsed.VCores, ns))
			}
		case "resources_memory":
			if queue.ResourcesUsed != nil {
				metrics = append(metrics, createMeasurement(mt, queue.ResourcesUsed.Memory, ns))
			}
		}
	}
	return metrics
}
