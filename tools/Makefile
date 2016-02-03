

all: unzlib gensha1

unzlib: unzlib.go utils.go
	go build unzlib.go utils.go

gensha1: gensha1.go utils.go
	go build gensha1.go utils.go

install: unzlib
	cp unzlib ${GOPATH}/bin
	cp gensha1 ${GOPATH}/bin
	
clean:
	go clean
