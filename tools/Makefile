COMMON_FILES = \
	utils.go

all: unzlib gensha1 obj2file

unzlib: unzlib.go common_files
	go build unzlib.go ${COMMON_FILES}

gensha1: gensha1.go common_files
	go build gensha1.go ${COMMON_FILES}

obj2file: obj2file.go common_files
	go build obj2file.go ${COMMON_FILES}

common_files:${COMMON_FILES}

install: unzlib
	cp unzlib ${GOPATH}/bin
	cp gensha1 ${GOPATH}/bin
	cp obj2file ${GOPATH}/bin
	
clean:
	go clean
