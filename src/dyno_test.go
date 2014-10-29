package main

import (
	"testing"
)

func Test_SystemDynoState(t *testing.T) {
	systemDynoState := NewSystemDynoState()
	err := systemDynoState.ProcessUpdate("1.2.3.4", "my-app_v23_worker_10005", "starting")
	if err != nil {
		t.Errorf("error processing state update: %v", err)
	}
	t.Logf("%v", systemDynoState)
	err = systemDynoState.ProcessUpdate("1.2.3.4", "my-app_v23_worker_10005", "stopped")
	if err != nil {
		t.Errorf("error processing state update: %v", err)
	}
	t.Logf("%v", systemDynoState)
}
