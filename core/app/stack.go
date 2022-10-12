// Copyright (C) 2020 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/google/agi/core/app/crash"
)

// Adds a signal handler for SIGQUIT to dump all go-routine stacks.
func init() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGQUIT)
	crash.Go(func() {
		for {
			<-sigchan

			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, true)]
			fmt.Println("------------------ Stack Dump: ------------------")
			fmt.Println(string(buf))
			fmt.Println("-------------------------------------------------")
		}
	})
}
