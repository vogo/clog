// Copyright 2018 The Vogo Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vogo/clog"
	"gopkg.in/natefinch/lumberjack.v2"
)

//KeyRequestID request id key
const KeyRequestID = "rid"

func main() {
	file, err := ioutil.TempFile("", "lumberjack_test_")
	if err != nil {

		fmt.Println("failed to create file")
		return
	}
	defer os.Remove(file.Name())
	fmt.Printf("temp log file %s\n", file.Name())
	clog.SetOutput(&lumberjack.Logger{
		Filename:   file.Name(),
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     1,    //days
		Compress:   true, // disabled by default
	})

	clog.Info(nil, "test clog")
	clog.Log(clog.InfoLevel, "r123", "request receive")
	clog.Log(clog.InfoLevel, "r123", "request process")
	clog.Logf(clog.InfoLevel, "r123", "request response: %s", "hello")

	clog.SetContextFommatter(func(ctx context.Context) string {
		if s, ok := ctx.Value(KeyRequestID).(string); ok {
			return s
		}
		return "--"
	})
	ctx := context.Background()
	clog.Warn(ctx, "cant get ctx info")

	var key interface{}
	key = KeyRequestID
	ctx = context.WithValue(ctx, key, "test-id")
	clog.Info(ctx, "context info")

	log, err := ioutil.ReadFile(file.Name())
	if err != nil {
		fmt.Println("failed to read log file content")
		return
	}
	content := string(log)
	fmt.Println(content)
}
