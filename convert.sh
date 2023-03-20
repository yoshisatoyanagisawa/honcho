#!/bin/sh

DIRNAME=$(cd $(dirname $0); pwd)
ROSTER="B2022.csv"
HISTORY="r3data.csv"
DATAFILE="out.json"
ADDRESS="address.csv"
CANDIDATES="duty.csv"
DONE="noduty.csv"
FINISHED="2022finished.csv"
MORNING="morning_oct.csv"
AFTERNOON="afternoon.csv"

(cd $DIRNAME; make)

"$DIRNAME/to_json" --roster="$ROSTER" --history="$HISTORY" --output="$DATAFILE"
"$DIRNAME/gen_address" --input="$DATAFILE" --output="$ADDRESS"
"$DIRNAME/gen_duty" --input="$DATAFILE" --candidates="$CANDIDATES" --done="$DONE"
"$DIRNAME/gen_patrol" --input="$DATAFILE" --finished="$FINISHED" --morning="$MORNING" --afternoon="$AFTERNOON" --start-morning="2022-10-03" --start-afternoon="2022-08-29"
