TARGET = to_json gen_address gen_duty gen_patrol

all: ${TARGET}

to_json: to_json.go data_structure.go
	go build -o $@ $?

gen_address: gen_address.go data_structure.go
	go build -o $@ $?

gen_duty: gen_duty.go data_structure.go
	go build -o $@ $?

gen_patrol: gen_patrol.go data_structure.go
	go build -o $@ $?

clean:
	rm -rf ${TARGET}
