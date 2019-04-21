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
	"fmt"
	"github.com/templexxx/zap"
	"github.com/templexxx/zap/zapcore"
	"google.golang.org/grpc/grpclog"
)

// NewGRPCLoggerV2 converts "*zap.Logger" to "grpclog.LoggerV2".
// It discards all INFO level logging in gRPC, if debug level
// is not enabled in "*zap.Logger".
func NewGRPCLoggerV2(lcfg zap.Config) (grpclog.LoggerV2, error) {
	lg, err := lcfg.Build() // to annotate caller outside of "logutil"
	if err != nil {
		return nil, err
	}
	return &zapGRPCLogger{lg: lg}, nil
}

// NewGRPCLoggerV2FromZapCore creates "grpclog.LoggerV2" from "zap.Core"
// It discards all INFO level logging in gRPC,
// if debug level is not enabled in "*zap.Logger".
func NewGRPCLoggerV2FromZapCore(cr zapcore.Core) grpclog.LoggerV2 {
	lg := zap.New(cr)
	return &zapGRPCLogger{lg: lg}
}

type zapGRPCLogger struct {
	lg    *zap.Logger
}

func (zl *zapGRPCLogger) Info(args ...interface{}) {
	if !zl.lg.Core().Enabled(zapcore.DebugLevel) {
		return
	}
	msg := fmt.Sprint(args...)
	zl.lg.Info(msg)
}

func (zl *zapGRPCLogger) Infoln(args ...interface{}) {
	if !zl.lg.Core().Enabled(zapcore.DebugLevel) {
		return
	}
	msg := fmt.Sprint(args...)
	zl.lg.Info(msg)
}

func (zl *zapGRPCLogger) Infof(format string, args ...interface{}) {
	if !zl.lg.Core().Enabled(zapcore.DebugLevel) {
		return
	}
	msg := fmt.Sprintf(format, args...)
	zl.lg.Info(msg)
}

func (zl *zapGRPCLogger) Warning(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Warn(msg)
}

func (zl *zapGRPCLogger) Warningln(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Warn(msg)
}

func (zl *zapGRPCLogger) Warningf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	zl.lg.Warn(msg)
}

func (zl *zapGRPCLogger) Error(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Error(msg)
}

func (zl *zapGRPCLogger) Errorln(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Error(msg)
}

func (zl *zapGRPCLogger) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	zl.lg.Error(msg)
}

func (zl *zapGRPCLogger) Fatal(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Fatal(msg)
}

func (zl *zapGRPCLogger) Fatalln(args ...interface{}) {
	msg := fmt.Sprint(args...)
	zl.lg.Fatal(msg)
}

func (zl *zapGRPCLogger) Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	zl.lg.Fatal(msg)
}

func (zl *zapGRPCLogger) V(l int) bool {
	// infoLog == 0
	if l <= 0 { // debug level, then we ignore info level in gRPC
		return !zl.lg.Core().Enabled(zapcore.DebugLevel)
	}
	return true
}
