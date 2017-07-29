package pb 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "pb/service.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/containers": {
      "post": {
        "summary": "Like 'docker run' command",
        "description": "Input/Output is a same protobuf/json object. For input:\n{\n  \"config\":\n    {\n      \"image\": \"nginx\",\n      \"exposed_ports\":\n        {\n          \"value\": \"webui\"\n        }\n    },\n  \"host_config\":\n    {\n      \"port_bindings\":\n        {\n          \"value\":\n            {\n              \"80\":\n                {\n                  \"host_port\": \"80\"\n                }\n            }\n        }\n    },\n  \"network_config\":\n    {\n    },\n  \"container_name\": \"nginx\"\n}\nAnd returning information append this object for output:\n{\n  \"state_code\": 0,  // succeeded, otherwise none zero\n  \"state_message\": \"if failed, provide error information\",\n  \"container_id\": \"regturned from docker engine\"  \n}",
        "operationId": "RunContainer",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerRunData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerRunData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/echo/{value}": {
      "get": {
        "operationId": "Echo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbEchoMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "value",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/provisions": {
      "post": {
        "summary": "Run containers with same user namespace information",
        "description": "Input/Output is a same protobuf/json object. For input:\n{\n  \"name\": \"fighter and target\"\n  \"namespace\": \"default\"\n  \"metadata\":\n    {\n      \"categroy_name\": \"basic-web-security\",\n      \"class_name\": \"http-protocol\"\n      \"field_name\": \"http-method\"\n    },\n  \"provisionings\": [\n    list of DockerRunData type, see previous\n  ]\n}",
        "operationId": "ProvisionContainers",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbProvisioningsData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbProvisioningsData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/pull": {
      "post": {
        "summary": "Like 'docker pull' command",
        "description": "Input/Output is a same protobuf/json object. For input:\n{\n  \"image\": \"tomcat:8\"\n}",
        "operationId": "PullImage",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerPullData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerPullData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/terminations": {
      "post": {
        "summary": "Delete containers with same user namespace information",
        "description": "Input/Output is a same protobuf/json object. For input:\n{\n  \"name\": \"fighter and target\"\n  \"namespace\": \"default\"\n  \"metadata\":\n    {\n      \"categroy_name\": \"basic-web-security\",\n      \"class_name\": \"http-protocol\"\n      \"field_name\": \"http-method\"\n    },\n}\nAnd returning information append this object for output:\n{\n  \"provisionings\": [\n    list of DockerRunData type, see previous\n  ]\n}",
        "operationId": "TerminationContainers",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbProvisioningsData"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbProvisioningsData"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    }
  },
  "definitions": {
    "ProvisioningsDataMetadata": {
      "type": "object",
      "properties": {
        "category_name": {
          "type": "string",
          "format": "string"
        },
        "class_name": {
          "type": "string",
          "format": "string"
        },
        "field_name": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "mobyBindOptions": {
      "type": "object",
      "properties": {
        "propagation": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "BindOptions defines options specific to mounts of type \"bind\".\nto see https://github.com/moby/moby/blob/master/api/types/mount/mount.go"
    },
    "mobyConfig": {
      "type": "object",
      "properties": {
        "args_escaped": {
          "type": "boolean",
          "format": "boolean"
        },
        "attach_stderr": {
          "type": "boolean",
          "format": "boolean"
        },
        "attach_stdin": {
          "type": "boolean",
          "format": "boolean"
        },
        "attach_stdout": {
          "type": "boolean",
          "format": "boolean"
        },
        "cmd": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "domainname": {
          "type": "string",
          "format": "string"
        },
        "entrypoint": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "env": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "exposed_ports": {
          "$ref": "#/definitions/mobyPortSet"
        },
        "healthcheck": {
          "$ref": "#/definitions/mobyHealthConfig"
        },
        "hostname": {
          "type": "string",
          "format": "string"
        },
        "image": {
          "type": "string",
          "format": "string"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "mac_address": {
          "type": "string",
          "format": "string"
        },
        "network_disabled": {
          "type": "boolean",
          "format": "boolean"
        },
        "on_build": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "open_stdin": {
          "type": "boolean",
          "format": "boolean"
        },
        "shell": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "stdin_once": {
          "type": "boolean",
          "format": "boolean"
        },
        "stop_signal": {
          "type": "string",
          "format": "string"
        },
        "stop_timeout": {
          "type": "integer",
          "format": "int32"
        },
        "tty": {
          "type": "boolean",
          "format": "boolean"
        },
        "user": {
          "type": "string",
          "format": "string"
        },
        "volumes": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/mobyVolumeMount"
          }
        },
        "working_dir": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "Config contains the configuration data about a container.\nIt should hold only portable information about the container.\nHere, \"portable\" means \"independent from the host we are running on\".\nNon-portable information *should* appear in HostConfig.\nAll fields added to this struct must be marked 'omitempty' to keep getting\npredictable hashes from the old 'v1Compatibility' configuration.\nto see https://github.com/moby/moby/blob/master/api/types/container/config.go"
    },
    "mobyDeviceMapping": {
      "type": "object",
      "properties": {
        "cgroup_permissions": {
          "type": "string",
          "format": "string"
        },
        "path_in_container": {
          "type": "string",
          "format": "string"
        },
        "path_on_host": {
          "type": "string",
          "format": "string"
        }
      },
      "description": "DeviceMapping represents the device mapping between the host and the container."
    },
    "mobyDriverConfig": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        }
      },
      "description": "Driver represents a volume driver."
    },
    "mobyEndpointIPAMConfig": {
      "type": "object",
      "properties": {
        "ipv4_address": {
          "type": "string",
          "format": "string"
        },
        "ipv6_address": {
          "type": "string",
          "format": "string"
        },
        "link_local_ips": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        }
      },
      "title": "EndpointIPAMConfig represents IPAM configurations for the endpoint"
    },
    "mobyEndpointSettings": {
      "type": "object",
      "properties": {
        "aliases": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "driver_opts": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "endpoint_id": {
          "type": "string",
          "format": "string"
        },
        "gateway": {
          "type": "string",
          "format": "string"
        },
        "global_ipv6_address": {
          "type": "string",
          "format": "string"
        },
        "global_ipv6_prefix_len": {
          "type": "integer",
          "format": "int32"
        },
        "ip_address": {
          "type": "string",
          "format": "string"
        },
        "ip_prefix_len": {
          "type": "integer",
          "format": "int32"
        },
        "ipam_config": {
          "$ref": "#/definitions/mobyEndpointIPAMConfig",
          "title": "Configurations"
        },
        "ipv6_gateway": {
          "type": "string",
          "format": "string"
        },
        "links": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "mac_address": {
          "type": "string",
          "format": "string"
        },
        "network_id": {
          "type": "string",
          "format": "string",
          "title": "Operational data"
        }
      },
      "title": "EndpointSettings stores the network endpoint details"
    },
    "mobyHealthConfig": {
      "type": "object",
      "properties": {
        "interval_seconds": {
          "type": "string",
          "format": "int64",
          "description": "Zero means to inherit. Durations are expressed as integer nanoseconds.\nGolang    time.Duration\nInterval is the time to wait between checks."
        },
        "retries": {
          "type": "integer",
          "format": "int32",
          "description": "Retries is the number of consecutive failures needed to consider a container as unhealthy.\nZero means inherit."
        },
        "start_period": {
          "type": "string",
          "format": "int64",
          "description": "Golang time.Duration\nThe start period for the container to initialize before the retries starts to count down."
        },
        "test": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Test is the test to perform to check that the container is healthy.\nAn empty slice means to inherit the default.\nThe options are:\n{} : inherit healthcheck\n{\"NONE\"} : disable healthcheck\n{\"CMD\", args...} : exec arguments directly\n{\"CMD-SHELL\", command} : run command with system's default shell"
        },
        "timeout_seconds": {
          "type": "string",
          "format": "int64",
          "description": "Golang     time.Duration\nTimeout is the time to wait before considering the check to have hung."
        }
      },
      "title": "HealthConfig holds configuration settings for the HEALTHCHECK feature.\nto see https://github.com/moby/moby/blob/master/api/types/container/config.go"
    },
    "mobyHostConfig": {
      "type": "object",
      "properties": {
        "auto_remove": {
          "type": "boolean",
          "format": "boolean"
        },
        "binds": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Applicable to all platforms"
        },
        "cap_add": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Applicable to UNIX platforms"
        },
        "cap_drop": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "cgroup": {
          "type": "string",
          "format": "string"
        },
        "console_size_height": {
          "type": "integer",
          "format": "int64",
          "title": "Applicable to Windows"
        },
        "console_size_width": {
          "type": "integer",
          "format": "int64"
        },
        "container_id_file": {
          "type": "string",
          "format": "string"
        },
        "dns": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "dns_options": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "dns_search": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "extra_hosts": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "group_add": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "init": {
          "type": "boolean",
          "format": "boolean",
          "title": "Run a custom init inside the container, if null, use the daemon's configured settings"
        },
        "ipc_mode": {
          "type": "string",
          "format": "string"
        },
        "isolation": {
          "type": "string",
          "format": "string"
        },
        "links": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "log_config": {
          "$ref": "#/definitions/mobyLogConfig"
        },
        "mounts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyVolumeMount"
          },
          "title": "Mounts specs used by the container"
        },
        "network_mode": {
          "type": "string",
          "format": "string"
        },
        "oom_score_adj": {
          "type": "integer",
          "format": "int32"
        },
        "pid_mode": {
          "type": "string",
          "format": "string"
        },
        "port_bindings": {
          "$ref": "#/definitions/mobyPortMap"
        },
        "privileged": {
          "type": "boolean",
          "format": "boolean"
        },
        "publish_all_ports": {
          "type": "boolean",
          "format": "boolean"
        },
        "readonly_rootfs": {
          "type": "boolean",
          "format": "boolean"
        },
        "resources": {
          "$ref": "#/definitions/mobyResources",
          "title": "Contains container's resources (cgroups, ulimits)"
        },
        "restart_policy": {
          "$ref": "#/definitions/mobyRestartPolicy"
        },
        "runtime": {
          "type": "string",
          "format": "string"
        },
        "security_opt": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "shm_size": {
          "type": "string",
          "format": "int64"
        },
        "storage_opt": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "sysctls": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "tmpfs": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "userns_mode": {
          "type": "string",
          "format": "string"
        },
        "uts_mode": {
          "type": "string",
          "format": "string"
        },
        "volume_driver": {
          "type": "string",
          "format": "string"
        },
        "volumes_from": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        }
      },
      "title": "HostConfig the non-portable Config structure of a container.\nHere, \"non-portable\" means \"dependent of the host we are running on\".\nPortable information *should* appear in Config.\nto see https://github.com/moby/moby/blob/master/api/types/container/host_config.go"
    },
    "mobyLogConfig": {
      "type": "object",
      "properties": {
        "config": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "config_type": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "LogConfig represents the logging configuration of the container.\nto see https://github.com/moby/moby/blob/master/api/types/container/host_config.go"
    },
    "mobyNetworkingConfig": {
      "type": "object",
      "properties": {
        "endpoints_config": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/mobyEndpointSettings"
          }
        }
      },
      "title": "NetworkingConfig represents the container's networking configuration for each of its interfaces\nCarries the networking configs specified in the 'docker run' and 'docker network connect' commands\nto see https://github.com/moby/moby/blob/master/api/types/network/network.go"
    },
    "mobyPortBinding": {
      "type": "object",
      "properties": {
        "host_ip": {
          "type": "string",
          "format": "string",
          "title": "HostIP is the host IP Address"
        },
        "host_port": {
          "type": "string",
          "format": "string"
        }
      },
      "title": "PortBinding represents a binding between a Host IP address and a Host Port\nto see https://github.com/docker/go-connections/blob/master/nat/nat.go"
    },
    "mobyPortMap": {
      "type": "object",
      "properties": {
        "value": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/mobyPortBinding"
          }
        }
      },
      "title": "PortMap is a collection of PortBinding indexed by Port\nto see https://github.com/docker/go-connections/blob/master/nat/nat.go"
    },
    "mobyPortSet": {
      "type": "object",
      "properties": {
        "value": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        }
      },
      "title": "PortSet is a collection of structs indexed by Port\nto see https://github.com/docker/go-connections/blob/master/nat/nat.go"
    },
    "mobyResources": {
      "type": "object",
      "properties": {
        "blkio_device_read_bps": {
          "$ref": "#/definitions/mobyThrottleDevice"
        },
        "blkio_device_read_iops": {
          "$ref": "#/definitions/mobyThrottleDevice"
        },
        "blkio_device_write_bps": {
          "$ref": "#/definitions/mobyThrottleDevice"
        },
        "blkio_device_write_iops": {
          "$ref": "#/definitions/mobyThrottleDevice"
        },
        "blkio_weight": {
          "type": "integer",
          "format": "int32"
        },
        "blkio_weight_device": {
          "$ref": "#/definitions/mobyWeightDevice"
        },
        "cgroup_parent": {
          "type": "string",
          "format": "string",
          "title": "Applicable to UNIX platforms"
        },
        "cpu_count": {
          "type": "string",
          "format": "int64",
          "title": "Applicable to Windows"
        },
        "cpu_percent": {
          "type": "string",
          "format": "int64"
        },
        "cpu_period": {
          "type": "string",
          "format": "int64"
        },
        "cpu_quota": {
          "type": "string",
          "format": "int64"
        },
        "cpu_realtime_period": {
          "type": "string",
          "format": "int64"
        },
        "cpu_realtime_runtime": {
          "type": "string",
          "format": "int64"
        },
        "cpu_shares": {
          "type": "string",
          "format": "int64",
          "title": "Applicable to all platforms"
        },
        "cpuset_cpus": {
          "type": "string",
          "format": "string"
        },
        "cpuset_mems": {
          "type": "string",
          "format": "string"
        },
        "device_cgroup_rules": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          }
        },
        "devices": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyDeviceMapping"
          }
        },
        "disk_quota": {
          "type": "string",
          "format": "int64"
        },
        "io_maximum_bandwidth": {
          "type": "string",
          "format": "uint64"
        },
        "io_maximum_iops": {
          "type": "string",
          "format": "uint64"
        },
        "kernel_memory": {
          "type": "string",
          "format": "int64"
        },
        "memory": {
          "type": "string",
          "format": "int64"
        },
        "memory_reservation": {
          "type": "string",
          "format": "int64"
        },
        "memory_swap": {
          "type": "string",
          "format": "int64"
        },
        "memory_swappiness": {
          "type": "string",
          "format": "int64"
        },
        "nano_cpus": {
          "type": "string",
          "format": "int64"
        },
        "oom_kill_disable": {
          "type": "string",
          "format": "int64"
        },
        "pids_limit": {
          "type": "string",
          "format": "int64"
        },
        "ulimits": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyUlimit"
          }
        }
      },
      "title": "Resources contains container's resources (cgroups config, ulimits...)\nto see https://github.com/moby/moby/blob/master/api/types/container/host_config.go"
    },
    "mobyRestartPolicy": {
      "type": "object",
      "properties": {
        "maximum_retry_count": {
          "type": "integer",
          "format": "int32"
        },
        "name": {
          "type": "string",
          "format": "string"
        }
      },
      "description": "RestartPolicy represents the restart policies of the container."
    },
    "mobyThrottleDevice": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string",
          "format": "string"
        },
        "rate": {
          "type": "string",
          "format": "uint64"
        }
      },
      "title": "ThrottleDevice is a structure that holds device:rate_per_second pair\nto see http://github.com/moby/moby/blob/master/api/types/blkiodev/blkio.go"
    },
    "mobyTmpfsOptions": {
      "type": "object",
      "properties": {
        "mode": {
          "type": "integer",
          "format": "int64",
          "title": "Mode of the tmpfs upon creation"
        },
        "size_bytes": {
          "type": "string",
          "format": "int64",
          "description": "Size sets the size of the tmpfs, in bytes.\n\nThis will be converted to an operating system specific value\ndepending on the host. For example, on linux, it will be converted to\nuse a 'k', 'm' or 'g' syntax. BSD, though not widely supported with\ndocker, uses a straight byte value.\n\nPercentages are not supported."
        }
      },
      "description": "TmpfsOptions defines options specific to mounts of type \"tmpfs\"."
    },
    "mobyUlimit": {
      "type": "object",
      "properties": {
        "hard": {
          "type": "string",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "soft": {
          "type": "string",
          "format": "int64"
        }
      },
      "title": "Ulimit is a human friendly version of Rlimit.\nto see https://github.com/docker/go-units/blob/master/ulimit.go"
    },
    "mobyVolumeMount": {
      "type": "object",
      "properties": {
        "bind_options": {
          "$ref": "#/definitions/mobyBindOptions"
        },
        "consistency": {
          "type": "string",
          "format": "string"
        },
        "mount_type": {
          "type": "string",
          "format": "string"
        },
        "read_only": {
          "type": "string",
          "format": "string"
        },
        "source": {
          "type": "string",
          "format": "string",
          "title": "Source specifies the name of the mount. Depending on mount type, this\nmay be a volume name or a host path, or even ignored.\nSource is not supported for tmpfs (must be an empty value)"
        },
        "target": {
          "type": "string",
          "format": "string"
        },
        "tmpfs_options": {
          "$ref": "#/definitions/mobyTmpfsOptions"
        },
        "volume_options": {
          "$ref": "#/definitions/mobyVolumeOptions"
        }
      },
      "title": "Mount represents a mount (volume).\nto see https://github.com/moby/moby/blob/master/api/types/mount/mount.go"
    },
    "mobyVolumeOptions": {
      "type": "object",
      "properties": {
        "driver": {
          "$ref": "#/definitions/mobyDriverConfig"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          }
        },
        "no_copy": {
          "type": "boolean",
          "format": "boolean"
        }
      },
      "description": "VolumeOptions represents the options for a mount of type volume."
    },
    "mobyWeightDevice": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string",
          "format": "string"
        },
        "weight": {
          "type": "integer",
          "format": "int32"
        }
      },
      "title": "WeightDevice is a structure that holds device:weight pair\nto see http://github.com/moby/moby/blob/master/api/types/blkiodev/blkio.go"
    },
    "pbDockerPullData": {
      "type": "object",
      "properties": {
        "image": {
          "type": "string",
          "format": "string"
        },
        "image_id": {
          "type": "string",
          "format": "string"
        },
        "progress_report": {
          "type": "string",
          "format": "string"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbDockerRunData": {
      "type": "object",
      "properties": {
        "config": {
          "$ref": "#/definitions/mobyConfig"
        },
        "container_id": {
          "type": "string",
          "format": "string"
        },
        "container_name": {
          "type": "string",
          "format": "string"
        },
        "host_config": {
          "$ref": "#/definitions/mobyHostConfig"
        },
        "network_config": {
          "$ref": "#/definitions/mobyNetworkingConfig"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbEchoMessage": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "pbProvisioningsData": {
      "type": "object",
      "properties": {
        "metadata": {
          "$ref": "#/definitions/ProvisioningsDataMetadata"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "namespace": {
          "type": "string",
          "format": "string"
        },
        "provisionings": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/pbDockerRunData"
          }
        }
      }
    }
  }
}
`
)
