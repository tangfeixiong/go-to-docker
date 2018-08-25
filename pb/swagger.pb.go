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
    "/v1/docker-network-create": {
      "post": {
        "summary": "Create Docker network",
        "operationId": "CreateNetwork",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkCreateReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkCreateReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-pull": {
      "post": {
        "summary": "Pull Docker image",
        "description": "Like command of 'docker pull', Input/Output is a same Protobuf/JSON object. For input example:\n{ \"image_ref\": \"docker.io/nginx\" }\nFor output example:\t\t\n{ \"image_ref\": \"docker.io/nginx\", \"state_code\": 0, \"state_message\": \"SUCCEEDED\" }",
        "operationId": "PullImage",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerImagePullReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerImagePullReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-run": {
      "post": {
        "summary": "Run Docker container",
        "description": "For output, plus result fileds:\n{ ..., \"state_code\": 0, \"state_message\": \"RUNNING\" }",
        "operationId": "RunContainer",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerContainerRunReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerContainerRunReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    }
  },
  "definitions": {
    "PortMapPortBindingSlice": {
      "type": "object",
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/natPortBinding"
          }
        }
      }
    },
    "blkiodevThrottleDevice": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string",
          "format": "string",
          "title": "Path string"
        },
        "rate": {
          "type": "string",
          "format": "uint64",
          "title": "Rate uint64"
        }
      },
      "title": "ThrottleDevice is a structure that holds device:rate_per_second pair\ntype ThrottleDevice struct"
    },
    "blkiodevWeightDevice": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string",
          "format": "string",
          "title": "Path string"
        },
        "weight": {
          "type": "integer",
          "format": "int32",
          "title": "Weight uint16"
        }
      },
      "title": "WeightDevice is a structure that holds device:weight pair\ntype WeightDevice struct"
    },
    "containerConfig": {
      "type": "object",
      "properties": {
        "args_escaped": {
          "type": "boolean",
          "format": "boolean",
          "title": "ArgsEscaped bool 0x60json\",omitempty\"0x60 // True if command is already escaped (Windows specific)"
        },
        "attach_stderr": {
          "type": "boolean",
          "format": "boolean",
          "title": "AttachStderr bool // Attach the standard error"
        },
        "attach_stdin": {
          "type": "boolean",
          "format": "boolean",
          "title": "AttachStdin bool // Attach the standard input, makes possible user interaction"
        },
        "attach_stdout": {
          "type": "boolean",
          "format": "boolean",
          "title": "AttachStdout bool // Attach the standard output"
        },
        "cmd": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Cmd strslice.StrSlice // Command to run when starting the container\nStrSlice represents a string or an array of strings. We need to override the json decoder to accept both options. // type StrSlice []string"
        },
        "domainname": {
          "type": "string",
          "format": "string",
          "title": "Domainname string // Domainname"
        },
        "entrypoint": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Entrypoint strslice.StrSlice // Entrypoint to run when starting the container"
        },
        "env": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Env []string  // List of environment variable to set in the container"
        },
        "exposed_ports": {
          "$ref": "#/definitions/natPortSet",
          "title": "ExposedPorts nat.PortSet 0x60json\",omitempty\"0x60 // List of exposed ports"
        },
        "healthcheck": {
          "$ref": "#/definitions/containerHealthConfig",
          "title": "Healthcheck *HealthConfig 0x60json:\",omitempty\"0x60 // Healthcheck describes how to check the container is healthy"
        },
        "hostname": {
          "type": "string",
          "format": "string",
          "title": "Hostname string // Hostname"
        },
        "image": {
          "type": "string",
          "format": "string",
          "title": "Image string // Name of the image as it was passed by the operator (e.g. could be symbolic)"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "Labels map[string]string // List of labels set to this container"
        },
        "mac_address": {
          "type": "string",
          "format": "string",
          "title": "MacAddress string 0x60json:\",omitempty\"0x60 // Mac Address of the container"
        },
        "network_disabled": {
          "type": "boolean",
          "format": "boolean",
          "title": "NetworkDisabled bool 0x60json:\",omitempty\"0x60 // Is network disabled"
        },
        "on_build": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "OnBuild []string // ONBUILD metadata that were defined on the image Dockerfile"
        },
        "open_stdin": {
          "type": "boolean",
          "format": "boolean",
          "title": "OpenStdin // Open stdin"
        },
        "shell": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Shell strslice.StrSlice 0x60json:\",omitempty\"0x60 // Shell for shell-form of RUN, CMD, ENTRYPOINT"
        },
        "stdin_once": {
          "type": "boolean",
          "format": "boolean",
          "description": "StdinOnce bool // If true, close stdin after the 1 attached client disconnects."
        },
        "stop_signal": {
          "type": "string",
          "format": "string",
          "title": "StopSignal string 0x60json:\",omitempty\"0x60 // Signal to stop a container"
        },
        "stop_timeout": {
          "type": "integer",
          "format": "int32",
          "title": "StopTimeout 0x60json:\",omitempty\"0x60 // Timeout (in seconds) to stop a container"
        },
        "tty": {
          "type": "boolean",
          "format": "boolean",
          "description": "Tty bool // Attach standard streams to a tty, including stdin if it is not closed."
        },
        "user": {
          "type": "string",
          "format": "string",
          "title": "User string // User that will run the command(s) inside the container, also support user:group"
        },
        "volumes": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/containerConfigVoidStruct"
          },
          "title": "Volumes map[string]struct{} // List of volumes (mounts) used for the container"
        },
        "working_dir": {
          "type": "string",
          "format": "string",
          "title": "WorkingDir string // Current directory (PWD) in the command will be launched"
        }
      },
      "title": "Config contains the configuration data about a container.\nIt should hold only portable information about the container.\nHere, \"portable\" means \"independent from the host we are running on\".\nNon-portable information *should* appear in HostConfig.\nAll fields added to this struct must be marked 'omitempty' to keep getting\npredictable hashes from the old 'v1Compatibility' configuration.\ntype Config struct"
    },
    "containerConfigVoidStruct": {
      "type": "object"
    },
    "containerContainerCreateCreatedBody": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string",
          "title": "The ID of the created container. Required: true\nID string 0x60json:\"Id\"0x60"
        },
        "warnings": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Warnings encountered when creating the container. Required: true\nWarnings []string 0x60json:\"Warnings\"0x60"
        }
      },
      "title": "ContainerCreateCreatedBody OK response to ContainerCreate operation\ntype ContainerCreateCreatedBody struct"
    },
    "containerDeviceMapping": {
      "type": "object",
      "properties": {
        "cgroup_permissions": {
          "type": "string",
          "format": "string",
          "title": "CgroupPermissions string"
        },
        "path_in_container": {
          "type": "string",
          "format": "string",
          "title": "PathInContainer string"
        },
        "path_on_host": {
          "type": "string",
          "format": "string",
          "title": "PathOnHost string"
        }
      },
      "title": "DeviceMapping represents the device mapping between the host and the container.\ntype DeviceMapping struct"
    },
    "containerHealthConfig": {
      "type": "object",
      "properties": {
        "interval_seconds": {
          "type": "string",
          "format": "int64",
          "description": "Zero means to inherit. Durations are expressed as integer nanoseconds.\nInterval time.Duration 0x60json:\",omitempty\"0x60 // Interval is the time to wait between checks."
        },
        "retries": {
          "type": "integer",
          "format": "int32",
          "description": "Retries int 0x60json:\",omitempty\"0x60// Retries is the number of consecutive failures needed to consider a container as unhealthy. Zero means inherit."
        },
        "start_period": {
          "type": "string",
          "format": "int64",
          "description": "StartPeriod time.Duration 0x60json:\",omitempty\"0x60 // The start period for the container to initialize before the retries starts to count down."
        },
        "test": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Test []string 0x60json:\",omitempty\"0x60\nTest is the test to perform to check that the container is healthy.\nAn empty slice means to inherit the default.\nThe options are:\n{} : inherit healthcheck\n{\"NONE\"} : disable healthcheck\n{\"CMD\", args...} : exec arguments directly\n{\"CMD-SHELL\", command} : run command with system's default shell"
        },
        "timeout_seconds": {
          "type": "string",
          "format": "int64",
          "description": "Timeout time.Duration 0x60json:\",omitempty\"0x60 // Timeout is the time to wait before considering the check to have hung."
        }
      },
      "title": "HealthConfig holds configuration settings for the HEALTHCHECK feature.\ntype HealthConfig struct"
    },
    "containerHostConfig": {
      "type": "object",
      "properties": {
        "auto_remove": {
          "type": "boolean",
          "format": "boolean",
          "title": "AutoRemove bool // Automatically remove container when it exits"
        },
        "binds": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "description": "Binds []string // List of volume bindings for this container",
          "title": "Applicable to all platforms"
        },
        "cap_add": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "description": "CapAdd strslice.StrSlice // List of kernel capabilities to add to the container",
          "title": "Applicable to UNIX platforms"
        },
        "cap_drop": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "CapDrop strslice.StrSlice // List of kernel capabilities to remove from the container"
        },
        "cgroup": {
          "type": "string",
          "format": "string",
          "title": "Cgroup CgroupSpec // Cgroup to use for the container, \"container:\u003cid\u003e\"\nCgroupSpec represents the cgroup to use for the container. // type CgroupSpec string"
        },
        "console_size_height": {
          "type": "integer",
          "format": "int64",
          "description": "ConsoleSize [2]uint // Initial console size (height,width)",
          "title": "Applicable to Windows"
        },
        "console_size_width": {
          "type": "integer",
          "format": "int64"
        },
        "container_id_file": {
          "type": "string",
          "format": "string",
          "title": "ContainerIDFile // File (path) where the containerId is written"
        },
        "dns": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "DNS []string 0x60json:\"Dns\"0x60 // List of DNS server to lookup"
        },
        "dns_options": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "DNSOptions []string 0x60json:\"DnsOptions\"0x60 // List of DNSOption to look for"
        },
        "dns_search": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "DNSSearch []string 0x60json:\"DnsSearch\"0x60 // List of DNSSearch to look for"
        },
        "extra_hosts": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "ExtraHosts []string // List of extra hosts"
        },
        "group_add": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "GroupAdd []string // List of additional groups that the container process will run as"
        },
        "init": {
          "type": "boolean",
          "format": "boolean",
          "title": "Run a custom init inside the container, if null, use the daemon's configured settings\nInit *bool 0x60json:\",omitempty\"0x60"
        },
        "ipc_mode": {
          "type": "string",
          "format": "string",
          "title": "IpcMode IpcMode // IPC namespace to use for the container, \"\", \"host\", \"container\"\nIpcMode represents the container ipc stack. // type IpcMode string"
        },
        "isolation": {
          "type": "string",
          "format": "string",
          "title": "Isolation Isolation // Isolation technology of the container (e.g. default, hyperv)\nIsolation represents the isolation technology of a container. The supported values are platform specific // type Isolation string"
        },
        "links": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Links []string // List of links (in the name:alias form)"
        },
        "log_config": {
          "$ref": "#/definitions/containerLogConfig",
          "title": "LogConfig LogConfig // Configuration of the logs for this container"
        },
        "masked_paths": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "MaskedPaths is the list of paths to be masked inside the container (this overrides the default set of paths)\nMaskedPaths []string"
        },
        "mounts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mountMount"
          },
          "title": "Mounts specs used by the container\nMounts []mount.Mount 0x60json:\",omitempty\"0x60"
        },
        "network_mode": {
          "type": "string",
          "format": "string",
          "title": "NetworkMode NetworkMode // Network mode to use for the container, \"none\", \"default\", \"container:\u003cid\u003e\"\nNetworkMode represents the container network stack. // type NetworkMode string"
        },
        "oom_score_adj": {
          "type": "integer",
          "format": "int32",
          "title": "OomScoreAdj int // Container preference for OOM-killing"
        },
        "pid_mode": {
          "type": "string",
          "format": "string",
          "title": "PidMode PidMode // PID namespace to use for the container\nPidMode represents the pid namespace of the container. // type PidMode string"
        },
        "port_bindings": {
          "$ref": "#/definitions/natPortMap",
          "title": "PortBindings nat.PortMap // Port mapping between the exposed port (container) and the host"
        },
        "privileged": {
          "type": "boolean",
          "format": "boolean",
          "title": "Privieged bool // Is the container in privileged mode"
        },
        "publish_all_ports": {
          "type": "boolean",
          "format": "boolean",
          "title": "PublishAllPorts bool // Should docker publish all exposed port for the container"
        },
        "readonly_paths": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "ReadonlyPaths is the list of paths to be set as read-only inside the container (this overrides the default set of paths)\nReadonlyPaths []string"
        },
        "readonly_rootfs": {
          "type": "boolean",
          "format": "boolean",
          "title": "ReadonlyRootfs bool // Is the container root filesystem in read-only"
        },
        "resources": {
          "$ref": "#/definitions/containerResources",
          "title": "Contains container's resources (cgroups, ulimits)\nResources"
        },
        "restart_policy": {
          "$ref": "#/definitions/containerRestartPolicy",
          "title": "RestartPolicy RestartPolicy // Restart policy to be used for the container"
        },
        "runtime": {
          "type": "string",
          "format": "string",
          "title": "Runtime string 0x60json:\",omitempty\"0x60\nRuntime to use with this container"
        },
        "security_opt": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "description": "SecurityOpt []string // List of string values to customize labels for MLS systems, such as SELinux."
        },
        "shm_size": {
          "type": "string",
          "format": "int64",
          "title": "ShmSize int64\nTotal shm memory usage"
        },
        "storage_opt": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "Storage driver options per container.\nStorageOpt map[string]string 0x60json:\",omitempty\"0x60"
        },
        "sysctls": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "Sysctls map[string]string 0x60json:\",omitempty\"0x60\nList of Namespaced sysctls used for the container"
        },
        "tmpfs": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "List of tmpfs (mounts) used for the container\nTmpfs map[string]string 0x60json:\",omitempty\"0x60"
        },
        "userns_mode": {
          "type": "string",
          "format": "string",
          "title": "UsernsMode UsernsMode // UsernsMode represents userns mode in the container. // type UsernsMode string\nThe user namespace to use for the container"
        },
        "uts_mode": {
          "type": "string",
          "format": "string",
          "title": "UTSMode UTSMode // UTSMode represents the UTS namespace of the container. // type UTSMode string\nUTS namespace to use for the container"
        },
        "volume_driver": {
          "type": "string",
          "format": "string",
          "title": "VolumeDriver string // Name of the volume driver used to mount volumes"
        },
        "volumes_from": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "VolumesFrom []string // List of volumes to take from other container"
        }
      },
      "title": "HostConfig the non-portable Config structure of a container.\nHere, \"non-portable\" means \"dependent of the host we are running on\".\nPortable information *should* appear in Config.\ntype HostConfig struct"
    },
    "containerLogConfig": {
      "type": "object",
      "properties": {
        "config": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "Config map[string]string"
        },
        "type": {
          "type": "string",
          "format": "string",
          "title": "Type string // \"\", \"blocking\", \"non-blocking\""
        }
      },
      "title": "LogConfig represents the logging configuration of the container.\ntype LogConfig struct"
    },
    "containerResources": {
      "type": "object",
      "properties": {
        "blkio_device_read_bps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/blkiodevThrottleDevice"
          },
          "title": "BlkioDeviceReadBps []*blkiodev.ThrottleDevice"
        },
        "blkio_device_read_iops": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/blkiodevThrottleDevice"
          },
          "title": "BlkioDeviceReadIOps []*blkiodev.ThrottleDevice"
        },
        "blkio_device_write_bps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/blkiodevThrottleDevice"
          },
          "title": "BlkioDeviceWriteBps []*blkiodev.ThrottleDevice"
        },
        "blkio_device_write_iops": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/blkiodevThrottleDevice"
          },
          "title": "BlkioDeviceWriteIOps []*blkiodev.ThrottleDevice"
        },
        "blkio_weight": {
          "type": "integer",
          "format": "int32",
          "title": "BlkioWeight uint16 // Block IO weight (relative weight vs. other containers)"
        },
        "blkio_weight_device": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/blkiodevWeightDevice"
          },
          "title": "BlkioWeightDevice []*blkiodev.WeightDevice"
        },
        "cgroup_parent": {
          "type": "string",
          "format": "string",
          "description": "CgroupParent string // Parent cgroup.",
          "title": "Applicable to UNIX platforms"
        },
        "cpu_count": {
          "type": "string",
          "format": "int64",
          "description": "CPUCount int64 // CPU count",
          "title": "Applicable to Windows"
        },
        "cpu_percent": {
          "type": "string",
          "format": "int64",
          "title": "CPUPercent int64 // CPU percent"
        },
        "cpu_period": {
          "type": "string",
          "format": "int64",
          "title": "CPUPeriod int64 0x60json:\"CpuPeriod\"0x60 // CPU CFS (Completely Fair Scheduler) period"
        },
        "cpu_quota": {
          "type": "string",
          "format": "int64",
          "title": "CPUQuota int64 0x60json:\"CpuQuota\"0x60 // CPU CFS (Completely Fair Scheduler) quota"
        },
        "cpu_realtime_period": {
          "type": "string",
          "format": "int64",
          "title": "CPURealtimePeriod int64 0x60json:\"CpuRealtimePeriod\"0x60 // CPU real-time period"
        },
        "cpu_realtime_runtime": {
          "type": "string",
          "format": "int64",
          "title": "CPURealtimeRuntime int64 0x60json:\"CpuRealtimeRuntime\"0x60 // CPU real-time runtime"
        },
        "cpu_shares": {
          "type": "string",
          "format": "int64",
          "description": "CPUShares int64 0x60json:\"CpuShares\"0x60 // CPU shares (relative weight vs. other containers)",
          "title": "Applicable to all platforms"
        },
        "cpuset_cpus": {
          "type": "string",
          "format": "string",
          "title": "CpusetCpus string // CpusetCpus 0-2, 0,1"
        },
        "cpuset_mems": {
          "type": "string",
          "format": "string",
          "title": "CpusetMems string // CpusetMems 0-2, 0,1"
        },
        "device_cgroup_rules": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "DeviceCgroupRules []string // List of rule to be added to the device cgroup"
        },
        "devices": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/containerDeviceMapping"
          },
          "title": "Devices []DeviceMapping // List of devices to map inside the container"
        },
        "disk_quota": {
          "type": "string",
          "format": "int64",
          "title": "DiskQuota int64 // Disk limit (in bytes)"
        },
        "io_maximum_bandwidth": {
          "type": "string",
          "format": "uint64",
          "title": "IOMaximumBandwidth // Maximum IO in bytes per second for the container system drive"
        },
        "io_maximum_iops": {
          "type": "string",
          "format": "uint64",
          "title": "IOMaximumIOps uint64 // Maximum IOps for the container system drive"
        },
        "kernel_memory": {
          "type": "string",
          "format": "int64",
          "title": "KernelMemory int64 // Kernel memory limit (in bytes)"
        },
        "memory": {
          "type": "string",
          "format": "int64",
          "title": "Memory int64 // Memory limit (in bytes)"
        },
        "memory_reservation": {
          "type": "string",
          "format": "int64",
          "title": "MemoryReservation int64 // Memory soft limit (in bytes)"
        },
        "memory_swap": {
          "type": "string",
          "format": "int64",
          "title": "MemorySwap int64 // Total memory usage (memory + swap); set `-1` to enable unlimited swap"
        },
        "memory_swappiness": {
          "type": "string",
          "format": "int64",
          "title": "MemorySwappiness *int64 // Tuning container memory swappiness behaviour"
        },
        "nano_cpus": {
          "type": "string",
          "format": "int64",
          "description": "NanoCPUs int64 0x60json:\"NonoCpus\"0x60 // CPU quota in units of 10\u003csup\u003e-9\u003c/sup\u003e CPUs."
        },
        "oom_kill_disable": {
          "type": "boolean",
          "format": "boolean",
          "title": "OomKillDisable *bool // Whether to disable OOM Killer or not"
        },
        "pids_limit": {
          "type": "string",
          "format": "int64",
          "title": "PidsLimit int64 // Setting pids limit for a container"
        },
        "ulimits": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/unitsUlimit"
          },
          "title": "Ulimits []*units.Ulimit // List of ulimits to be set in the container"
        }
      },
      "title": "Resources contains container's resources (cgroups config, ulimits...)\nto see https://github.com/moby/moby/blob/master/api/types/container/host_config.go"
    },
    "containerRestartPolicy": {
      "type": "object",
      "properties": {
        "maximum_retry_count": {
          "type": "integer",
          "format": "int32",
          "title": "MaximumRetryCount int"
        },
        "name": {
          "type": "string",
          "format": "string",
          "title": "Name string"
        }
      },
      "title": "RestartPolicy represents the restart policies of the container.\ntype RestartPolicy struct"
    },
    "mobyImagePullOptions": {
      "type": "object",
      "properties": {
        "all": {
          "type": "boolean",
          "format": "boolean",
          "title": "All bool"
        },
        "platform": {
          "type": "string",
          "format": "string",
          "title": "Platform string"
        },
        "privilege_func": {
          "type": "string",
          "format": "string",
          "title": "PrivilegeFunc RequestPrivilegeFunc"
        },
        "registry_auth": {
          "type": "string",
          "format": "string",
          "title": "RegistryAuth string // RegistryAuth is the base64 encoded credentials for the registry"
        }
      },
      "title": "ImagePullOptions holds information to pull images.\ntype ImagePullOptions struct"
    },
    "mobyNetworkCreate": {
      "type": "object",
      "properties": {
        "attachable": {
          "type": "boolean",
          "format": "boolean",
          "title": "Attachable bool"
        },
        "check_duplicate": {
          "type": "boolean",
          "format": "boolean",
          "description": "CheckDuplicate bool // Check for networks with duplicate names. Network is primarily keyed based on a random ID and not on the name. Network name is strictly a user-friendly alias to the network which is uniquely identified using IO. And there is no gauranteed way to check for duplicates. Option CheckDuplicate is there to provide a best effort checking of any networks which had the same name but it is not guaranteed to catch all name collisions."
        },
        "config_from": {
          "$ref": "#/definitions/networkConfigReference",
          "title": "ConfigFrom *network.ConfigReference"
        },
        "config_only": {
          "type": "boolean",
          "format": "boolean",
          "title": "ConfigOnly bool"
        },
        "driver": {
          "type": "string",
          "format": "string",
          "title": "Driver string"
        },
        "enable_ipv6": {
          "type": "boolean",
          "format": "boolean",
          "title": "EnableIPv6 bool"
        },
        "ingress": {
          "type": "boolean",
          "format": "boolean",
          "title": "Ingress bool"
        },
        "internal": {
          "type": "boolean",
          "format": "boolean",
          "title": "Internal bool"
        },
        "ipam": {
          "$ref": "#/definitions/networkIPAM",
          "title": "IPAM *network.IPAM"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "Labels map[string]string"
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "Options map[string]string"
        },
        "scope": {
          "type": "string",
          "format": "string",
          "title": "Scope string"
        }
      },
      "title": "NetworkCreate is the expected body of the \"create network\" http request message\ntype NetworkCreate struct"
    },
    "mobyNetworkCreateResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "string",
          "title": "ID string 0x60json:\"Id\"0x60"
        },
        "warning": {
          "type": "string",
          "format": "string",
          "title": "Warning string"
        }
      },
      "title": "NetworkCreateResponse is the response message sent by the server for network create call\ntype NetworkCreateResponse struct"
    },
    "mountBindOptions": {
      "type": "object",
      "properties": {
        "propagation": {
          "type": "string",
          "format": "string",
          "title": "Propagation Propagation 0x60json\",omitempty\"0x60\nPropagation represents the propagation of a mount. // type Propagation string"
        }
      },
      "title": "BindOptions defines options specific to mounts of type \"bind\".\ntype BindOptions struct"
    },
    "mountDriver": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string",
          "title": "Name string 0x60json:\",omitempty\"0x60"
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "Options map[string]string 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "Driver represents a volume driver.\ntype Driver struct"
    },
    "mountMount": {
      "type": "object",
      "properties": {
        "bind_options": {
          "$ref": "#/definitions/mountBindOptions",
          "title": "BindOptions *BindOptions 0x60json:\",omitempty\"0x60"
        },
        "consistency": {
          "type": "string",
          "format": "string",
          "title": "Consistency consistency 0x60json:\",omitempty\"0x60\nConsistency represents the consistency requirements of a mount. // type Consistency string"
        },
        "read_only": {
          "type": "string",
          "format": "string",
          "title": "ReadOnly bool 0x60json:\",omitempty\"0x60"
        },
        "source": {
          "type": "string",
          "format": "string",
          "title": "Source specifies the name of the mount. Depending on mount type, this may be a volume name or a host path, or even ignored.\nSource is not supported for tmpfs (must be an empty value)\nSource string 0x60json:\",omitempty\"0x60"
        },
        "target": {
          "type": "string",
          "format": "string",
          "title": "Target string 0x60json:\",omitempty\"0x60"
        },
        "tmpfs_options": {
          "$ref": "#/definitions/mountTmpfsOptions",
          "title": "TmpfsOptions *TmpfsOptions 0x60json:\",omitempty\"0x60"
        },
        "type": {
          "type": "string",
          "format": "string",
          "title": "Type Type 0x60json:\",omitempty\"0x60"
        },
        "volume_options": {
          "$ref": "#/definitions/mountVolumeOptions",
          "title": "VolumeOptions *VolumeOptions 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "Mount represents a mount (volume).\ntype Mount struct"
    },
    "mountTmpfsOptions": {
      "type": "object",
      "properties": {
        "mode": {
          "type": "integer",
          "format": "int64",
          "title": "Mode of the tmpfs upon creation\nMode os.FileMode 0x60json:\",omitempty\"0x60"
        },
        "size_bytes": {
          "type": "string",
          "format": "int64",
          "description": "Size sets the size of the tmpfs, in bytes.\n\nThis will be converted to an operating system specific value\ndepending on the host. For example, on linux, it will be converted to\nuse a 'k', 'm' or 'g' syntax. BSD, though not widely supported with\ndocker, uses a straight byte value.\n\nPercentages are not supported.\nSizeBytes int84 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "TmpfsOptions defines options specific to mounts of type \"tmpfs\".\ntype TmpfsOptions struct"
    },
    "mountVolumeOptions": {
      "type": "object",
      "properties": {
        "driver_config": {
          "$ref": "#/definitions/mountDriver",
          "title": "DriverConfig *Driver 0x60json:\",omitempty\"0x60"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "Labels map[string]string 0x60json:\",omitempty\"0x60"
        },
        "no_copy": {
          "type": "boolean",
          "format": "boolean",
          "title": "NoCopy bool 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "VolumeOptions represents the options for a mount of type volume.\ntype VolumeOptions struct"
    },
    "natPortBinding": {
      "type": "object",
      "properties": {
        "host_ip": {
          "type": "string",
          "format": "string",
          "title": "HostIP is the host IP Address\nHostIP string 0x60json:\"HostIp\"0x60"
        },
        "host_port": {
          "type": "string",
          "format": "string",
          "title": "HostPort is the host port number\nHostPort string"
        }
      },
      "title": "PortBinding represents a binding between a Host IP address and a Host Port\ntype PortBinding struct"
    },
    "natPortMap": {
      "type": "object",
      "properties": {
        "value": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/PortMapPortBindingSlice"
          }
        }
      },
      "title": "PortMap is a collection of PortBinding indexed by Port\nPort is a string containing port number and protocol in the format \"80/tcp\" // type Port string\ntype PortMap map[Port][]PortBinding"
    },
    "natPortSet": {
      "type": "object",
      "properties": {
        "value": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/natPortSetVoidStruct"
          }
        }
      },
      "title": "PortSet is a collection of structs indexed by Port\ntype PortSet map[Port]struct{}"
    },
    "natPortSetVoidStruct": {
      "type": "object"
    },
    "networkConfigReference": {
      "type": "object",
      "properties": {
        "network": {
          "type": "string",
          "format": "string",
          "title": "Network string"
        }
      },
      "title": "ConfigReference specifies the source which provides a network's cconfiguration\ntype ConfigReference struct"
    },
    "networkEndpointIPAMConfig": {
      "type": "object",
      "properties": {
        "ipv4_address": {
          "type": "string",
          "format": "string",
          "title": "IPv4Address string 0x60json:\",omitempty\"0x60"
        },
        "ipv6_address": {
          "type": "string",
          "format": "string",
          "title": "IPv6Address string 0x60json:\",omitempty\"0x60"
        },
        "link_local_ips": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "LinkLocalIPs []string 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "EndpointIPAMConfig represents IPAM configurations for the endpoint\ntype EndpointIPAMConfig struct"
    },
    "networkEndpointSettings": {
      "type": "object",
      "properties": {
        "aliases": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Aliases []string"
        },
        "driver_opts": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "DriverOpts map[string]string"
        },
        "endpoint_id": {
          "type": "string",
          "format": "string",
          "title": "EndpointID string"
        },
        "gateway": {
          "type": "string",
          "format": "string",
          "title": "Gateway string"
        },
        "global_ipv6_address": {
          "type": "string",
          "format": "string",
          "title": "GlobalIPv6Address string"
        },
        "global_ipv6_prefix_len": {
          "type": "integer",
          "format": "int32",
          "title": "GlobalIPv6PrefixLen int"
        },
        "ip_address": {
          "type": "string",
          "format": "string",
          "title": "IPAddress string"
        },
        "ip_prefix_len": {
          "type": "integer",
          "format": "int32",
          "title": "IPPrefixLen int"
        },
        "ipam_config": {
          "$ref": "#/definitions/networkEndpointIPAMConfig",
          "title": "IPAMConfig *EndpointIPAMConfig\t// Configurations"
        },
        "ipv6_gateway": {
          "type": "string",
          "format": "string",
          "title": "IPv6Gateway string"
        },
        "links": {
          "type": "array",
          "items": {
            "type": "string",
            "format": "string"
          },
          "title": "Links []string"
        },
        "mac_address": {
          "type": "string",
          "format": "string",
          "title": "MacAddress string"
        },
        "network_id": {
          "type": "string",
          "format": "string",
          "title": "NetworkID string // Operational data"
        }
      },
      "title": "EndpointSettings stores the network endpoint details\ntype EndpointSettings struct"
    },
    "networkIPAM": {
      "type": "object",
      "properties": {
        "config": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/networkIPAMConfig"
          },
          "title": "Config []IPAMConfig"
        },
        "driver": {
          "type": "string",
          "format": "string",
          "title": "Driver string"
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "Options map[string]string //Per network IPAM driver options"
        }
      },
      "title": "IPAM represents IP Address Management\ntype IPAM struct"
    },
    "networkIPAMConfig": {
      "type": "object",
      "properties": {
        "aux_address": {
          "type": "object",
          "additionalProperties": {
            "type": "string",
            "format": "string"
          },
          "title": "AuxAddress map[string]string 0x60json:\"AuxiliaryAddresses,omitempty\""
        },
        "gateway": {
          "type": "string",
          "format": "string",
          "title": "Gateway string 0x60json:\",omitempty\"0x60"
        },
        "ip_range": {
          "type": "string",
          "format": "string",
          "title": "IPRange string 0x60json:\",omitempty\"0x60"
        },
        "subnet": {
          "type": "string",
          "format": "string",
          "title": "Subnet string 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "IPAMConfig represents IPAM configurations\ntype IPAMConfig struct"
    },
    "networkNetworkingConfig": {
      "type": "object",
      "properties": {
        "endpoints_config": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/networkEndpointSettings"
          },
          "title": "EndpointsConfig map[string]*EndpointSettings // Endpoint configs for each connecting network"
        }
      },
      "title": "NetworkingConfig represents the container's networking configuration for each of its interfaces Carries the networking configs specified in the 'docker run' and 'docker network connect' commands\ntype NetworkingConfig struct"
    },
    "pbDockerContainerRunReqResp": {
      "type": "object",
      "properties": {
        "config": {
          "$ref": "#/definitions/containerConfig"
        },
        "container_create_response": {
          "$ref": "#/definitions/containerContainerCreateCreatedBody"
        },
        "host_config": {
          "$ref": "#/definitions/containerHostConfig"
        },
        "name": {
          "type": "string",
          "format": "string"
        },
        "networking_config": {
          "$ref": "#/definitions/networkNetworkingConfig"
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
    "pbDockerImagePullReqResp": {
      "type": "object",
      "properties": {
        "image_pull_options": {
          "$ref": "#/definitions/mobyImagePullOptions"
        },
        "ref_str": {
          "type": "string",
          "format": "string"
        },
        "resp_body": {
          "type": "string",
          "format": "byte"
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
    "pbDockerNetworkCreateReqResp": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "format": "string"
        },
        "network_create": {
          "$ref": "#/definitions/mobyNetworkCreate"
        },
        "network_create_response": {
          "$ref": "#/definitions/mobyNetworkCreateResponse"
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
    "unitsUlimit": {
      "type": "object",
      "properties": {
        "hard": {
          "type": "string",
          "format": "int64",
          "title": "Hard int64"
        },
        "name": {
          "type": "string",
          "format": "string",
          "title": "Name sstring"
        },
        "soft": {
          "type": "string",
          "format": "int64",
          "title": "Soft int64"
        }
      },
      "title": "Ulimit is a human friendly version of Rlimit.\ntype Ulimit struct"
    }
  }
}
`
)
