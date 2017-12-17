package main

import "testing"

func TestFirewallLayer_Move(t *testing.T) {
	tests := []struct {
		name          string
		firewallLayer *FirewallLayer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.firewallLayer.Move()
		})
	}
}
