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
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Created so that multiple inputs can be accecpted
type arrayField []string

// configMap
type configMap map[string]string

// content
type contentArray []float64

// OutRow is a output line
type OutRow struct {
	Filename string
	Center   string
	Fields   []float64
}

const (
	// VERSION info
	VERSION = "v0.0.0"
)

var (
	// version flag
	showVersion bool
	// field code
	field arrayField
	// usecol is column of using calculation
	usecol int
	// debug mode
	debug bool
)

func main() {
	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.Var(&field, "f", "Field range such as -f 50-100")
	flag.IntVar(&usecol, "c", 1, "Column of using calculation")
	flag.BoolVar(&debug, "debug", false, "Debug mode")
	flag.Parse()
	if showVersion {
		fmt.Println("elen version:", VERSION)
		return // Exit with version info
	}

	files := flag.Args()
	out := make([]OutRow, len(files))
	for i, filename := range files {
		out[i].Filename = filename
		config, content, err := readTrace(filename)
		if err != nil {
			panic(err)
		}
		if debug {
			fmt.Printf("[ CONFIG ]:%v\n", config)
			fmt.Printf("[ CONTENT ]:%v\n", content)
			fmt.Printf("[ FIELD ]:%v\n", field)
		}
		out[i].Center = config[":FREQ:CENT"]
		for _, f := range field {
			m, n, err := parseField(f)
			if err != nil {
				panic(err)
			}
			mw := content.signalBand(m, n)
			out[i].Fields = append(out[i].Fields, mw)
		}
	}
	fmt.Printf("%v", out)
}

// signalBand convert mWatt then sum between band
func (c contentArray) signalBand(m, n int) (mw float64) {
	for i := m; i <= n; i++ {
		mw += db2mw(c[i])
	}
	return
}

// parseConfig convert first line of data to config map
func parseConfig(b []byte) configMap {
	config := make(configMap)
	sarray := bytes.Split(b, []byte(";"))
	// snip # 20200627_180505 *RST & *CLS
	// chomp last new line
	sa := sarray[2 : len(sarray)-1]
	for _, e := range sa {
		kv := strings.Fields(string(e))
		config[kv[0]] = strings.Join(kv[1:], " ")
	}
	return config
}

// parseField convert -f option to 2 int pair
func parseField(s string) (i0, i1 int, err error) {
	if !strings.Contains(s, "-") {
		err = errors.New("Error: Field flag -f " + s +
			" not contains range \"-\", use int-int")
		return
	}
	ss := strings.Split(s, "-")
	i0, err = strconv.Atoi(ss[0])
	i1, err = strconv.Atoi(ss[1])
	if i0 > i1 {
		err = fmt.Errorf("Error: Must be lower %d than %d", i0, i1)
	}
	return
}

// arrayField.String sets multiple -f flag
func (i *arrayField) String() string {
	// change this, this is just can example to satisfy the interface
	return "my string representation"
}

// arrayField.Set sets multiple -f flag
func (i *arrayField) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}

// db2mw returns dB convert to mWatt
func db2mw(db float64) float64 {
	return math.Pow(10, db/10)
}

// readTrace read from a filename to `config` from first line,
// `content` from no # line.
func readTrace(filename string) (config configMap, content contentArray, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var (
		line   []byte
		isConf = true
		f      float64
	)
	for {
		line, _, err = reader.ReadLine()
		if isConf { // First line is configure
			config = parseConfig(line)
			isConf = false
			continue
		}
		if bytes.HasPrefix(line, []byte("#")) { // # <eof> then break
			return
		}
		if err == io.EOF { // if EOF then finish func
			return // might not work because HasPrefix([]byte("#"))
		}
		if err != nil { // if error at ReadLine then finish func
			return
		}
		// Trim Prefix/Surfix/Middle whitespace
		bb := bytes.Fields(bytes.TrimSpace(line))
		f, err = strconv.ParseFloat(string(bb[usecol]), 64)
		if debug {
			fmt.Printf("[ USECOL ]%v\n", usecol)
		}
		if err != nil {
			return
		}
		content = append(content, f)
	}
}
