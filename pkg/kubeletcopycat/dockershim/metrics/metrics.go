/*
  Inspired by
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockershim/metrics/metrics.go
*/

package metrics

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// DockerOperationsKey is the key for docker operation metrics.
	DockerOperationsKey = "docker_operations"
	// DockerOperationsLatencyKey is the key for the operation latency metrics.
	DockerOperationsLatencyKey = "docker_operations_latency_microseconds"
	// DockerOperationsErrorsKey is the key for the operation error metrics.
	DockerOperationsErrorsKey = "docker_operations_errors"
	// DockerOperationsTimeoutKey is the key for the operation timoeut metrics.
	DockerOperationsTimeoutKey = "docker_operations_timeout"

	// Keep the "kubelet" subsystem for backward compatibility.
	kubeletSubsystem = "kubelet"
)

var (
	// DockerOperationsLatency collects operation latency numbers by operation
	// type.
	DockerOperationsLatency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Subsystem: kubeletSubsystem,
			Name:      DockerOperationsLatencyKey,
			Help:      "Latency in microseconds of Docker operations. Broken down by operation type.",
		},
		[]string{"operation_type"},
	)
	// DockerOperations collects operation counts by operation type.
	DockerOperations = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: kubeletSubsystem,
			Name:      DockerOperationsKey,
			Help:      "Cumulative number of Docker operations by operation type.",
		},
		[]string{"operation_type"},
	)
	// DockerOperationsErrors collects operation errors by operation
	// type.
	DockerOperationsErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: kubeletSubsystem,
			Name:      DockerOperationsErrorsKey,
			Help:      "Cumulative number of Docker operation errors by operation type.",
		},
		[]string{"operation_type"},
	)
	// DockerOperationsTimeout collects operation timeouts by operation type.
	DockerOperationsTimeout = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: kubeletSubsystem,
			Name:      DockerOperationsTimeoutKey,
			Help:      "Cumulative number of Docker operation timeout by operation type.",
		},
		[]string{"operation_type"},
	)
)

var registerMetrics sync.Once

// Register all metrics.
func Register() {
	registerMetrics.Do(func() {
		prometheus.MustRegister(DockerOperationsLatency)
		prometheus.MustRegister(DockerOperations)
		prometheus.MustRegister(DockerOperationsErrors)
		prometheus.MustRegister(DockerOperationsTimeout)
	})
}

// SinceInMicroseconds gets the time since the specified start in microseconds.
func SinceInMicroseconds(start time.Time) float64 {
	return float64(time.Since(start).Nanoseconds() / time.Microsecond.Nanoseconds())
}
