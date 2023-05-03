
PREFIX		=	${HOME}/.local
SHAREDIR	=	${PREFIX}/share/wedpad
FONTDIR		=	${PREFIX}/share/fonts
CONFDIR		=	${PREFIX}/etc
GOBIN		=	${PREFIX}/bin

GO111MODULE	=	auto

all: build

build:
	go build ./cmd/wedpad

#install:
#	go env -w GOBIN=${GOBIN}
#	go install ./cmd/wedpad
#	mkdir -p ${SHAREDIR}
#	cp -f res/*.wav ${SHAREDIR}
#	cp -f res/edpad2.png ${SHAREDIR}/edpad2.png
#	mkdir -p ${CONFDIR}
#	cp -r etc/edpad2.conf ${CONFDIR}/edpad2.conf
#	mkdir -p ${FONTDIR}
#	cp -f res/*.ttf ${FONTDIR}
#	fc-cache -f

