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
    "/v1/docker-container-inspect": {
      "post": {
        "summary": "Inspect Docker container",
        "description": "For output, plus result fileds:\n{ ..., \"state_code\": 0, \"state_message\": \"RUNNING\" }",
        "operationId": "InspectDockerContainer",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerContainerInspectReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerContainerInspectReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-container-ls": {
      "post": {
        "summary": "List Docker containers",
        "description": "For output, plus result fileds:\n{ ..., \"state_code\": 0, \"state_message\": \"RUNNING\" }",
        "operationId": "ListDockerContainers",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerContainerListReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerContainerListReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-container-prune": {
      "post": {
        "summary": "Prune Docker containers",
        "description": "For output, plus result fileds:\n{ ..., \"state_code\": 0, \"state_message\": \"RUNNING\" }",
        "operationId": "PruneDockerContainers",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerContainerPruneReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerContainerPruneReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-container-rm": {
      "post": {
        "summary": "Remove Docker containers",
        "description": "For output, plus result fileds:\n{ ..., \"state_code\": 0, \"state_message\": \"RUNNING\" }",
        "operationId": "RemoveDockerContainer",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerContainerRemoveReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerContainerRemoveReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-container-run": {
      "post": {
        "summary": "Run Docker container",
        "description": "For output, plus result fileds:\n{ ..., \"state_code\": 0, \"state_message\": \"RUNNING\" }",
        "operationId": "RunDockerContainer",
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
    },
    "/v1/docker-image-build": {
      "post": {
        "operationId": "BuildDockerImage",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerImageBuildReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerImageBuildReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-image-inspect": {
      "post": {
        "operationId": "InspectDockerImage",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerImageInspectReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerImageInspectReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-image-list": {
      "post": {
        "operationId": "ListDockerImages",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerImageListReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerImageListReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-image-prune": {
      "post": {
        "operationId": "PruneDockerImages",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerImagePruneReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerImagePruneReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-image-pull": {
      "post": {
        "summary": "Pull Docker image",
        "description": "Like command of 'docker pull', Input/Output is a same Protobuf/JSON object. For input example:\n{ \"image_ref\": \"docker.io/nginx\" }\nFor output example:\t\t\n{ \"image_ref\": \"docker.io/nginx\", \"state_code\": 0, \"state_message\": \"SUCCEEDED\" }",
        "operationId": "PullDockerImage",
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
    "/v1/docker-image-push": {
      "post": {
        "operationId": "PushDockerImage",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerImagePushReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerImagePushReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-image-remove": {
      "post": {
        "operationId": "RemoveDockerImage",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerImageRemoveReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerImageRemoveReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-network-create": {
      "post": {
        "summary": "Create Docker network - like CLI: docker network create",
        "description": "create with IPAM config, for example request:\n{} will create bridged network with random name\n{\"name\": \"foo\"} will create bridged network named foo\n{\"name\": \"foo\", \"networ_create\": {\"ipam\": { \"config\": {\"subnet\": \"172.30.1.128/25\", \"172.30.1.\"}}}}",
        "operationId": "CreateDockerNetwork",
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
    "/v1/docker-network-inspect": {
      "post": {
        "summary": "Inspect Docker network\nreturn detailed content",
        "operationId": "InspectDockerNetwork",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkInspectReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkInspectReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-network-ls": {
      "post": {
        "summary": "List Docker networks\nreturn all networks",
        "operationId": "ListDockerNetworks",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkListReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkListReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-network-prune": {
      "post": {
        "summary": "Prune Docker networks\nclear unused networks",
        "operationId": "PruneDockerNetworks",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkPruneReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkPruneReqResp"
            }
          }
        ],
        "tags": [
          "SimpleService"
        ]
      }
    },
    "/v1/docker-network-rm": {
      "post": {
        "summary": "Remove Docker network\nremove network with ID or name",
        "operationId": "RemoveDockerNetwork",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkRemoveReqResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/pbDockerNetworkRemoveReqResp"
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
    "ArgsValue": {
      "type": "object",
      "properties": {
        "value": {
          "type": "object",
          "additionalProperties": {
            "type": "boolean",
            "format": "boolean"
          }
        }
      }
    },
    "ConfigInt32Struct": {
      "type": "object",
      "properties": {
        "value": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "ImageBuildOptionsStringStruct": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      }
    },
    "PortMapPortBindingSlice": {
      "type": "object",
      "properties": {
        "internal_list": {
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
        "hostname": {
          "type": "string",
          "title": "Hostname string // Hostname"
        },
        "domainname": {
          "type": "string",
          "title": "Domainname string // Domainname"
        },
        "user": {
          "type": "string",
          "title": "User string // User that will run the command(s) inside the container, also support user:group"
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
        "attach_stderr": {
          "type": "boolean",
          "format": "boolean",
          "title": "AttachStderr bool // Attach the standard error"
        },
        "exposed_ports": {
          "$ref": "#/definitions/natPortSet",
          "title": "ExposedPorts nat.PortSet 0x60json\",omitempty\"0x60 // List of exposed ports"
        },
        "tty": {
          "type": "boolean",
          "format": "boolean",
          "description": "Tty bool // Attach standard streams to a tty, including stdin if it is not closed."
        },
        "open_stdin": {
          "type": "boolean",
          "format": "boolean",
          "title": "OpenStdin // Open stdin"
        },
        "stdin_once": {
          "type": "boolean",
          "format": "boolean",
          "description": "StdinOnce bool // If true, close stdin after the 1 attached client disconnects."
        },
        "env": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Env []string  // List of environment variable to set in the container"
        },
        "cmd": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Cmd strslice.StrSlice // Command to run when starting the container\nStrSlice represents a string or an array of strings. We need to override the json decoder to accept both options. // type StrSlice []string"
        },
        "healthcheck": {
          "$ref": "#/definitions/containerHealthConfig",
          "title": "Healthcheck *HealthConfig 0x60json:\",omitempty\"0x60 // Healthcheck describes how to check the container is healthy"
        },
        "args_escaped": {
          "type": "boolean",
          "format": "boolean",
          "title": "ArgsEscaped bool 0x60json\",omitempty\"0x60 // True if command is already escaped (Windows specific)"
        },
        "image": {
          "type": "string",
          "title": "Image string // Name of the image as it was passed by the operator (e.g. could be symbolic)"
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
          "title": "WorkingDir string // Current directory (PWD) in the command will be launched"
        },
        "entrypoint": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Entrypoint strslice.StrSlice // Entrypoint to run when starting the container"
        },
        "network_disabled": {
          "type": "boolean",
          "format": "boolean",
          "title": "NetworkDisabled bool 0x60json:\",omitempty\"0x60 // Is network disabled"
        },
        "mac_address": {
          "type": "string",
          "title": "MacAddress string 0x60json:\",omitempty\"0x60 // Mac Address of the container"
        },
        "on_build": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "OnBuild []string // ONBUILD metadata that were defined on the image Dockerfile"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Labels map[string]string // List of labels set to this container"
        },
        "stop_signal": {
          "type": "string",
          "title": "StopSignal string 0x60json:\",omitempty\"0x60 // Signal to stop a container"
        },
        "stop_timeout": {
          "$ref": "#/definitions/ConfigInt32Struct",
          "title": "StopTimeout *int 0x60json:\",omitempty\"0x60 // Timeout (in seconds) to stop a container"
        },
        "shell": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Shell strslice.StrSlice 0x60json:\",omitempty\"0x60 // Shell for shell-form of RUN, CMD, ENTRYPOINT"
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
          "title": "The ID of the created container. Required: true\nID string 0x60json:\"Id\"0x60"
        },
        "warnings": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Warnings encountered when creating the container. Required: true\nWarnings []string 0x60json:\"Warnings\"0x60"
        }
      },
      "title": "ContainerCreateCreatedBody OK response to ContainerCreate operation\ntype ContainerCreateCreatedBody struct"
    },
    "containerDeviceMapping": {
      "type": "object",
      "properties": {
        "path_on_host": {
          "type": "string",
          "title": "PathOnHost string"
        },
        "path_in_container": {
          "type": "string",
          "title": "PathInContainer string"
        },
        "cgroup_permissions": {
          "type": "string",
          "title": "CgroupPermissions string"
        }
      },
      "title": "DeviceMapping represents the device mapping between the host and the container.\ntype DeviceMapping struct"
    },
    "containerHealthConfig": {
      "type": "object",
      "properties": {
        "test": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Test []string 0x60json:\",omitempty\"0x60\nTest is the test to perform to check that the container is healthy.\nAn empty slice means to inherit the default.\nThe options are:\n{} : inherit healthcheck\n{\"NONE\"} : disable healthcheck\n{\"CMD\", args...} : exec arguments directly\n{\"CMD-SHELL\", command} : run command with system's default shell"
        },
        "interval": {
          "$ref": "#/definitions/protobufDuration",
          "description": "Zero means to inherit. Durations are expressed as integer nanoseconds.\nInterval time.Duration 0x60json:\",omitempty\"0x60 // Interval is the time to wait between checks."
        },
        "timeout": {
          "$ref": "#/definitions/protobufDuration",
          "description": "Timeout time.Duration 0x60json:\",omitempty\"0x60 // Timeout is the time to wait before considering the check to have hung."
        },
        "start_period": {
          "$ref": "#/definitions/protobufDuration",
          "description": "StartPeriod time.Duration 0x60json:\",omitempty\"0x60 // The start period for the container to initialize before the retries starts to count down."
        },
        "retries": {
          "type": "integer",
          "format": "int32",
          "description": "Retries int 0x60json:\",omitempty\"0x60// Retries is the number of consecutive failures needed to consider a container as unhealthy. Zero means inherit."
        }
      },
      "title": "HealthConfig holds configuration settings for the HEALTHCHECK feature.\ntype HealthConfig struct"
    },
    "containerHostConfig": {
      "type": "object",
      "properties": {
        "binds": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Binds []string // List of volume bindings for this container",
          "title": "Applicable to all platforms"
        },
        "container_id_file": {
          "type": "string",
          "title": "ContainerIDFile // File (path) where the containerId is written"
        },
        "log_config": {
          "$ref": "#/definitions/containerLogConfig",
          "title": "LogConfig LogConfig // Configuration of the logs for this container"
        },
        "network_mode": {
          "type": "string",
          "title": "NetworkMode NetworkMode // Network mode to use for the container, \"none\", \"default\", \"container:\u003cid\u003e\"\nNetworkMode represents the container network stack. // type NetworkMode string"
        },
        "port_bindings": {
          "$ref": "#/definitions/natPortMap",
          "title": "PortBindings nat.PortMap // Port mapping between the exposed port (container) and the host"
        },
        "restart_policy": {
          "$ref": "#/definitions/containerRestartPolicy",
          "title": "RestartPolicy RestartPolicy // Restart policy to be used for the container"
        },
        "auto_remove": {
          "type": "boolean",
          "format": "boolean",
          "title": "AutoRemove bool // Automatically remove container when it exits"
        },
        "volume_driver": {
          "type": "string",
          "title": "VolumeDriver string // Name of the volume driver used to mount volumes"
        },
        "volumes_from": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "VolumesFrom []string // List of volumes to take from other container"
        },
        "cap_add": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "CapAdd strslice.StrSlice // List of kernel capabilities to add to the container",
          "title": "Applicable to UNIX platforms"
        },
        "cap_drop": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "CapDrop strslice.StrSlice // List of kernel capabilities to remove from the container"
        },
        "dns": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "DNS []string 0x60json:\"Dns\"0x60 // List of DNS server to lookup"
        },
        "dns_options": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "DNSOptions []string 0x60json:\"DnsOptions\"0x60 // List of DNSOption to look for"
        },
        "dns_search": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "DNSSearch []string 0x60json:\"DnsSearch\"0x60 // List of DNSSearch to look for"
        },
        "extra_hosts": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "ExtraHosts []string // List of extra hosts"
        },
        "group_add": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "GroupAdd []string // List of additional groups that the container process will run as"
        },
        "ipc_mode": {
          "type": "string",
          "title": "IpcMode IpcMode // IPC namespace to use for the container, \"\", \"host\", \"container\"\nIpcMode represents the container ipc stack. // type IpcMode string"
        },
        "cgroup": {
          "type": "string",
          "title": "Cgroup CgroupSpec // Cgroup to use for the container, \"container:\u003cid\u003e\"\nCgroupSpec represents the cgroup to use for the container. // type CgroupSpec string"
        },
        "links": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Links []string // List of links (in the name:alias form)"
        },
        "oom_score_adj": {
          "type": "integer",
          "format": "int32",
          "title": "OomScoreAdj int // Container preference for OOM-killing"
        },
        "pid_mode": {
          "type": "string",
          "title": "PidMode PidMode // PID namespace to use for the container\nPidMode represents the pid namespace of the container. // type PidMode string"
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
        "readonly_rootfs": {
          "type": "boolean",
          "format": "boolean",
          "title": "ReadonlyRootfs bool // Is the container root filesystem in read-only"
        },
        "security_opt": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "SecurityOpt []string // List of string values to customize labels for MLS systems, such as SELinux."
        },
        "storage_opt": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Storage driver options per container.\nStorageOpt map[string]string 0x60json:\",omitempty\"0x60"
        },
        "tmpfs": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "List of tmpfs (mounts) used for the container\nTmpfs map[string]string 0x60json:\",omitempty\"0x60"
        },
        "uts_mode": {
          "type": "string",
          "title": "UTSMode UTSMode // UTSMode represents the UTS namespace of the container. // type UTSMode string\nUTS namespace to use for the container"
        },
        "userns_mode": {
          "type": "string",
          "title": "UsernsMode UsernsMode // UsernsMode represents userns mode in the container. // type UsernsMode string\nThe user namespace to use for the container"
        },
        "shm_size": {
          "type": "string",
          "format": "int64",
          "title": "ShmSize int64\nTotal shm memory usage"
        },
        "sysctls": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Sysctls map[string]string 0x60json:\",omitempty\"0x60\nList of Namespaced sysctls used for the container"
        },
        "runtime": {
          "type": "string",
          "title": "Runtime string 0x60json:\",omitempty\"0x60\nRuntime to use with this container"
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
        "isolation": {
          "type": "string",
          "title": "Isolation Isolation // Isolation technology of the container (e.g. default, hyperv)\nIsolation represents the isolation technology of a container. The supported values are platform specific // type Isolation string"
        },
        "resources": {
          "$ref": "#/definitions/containerResources",
          "title": "Contains container's resources (cgroups, ulimits)\nResources"
        },
        "mounts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mountMount"
          },
          "title": "Mounts specs used by the container\nMounts []mount.Mount 0x60json:\",omitempty\"0x60"
        },
        "masked_paths": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "MaskedPaths is the list of paths to be masked inside the container (this overrides the default set of paths)\nMaskedPaths []string"
        },
        "readonly_paths": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "ReadonlyPaths is the list of paths to be set as read-only inside the container (this overrides the default set of paths)\nReadonlyPaths []string"
        },
        "init": {
          "$ref": "#/definitions/containerHostConfigBoolStruct",
          "title": "Run a custom init inside the container, if null, use the daemon's configured settings\nInit *bool 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "HostConfig the non-portable Config structure of a container.\nHere, \"non-portable\" means \"dependent of the host we are running on\".\nPortable information *should* appear in Config.\ntype HostConfig struct"
    },
    "containerHostConfigBoolStruct": {
      "type": "object",
      "properties": {
        "value": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "containerLogConfig": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "title": "Type string // \"\", \"blocking\", \"non-blocking\""
        },
        "config": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Config map[string]string"
        }
      },
      "title": "LogConfig represents the logging configuration of the container.\ntype LogConfig struct"
    },
    "containerResources": {
      "type": "object",
      "properties": {
        "cpu_shares": {
          "type": "string",
          "format": "int64",
          "description": "CPUShares int64 0x60json:\"CpuShares\"0x60 // CPU shares (relative weight vs. other containers)",
          "title": "Applicable to all platforms"
        },
        "memory": {
          "type": "string",
          "format": "int64",
          "title": "Memory int64 // Memory limit (in bytes)"
        },
        "nano_cpus": {
          "type": "string",
          "format": "int64",
          "description": "NanoCPUs int64 0x60json:\"NonoCpus\"0x60 // CPU quota in units of 10\u003csup\u003e-9\u003c/sup\u003e CPUs."
        },
        "cgroup_parent": {
          "type": "string",
          "description": "CgroupParent string // Parent cgroup.",
          "title": "Applicable to UNIX platforms"
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
        "blkio_device_read_bps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/blkiodevThrottleDevice"
          },
          "title": "BlkioDeviceReadBps []*blkiodev.ThrottleDevice"
        },
        "blkio_device_write_bps": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/blkiodevThrottleDevice"
          },
          "title": "BlkioDeviceWriteBps []*blkiodev.ThrottleDevice"
        },
        "blkio_device_read_iops": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/blkiodevThrottleDevice"
          },
          "title": "BlkioDeviceReadIOps []*blkiodev.ThrottleDevice"
        },
        "blkio_device_write_iops": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/blkiodevThrottleDevice"
          },
          "title": "BlkioDeviceWriteIOps []*blkiodev.ThrottleDevice"
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
        "cpuset_cpus": {
          "type": "string",
          "title": "CpusetCpus string // CpusetCpus 0-2, 0,1"
        },
        "cpuset_mems": {
          "type": "string",
          "title": "CpusetMems string // CpusetMems 0-2, 0,1"
        },
        "devices": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/containerDeviceMapping"
          },
          "title": "Devices []DeviceMapping // List of devices to map inside the container"
        },
        "device_cgroup_rules": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "DeviceCgroupRules []string // List of rule to be added to the device cgroup"
        },
        "disk_quota": {
          "type": "string",
          "format": "int64",
          "title": "DiskQuota int64 // Disk limit (in bytes)"
        },
        "kernel_memory": {
          "type": "string",
          "format": "int64",
          "title": "KernelMemory int64 // Kernel memory limit (in bytes)"
        },
        "memory_reservation": {
          "type": "string",
          "format": "int64",
          "title": "MemoryReservation int64 // Memory soft limit (in bytes)"
        },
        "memory_swap": {
          "type": "string",
          "format": "int64",
          "title": "MemorySwap int64 // Total memory usage (memory + swap); set -1 to enable unlimited swap"
        },
        "memory_swappiness": {
          "$ref": "#/definitions/containerResourcesInt64Struct",
          "title": "MemorySwappiness *int64 // Tuning container memory swappiness behaviour"
        },
        "oom_kill_disable": {
          "$ref": "#/definitions/containerResourcesBoolStruct",
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
        "io_maximum_iops": {
          "type": "string",
          "format": "uint64",
          "title": "IOMaximumIOps uint64 // Maximum IOps for the container system drive"
        },
        "io_maximum_bandwidth": {
          "type": "string",
          "format": "uint64",
          "title": "IOMaximumBandwidth // Maximum IO in bytes per second for the container system drive"
        }
      },
      "title": "Resources contains container's resources (cgroups config, ulimits...)\nto see https://github.com/moby/moby/blob/master/api/types/container/host_config.go"
    },
    "containerResourcesBoolStruct": {
      "type": "object",
      "properties": {
        "value": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "containerResourcesInt64Struct": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "containerRestartPolicy": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "Name string"
        },
        "maximum_retry_count": {
          "type": "integer",
          "format": "int32",
          "title": "MaximumRetryCount int"
        }
      },
      "title": "RestartPolicy represents the restart policies of the container.\ntype RestartPolicy struct"
    },
    "filtersArgs": {
      "type": "object",
      "properties": {
        "fields": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/ArgsValue"
          },
          "title": "Fields map[string]map[string]bool"
        }
      },
      "title": "Args stores filter arguments as map key:{map key: bool}.\nIt contains an aggregation of the map of arguments (which are in the form\nof -f 'key=value') based on the key, and stores values for the same key\nin a map with string keys and boolean values.\ne.g given -f 'label=label1=1' -f 'label=label2=2' -f 'image.name=ubuntu'\nthe args will be {\"image.name\":{\"ubuntu\":true},\"label\":{\"label1=1\":true,\"label2=2\":true}}\ntype Args struct"
    },
    "mobyAuthConfig": {
      "type": "object",
      "properties": {
        "username": {
          "type": "string",
          "title": "Username string 0x60json:\"username,omitempty\"0x60"
        },
        "password": {
          "type": "string",
          "title": "Password string 0x60json:\"password,omitempty\"0x60"
        },
        "auth": {
          "type": "string",
          "title": "Auth string  0x60json:\"auth,omitempty\"0x60"
        },
        "email": {
          "type": "string",
          "description": "Email string 0x60json:\"email,omitempty\"0x60 // Email is an Operational value associated with the username. This field is deprecated and will be removed in a later version of docker."
        },
        "server_address": {
          "type": "string",
          "title": "ServerAddress string 0x60json:\"serveraddress,omitempty\"0x60"
        },
        "identity_token": {
          "type": "string",
          "description": "IdentityToken string 0x60json:\"identitytoken,omitempty\"0x60 // IdentityToken is used to authenticate the user and get an access token for the registry."
        },
        "registry_token": {
          "type": "string",
          "title": "RegistryToken string 0x60json:\"registrytoken,omitempty\"0x60 // RegistryToken is a bearer token to be sent to a registry"
        }
      },
      "description": "AuthConfig contains authorization information for connecting to a Registry\ntype AuthConfig struct",
      "title": "https://github.com/moby/moby/blob/master/api/types/auth.go"
    },
    "mobyContainer": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "ID string 0x60json:\"Id\"0x60"
        },
        "names": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Names []string"
        },
        "image": {
          "type": "string",
          "title": "Image string"
        },
        "image_id": {
          "type": "string",
          "title": "ImageID string"
        },
        "command": {
          "type": "string",
          "title": "Command string"
        },
        "created": {
          "type": "string",
          "format": "int64",
          "title": "Created int64"
        },
        "Ports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyPort"
          },
          "title": "Ports []Port"
        },
        "size_rw": {
          "type": "string",
          "format": "int64",
          "title": "SizeRw int64 0x60json:\",omitempty\"0x60"
        },
        "size_root_fs": {
          "type": "string",
          "format": "int64",
          "title": "SizeRootFs int64 0x60json:\",omitempty\"0x60"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Labels map[string]string"
        },
        "state": {
          "type": "string",
          "title": "State string"
        },
        "status": {
          "type": "string",
          "title": "Status string"
        },
        "host_config": {
          "$ref": "#/definitions/mobyContainerHostConfig",
          "title": "HostConfig struct"
        },
        "network_settings": {
          "$ref": "#/definitions/mobySummaryNetworkSettings",
          "title": "NetworkSettings *SummaryNetworkSettings"
        },
        "mounts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyMountPoint"
          },
          "title": "Mounts []MountPoint"
        }
      },
      "title": "Container contains response of Remote API:\nGET  \"/containers/json\"\ntype Container struct"
    },
    "mobyContainerHostConfig": {
      "type": "object",
      "properties": {
        "network_mode": {
          "type": "string",
          "title": "NetworkMode string 0x60json:\",omitempty\""
        }
      }
    },
    "mobyContainerJSON": {
      "type": "object",
      "properties": {
        "container_json_base": {
          "$ref": "#/definitions/mobyContainerJSONBase",
          "title": "ContainerJSONBase"
        },
        "mounts": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyMountPoint"
          },
          "title": "Mounts []MountPoint"
        },
        "config": {
          "$ref": "#/definitions/containerConfig",
          "title": "Config *container.Config"
        },
        "network_settings": {
          "$ref": "#/definitions/mobyNetworkSettings",
          "title": "NetworkSettings *NetworkSettings"
        }
      },
      "title": "ContainerJSON is newly used struct along with MountPoint\ntype ContainerJSON struct"
    },
    "mobyContainerJSONBase": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "ID string 0x60json:\"Id\"0x60"
        },
        "created": {
          "type": "string",
          "title": "Created string"
        },
        "path": {
          "type": "string",
          "title": "Path string"
        },
        "args": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Args []string"
        },
        "state": {
          "$ref": "#/definitions/mobyContainerState",
          "title": "State *ContainerState"
        },
        "image": {
          "type": "string",
          "title": "Image string"
        },
        "resolv_conf_path": {
          "type": "string",
          "title": "ResolvConfPath string"
        },
        "hostname_path": {
          "type": "string",
          "title": "HostnamePath string"
        },
        "hosts_path": {
          "type": "string",
          "title": "HostsPath string"
        },
        "log_path": {
          "type": "string",
          "title": "LogPath string"
        },
        "node": {
          "$ref": "#/definitions/mobyContainerNode",
          "title": "Node *ContainerNode 0x60json:\",omitempty\"0x60"
        },
        "name": {
          "type": "string",
          "title": "Name string"
        },
        "restart_count": {
          "type": "integer",
          "format": "int32",
          "title": "RestartCount int"
        },
        "driver": {
          "type": "string",
          "title": "Driver string"
        },
        "platform": {
          "type": "string",
          "title": "Platform string"
        },
        "mount_label": {
          "type": "string",
          "title": "MountLabel string"
        },
        "process_label": {
          "type": "string",
          "title": "ProcessLabel string"
        },
        "app_armor_profile": {
          "type": "string",
          "title": "AppArmorProfile string"
        },
        "exec_ids": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "ExecIDs []string"
        },
        "host_config": {
          "$ref": "#/definitions/containerHostConfig",
          "title": "HostConfig *container.HostConfig"
        },
        "graph_driver": {
          "$ref": "#/definitions/mobyGraphDriverData",
          "title": "GraphDriver GraphDriverData"
        },
        "size_rw": {
          "$ref": "#/definitions/mobyContainerJSONBaseInt64Struct",
          "title": "SizeRw *int64 0x60json:\",omitempty\"0x60"
        },
        "size_root_fs": {
          "$ref": "#/definitions/mobyContainerJSONBaseInt64Struct",
          "title": "SizeRootFs *int64 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "ContainerJSONBase contains response of Remote API:\nGET \"/containers/{name:.*}/json\"\ntype ContainerJSONBase struct"
    },
    "mobyContainerJSONBaseInt64Struct": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "mobyContainerListOptions": {
      "type": "object",
      "properties": {
        "quiet": {
          "type": "boolean",
          "format": "boolean",
          "title": "Quiet bool"
        },
        "size": {
          "type": "boolean",
          "format": "boolean",
          "title": "Size bool"
        },
        "all": {
          "type": "boolean",
          "format": "boolean",
          "title": "All bool"
        },
        "latest": {
          "type": "boolean",
          "format": "boolean",
          "title": "Latest bool"
        },
        "since": {
          "type": "string",
          "title": "Since string"
        },
        "before": {
          "type": "string",
          "title": "Before string"
        },
        "limit": {
          "type": "integer",
          "format": "int32",
          "title": "Limit int"
        },
        "filters": {
          "$ref": "#/definitions/filtersArgs",
          "title": "Filters filters.Args"
        }
      },
      "description": "ContainerListOptions holds parameters to list containers with.\ntype ContainerListOptions struct",
      "title": "https://github.com/moby/moby/blob/master/api/types/client.go"
    },
    "mobyContainerNode": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "ID string"
        },
        "ip_address": {
          "type": "string",
          "title": "IPAddress string"
        },
        "addr": {
          "type": "string",
          "title": "Addr string"
        },
        "name": {
          "type": "string",
          "title": "Name string"
        },
        "cpus": {
          "type": "integer",
          "format": "int32",
          "title": "Cpus int"
        },
        "memory": {
          "type": "string",
          "format": "int64",
          "title": "Memory int64"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Labels map[string]string"
        }
      },
      "title": "ContainerNode stores information about the node that a container is running on. It's only available in Docker Swarm\ntype ContainerNode struct"
    },
    "mobyContainerRemoveOptions": {
      "type": "object",
      "properties": {
        "remove_volumes": {
          "type": "boolean",
          "format": "boolean",
          "title": "RemoveVolumes bool"
        },
        "remove_links": {
          "type": "boolean",
          "format": "boolean",
          "title": "RemoveLinks bool"
        },
        "force": {
          "type": "boolean",
          "format": "boolean",
          "title": "Force bool"
        }
      },
      "title": "ContainerRemoveOptions holds parameters to remove containers\ntype ContainerRemoveOptions struct"
    },
    "mobyContainerState": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string",
          "title": "Status string // string representation of the container state. Can be one of \"created\", \"running\", \"paused\", \"restarting\", \"removing\", \"exited\", of \"dead\""
        },
        "running": {
          "type": "boolean",
          "format": "boolean",
          "title": "Running bool"
        },
        "paused": {
          "type": "boolean",
          "format": "boolean",
          "title": "Paused bool"
        },
        "restarting": {
          "type": "boolean",
          "format": "boolean",
          "title": "Restarting bool"
        },
        "oom_killed": {
          "type": "boolean",
          "format": "boolean",
          "title": "OOMKilled bool"
        },
        "dead": {
          "type": "boolean",
          "format": "boolean",
          "title": "Dead bool"
        },
        "pid": {
          "type": "integer",
          "format": "int32",
          "title": "Pid int"
        },
        "exit_code": {
          "type": "integer",
          "format": "int32",
          "title": "ExitCode int"
        },
        "error": {
          "type": "string",
          "title": "Error string"
        },
        "started_at": {
          "type": "string",
          "title": "StartedAt string"
        },
        "finished_at": {
          "type": "string",
          "title": "FinishedAt string"
        },
        "health": {
          "$ref": "#/definitions/mobyHealth",
          "title": "Health *Health 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "ContainerState stores container's running state\nit's part of ContainerJSONBase and will return by \"inspect\" command\ntype ContainerState struct"
    },
    "mobyContainersPruneReport": {
      "type": "object",
      "properties": {
        "containers_deleted": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "ContainersDeleted []string"
        },
        "space_reclaimed": {
          "type": "string",
          "format": "uint64",
          "title": "SpaceReclaimed uint64"
        }
      },
      "title": "ContainersPruneReport contains the response for Engine API:\nPOST \"/conainers/prune\"\ntype ContainersPruneReport struct"
    },
    "mobyDefaultNetworkSettings": {
      "type": "object",
      "properties": {
        "endpoint_id": {
          "type": "string",
          "title": "EndpointID string // EndpointID uniquely represents a service endpoint in a sandbox"
        },
        "gateway": {
          "type": "string",
          "title": "Gateway string // Gateway holds the gateway address for the network"
        },
        "global_ipv6_address": {
          "type": "string",
          "title": "GlobalIPv6Address string // GlobalIPv6Address holds network's global IPv6 address"
        },
        "global_ipv6_prefix_len": {
          "type": "integer",
          "format": "int32",
          "title": "GlobalIPv6PrefixLen int // GlobalIPv6PrefixLen represents mask length of network's global IPv6 address"
        },
        "ip_address": {
          "type": "string",
          "title": "IPAddress string // IPAddress holds the IPv4 address for the network"
        },
        "ip_prefix_len": {
          "type": "integer",
          "format": "int32",
          "title": "IPPrefixLen int // IPPrefixLen represents mask length of network's IPv4 address"
        },
        "ipv6_gateway": {
          "type": "string",
          "title": "IPv6Gateway string // IPv6Gateway holds gateway address specific for IPv6"
        },
        "mac_address": {
          "type": "string",
          "title": "MacAddress string // MacAddress holds the MAC address for the network"
        }
      },
      "title": "DefaultNetworkSettings holds network information during the 2 release deprecation period.\nIt will be removed in Docker 1.11.\ntype DefaultNetworkSettings struct"
    },
    "mobyEndpointResource": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "Name string"
        },
        "endpoint_id": {
          "type": "string",
          "title": "EndpointID string"
        },
        "mac_address": {
          "type": "string",
          "title": "MacAddress string"
        },
        "ipv4_address": {
          "type": "string",
          "title": "IPv4Address string"
        },
        "ipv6_address": {
          "type": "string",
          "title": "IPv6Address string"
        }
      },
      "title": "EndpointResource contains network resources allocated and used for a container in a network\ntype EndpointResource struct"
    },
    "mobyGraphDriverData": {
      "type": "object",
      "properties": {
        "data": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "data\nRequired: true\nData map[string]string 0x60json:\"Data\"0x60"
        },
        "name": {
          "type": "string",
          "title": "name\nRequired: true\nName string 0x60json:\"Name\"0x60"
        }
      },
      "description": "GraphDriverData Information about a container's graph driver.\nswagger:model GraphDriverData\ntype GraphDriverData struct",
      "title": "https://github.com/moby/moby/blob/master/api/types/graph_driver_data.go"
    },
    "mobyHealth": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string",
          "title": "Status string // Status is one of Starting, Healthy or Unhealthy"
        },
        "failing_streak": {
          "type": "integer",
          "format": "int32",
          "title": "FailingStreak int // FailingStreak is the number of consecutive failures"
        },
        "log": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyHealthcheckResult"
          },
          "title": "Log []*HealthcheckResult // Log contains the last few results (oldest first)"
        }
      },
      "title": "Health stores information about the Container's healthcheck results\ntype Health struct"
    },
    "mobyHealthcheckResult": {
      "type": "object",
      "properties": {
        "start": {
          "type": "string",
          "format": "date-time",
          "title": "Start time.Time // Start is the time this check started"
        },
        "end": {
          "type": "string",
          "format": "date-time",
          "title": "End time.Time // End is the time this check ended"
        },
        "exit_code": {
          "type": "integer",
          "format": "int32",
          "title": "ExitCode int // ExitCode meanings: 0=healthy, 1=unhealthy, 2=reserved (considered unhealthy), else=error running probe"
        },
        "output": {
          "type": "string",
          "title": "Output string // Output from last check"
        }
      },
      "title": "HealthcheckResult stores information about a single run of a healthcheck probe\ntype HealthcheckResult struct"
    },
    "mobyImageBuildOptions": {
      "type": "object",
      "properties": {
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Tags []string"
        },
        "suppress_output": {
          "type": "boolean",
          "format": "boolean",
          "title": "SuppressOutput bool"
        },
        "remote_context": {
          "type": "string",
          "title": "RemoteContext string"
        },
        "no_cache": {
          "type": "boolean",
          "format": "boolean",
          "title": "NoCache bool"
        },
        "remove": {
          "type": "boolean",
          "format": "boolean",
          "title": "Remove bool"
        },
        "force_remove": {
          "type": "boolean",
          "format": "boolean",
          "title": "ForceRemove bool"
        },
        "pull_parent": {
          "type": "boolean",
          "format": "boolean",
          "title": "PullParent bool"
        },
        "isolation": {
          "type": "string",
          "title": "Isolation container.Isolation // Isolation represents the isolation technology of a container. The supported value are platform specific. // type Isolation string"
        },
        "cpu_set_cpus": {
          "type": "string",
          "title": "CPUSetCPUs string"
        },
        "cpu_set_mems": {
          "type": "string",
          "title": "CPUSetMems string"
        },
        "cpu_shares": {
          "type": "string",
          "format": "int64",
          "title": "CPUShares int64"
        },
        "cpu_quota": {
          "type": "string",
          "format": "int64",
          "title": "CPUQuota int64"
        },
        "cpu_period": {
          "type": "string",
          "format": "int64",
          "title": "CPUPeriod int64"
        },
        "memory": {
          "type": "string",
          "format": "int64",
          "title": "Memory int64"
        },
        "memory_swap": {
          "type": "string",
          "format": "int64",
          "title": "MemorySwap int64"
        },
        "cgroup_parent": {
          "type": "string",
          "title": "CgroupParent string"
        },
        "network_mode": {
          "type": "string",
          "title": "NetworkMode string"
        },
        "shm_size": {
          "type": "string",
          "format": "int64",
          "title": "ShmSize int64"
        },
        "dockerfile": {
          "type": "string",
          "title": "Dockerfile string"
        },
        "ulimits": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/unitsUlimit"
          },
          "title": "Ulimits []*units.Ulimit"
        },
        "build_args": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/ImageBuildOptionsStringStruct"
          },
          "description": "BuildArgs map[string]*string // BuildArgs needs to be a *string instead of just a string so that we can tell the difference between \"\" (empty string) and no value at all (nil). See the parsing of buildArgs in api/server/router/build/build_routers.go for even more info."
        },
        "auth_configs": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/mobyAuthConfig"
          },
          "title": "AuthConfigs map[string]AuthConfig"
        },
        "context": {
          "type": "string",
          "format": "byte",
          "title": "Context io.Reader"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Labels map[string]string"
        },
        "squash": {
          "type": "boolean",
          "format": "boolean",
          "title": "Squash bool // squash the resulting image's layers to the parent preserve the original image and creates a new one from the parent with all the changes applied to a single layer"
        },
        "cache_from": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "CacheFrom []string // CacheFrom specifies images that are used for matching cache. Images specified here do not need to have a valid parent chain to match cache."
        },
        "security_opt": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "SecurityOpt []string"
        },
        "extra_hosts": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "ExtraHosts []string // List of extra hosts"
        },
        "target": {
          "type": "string",
          "title": "Target string"
        },
        "session_id": {
          "type": "string",
          "title": "SessionID string"
        },
        "platform": {
          "type": "string",
          "title": "Platform string"
        },
        "version": {
          "type": "string",
          "title": "Version BuildVersion // Version specifies the version of the underlying builder to use // type BuilderVersion string // BuilderV1 is the first generation builder in docker daemon. BuilderV1 BuilderVersion = \"1\" // BuilderBuildKit is builder based on moby/buildkit project. BuilderBuildKit = \"2\""
        },
        "build_id": {
          "type": "string",
          "description": "BuildID string // BuildID is an Operational identifier that can be passed together with the build request. Tha same identifier can be used to gracefully cancel the build with the cancel request."
        }
      },
      "title": "ImageBuildOptions holds the information necessary to build images.\ntype ImageBuildOptions struct"
    },
    "mobyImageBuildResponse": {
      "type": "object",
      "properties": {
        "body": {
          "type": "string",
          "format": "byte",
          "title": "Body io.ReadCloser"
        },
        "os_type": {
          "type": "string",
          "title": "OSType string"
        }
      },
      "title": "ImageBuildResponse holds information returned by a server after building an image\ntype ImageBuildResponse struct"
    },
    "mobyImageDeleteResponseItem": {
      "type": "object",
      "properties": {
        "deleted": {
          "type": "string",
          "title": "The image ID of an image that was deleted\nDeleted string 0x60json:\"Deleted,omitempty\"0x60"
        },
        "untagged": {
          "type": "string",
          "title": "The image ID of an image that was untagged\nUntagged string 0x60json:\"Untagged,omitempty\"0x60"
        }
      },
      "description": "ImageDeleteResponseItem image delete response item\nswagger:model ImageDeleteResponseItem\ntype ImageDeleteResponseItem struct",
      "title": "https://github.com/moby/moby/blob/master/api/types/image_delete_response_item.go"
    },
    "mobyImageInspect": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "ID string 0x60json:\"Id\"0x60"
        },
        "repo_tags": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "RepoTags []string"
        },
        "repo_digests": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "RepoDigests []string"
        },
        "parent": {
          "type": "string",
          "title": "Parent string"
        },
        "comment": {
          "type": "string",
          "title": "Comment string"
        },
        "created": {
          "type": "string",
          "title": "Created string"
        },
        "container": {
          "type": "string",
          "title": "Container string"
        },
        "container_config": {
          "$ref": "#/definitions/containerConfig",
          "title": "ContainerConfig *container.Config"
        },
        "docker_version": {
          "type": "string",
          "title": "DockerVersion string"
        },
        "author": {
          "type": "string",
          "title": "Author string"
        },
        "config": {
          "$ref": "#/definitions/containerConfig",
          "title": "Config *container.Config"
        },
        "architecture": {
          "type": "string",
          "title": "Architecture string"
        },
        "os": {
          "type": "string",
          "title": "Os string"
        },
        "os_version": {
          "type": "string",
          "title": "OsVersion string 0x60json:\",omitempty\"0x60"
        },
        "size": {
          "type": "string",
          "format": "int64",
          "title": "Size int64"
        },
        "virtual_size": {
          "type": "string",
          "format": "int64",
          "title": "VirtualSize int64"
        },
        "graph_driver": {
          "$ref": "#/definitions/mobyGraphDriverData",
          "title": "GraphDriver GraphDriverData"
        },
        "root_fs": {
          "$ref": "#/definitions/mobyRootFS",
          "title": "RootFS RootFS"
        },
        "metadata": {
          "$ref": "#/definitions/mobyImageMetadata",
          "title": "Metadata ImageMetadata"
        }
      },
      "title": "ImageInspect contains response of Engine API:\nGET \"/images/{name:.*}/json\"\ntype ImageInspect struct"
    },
    "mobyImageListOptions": {
      "type": "object",
      "properties": {
        "all": {
          "type": "boolean",
          "format": "boolean",
          "title": "All bool"
        },
        "filters": {
          "$ref": "#/definitions/filtersArgs",
          "title": "Filters filters.Args"
        }
      },
      "title": "ImageListOptions holds parameters to filter the list of images with.\ntype ImageListOptions struct"
    },
    "mobyImageMetadata": {
      "type": "object",
      "properties": {
        "last_tag_time": {
          "type": "string",
          "format": "date-time",
          "title": "LastTagTime time.Time 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "ImageMetadata contains engine-local data about the image\ntype ImageMetadata struct"
    },
    "mobyImagePullOptions": {
      "type": "object",
      "properties": {
        "all": {
          "type": "boolean",
          "format": "boolean",
          "title": "All bool"
        },
        "registry_auth": {
          "type": "string",
          "title": "RegistryAuth string // RegistryAuth is the base64 encoded credentials for the registry"
        },
        "privilege_func": {
          "type": "string",
          "title": "PrivilegeFunc RequestPrivilegeFunc"
        },
        "platform": {
          "type": "string",
          "title": "Platform string"
        }
      },
      "title": "ImagePullOptions holds information to pull images.\ntype ImagePullOptions struct"
    },
    "mobyImagePushOptions": {
      "type": "object",
      "properties": {
        "all": {
          "type": "boolean",
          "format": "boolean",
          "title": "All bool"
        },
        "registry_auth": {
          "type": "string",
          "title": "RegistryAuth string // RegistryAuth is the base64 encoded credentials for the registry"
        },
        "privilege_func": {
          "type": "string",
          "title": "PrivilegeFunc RequestPrivilegeFunc"
        },
        "platform": {
          "type": "string",
          "title": "Platform string"
        }
      },
      "title": "ImagePushOptions holds information to push images.\ntype ImagePushOptions ImagePullOptions"
    },
    "mobyImageRemoveOptions": {
      "type": "object",
      "properties": {
        "force": {
          "type": "boolean",
          "format": "boolean",
          "title": "Force bool"
        },
        "prune_children": {
          "type": "boolean",
          "format": "boolean",
          "title": "PruneChildren bool"
        }
      },
      "title": "ImageRemoveOptions holds parameters to remove images.\ntype ImageRemoveOptions struct"
    },
    "mobyImageSummary": {
      "type": "object",
      "properties": {
        "containers": {
          "type": "string",
          "format": "int64",
          "title": "containers int64 0x60json:\"Containers\"0x60\nRequired: true"
        },
        "created": {
          "type": "string",
          "format": "int64",
          "title": "created int64 0x60json:\"Created\"0x60\nRequired: true"
        },
        "id": {
          "type": "string",
          "title": "Id string 0x60json:\"Id\"0x60\nRequired: true"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "labels map[string]string 0x60json:\"Labels\"0x60\nRequired: true"
        },
        "parent_id": {
          "type": "string",
          "title": "parent Id\nRequired: true\nParentID string 0x60json:\"ParentId\"0x60"
        },
        "repo_digests": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "repo digests\nRequired: true\nRepoDigests []string 0x60json:\"RepoDigests\"0x60"
        },
        "repo_tags": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "repo tags\nRequired: true\nRepoTags []string 0x60json:\"RepoTags\"0x60"
        },
        "shared_size": {
          "type": "string",
          "format": "int64",
          "title": "shared size\nRequired: true\nSharedSize int64 0x60json:\"SharedSize\"0x60"
        },
        "size": {
          "type": "string",
          "format": "int64",
          "title": "size int64 0x60json:\"Size\"0x60\nRequired: true"
        },
        "virtual_size": {
          "type": "string",
          "format": "int64",
          "title": "virtual size\nRequired: true\nVirtualSize int64 0x60json:\"VirtualSize\"0x60"
        }
      },
      "description": "ImageSummary image summary\nswagger:model ImageSummary\ntype ImageSummary struct",
      "title": "https://github.com/moby/moby/blob/master/api/types/image_summary.go"
    },
    "mobyImagesPruneReport": {
      "type": "object",
      "properties": {
        "images_deleted": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyImageDeleteResponseItem"
          },
          "title": "ImagesDeleted []ImageDeleteResponseItem"
        },
        "space_reclaimed": {
          "type": "string",
          "format": "uint64",
          "title": "SpaceReclaimed uint64"
        }
      },
      "title": "ImagesPruneReport contains the response for Engine API:\nPOST \"/images/prune\"\ntype ImagesPruneReport struct"
    },
    "mobyMountPoint": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "title": "Type mount.Type 0x60json:\",omitempty\"0x60\nType represents the type of mount. // type Type string"
        },
        "name": {
          "type": "string",
          "title": "Name string 0x60json:\",omitempty\"0x60"
        },
        "source": {
          "type": "string",
          "title": "Source string"
        },
        "destination": {
          "type": "string",
          "title": "Destination string"
        },
        "driver": {
          "type": "string",
          "title": "Driver string 0x60json:\",omitempty\"0x60"
        },
        "mode": {
          "type": "string",
          "title": "Mode string"
        },
        "rw": {
          "type": "boolean",
          "format": "boolean",
          "title": "RW bool"
        },
        "propagation": {
          "type": "string",
          "title": "Propagation mount.Propagation\nPropagation represents the propagation of a mount. // type Propagation string"
        }
      },
      "title": "MountPoint represents a mount point configuration inside the container.\nThis is used for reporting the mountpoints in used by a container.\ntype MountPoint struct"
    },
    "mobyNetworkCreate": {
      "type": "object",
      "properties": {
        "check_duplicate": {
          "type": "boolean",
          "format": "boolean",
          "description": "CheckDuplicate bool // Check for networks with duplicate names. Network is primarily keyed based on a random ID and not on the name. Network name is strictly a user-friendly alias to the network which is uniquely identified using IO. And there is no gauranteed way to check for duplicates. Option CheckDuplicate is there to provide a best effort checking of any networks which had the same name but it is not guaranteed to catch all name collisions."
        },
        "driver": {
          "type": "string",
          "title": "Driver string"
        },
        "scope": {
          "type": "string",
          "title": "Scope string"
        },
        "enable_ipv6": {
          "type": "boolean",
          "format": "boolean",
          "title": "EnableIPv6 bool"
        },
        "ipam": {
          "$ref": "#/definitions/networkIPAM",
          "title": "IPAM *network.IPAM"
        },
        "internal": {
          "type": "boolean",
          "format": "boolean",
          "title": "Internal bool"
        },
        "attachable": {
          "type": "boolean",
          "format": "boolean",
          "title": "Attachable bool"
        },
        "ingress": {
          "type": "boolean",
          "format": "boolean",
          "title": "Ingress bool"
        },
        "config_only": {
          "type": "boolean",
          "format": "boolean",
          "title": "ConfigOnly bool"
        },
        "config_from": {
          "$ref": "#/definitions/networkConfigReference",
          "title": "ConfigFrom *network.ConfigReference"
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Options map[string]string"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Labels map[string]string"
        }
      },
      "title": "NetworkCreate is the expected body of the \"create network\" http request message\ntype NetworkCreate struct"
    },
    "mobyNetworkCreateResponse": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "ID string 0x60json:\"Id\"0x60"
        },
        "warning": {
          "type": "string",
          "title": "Warning string"
        }
      },
      "title": "NetworkCreateResponse is the response message sent by the server for network create call\ntype NetworkCreateResponse struct"
    },
    "mobyNetworkInspectOptions": {
      "type": "object",
      "properties": {
        "scope": {
          "type": "string",
          "title": "Scope string"
        },
        "verbose": {
          "type": "boolean",
          "format": "boolean",
          "title": "Verbose bool"
        }
      },
      "title": "NetworkInspectOptions holds parameters to inspect network\ntype NetworkInspectOptions struct"
    },
    "mobyNetworkListOptions": {
      "type": "object",
      "properties": {
        "filters": {
          "$ref": "#/definitions/filtersArgs",
          "title": "Filters filters.Args"
        }
      },
      "title": "NetworkListOptions holds parameters to filter the list of networks with.\ntype NetworkListOptions struct"
    },
    "mobyNetworkResource": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "Name string // Name is the requested name of the network"
        },
        "id": {
          "type": "string",
          "title": "ID string 0x60json:\"Id\"0x60 // ID uniquely identifies a network on a single machine"
        },
        "created": {
          "type": "string",
          "format": "date-time",
          "title": "Created time.Time // Created is the time the network created"
        },
        "scope": {
          "type": "string",
          "title": "Scope string // Scope describes the level at which the network exists (e.g. 'swarm' for cluster-wide or 'local' for machine level)"
        },
        "driver": {
          "type": "string",
          "title": "Driver string // Driver is the Driver name used to create the network (e.g. 'bridge', 'overlay')"
        },
        "enable_ipv6": {
          "type": "boolean",
          "format": "boolean",
          "title": "EnableIPv6 bool //EnableIPv6 represents whether to enable IPv6"
        },
        "ipam": {
          "$ref": "#/definitions/networkIPAM",
          "title": "IPAM network.IPAM // IPAM is the network's IP Address Management"
        },
        "internal": {
          "type": "boolean",
          "format": "boolean",
          "title": "Internal bool // Internal represents if the network is used internal only"
        },
        "attachable": {
          "type": "boolean",
          "format": "boolean",
          "description": "Attachable bool // Attachable represents if the global scope is manually attachable by regular containers from workers in swarm mode."
        },
        "ingress": {
          "type": "boolean",
          "format": "boolean",
          "description": "Ingress bool // Ingress indicates the network is providing the routing-mesh for the swarm custer."
        },
        "config_from": {
          "$ref": "#/definitions/networkConfigReference",
          "description": "ConfigFrom network.ConfigReference // ConfigFrom speifies the source which will provide the cconfiguration for this network."
        },
        "config_only": {
          "type": "boolean",
          "format": "boolean",
          "description": "ConfigOnly bool // ConfigOnly networks are place-holder networks for network configurations to be used by other network. ConfigOnly networks cannot be used directly to run containers or services."
        },
        "containers": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/mobyEndpointResource"
          },
          "title": "Containers map[string]EndpointResource // Containers contains endpoints belonging to the network"
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Options map[string]string // Options holds the network specific options to use for when creating the network"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Labels map[string]string // Labels holds metadata specific to the network being created"
        },
        "peers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/networkPeerInfo"
          },
          "title": "Peers []network.PeefInfo 0x60json:\",omitempty\"0x60 // List of peer nodes for an overlay network"
        },
        "services": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/networkServiceInfo"
          },
          "title": "Services map[string]network.ServiceInfo 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "NetworkResource is the body of the \"get network\" http response message\ntype NetworkResource struct"
    },
    "mobyNetworkSettings": {
      "type": "object",
      "properties": {
        "network_settings_base": {
          "$ref": "#/definitions/mobyNetworkSettingsBase",
          "title": "NetworkSettingsBase"
        },
        "default_network_settings": {
          "$ref": "#/definitions/mobyDefaultNetworkSettings",
          "title": "DefaultNetworkSettings"
        },
        "networks": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/networkEndpointSettings"
          },
          "title": "Networks map[string]*network.EndpointSettings"
        }
      },
      "title": "NetworkSettings exposes the network settings in the api\ntype NetworkSettings struct"
    },
    "mobyNetworkSettingsBase": {
      "type": "object",
      "properties": {
        "bridge": {
          "type": "string",
          "title": "Bridge string // Bridge is the Bridge name the network uses(e.g. 'docker0')"
        },
        "sandbox_id": {
          "type": "string",
          "title": "SandboxID string // SandboxID uniquely represents a container's network stack"
        },
        "hairpin_mode": {
          "type": "boolean",
          "format": "boolean",
          "title": "HairpinMode bool // HairpinMode specifies if hairpin NAT should be enabled on the virtual interface"
        },
        "link_local_ipv6_address": {
          "type": "string",
          "title": "LinkLocalIPv6Address string // LinkLocalIPv6Address is an IPv6 unicast address using the link-local prefix"
        },
        "link_local_ipv6_prefix_len": {
          "type": "integer",
          "format": "int32",
          "title": "LinkLocalIPv6PrefixLen int // LinkLocalIPv6PrefixLen is the prefix length of an IPv6 unicast address"
        },
        "ports": {
          "$ref": "#/definitions/natPortMap",
          "title": "Ports nat.PortMap // Ports is a collection of PortBinding indexed by Port"
        },
        "sandbox_key": {
          "type": "string",
          "title": "SandboxKey string //SandboxKey identifies the sandbox"
        },
        "secondary_ip_addresses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/networkAddress"
          },
          "title": "SecondaryIPAddresses []network.Address"
        },
        "secondary_ipv6_addresses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/networkAddress"
          },
          "title": "SecondaryIPv6Addresses []network.Address"
        }
      },
      "title": "NetworkSettingsBase holds basic information about networks\ntype NetworkSettingsBase"
    },
    "mobyNetworksPruneReport": {
      "type": "object",
      "properties": {
        "networks_deleted": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "NetworksDeleted []string"
        }
      },
      "title": "NetworksPruneReport contains the response for Engine API:\nPOST \"/networks/prune\"\ntype NetworksPruneReport struct"
    },
    "mobyPort": {
      "type": "object",
      "properties": {
        "ip": {
          "type": "string",
          "title": "Host IP address that the container's port is mapped to\nIP string 0x60json:\"IP,omitempty\"0x60"
        },
        "private_port": {
          "type": "integer",
          "format": "int32",
          "title": "Port on the container\nRequired: true\nPrivatePort uint16 0x60json:\"PrivatePort\"0x60"
        },
        "public_port": {
          "type": "integer",
          "format": "int32",
          "title": "Port exposed on the host\nPublicPort uint16 0x60json:\"PublicPort,omitempty\"0x60"
        },
        "type": {
          "type": "string",
          "title": "type\nRequired: true\nType string 0x60json:\"Type\"0x60"
        }
      },
      "description": "Port stores open ports info of container\ne.g. {\"PrivatePort\": 8080, \"PublicPort\": 80, \"Type\": \"tcp\"}\ntype Port struct",
      "title": "https://github.com/moby/moby/blob/master/api/types/port.go"
    },
    "mobyRootFS": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "title": "Type string"
        },
        "layers": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Layers []string 0x60json:\",omitempty\"0x60"
        },
        "base_layer": {
          "type": "string",
          "title": "BaseLayer string 0x60json:\",omitempty\"0x60"
        }
      },
      "description": "RootFS returns Image's RootFS description including the layer IDs.\ntype RootFS struct",
      "title": "https://github.com/moby/moby/blob/master/api/types/types.go"
    },
    "mobySummaryNetworkSettings": {
      "type": "object",
      "properties": {
        "networks": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/networkEndpointSettings"
          },
          "title": "Networks map[string]*network.EndpointSettings"
        }
      },
      "title": "SummaryNetworkSettings provides a summary of container's networks\nin /containers/json\ntype SummaryNetworkSettings struct"
    },
    "mountBindOptions": {
      "type": "object",
      "properties": {
        "propagation": {
          "type": "string",
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
          "title": "Name string 0x60json:\",omitempty\"0x60"
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Options map[string]string 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "Driver represents a volume driver.\ntype Driver struct"
    },
    "mountMount": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "title": "Type Type 0x60json:\",omitempty\"0x60"
        },
        "source": {
          "type": "string",
          "title": "Source specifies the name of the mount. Depending on mount type, this may be a volume name or a host path, or even ignored.\nSource is not supported for tmpfs (must be an empty value)\nSource string 0x60json:\",omitempty\"0x60"
        },
        "target": {
          "type": "string",
          "title": "Target string 0x60json:\",omitempty\"0x60"
        },
        "read_only": {
          "type": "boolean",
          "format": "boolean",
          "title": "ReadOnly bool 0x60json:\",omitempty\"0x60"
        },
        "consistency": {
          "type": "string",
          "title": "Consistency consistency 0x60json:\",omitempty\"0x60\nConsistency represents the consistency requirements of a mount. // type Consistency string"
        },
        "bind_options": {
          "$ref": "#/definitions/mountBindOptions",
          "title": "BindOptions *BindOptions 0x60json:\",omitempty\"0x60"
        },
        "volume_options": {
          "$ref": "#/definitions/mountVolumeOptions",
          "title": "VolumeOptions *VolumeOptions 0x60json:\",omitempty\"0x60"
        },
        "tmpfs_options": {
          "$ref": "#/definitions/mountTmpfsOptions",
          "title": "TmpfsOptions *TmpfsOptions 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "Mount represents a mount (volume).\ntype Mount struct"
    },
    "mountTmpfsOptions": {
      "type": "object",
      "properties": {
        "size_bytes": {
          "type": "string",
          "format": "int64",
          "description": "Size sets the size of the tmpfs, in bytes.\n\nThis will be converted to an operating system specific value\ndepending on the host. For example, on linux, it will be converted to\nuse a 'k', 'm' or 'g' syntax. BSD, though not widely supported with\ndocker, uses a straight byte value.\n\nPercentages are not supported.\nSizeBytes int84 0x60json:\",omitempty\"0x60"
        },
        "mode": {
          "type": "integer",
          "format": "int64",
          "title": "Mode of the tmpfs upon creation\nMode os.FileMode 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "TmpfsOptions defines options specific to mounts of type \"tmpfs\".\ntype TmpfsOptions struct"
    },
    "mountVolumeOptions": {
      "type": "object",
      "properties": {
        "no_copy": {
          "type": "boolean",
          "format": "boolean",
          "title": "NoCopy bool 0x60json:\",omitempty\"0x60"
        },
        "labels": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Labels map[string]string 0x60json:\",omitempty\"0x60"
        },
        "driver_config": {
          "$ref": "#/definitions/mountDriver",
          "title": "DriverConfig *Driver 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "VolumeOptions represents the options for a mount of type volume.\ntype VolumeOptions struct"
    },
    "natPortBinding": {
      "type": "object",
      "properties": {
        "host_ip": {
          "type": "string",
          "title": "HostIP is the host IP Address\nHostIP string 0x60json:\"HostIp\"0x60"
        },
        "host_port": {
          "type": "string",
          "title": "HostPort is the host port number\nHostPort string"
        }
      },
      "title": "PortBinding represents a binding between a Host IP address and a Host Port\ntype PortBinding struct"
    },
    "natPortMap": {
      "type": "object",
      "properties": {
        "internal_map": {
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
        "internal_map": {
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
    "networkAddress": {
      "type": "object",
      "properties": {
        "addr": {
          "type": "string",
          "title": "Addr string"
        },
        "prefix_len": {
          "type": "integer",
          "format": "int32",
          "title": "PrefixLen int"
        }
      },
      "title": "Address represents an IP address\ntype Address struct"
    },
    "networkConfigReference": {
      "type": "object",
      "properties": {
        "network": {
          "type": "string",
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
          "title": "IPv4Address string 0x60json:\",omitempty\"0x60"
        },
        "ipv6_address": {
          "type": "string",
          "title": "IPv6Address string 0x60json:\",omitempty\"0x60"
        },
        "link_local_ips": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "LinkLocalIPs []string 0x60json:\",omitempty\"0x60"
        }
      },
      "title": "EndpointIPAMConfig represents IPAM configurations for the endpoint\ntype EndpointIPAMConfig struct"
    },
    "networkEndpointSettings": {
      "type": "object",
      "properties": {
        "ipam_config": {
          "$ref": "#/definitions/networkEndpointIPAMConfig",
          "title": "IPAMConfig *EndpointIPAMConfig\t// Configurations"
        },
        "links": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Links []string"
        },
        "aliases": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Aliases []string"
        },
        "network_id": {
          "type": "string",
          "title": "NetworkID string // Operational data"
        },
        "endpoint_id": {
          "type": "string",
          "title": "EndpointID string"
        },
        "gateway": {
          "type": "string",
          "title": "Gateway string"
        },
        "ip_address": {
          "type": "string",
          "title": "IPAddress string"
        },
        "ip_prefix_len": {
          "type": "integer",
          "format": "int32",
          "title": "IPPrefixLen int"
        },
        "ipv6_gateway": {
          "type": "string",
          "title": "IPv6Gateway string"
        },
        "global_ipv6_address": {
          "type": "string",
          "title": "GlobalIPv6Address string"
        },
        "global_ipv6_prefix_len": {
          "type": "integer",
          "format": "int32",
          "title": "GlobalIPv6PrefixLen int"
        },
        "mac_address": {
          "type": "string",
          "title": "MacAddress string"
        },
        "driver_opts": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "DriverOpts map[string]string"
        }
      },
      "title": "EndpointSettings stores the network endpoint details\ntype EndpointSettings struct"
    },
    "networkIPAM": {
      "type": "object",
      "properties": {
        "driver": {
          "type": "string",
          "title": "Driver string"
        },
        "options": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Options map[string]string //Per network IPAM driver options"
        },
        "config": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/networkIPAMConfig"
          },
          "title": "Config []IPAMConfig"
        }
      },
      "title": "IPAM represents IP Address Management\ntype IPAM struct"
    },
    "networkIPAMConfig": {
      "type": "object",
      "properties": {
        "subnet": {
          "type": "string",
          "title": "Subnet string 0x60json:\",omitempty\"0x60"
        },
        "ip_range": {
          "type": "string",
          "title": "IPRange string 0x60json:\",omitempty\"0x60"
        },
        "gateway": {
          "type": "string",
          "title": "Gateway string 0x60json:\",omitempty\"0x60"
        },
        "aux_address": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "AuxAddress map[string]string 0x60json:\"AuxiliaryAddresses,omitempty\""
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
    "networkPeerInfo": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "Name string"
        },
        "ip": {
          "type": "string",
          "title": "IP string"
        }
      },
      "title": "PeerInfo represents one peer of an overlay network\ntype PeerInfo struct"
    },
    "networkServiceInfo": {
      "type": "object",
      "properties": {
        "vip": {
          "type": "string",
          "title": "VIP string"
        },
        "ports": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "Ports []string"
        },
        "local_lb_index": {
          "type": "integer",
          "format": "int32",
          "title": "LocalLBIndex int"
        },
        "tasks": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/networkTask"
          },
          "title": "Tasks []Task"
        }
      },
      "title": "ServiceInfo represents service parameters with the list of service's tasks\ntype ServiceInfo struct"
    },
    "networkTask": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "Name string"
        },
        "endpoint_id": {
          "type": "string",
          "title": "EndpointID string"
        },
        "endpoint_ip": {
          "type": "string",
          "title": "EndpointIP string"
        },
        "info": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Info map[string]string"
        }
      },
      "title": "Task carries the information about one backend task\ntype Task struct"
    },
    "pbDockerContainerInspectReqResp": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "key_type": {
          "$ref": "#/definitions/pbDockerContainerInspectReqRespKeyType"
        },
        "container_json": {
          "$ref": "#/definitions/mobyContainerJSON"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerContainerInspectReqRespKeyType": {
      "type": "string",
      "enum": [
        "ID",
        "NAME"
      ],
      "default": "ID"
    },
    "pbDockerContainerListReqResp": {
      "type": "object",
      "properties": {
        "container_list_options": {
          "$ref": "#/definitions/mobyContainerListOptions"
        },
        "containers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyContainer"
          }
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerContainerPruneReqResp": {
      "type": "object",
      "properties": {
        "filters": {
          "$ref": "#/definitions/filtersArgs"
        },
        "containers_prune_report": {
          "$ref": "#/definitions/mobyContainersPruneReport"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerContainerRemoveReqResp": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "key_type": {
          "$ref": "#/definitions/pbDockerContainerRemoveReqRespKeyType"
        },
        "container_remove_options": {
          "$ref": "#/definitions/mobyContainerRemoveOptions"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerContainerRemoveReqRespKeyType": {
      "type": "string",
      "enum": [
        "ID",
        "NAME"
      ],
      "default": "ID"
    },
    "pbDockerContainerRunReqResp": {
      "type": "object",
      "properties": {
        "config": {
          "$ref": "#/definitions/containerConfig"
        },
        "host_config": {
          "$ref": "#/definitions/containerHostConfig"
        },
        "networking_config": {
          "$ref": "#/definitions/networkNetworkingConfig"
        },
        "name": {
          "type": "string"
        },
        "image_pull_options": {
          "$ref": "#/definitions/mobyImagePullOptions"
        },
        "container_create_created_body": {
          "$ref": "#/definitions/containerContainerCreateCreatedBody"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerImageBuildReqResp": {
      "type": "object",
      "properties": {
        "build_context": {
          "type": "string",
          "format": "byte"
        },
        "image_build_options": {
          "$ref": "#/definitions/mobyImageBuildOptions"
        },
        "image_build_response": {
          "$ref": "#/definitions/mobyImageBuildResponse"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerImageInspectReqResp": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "ref": {
          "type": "string"
        },
        "key_type": {
          "$ref": "#/definitions/pbDockerImageInspectReqRespKeyType"
        },
        "image_inspect": {
          "$ref": "#/definitions/mobyImageInspect"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerImageInspectReqRespKeyType": {
      "type": "string",
      "enum": [
        "ID",
        "REF"
      ],
      "default": "ID"
    },
    "pbDockerImageListReqResp": {
      "type": "object",
      "properties": {
        "image_list_options": {
          "$ref": "#/definitions/mobyImageListOptions"
        },
        "image_summaries": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyImageSummary"
          }
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerImagePruneReqResp": {
      "type": "object",
      "properties": {
        "filters": {
          "$ref": "#/definitions/filtersArgs"
        },
        "images_prune_report": {
          "$ref": "#/definitions/mobyImagesPruneReport"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerImagePullReqResp": {
      "type": "object",
      "properties": {
        "ref_str": {
          "type": "string"
        },
        "image_pull_options": {
          "$ref": "#/definitions/mobyImagePullOptions"
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
          "type": "string"
        }
      }
    },
    "pbDockerImagePushReqResp": {
      "type": "object",
      "properties": {
        "ref_str": {
          "type": "string"
        },
        "image_push_options": {
          "$ref": "#/definitions/mobyImagePushOptions"
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
          "type": "string"
        }
      }
    },
    "pbDockerImageRemoveReqResp": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "ref": {
          "type": "string"
        },
        "key_type": {
          "$ref": "#/definitions/pbDockerImageRemoveReqRespKeyType"
        },
        "image_remove_options": {
          "$ref": "#/definitions/mobyImageRemoveOptions"
        },
        "image_delete_response_items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyImageDeleteResponseItem"
          }
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerImageRemoveReqRespKeyType": {
      "type": "string",
      "enum": [
        "ID",
        "REF"
      ],
      "default": "ID"
    },
    "pbDockerNetworkCreateReqResp": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
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
          "type": "string"
        }
      }
    },
    "pbDockerNetworkInspectReqResp": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "key_type": {
          "$ref": "#/definitions/pbDockerNetworkInspectReqRespKeyType"
        },
        "network_inspect_options": {
          "$ref": "#/definitions/mobyNetworkInspectOptions"
        },
        "network_resource": {
          "$ref": "#/definitions/mobyNetworkResource"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerNetworkInspectReqRespKeyType": {
      "type": "string",
      "enum": [
        "ID",
        "NAME"
      ],
      "default": "ID"
    },
    "pbDockerNetworkListReqResp": {
      "type": "object",
      "properties": {
        "network_list_options": {
          "$ref": "#/definitions/mobyNetworkListOptions"
        },
        "network_resources": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/mobyNetworkResource"
          }
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerNetworkPruneReqResp": {
      "type": "object",
      "properties": {
        "filters": {
          "$ref": "#/definitions/filtersArgs"
        },
        "networks_prune_report": {
          "$ref": "#/definitions/mobyNetworksPruneReport"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerNetworkRemoveReqResp": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "key_type": {
          "$ref": "#/definitions/pbDockerNetworkRemoveReqRespKeyType"
        },
        "state_code": {
          "type": "integer",
          "format": "int32"
        },
        "state_message": {
          "type": "string"
        }
      }
    },
    "pbDockerNetworkRemoveReqRespKeyType": {
      "type": "string",
      "enum": [
        "ID",
        "NAME"
      ],
      "default": "ID"
    },
    "protobufDuration": {
      "type": "object",
      "properties": {
        "seconds": {
          "type": "string",
          "format": "int64",
          "title": "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years"
        },
        "nanos": {
          "type": "integer",
          "format": "int32",
          "description": "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n0x60seconds0x60 field and a positive or negative 0x60nanos0x60 field. For durations\nof one second or more, a non-zero value for the 0x60nanos0x60 field must be\nof the same sign as the 0x60seconds0x60 field. Must be from -999,999,999\nto +999,999,999 inclusive."
        }
      },
      "description": "A Duration represents a signed, fixed-length span of time represented\nas a count of seconds and fractions of seconds at nanosecond\nresolution. It is independent of any calendar and concepts like \"day\"\nor \"month\". It is related to Timestamp in that the difference between\ntwo Timestamp values is a Duration and it can be added or subtracted\nfrom a Timestamp. Range is approximately +-10,000 years.\n\n# Examples\n\nExample 1: Compute Duration from two Timestamps in pseudo code.\n\n    Timestamp start = ...;\n    Timestamp end = ...;\n    Duration duration = ...;\n\n    duration.seconds = end.seconds - start.seconds;\n    duration.nanos = end.nanos - start.nanos;\n\n    if (duration.seconds \u003c 0 \u0026\u0026 duration.nanos \u003e 0) {\n      duration.seconds += 1;\n      duration.nanos -= 1000000000;\n    } else if (durations.seconds \u003e 0 \u0026\u0026 duration.nanos \u003c 0) {\n      duration.seconds -= 1;\n      duration.nanos += 1000000000;\n    }\n\nExample 2: Compute Timestamp from Timestamp + Duration in pseudo code.\n\n    Timestamp start = ...;\n    Duration duration = ...;\n    Timestamp end = ...;\n\n    end.seconds = start.seconds + duration.seconds;\n    end.nanos = start.nanos + duration.nanos;\n\n    if (end.nanos \u003c 0) {\n      end.seconds -= 1;\n      end.nanos += 1000000000;\n    } else if (end.nanos \u003e= 1000000000) {\n      end.seconds += 1;\n      end.nanos -= 1000000000;\n    }\n\nExample 3: Compute Duration from datetime.timedelta in Python.\n\n    td = datetime.timedelta(days=3, minutes=10)\n    duration = Duration()\n    duration.FromTimedelta(td)\n\n# JSON Mapping\n\nIn JSON format, the Duration type is encoded as a string rather than an\nobject, where the string ends in the suffix \"s\" (indicating seconds) and\nis preceded by the number of seconds, with nanoseconds expressed as\nfractional seconds. For example, 3 seconds with 0 nanoseconds should be\nencoded in JSON format as \"3s\", while 3 seconds and 1 nanosecond should\nbe expressed in JSON format as \"3.000000001s\", and 3 seconds and 1\nmicrosecond should be expressed in JSON format as \"3.000001s\"."
    },
    "unitsUlimit": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "Name sstring"
        },
        "hard": {
          "type": "string",
          "format": "int64",
          "title": "Hard int64"
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
