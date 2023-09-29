// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package context

import (
	ctx2 "context"

	"github.com/ysicing/tiga/pkg/log"
)

type Context interface {
	Context() ctx2.Context
	Log() log.Logger
}

func NewContext(ctx ctx2.Context, log log.Logger) Context {
	return &context{
		context: ctx,
		log:     log,
	}
}

type context struct {
	context ctx2.Context
	log     log.Logger
}

func (c *context) Context() ctx2.Context {
	return c.context
}

func (c *context) Log() log.Logger {
	return c.log
}
