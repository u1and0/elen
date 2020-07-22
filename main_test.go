package main

import (
	"testing"
)

func Test_readTrace(t *testing.T) {
	filename := "data/20200627_180505.txt"
	actualConf, _, err := readTrace(filename)
	// actualConf, actualCont, err := readTrace(filename)
	if err != nil {
		panic(err)
	}

	// Config test
	config := map[string]string{
		":INP:COUP":              "DC",
		":BAND:RES":              "1 Hz",
		":AVER:COUNT":            "10",
		":SWE:POIN":              "1001",
		":FREQ:CENT":             "22.2 kHz",
		":FREQ:SPAN":             "2 kHz",
		":TRAC1:TYPE":            "AVER",
		":INIT:CONT":             "0",
		":FORM":                  "REAL,32",
		":FORM:BORD":             "SWAP",
		":INIT:IMM":              "",
		":POW:ATT":               "0",
		":DISP:WIND:TRAC:Y:RLEV": "-30 dBm",
	}
	for k, v := range actualConf {
		if config[k] != v {
			t.Fatalf("got: %v want: %v", v, config[k])
		}
	}

	/* content test cannot run
	dump [0 1 2 3 ...]
	// Content test
	content0 := -93.21
	actualCont0 := actualCont[0]
	if actualCont0 != content0 {
		t.Fatalf("got: %v want: %v\ndump all: %v", actualCont0, content0, actualCont)
	}
	content1 := -90.08
	actualCont1 := actualCont[len(actualCont)-1]
	if actualCont1 != content1 {
		t.Fatalf("got: %v want: %v\ndump all: %v", actualCont1, content1, actualCont)
	}
	*/
}
