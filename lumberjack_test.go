// Copyright 2018 The Vogo Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package clog

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/natefinch/lumberjack.v2"
)

const KeyRequestID = "rid"

func TestLumberjack(t *testing.T) {
	file, err := ioutil.TempFile("", "lumberjack_test_")
	if err != nil {

		t.Error("failed to create file")
		t.FailNow()
	}
	defer os.Remove(file.Name())
	t.Logf("temp log file %s", file.Name())
	SetOutput(&lumberjack.Logger{
		Filename:   file.Name(),
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     1,    //days
		Compress:   true, // disabled by default
	})

	Info(nil, "test clog")

	SetContextFommatter(func(ctx context.Context) string {
		if s, ok := ctx.Value(KeyRequestID).(string); ok {
			return s
		}
		return "--"
	})
	ctx := context.Background()
	Warn(ctx, "cant get ctx info")

	var key interface{}
	key = KeyRequestID
	ctx = context.WithValue(ctx, key, "hello")
	Info(ctx, "log with context info")

	log, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(string(log))
}
