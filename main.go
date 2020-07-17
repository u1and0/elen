/*
elen - Electric Energy converter

Usage:
```
$ elen --f 425-575 *.txt
```

-f: Filed range (multiple OK)
-o: Output to file

```
$ elen -f -75 -f 205-280 -f 425-575 -f 725-800 -f 925- \
    -o ../stats/elen.csv *.txt
```
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Created so that multiple inputs can be accecpted
type arrayField []string

const (
	// VERSION info
	VERSION = "v0.0.0"
)

var (
	// version flag
	showVersion bool
	// field code
	field arrayField
)

func (i *arrayField) String() string {
	// change this, this is just can example to satisfy the interface
	return "my string representation"
}

func (i *arrayField) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}

func main() {
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.Var(&field, "f", "Field range such as (50-100)")
	flag.Parse()

	if showVersion {
		fmt.Println("elen version:", VERSION)
		return // Exit with version info
	}

	b, err := ioutil.ReadFile("data/20200627_180505.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read a file: %s", err)
		os.Exit(1)
	}
	lines := string(b)
	fmt.Println(lines)
	fmt.Println(field)
}
