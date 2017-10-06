package exporter

import (
	"testing"
)

func TestReapMetrics_cAdvisor(t *testing.T) {
	d := NewCAdvisorManager([]string{"http://172.17.4.50:38080"})
	if result, err := d.ReapMetrics("http://172.17.4.50:38080"); err != nil {
		t.Fail()
	} else {
		t.Logf("%q", result)
	}
}

func TestStartMetering_cAdvisor(t *testing.T) {
	d := NewCAdvisorManager([]string{"http://172.17.4.50:38080"})
	resp := d.StartMetering()
	t.Log(resp)
}
