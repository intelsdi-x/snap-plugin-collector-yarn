
# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**




[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-collector-yarn.svg)](https://travis-ci.org/intelsdi-x/snap-plugin-collector-yarn )
[![Go Report Card](http://goreportcard.com/badge/intelsdi-x/snap-plugin-collector-yarn)](http://goreportcard.com/report/intelsdi-x/snap-plugin-collector-yarn)

This plugin collects metrics from Yarn hadoop scheduler.  

It's used in the [Snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Operating systems](#operating-systems)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Global Config](#global-config)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements
* [golang 1.7+](https://golang.org/dl/)  - needed only for building
* [hadoop 3.2.x](http://hadoop.apache.org/) 
### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation


#### Download the plugin binary:

You can get the pre-built binaries for your OS and architecture from the plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-collector-yarn/releases) page. Download the plugin from the latest release and load it into `snapteld` (`/opt/snap/plugins` is the default location for snap packages).


#### To build the plugin binary:

Fork https://github.com/intelsdi-x/snap-plugin-collector-yarn
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-yarn.git
```

Build the snap yarn plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build/`

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started).
* Configure connection strings to Hadoop resource manager which includes hostname and port.

* Load the plugin and create a task, see example in [Examples](#examples).

## Documentation

This collector gathers metrics from yarn server status command. 

### Global config
Global configuration files are described in [snap's documentation](https://github.com/intelsdi-x/snap/blob/master/docs/snapteld_CONFIGURATION.md). You have to add `"yarn"` section with following entries:

 - `"hostname"` - hostname of the hadoop cluster
 - `"port"` - resource manager port number

See exemplary Global configuration files in [examplary config files] (examples/configs/).

### Collected Metrics

List of collected metrics is described in [METRICS.md](METRICS.md).

### Examples

Example of running snap yarn collector and writing data to file.

Ensure [snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
* initd: `service snap-telemetry start`
* systemd: `systemctl start snap-telemetry`
* command line: `snapteld -l 1 -t 0 &`

Download and load snap plugins:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-yarn/latest/linux/x86_64/snap-plugin-collector-yarn
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ chmod 755 snap-plugin-*
$ snaptel plugin load snap-plugin-collector-yarn
$ snaptel plugin load snap-plugin-publisher-file

Create a task manifest file  (exemplary files in [examples/tasks/] (examples/tasks/)):
```yaml
---
  version: 1
  schedule:
    type: "simple"
    interval: "1s"
  max-failures: 10
  workflow:
    collect:
      metrics:
        /intel/yarn/queue/*/absolutecapacity: {}
        /intel/yarn/queue/*/absolutemaxcapacity: {}
        /intel/yarn/queue/*/absoluteusedcapacity: {}
        /intel/yarn/queue/*/capacity: {}
        /intel/yarn/queue/*/maxactiveapplications: {}
        /intel/yarn/queue/*/maxactiveapplicationsperuser: {}
        /intel/yarn/queue/*/maxapplications: {}
        /intel/yarn/queue/*/maxapplicationsperuser: {}
        /intel/yarn/queue/*/maxcapacity: {}
        /intel/yarn/queue/*/numactiveapplications: {}
        /intel/yarn/queue/*/numapplications: {}
        /intel/yarn/queue/*/numcontainers: {}
        /intel/yarn/queue/*/numpendingapplications: {}
        /intel/yarn/queue/*/usedcapacity: {}
        /intel/yarn/queue/*/usedresources: {}
        /intel/yarn/queue/*/userlimit: {}
        /intel/yarn/queue/*/userlimitfactor: {}
        /intel/yarn/queue/*/resources_memory: {}
        /intel/yarn/queue/*/resources_vcores: {}
        /intel/yarn/schedulerinfo/capacity: {}
        /intel/yarn/schedulerinfo/maxcapacity: {}
      publish:
        - plugin_name: "file"
          config:
            file: "/tmp/yarn_metrics"
```
Download an [example task file](https://github.com/intelsdi-x/snap-plugin-collector-yarn/blob/master/examples/tasks/) and load it:
```
$ curl -sfLO https://raw.githubusercontent.com/intelsdi-x/snap-plugin-collector-yarn/master/examples/tasks/yarn-file.yml
$ snaptel task create -t yarn-file.json
Using task manifest to create task
Task created
ID: 480323af-15b0-4af8-a526-eb2ca6d8ae67
Name: Task-480323af-15b0-4af8-a526-eb2ca6d8ae67
State: Running
```

See realtime output from `snaptel task watch <task_id>` (CTRL+C to exit)
```
$ snaptel task watch 480323af-15b0-4af8-a526-eb2ca6d8ae67
```

This data is published to a file `/tmp/yarn_metrics` per task specification

Stop task:
```
$ snaptel task stop 480323af-15b0-4af8-a526-eb2ca6d8ae67
Task stopped:
ID: 480323af-15b0-4af8-a526-eb2ca6d8ae67
```



### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. 

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-yarn/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-yarn/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap.

To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support) or visit [Slack](http://slack.snap-telemetry.io).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
Snap, along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@Marcin Spoczynski](https://github.com/sandlbn/)

This software has been contributed by MIKELANGELO, a Horizon 2020 project co-funded by the European Union. https://www.mikelangelo-project.eu/
## Thank You
And **thank you!** Your contribution, through code and participation, is incredibly important to us.
