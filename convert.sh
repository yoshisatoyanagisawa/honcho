#!/bin/sh

DIRNAME=$(cd $(dirname $0); pwd)
ROSTER="B2022.csv"
HISTORY="r3data.csv"
DATAFILE="out.json"
ADDRESS="address.csv"

(cd $DIRNAME; make)

"$DIRNAME/to_json" --roster="$ROSTER" --history="$HISTORY" --output="$DATAFILE"
"$DIRNAME/gen_address" --input="$DATAFILE" --output="$ADDRESS"
