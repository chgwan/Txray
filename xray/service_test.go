package xray

import (
	"net"
	"testing"
	"time"
)

func TestWaitForPort_Success(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to start listener: %v", err)
	}
	defer ln.Close()
	addr := ln.Addr().String()
	deadline := time.Now().Add(2 * time.Second)
	if !waitForPort(addr, deadline) {
		t.Errorf("waitForPort returned false for a listening port")
	}
}

func TestWaitForPort_Timeout(t *testing.T) {
	// Find a free port, then close the listener before testing.
	// Use a very short deadline so the test runs quickly.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to find a free port: %v", err)
	}
	addr := ln.Addr().String()
	ln.Close() // nothing is listening on this addr now
	deadline := time.Now().Add(300 * time.Millisecond)
	if waitForPort(addr, deadline) {
		t.Errorf("waitForPort returned true for a non-listening port")
	}
}
