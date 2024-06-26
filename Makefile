LANG=en_US.UTF-8
SHELL=/bin/bash
.SHELLFLAGS=--norc --noprofile -e -u -o pipefail -c

clean:
	find ./subs -type f -name "*.vtt" -delete
run:
	go build test_subs.go
	chmod +x ./test_subs
	./test_subs
build:
	find ./subs -type f -name "*.vtt" -delete
	go build test_subs.go

