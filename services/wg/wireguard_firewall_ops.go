//go:build !windows
// +build !windows

package wg

import (
	"errors"
	"fmt"
	"net/netip"
)

func (w *wireGuard) applyFirewallRulesLocked() error {
	if w.useGvisorNet || w.fwManager == nil {
		return nil
	}

	prefix, err := netip.ParsePrefix(w.ifce.GetLocalAddress())
	if err != nil {
		return errors.Join(fmt.Errorf("parse local address '%s' for firewall", w.ifce.GetLocalAddress()), err)
	}

	err = w.fwManager.ApplyRelayRules(w.ifce.GetInterfaceName(), prefix.Masked().String())
	if err != nil {
		// If iptables is not available (e.g., not installed in container), log warning and continue
		w.svcLogger.WithError(err).Warn("failed to apply firewall rules, continuing without iptables (may need manual routing configuration)")
		return nil
	}
	return nil
}

func (w *wireGuard) cleanupFirewallRulesLocked() error {
	if w.useGvisorNet || w.fwManager == nil {
		return nil
	}
	err := w.fwManager.Cleanup(w.ifce.GetInterfaceName())
	if err != nil {
		w.svcLogger.WithError(err).Warn("failed to cleanup firewall rules")
		return nil
	}
	return nil
}
