// Copyright 2018 The Vogo Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package clog

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var replacer = strings.NewReplacer("\r", "\\r", "\n", "\\n")

// formatOutput format output
func formatOutput(level Level, hideCallstack bool, ctxInfo, output string, callerDepth int) string {
	now := time.Now().Format("20060102 15:04:05.99999")

	output = replacer.Replace(output)

	if hideCallstack {
		return fmt.Sprintf("%-25s %-5s [%s] %s",
			now, strings.ToUpper(level.String()), ctxInfo, output)
	}
	_, file, line, ok := runtime.Caller(callerDepth)
	if !ok {
		file = "???"
		line = 0
	}
	// short file name
	file = filepath.Base(file)
	return fmt.Sprintf("%-25s %-5s [%s] %s (%s:%d)",
		now, strings.ToUpper(level.String()), ctxInfo, output, file, line)
}
