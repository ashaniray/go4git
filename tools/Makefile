COMMON_FILES = \
	utils.go \
	objects.go \
	fs.go \
	commit.go

all: unzlib zlib gensha1 obj2file lstree lstype lslobj initr ppcommit lscommits

unzlib: unzlib.go common_files
	go build unzlib.go ${COMMON_FILES}

zlib: zlib.go common_files
	go build zlib.go ${COMMON_FILES}

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

initr: initr.go common_files
	go build initr.go ${COMMON_FILES}

lscommits: lscommits.go common_files
	go build lscommits.go ${COMMON_FILES}

ppcommit: ppcommit.go common_files
	go build ppcommit.go ${COMMON_FILES}

common_files:${COMMON_FILES}

install: unzlib
	cp unzlib ${GOPATH}/bin
	cp zlib ${GOPATH}/bin
	cp gensha1 ${GOPATH}/bin
	cp obj2file ${GOPATH}/bin
	cp lstree ${GOPATH}/bin
	cp lstype ${GOPATH}/bin
	cp lslobj ${GOPATH}/bin
	cp initr  ${GOPATH}/bin
	cp lscommits  ${GOPATH}/bin
	cp ppcommit  ${GOPATH}/bin
	
clean:
	go clean
