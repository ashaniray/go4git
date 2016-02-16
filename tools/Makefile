
all: unzlib zlib gensha1 obj2file lstree lstype lslobj initr ppcommit lscommits \
 readpackedobject showindex repostat pptag

unzlib: unzlib.go 
	go build unzlib.go 

zlib: zlib.go 
	go build zlib.go 

showindex: showindex.go 
	go build showindex.go 

gensha1: gensha1.go 
	go build gensha1.go 

readpackedobject: readpackedobject.go 
	go build readpackedobject.go 

lstree: lstree.go 
	go build lstree.go 

lstype: lstype.go 
	go build lstype.go 

obj2file: obj2file.go 
	go build obj2file.go 

lslobj: lslobj.go 
	go build lslobj.go 

initr: initr.go 
	go build initr.go 

lscommits: lscommits.go 
	go build lscommits.go 

ppcommit: ppcommit.go 
	go build ppcommit.go 

repostat: repostat.go 
	go build repostat.go 

pptag: pptag.go 
	go build pptag.go 


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
