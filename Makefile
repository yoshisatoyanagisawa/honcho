TARGET = to_json gen_address gen_duty gen_patrol gen_roster
LIBS = ioutil.go datautil.go data_structure.go

all: ${TARGET}

to_json: to_json.go ${LIBS}
	go build -o $@ $?

gen_address: gen_address.go ${LIBS}
	go build -o $@ $?

gen_duty: gen_duty.go ${LIBS}
	go build -o $@ $?

gen_patrol: gen_patrol.go ${LIBS}
	go build -o $@ $?

gen_roster: gen_roster.go ${LIBS}
	go build -o $@ $?

clean:
	rm -rf ${TARGET}
