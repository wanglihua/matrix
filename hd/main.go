// Copyright 2013 bee authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"flag"
	"fmt"
	"github.com/revel/revel"
	"html/template"
	"os"
	"strings"
)

const version = "1.0.0"

type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string) int

	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'go help' output.
	Short template.HTML

	// Long is the long message shown in the 'go help <this-command>' output.
	Long template.HTML

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet

	// CustomFlags indicates that the command will do its own
	// flag parsing.
	CustomFlags bool
}

// Name returns the command's name: the first word in the usage line.
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(string(c.Long)))
	os.Exit(2)
}

// Runnable reports whether the command can be run; otherwise
// it is a documentation pseudo-command such as importpath.
func (c *Command) Runnable() bool {
	return c.Run != nil
}

var commands = []*Command{
	cmdRun,
	cmdRoute,
	cmdPackage,
	cmdLic,
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		os.Exit(cmdRun.Run(cmdRun, args))
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			cmd.Run(cmd, args[1:])
			return
		}
	}

	revel.ERROR.Println("unknow command: " + args[0])
}
