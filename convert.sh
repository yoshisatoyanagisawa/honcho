#!/bin/sh

DIRNAME=$(cd $(dirname $0); pwd)
# input files.
ROSTER="B2022.csv"
HISTORY="r3data.csv"
FINISHED="2022finished.csv"

# intermediate file.
DATAFILE="out.json"

# output files.
ADDRESS="address.csv"
CANDIDATES="duty.csv"
DONE="noduty.csv"
MORNING="morning_oct.csv"
AFTERNOON="afternoon.csv"
NEW_ROSTER="roster.csv"

(cd $DIRNAME; make)

rm -f "$DATAFILE" "$ADDRESS" "$CANDIDATES" "$MORNING" "$AFTERNOON" "$NEW_ROSTER"
"$DIRNAME/to_json" --roster="$ROSTER" --history="$HISTORY" --output="$DATAFILE" --finished="$FINISHED"
"$DIRNAME/gen_address" --input="$DATAFILE" --output="$ADDRESS"
"$DIRNAME/gen_duty" --input="$DATAFILE" --candidates="$CANDIDATES" --done="$DONE"
"$DIRNAME/gen_patrol" --input="$DATAFILE" --morning="$MORNING" --afternoon="$AFTERNOON" --start-morning="2022-10-03" --start-afternoon="2022-08-29"
"$DIRNAME/gen_roster" --input="$DATAFILE" --output="$NEW_ROSTER"
