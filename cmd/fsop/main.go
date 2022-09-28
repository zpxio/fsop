/*
 * Copyright 2022 zpxio (Jeff Sharpe)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"github.com/apex/log"
	"github.com/mkideal/cli"
	"github.com/zpxio/fsop/pkg/scan"
	"os"
)

type argT struct {
	cli.Helper
	Directories []string `cli:"d" usage:"Scan one or more directories"`
}

func main() {
	log.Debugf("Scanning arguments...")
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		log.Infof("Starting up.")

		ctx.JSONln(ctx.Argv())
		argv := ctx.Argv().(*argT)

		log.Infof("Creating scanner: %s", argv.Directories)
		scanner := scan.CreateScanner(argv.Directories...)

		log.Infof("Scanning...")
		scanErr := scanner.ScanTo()
		if scanErr != nil {
			return fmt.Errorf("error while scanning directory: %s", scanErr)
		}
		log.Infof("Scan Complete.")

		log.Infof("Shutting down.")
		return nil
	}))
}
