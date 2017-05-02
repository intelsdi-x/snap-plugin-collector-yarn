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

const (
	// Name of plugin
	Name = "yarn"
	// Vendor  prefix
	Vendor = "intel"
	// Plugin plugin name
	Plugin = "yarn"
	// Version of plugin
	Version         = 2
	nsMetricPostion = 2
	nsSubMetric     = 3
	nsSubSubMetric  = 4
)

// YarnCollector type
type YarnCollector struct {
}

// CollectMetrics returns collected metrics
func (YarnCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}

	_, schedulerInfoMetrics := filterNamespace("schedulerinfo", mts)
	_, queueMetrics := filterNamespace("queue", mts)

	hostname, err := mts[0].Config.GetString("hostname")
	if err != nil {
		return nil, err
	}
	port, err := mts[0].Config.GetInt("port")
	if err != nil {
		return nil, err
	}

	response, err := yarn.HadoopRequest(hostname, port, "ws/v1/cluster/scheduler")
	if err != nil {
		return nil, err
	}
	schedulerInfo, err := yarn.GetSchedulerInfo(response)

	for _, mt := range schedulerInfoMetrics {
		ns := mt.Namespace
		switch ns[nsSubMetric].Value {
		case "capacity":
			metrics = append(metrics, createMeasurement(mt, schedulerInfo.SchedulerInfo.Capacity, ns))

		case "maxcapacity":
			metrics = append(metrics, createMeasurement(mt, schedulerInfo.SchedulerInfo.MaxCapacity, ns))

		}
	}
	queueCounters, err := getQueueInfo(schedulerInfo.SchedulerInfo, queueMetrics)
	if err != nil {
		return nil, err
	}
	metrics = append(metrics, queueCounters...)

	return metrics, nil

}

// GetConfigPolicy returns a config policy
func (YarnCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	configKey := []string{"intel", "yarn"}
	policy.AddNewStringRule(configKey,
		"hostname",
		false)
	policy.AddNewIntRule(configKey,
		"port",
		false)
	return *policy, nil
}

// GetMetricTypes returns metric types that can be collected
func (YarnCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {

	var metrics []plugin.Metric
	ns := plugin.NewNamespace(Vendor, Plugin)
	for _, v := range nsTypes.queue {
		metrics = append(metrics, createMetric(ns.AddStaticElement("queue").AddDynamicElement("queue_id", "Id of the queue").AddStaticElement(v)))
	}
	for _, v := range nsTypes.schedulerInfo {
		metrics = append(metrics, createMetric(ns.AddStaticElements("schedulerinfo", v)))
	}

	return metrics, nil
}
