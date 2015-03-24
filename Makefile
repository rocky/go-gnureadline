# Supplying a make file prevents travis from trying to build
# using go get -t -v ../... which tries to build
# demo code.
.PHONY: all test

all:
	go build

#: Same as "check"
test: check

#: Run all tests (quick and interpreter)
check:
	go test -v
