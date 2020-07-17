package main

import (
	"testing"
)

func TestLocateStats(t *testing.T) {
	filename := "data/20200627_180505.txt"
	actualConf, actualCont, err := readTrace(filename)
	if err != nil {
		panic(err)
	}
	// Config test
	config := "# 20200627_180505 *RST;*CLS;:INP:COUP DC;:BAND:RES 1 Hz;:AVER:COUNT 10;:SWE:POIN 1001;:FREQ:CENT 22.2 kHz;:FREQ:SPAN 2 kHz;:TRAC1:TYPE AVER;:INIT:CONT 0;:FORM REAL,32;:FORM:BORD SWAP;:INIT:IMM;:POW:ATT 0;:DISP:WIND:TRAC:Y:RLEV -30 dBm;"
	if actualConf != config {
		t.Fatalf("got: %v want: %v", actualConf, config)
	}
	// Content test
	content0 := -93.21
	actualCont0 := actualCont[0]
	if actualCont0 != content0 {
		t.Fatalf("got: %v want: %v", actualCont0, content0)
	}
	content1 := -90.08
	actualCont1 := actualCont[len(actualCont)-1]
	if actualCont1 != content1 {
		t.Fatalf("got: %v want: %v", actualCont1, content1)
	}

}
