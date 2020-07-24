elen - Electric Energy converter

# Usage:
```
$ elen --f 425-575 *.txt
```

-f: Filed range (multiple OK)

```
$ elen -f -75 -f 205-280 -f 425-575 -f 725-800 -f 925- \
    -o ../stats/elen.csv *.txt
```

# Data Structure
```
# 20200627_180505 *RST;*CLS;:INP:COUP DC;:BAND:RES 1 Hz;:AVER:COUNT 10;:SWE:POIN 1001;:FREQ:CENT 22.2 kHz;:FREQ:SPAN 2 kHz;:TRAC1:TYPE AVER;:INIT:CONT 0;:FORM REAL,32;:FORM:BORD SWAP;:INIT:IMM;:POW:ATT 0;:DISP:WIND:TRAC:Y:RLEV -30 dBm;
   0     -93.21
   1     -93.97
   2     -94.93
   3     -84.87
   4     -96.31
   5     -95.23
   ...
```

0  raw : configure strings
1~ raw : data

0  columns: points
1~ columns: data


# Installation

```
$ go get github.com/u1and0/elen
```


# Licence
MIT

Copyright (c) 2020 u1and0
http://wisdommingle.com/

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
