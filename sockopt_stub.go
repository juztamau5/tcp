// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris,!windows

package tcp

import "time"

func (c *Conn) setKeepAliveIdleInterval(d time.Duration) error {
	return errOpNoSupport
}

func (c *Conn) setKeepAliveProbeInterval(d time.Duration) error {
	return errOpNoSupport
}

func (c *Conn) setKeepAliveProbes(probes int) error {
	return errOpNoSupport
}

func (c *Conn) setCork(on bool) error {
	return errOpNoSupport
}

func getInt(fd int, opt *sockOpt) (int, error) {
	return 0, errOpNoSupport
}

func setInt(fd int, opt *sockOpt, v int) error {
	return errOpNoSupport
}
