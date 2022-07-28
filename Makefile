TARGET = to_json to_id_history gen_address gen_duty gen_patrol

all: ${TARGET}

to_json: to_json.go data_structure.go
	go build -o $@ $?

to_id_history: to_id_history.go data_structure.go
	go build -o $@ $?

gen_address: gen_address.go data_structure.go
	go build -o $@ $?

gen_duty: gen_duty.go data_structure.go
	go build -o $@ $?

gen_patrol: gen_patrol.go data_structure.go
	go build -o $@ $?

clean:
	rm -rf ${TARGET}
