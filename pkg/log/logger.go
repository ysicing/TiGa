// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package log

import (
	"io"

	log "github.com/loft-sh/utils/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/ysicing/tiga/pkg/log/survey"
)

// logFunctionType type
type logFunctionType uint32

const (
	fatalFn logFunctionType = iota
	errorFn
	warnFn
	infoFn
	debugFn
	doneFn
)

// Logger defines the devspace common logging interface
type Logger interface {
	log.Logger
	// WithLevel creates a new logger with the given level
	WithLevel(level logrus.Level) Logger
	Question(params *survey.QuestionOptions) (string, error)
	ErrorStreamOnly() Logger
	WithPrefix(prefix string) Logger
	WithPrefixColor(prefix, color string) Logger
	WithSink(sink Logger) Logger
	AddSink(sink Logger)

	Writer(level logrus.Level, raw bool) io.WriteCloser
	WriteString(level logrus.Level, message string)
}
