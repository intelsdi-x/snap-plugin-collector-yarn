//
// +build small

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
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/intelsdi-x/snap-plugin-collector-yarn/yarn/httpmock"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	. "github.com/smartystreets/goconvey/convey"
)

func TestYarnPlugin(t *testing.T) {
	httpmock.Mock = true
	var port int64
	port = 80888
	config := plugin.Config{
		"hostname": "192.168.192.200",
		"port":     port,
	}
	Convey("Create Yarn Collector", t, func() {
		yarnCol := YarnCollector{}
		Convey("So Yarn should not be nil", func() {
			So(yarnCol, ShouldNotBeNil)
		})
		Convey("So Yarn should be of Yarn type", func() {
			So(yarnCol, ShouldHaveSameTypeAs, YarnCollector{})
		})
		Convey("yarnCol.GetConfigPolicy() should return a config policy", func() {
			configPolicy, _ := yarnCol.GetConfigPolicy()
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So config policy should be a plugin.ConfigPolicy", func() {
				So(configPolicy, ShouldHaveSameTypeAs, plugin.ConfigPolicy{})
			})
		})
	})
	Convey("Get Metric Yarn Types", t, func() {
		yarnCol := YarnCollector{}
		var cfg = plugin.Config{}
		metrics, err := yarnCol.GetMetricTypes(cfg)
		So(err, ShouldBeNil)
		So(len(metrics), ShouldResemble, 21)
	})

	Convey("Collect metrics should return schedulerInfo metrics", t, func() {
		yarnCol := YarnCollector{}

		defer httpmock.ResetResponders()
		buf := bytes.NewBuffer(nil)

		filename := "./response.json"
		f, _ := os.Open(filename)
		io.Copy(buf, f)
		f.Close()

		httpmock.RegisterResponder("GET", "http://192.168.192.200:80888/ws/v1/cluster/scheduler", buf.String(), 200)
		mts := []plugin.Metric{}
		for _, v := range nsTypes.schedulerInfo {
			mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin, "schedulerinfo", v), Config: config})
		}
		collectedMetrics, err := yarnCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(collectedMetrics), ShouldResemble, 2)

	})
	Convey("Collect metrics should return queue metrics", t, func() {
		yarnCol := YarnCollector{}

		defer httpmock.ResetResponders()
		buf := bytes.NewBuffer(nil)

		filename := "./response.json"
		f, _ := os.Open(filename)
		io.Copy(buf, f)
		f.Close()

		httpmock.RegisterResponder("GET", "http://192.168.192.200:80888/ws/v1/cluster/scheduler", buf.String(), 200)
		mts := []plugin.Metric{}
		for _, v := range nsTypes.queue {
			mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin, "queue").AddDynamicElement("queue_id", "Id of queqe").AddStaticElement(v), Config: config})
		}
		collectedMetrics, err := yarnCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(collectedMetrics), ShouldResemble, 112)
		So(collectedMetrics[0].Data, ShouldResemble, 10.5)
		So(collectedMetrics[10].Data, ShouldResemble, 0)

	})
	Convey("Collect metrics should not return queue metrics when queue dosn't exist", t, func() {
		yarnCol := YarnCollector{}

		defer httpmock.ResetResponders()
		buf := bytes.NewBuffer(nil)

		filename := "./response_wqueue.json"
		f, _ := os.Open(filename)
		io.Copy(buf, f)
		f.Close()

		httpmock.RegisterResponder("GET", "http://192.168.192.200:80888/ws/v1/cluster/scheduler", buf.String(), 200)
		mts := []plugin.Metric{}
		for _, v := range nsTypes.queue {
			mts = append(mts, plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Plugin, "queue").AddDynamicElement("queue_id", "Id of queqe").AddStaticElement(v), Config: config})
		}
		collectedMetrics, err := yarnCol.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(len(collectedMetrics), ShouldResemble, 0)

	})
}
