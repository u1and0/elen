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
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
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

func main() {
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.Var(&field, "f", "Field range such as (50-100)")
	flag.Parse()

	if showVersion {
		fmt.Println("elen version:", VERSION)
		return // Exit with version info
	}

	// b, err := ioutil.ReadFile("data/20200627_180505.txt")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "cannot read a file: %s", err)
	// 	os.Exit(1)
	// }

	filename := "data/20200627_180505.txt"
	config, content, err := readTrace(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	fmt.Println(content[0], content[len(content)-1])
	// c := content[1 : len(content)-1]
	// fmt.Println(c)

	// lines := string(b)
	// lines := strings.Split(string(b), "\n")
	// for _, l := range lines {
	// 	fmt.Println(l)
	// }
	// st := strings.SplitN(string(b), "\n", 1)
	// df := dataframe.ReadCSV(content, delimiter(`\t`))
	// fmt.Println(st[0])
	fmt.Println(field)
}
func (i *arrayField) String() string {
	// change this, this is just can example to satisfy the interface
	return "my string representation"
}

func (i *arrayField) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}

func readTrace(filename string) (config string, content []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var (
		line     []byte
		isPrefix bool
		isConf   = true
	)
	for {
		line, isPrefix, err = reader.ReadLine()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return
		}
		if isConf { // First line is configure
			config = string(line)
			isConf = false
			continue
		}
		s := string(bytes.TrimSpace(line))
		ss := strings.Replace(s, " ", ",", 1) // Trim whitespace, Middle space=>,
		sss := strings.Split(ss, ",")         // split delim ,
		ssss := strings.TrimSpace(sss[1])     // Get second column
		content = append(content, ssss)
		if !isPrefix {
			fmt.Println()
		}
	}
	content = content[:len(content)-1] // chomp #<eof>
	return
}
