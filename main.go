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
	"strconv"
	"strings"
)

// Created so that multiple inputs can be accecpted
type arrayField []string

// configMap
type configMap map[string]string

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

	filename := "data/20200627_180505.txt"
	config, content, err := readTrace(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	fmt.Println(field)
	for _, f := range field {
		mn, mx, err := parseField(f)
		if err != nil {
			panic(err)
		}
		for i := mn; i < mx; i++ {
			fmt.Println(content[i])
		}
	}
}

// parseConfig convert first line of data to config map
func parseConfig(s string) configMap {
	config := make(configMap)
	sarray := strings.Split(s, ";")
	// snip # 20200627_180505 *RST & *CLS
	// chomp last new line
	sa := sarray[2 : len(sarray)-1]
	for _, e := range sa {
		kv := strings.Fields(e)
		fmt.Println(kv[0])
		fmt.Println(strings.Join(kv[1:], " "))
		config[kv[0]] = strings.Join(kv[1:], " ")
	}
	return config
}

// parseField convert -f option to 2 int pair
func parseField(s string) (i0, i1 int, err error) {
	ss := strings.Split(s, "-")
	i0, err = strconv.Atoi(ss[0])
	i1, err = strconv.Atoi(ss[1])
	if i0 > i1 {
		err = fmt.Errorf("Error: Must be lower %d than %d", i0, i1)
	}
	return
}

func (i *arrayField) String() string {
	// change this, this is just can example to satisfy the interface
	return "my string representation"
}

func (i *arrayField) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}

func readTrace(filename string) (config configMap, content []float64, err error) {
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
		f        float64
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
			config = parseConfig(string(line))
			isConf = false
			continue
		}
		s := string(bytes.TrimSpace(line)) // Trim prefix/surfix-whitespace
		ss := strings.Fields(s)            // Trim Middle whitespace then return []string
		f, err = strconv.ParseFloat(ss[1], 64)
		if err != nil {
			err = nil
			break
		}
		content = append(content, f)
		if !isPrefix {
			fmt.Println()
		}
	}
	return
}
