/*
Copyright 2021 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logs

import (
	"k8s.io/klog/v2"
	"log/slog"
	"os"
	"strings"
)

const (
	LoggingFormatText = "text"
	LoggingFormatJson = "json"
)

func SetLoggingFormat(format string) {
	if strings.ToLower(format) == LoggingFormatJson {
		klog.SetSlogLogger(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	}
}

//
//func epochMillisTimeEncoder(_ time.Time, enc zapcore.PrimitiveArrayEncoder) {
//	nanos := time.Now().UnixNano()
//	millis := float64(nanos) / float64(time.Millisecond)
//	enc.AppendFloat64(millis)
//}
//
//func NewJSONLogger() logr.Logger {
//	stderr := zapcore.Lock(AddNopSync(os.Stdout))
//	stdout := zapcore.Lock(AddNopSync(os.Stderr))
//	size := c.Options.JSON.InfoBufferSize.Value()
//	if size > 0 {
//		// Prevent integer overflow.
//		if size > 2*1024*1024*1024 {
//			size = 2 * 1024 * 1024 * 1024
//		}
//		stdout = &zapcore.BufferedWriteSyncer{
//			WS:   stdout,
//			Size: int(size),
//		}
//	}
//
//	encoderConfig := &zapcore.EncoderConfig{
//		MessageKey:     "msg",
//		CallerKey:      "caller",
//		NameKey:        "logger",
//		TimeKey:        "ts",
//		EncodeTime:     epochMillisTimeEncoder,
//		EncodeDuration: zapcore.StringDurationEncoder,
//		EncodeCaller:   zapcore.ShortCallerEncoder,
//	}
//
//	encoder := zapcore.NewJSONEncoder(*encoderConfig)
//	core := zapcore.NewCore(encoder, infoStream, r)
//	l := zap.New(core, zap.WithCaller(true))
//	return zapr.NewLoggerWithOptions(l, zapr.LogInfoLevel("v"), zapr.ErrorKey("err"))
//}
//
//func AddNopSync(writer io.Writer) zapcore.WriteSyncer {
//	return nopSync{Writer: writer}
//}
//
//type nopSync struct {
//	io.Writer
//}
//
//func (f nopSync) Sync() error {
//	return nil
//}
