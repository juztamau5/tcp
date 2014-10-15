// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcp

import "time"

const (
	sysTCP_KEEPIDLE  = 0x100
	sysTCP_KEEPINTVL = 0x200
	sysTCP_KEEPCNT   = 0x400
	sysTCP_NOPUSH    = 0x4
	sysTCP_INFO      = 0x20
)

var sockOpts = [ssoMax]sockOpt{
	ssoKeepAliveIdleInterval:  {sysTCP_KEEPIDLE, ssoTypeInt, time.Second},
	ssoKeepAliveProbeInterval: {sysTCP_KEEPINTVL, ssoTypeInt, time.Second},
	ssoKeepAliveProbes:        {sysTCP_KEEPCNT, ssoTypeInt, 0},
	ssoCork:                   {sysTCP_NOPUSH, ssoTypeInt, 0},
	ssoInfo:                   {sysTCP_INFO, ssoTypeInfo, 0},
}
