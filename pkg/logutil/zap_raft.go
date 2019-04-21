// Copyright 2018 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logutil

import (
	"errors"
	"fmt"

	"github.com/templexxx/etcd/raft"

	"github.com/templexxx/zap"
	"github.com/templexxx/zap/zapcore"
)

// NewRaftLogger converts "*zap.Logger" to "raft.Logger".
func NewRaftLogger(lcfg *zap.Config) (raft.Logger, error) {
	if lcfg == nil {
		return nil, errors.New("nil zap.Config")
	}
	lg, err := lcfg.Build() // to annotate caller outside of "logutil"
	if err != nil {
		return nil, err
	}
	return &zapRaftLogger{lg: lg}, nil
}

// NewRaftLoggerFromZapCore creates "raft.Logger" from "zap.Core"
// and "zapcore.WriteSyncer".
func NewRaftLoggerFromZapCore(cr zapcore.Core, syncer zapcore.WriteSyncer) raft.Logger {
	// "AddCallerSkip" to annotate caller outside of "logutil"
	lg := zap.New(cr)
	return &zapRaftLogger{lg: lg}
}

type zapRaftLogger struct {
	lg *zap.Logger
}

func (zl *zapRaftLogger) Debug(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Debug(msg)
}

func (zl *zapRaftLogger) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	zl.lg.Debug(msg)
}

func (zl *zapRaftLogger) Error(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Error(msg)
}

func (zl *zapRaftLogger) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	zl.lg.Error(msg)
}

func (zl *zapRaftLogger) Info(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Info(msg)
}

func (zl *zapRaftLogger) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	zl.lg.Info(msg)
}

func (zl *zapRaftLogger) Warning(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Warn(msg)
}

func (zl *zapRaftLogger) Warningf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	zl.lg.Warn(msg)
}

func (zl *zapRaftLogger) Fatal(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Fatal(msg)
}

func (zl *zapRaftLogger) Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	zl.lg.Fatal(msg)
}

func (zl *zapRaftLogger) Panic(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Panic(msg)
}

func (zl *zapRaftLogger) Panicf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	zl.lg.Panic(msg)
}
