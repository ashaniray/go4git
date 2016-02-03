

all: unzlib gensha1 obj2file

unzlib: unzlib.go utils.go
	go build unzlib.go utils.go

gensha1: gensha1.go utils.go
	go build gensha1.go utils.go

obj2file: obj2file.go
	go build obj2file.go utils.go

install: unzlib
	cp unzlib ${GOPATH}/bin
	cp gensha1 ${GOPATH}/bin
	cp obj2file ${GOPATH}/bin
	
clean:
	go clean
