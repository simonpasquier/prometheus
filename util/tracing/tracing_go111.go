// Copyright 2018 The Prometheus Authors
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

// +build go1.11

package tracing

import (
	"context"
	"runtime/trace"
)

func newTask(ctx context.Context, name string) (context.Context, runtimeTracer) {
	if !trace.IsEnabled() {
		return ctx, noopTracer{}
	}
	return trace.NewTask(ctx, name)
}

func startRegion(ctx context.Context, regionType string) runtimeTracer {
	if !trace.IsEnabled() {
		return noopTracer{}
	}
	return trace.StartRegion(ctx, regionType)
}

func log(ctx context.Context, category, message string) {
	if trace.IsEnabled() {
		trace.Log(ctx, category, message)
	}
}
