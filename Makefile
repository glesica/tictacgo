GOC=go build
EXE=tictacgo
EXEDEST=/usr/bin

.PHONY: build get-deps install uninstall

build:
	${GOC}

get-deps:
	dep ensure

install:
	install ${EXE} ${EXEDEST}/

uninstall:
	rm ${EXEDEST}/${EXE}