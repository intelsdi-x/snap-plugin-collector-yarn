# Snap collector plugin - yarn

## Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Description
----------|-----------------------
/intel/yarn/queue/*/absolutecapacity | Absolute capacity percentage this queue can use of entire cluster
/intel/yarn/queue/*/absolutemaxcapacity | Absolute maximum capacity percentage this queue can use of the entire cluster
/intel/yarn/queue/*/absoluteusedcapacity | Absolute used capacity percentage this queue is using of the entire cluster
/intel/yarn/queue/*/capacity | Configured queue capacity in percentage relative to its parent queue
/intel/yarn/queue/*/maxactiveapplications | The maximum number of active applications this queue can have
/intel/yarn/queue/*/maxactiveapplicationsperuser | The maximum number of active applications this queue can have
/intel/yarn/queue/*/maxapplications | The maximum number of active applications this queue can have
/intel/yarn/queue/*/maxapplicationsperuser | The maximum number of active applications per user this queue can have
/intel/yarn/queue/*/maxcapacity | Configured maximum queue capacity in percentage relative to its parent queue
/intel/yarn/queue/*/numactiveapplications | The number of pending applications for this user in this queue
/intel/yarn/queue/*/numapplications | The number of applications currently in the queue
/intel/yarn/queue/*/numcontainers | The number of containers being used
/intel/yarn/queue/*/numpendingapplications | The number of pending applications for this user in this queue
/intel/yarn/queue/*/usedcapacity | Used queue capacity in percentage
/intel/yarn/queue/*/usedresources | Used queue resources
/intel/yarn/queue/*/userlimit | The minimum user limit percent set in the configuration
/intel/yarn/queue/*/userlimitfactor | The user limit factor set in the configuration
/intel/yarn/queue/*/resources_memory | The amount of memory used (in MB)
/intel/yarn/queue/*/resources_vcores | The number of virtual cores
/intel/yarn/schedulerinfo/capacity | Configured queue capacity in percentage for root queue
/intel/yarn/schedulerinfo/maxcapacity | Configured maximum queue capacity in percentage for root queue
