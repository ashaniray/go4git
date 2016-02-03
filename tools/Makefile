

all: unzlib gensha1

unzlib: unzlib.go
	go build unzlib.go utils.go

gensha1: gensha1.go
	go build gensha1.go utils.go

install: unzlib
	cp unzlib ${GOPATH}/bin
	cp gensha1 ${GOPATH}/bin
	
clean:
	go clean
