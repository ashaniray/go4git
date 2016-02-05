COMMON_FILES = \
	utils.go \
	objects.go \
	fs.go

all: unzlib gensha1 obj2file lstree lstype lslobj

unzlib: unzlib.go common_files
	go build unzlib.go ${COMMON_FILES}

gensha1: gensha1.go common_files
	go build gensha1.go ${COMMON_FILES}

lstree: lstree.go common_files
	go build lstree.go ${COMMON_FILES}

lstype: lstype.go common_files
	go build lstype.go ${COMMON_FILES}

obj2file: obj2file.go common_files
	go build obj2file.go ${COMMON_FILES}

lslobj: lslobj.go common_files
	go build lslobj.go ${COMMON_FILES}

common_files:${COMMON_FILES}

install: unzlib
	cp unzlib ${GOPATH}/bin
	cp gensha1 ${GOPATH}/bin
	cp obj2file ${GOPATH}/bin
	cp lstree ${GOPATH}/bin
	cp lstype ${GOPATH}/bin
	cp lslobj ${GOPATH}/bin
	
clean:
	go clean
