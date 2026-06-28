/*
Copyright (c) Tobias Schäfer. All rights reserved.
Licensed under the Apache-2.0 license, see LICENSE in the project root for details.
*/
package version

import (
	"fmt"
	"os"
)

var (
	GitCommit, Version string
)

func Release() string {
	if Version == "" {
		Version = "dev"
	}

	return Version
}

func Commit() string {
	return GitCommit
}

func Banner() string {
	return `
     _               _   _
  __| | _____  _____| |_| |
 / _  |/ _ \ \/ / __| __| |
| (_| |  __/>  < (__| |_| |
 \__,_|\___/_/\_\___|\__|_|
`
}

func Print() {
	no_color := os.Getenv("NO_COLOR")
	if no_color != "" {
		fmt.Printf("%s\n", Banner())
	} else {
		fmt.Printf("\033[34m%s\033[0m\n", Banner())
	}
	fmt.Printf("Release: %s\n", Release())
	fmt.Printf("Commit:  %s\n", Commit())
}
