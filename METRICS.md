# Snap collector plugin - yarn

## Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Data Type | Description
----------|-----------|-------------
/intel/yarn/queue/[queue_id]/absolutecapacity | float64 | Absolute capacity percentage this queue can use of entire cluster
/intel/yarn/queue/[queue_id]/absolutemaxcapacity | int | Absolute maximum capacity percentage this queue can use of the entire cluster
/intel/yarn/queue/[queue_id]/absoluteusedcapacity | int | Absolute used capacity percentage this queue is using of the entire cluster
/intel/yarn/queue/[queue_id]/capacity | int | Configured queue capacity in percentage relative to its parent queue
/intel/yarn/queue/[queue_id]/maxactiveapplications | int | The maximum number of active applications this queue can have
/intel/yarn/queue/[queue_id]/maxactiveapplicationsperuser | int | The maximum number of active applications this queue can have
/intel/yarn/queue/[queue_id]/maxapplications | int | The maximum number of active applications this queue can have
/intel/yarn/queue/[queue_id]/maxapplicationsperuser | int | The maximum number of active applications per user this queue can have
/intel/yarn/queue/[queue_id]/maxcapacity | int | Configured maximum queue capacity in percentage relative to its parent queue
/intel/yarn/queue/[queue_id]/numactiveapplications | int | The number of pending applications for this user in this queue
/intel/yarn/queue/[queue_id]/numapplications | int | The number of applications currently in the queue
/intel/yarn/queue/[queue_id]/numcontainers | int | The number of containers being used
/intel/yarn/queue/[queue_id]/numpendingapplications | int | The number of pending applications for this user in this queue
/intel/yarn/queue/[queue_id]/usedcapacity | int | Used queue capacity in percentage
/intel/yarn/queue/[queue_id]/usedresources | int | Used queue resources
/intel/yarn/queue/[queue_id]/userlimit | int | The minimum user limit percent set in the configuration
/intel/yarn/queue/[queue_id]/userlimitfactor | int | The user limit factor set in the configuration
/intel/yarn/queue/[queue_id]/resources_memory | int | The amount of memory used (in MB)
/intel/yarn/queue/[queue_id]/resources_vcores | int | The number of virtual cores
/intel/yarn/schedulerinfo/capacity | int | Configured queue capacity in percentage for root queue
/intel/yarn/schedulerinfo/maxcapacity | int | Configured maximum queue capacity in percentage for root queue
