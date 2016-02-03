

all: unzlib gen-sha1

unzlib: unzlib.go
	go build unzlib.go utils.go

gen-sha1: gen-sha1.go
	go build gen-sha1.go utils.go

install: unzlib
	cp unzlib ${GOPATH}/bin
	cp gen-sha1 ${GOPATH}/bin
	
clean:
	go clean
