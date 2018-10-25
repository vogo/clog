// Copyright 2018 The Vogo Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package clog

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readBuf(buf *bytes.Buffer) string {
	str := buf.String()
	buf.Reset()
	return str
}

var ctx = context.TODO()

func TestClog(t *testing.T) {
	buf := new(bytes.Buffer)
	SetOutput(buf)

	Debug(ctx, "debug log, should ignore by default")
	assert.Empty(t, readBuf(buf))

	Info(ctx, "info log, should visable")
	assert.Contains(t, readBuf(buf), "info log, should visable")

	Info(ctx, "format [%d]", 111)
	log := readBuf(buf)
	assert.Contains(t, log, "format [111]")
	t.Log(log)

	SetLevelByString("debug")
	Debug(ctx, "debug log, now it becomes visible")
	assert.Contains(t, readBuf(buf), "debug log, now it becomes visible")

	Logf(InfoLevel, "test-id", "hello %v %v", "world", "log")
	log = readBuf(buf)
	assert.Contains(t, log, "test-id")
	assert.Contains(t, log, "hello world")
	t.Log(log)

	logger = NewClog()
	logger.SetOutput(buf)

	logger.HideCallstack()
	logger.Warn(nil, "log_content")
	log = readBuf(buf)
	assert.Regexp(t, "WARN[ ]+\\[.*\\] log_content", log)
	t.Log(log)

}

func TestReplacer(t *testing.T) {
	input := `
x
y
z
`
	output := "\\nx\\ny\\nz\\n"
	assert.Equal(t, output, replacer.Replace(input))
}
