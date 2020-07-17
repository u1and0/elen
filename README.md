elen - Electric Energy converter

# Usage:
```
$ elen --f 425-575 *.txt
```

-f: Filed range (multiple OK)
-o: Output to file

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

# Licence
