COMMON_FILES = \
	utils.go \
	objects.go \
	fs.go \
	commit.go \
	pack.go \
	person.go \
	repository.go \
	packindex.go

all: unzlib zlib gensha1 obj2file lstree lstype lslobj initr ppcommit lscommits \
 readpackedobject showindex

unzlib: unzlib.go common_files
	go build unzlib.go ${COMMON_FILES}

zlib: zlib.go common_files
	go build zlib.go ${COMMON_FILES}

showindex: showindex.go common_files
	go build showindex.go ${COMMON_FILES}

gensha1: gensha1.go common_files
	go build gensha1.go ${COMMON_FILES}

readpackedobject: readpackedobject.go common_files
	go build readpackedobject.go ${COMMON_FILES}

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

repostat: repostat.go common_files
	go build repostat.go ${COMMON_FILES}

common_files:${COMMON_FILES}

install: unzlib
	cp unzlib ${GOPATH}/bin
	cp zlib ${GOPATH}/bin
	cp showindex ${GOPATH}/bin
	cp gensha1 ${GOPATH}/bin
	cp readpackedobject ${GOPATH}/bin
	cp obj2file ${GOPATH}/bin
	cp lstree ${GOPATH}/bin
	cp lstype ${GOPATH}/bin
	cp lslobj ${GOPATH}/bin
	cp initr  ${GOPATH}/bin
	cp lscommits  ${GOPATH}/bin
	cp ppcommit  ${GOPATH}/bin
	cp repostat  ${GOPATH}/bin
	
clean:
	go clean
