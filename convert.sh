#!/bin/sh

DIRNAME=$(cd $(dirname $0); pwd)
ROSTER="B2022.csv"
HISTORY="r3data.csv"
OUTPUT="out.json"

(cd $DIRNAME; make)

"$DIRNAME/to_json" --roster="$ROSTER" --history="$HISTORY" --output="$OUTPUT"
