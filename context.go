// Copyright 2018 The Vogo Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package clog

import "context"

//ContextFormatter format a context to string
type ContextFormatter func(ctx context.Context) string

//DefaultContextFormatter formatter
func DefaultContextFormatter(ctx context.Context) string {
	return "-"
}
