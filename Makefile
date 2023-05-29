
PREFIX		=	${HOME}/.local
SHAREDIR	=	${PREFIX}/share/wedpad
CONFDIR		=	${PREFIX}/etc
GOBIN		=	${PREFIX}/bin

GO111MODULE	=	auto

all: build

build:
	go build ./cmd/wedpad

install:
	go env -w GOBIN=${GOBIN}
	go install ./cmd/wedpad
	mkdir -p ${SHAREDIR}
	cp -a res/* ${SHAREDIR}
	mkdir -p ${CONFDIR}
	cp -a conf/wedpad.conf ${CONFDIR}
	cp -a wedpad.service ${HOME}/.config/systemd/user/
	systemctl --user daemon-reload

