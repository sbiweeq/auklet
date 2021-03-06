// Copyright (c) 2016-2018 iQIYI.com.  All rights reserved.
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
// 

package command

import (
	"log"
	"strings"
)

type StartCommand struct {
	Logger *log.Logger
}

func (c *StartCommand) Help() string {
	helpText := `
usage: auklet start [object | pack-auditor]

  start a server/daemon
`
	return strings.TrimSpace(helpText)
}

func (c *StartCommand) Run(args []string) int {
	if len(args) < 1 {
		c.Logger.Println(c.Help())
		return EXIT_USAGE
	}

	if err := startServer(args[0], args[1:]...); err != nil {
		c.Logger.Printf("unable to start %s: %v", args[0], err)
		return EXIT_START
	}

	c.Logger.Printf("%s started", args[0])
	return EXIT_OK
}

func (c *StartCommand) Synopsis() string {
	return "start a server/daemon"
}
