release = no

svn = yes
SVN_VER = 

GOROOT=/usr/lib/golang
GOBIN=$(GOROOT)/bin

BUILD_PATH=./build

DEBUG_NAME = fcas_server_d
RELEASE_NAME = fcas_server

PROJECT_NAME = ${DEBUG_NAME}

BUILD_FLAGS = -gcflags 'all=-N -l'

BUILD_VERSION=$(shell cat ../BUILD_VERSION)
BUILD_TIME=$(shell date +"%Y%m%d%H%M%S")

RELEASE_FLAG=""

.PHONY:build 


ifeq ($(release), yes)
RELEASE_FLAG = "-s -w"
PROJECT_NAME = ${RELEASE_NAME}
endif

LDFLAGS +=-ldflags "-X main.buildTime=${BUILD_TIME} -X main.buildVer=${BUILD_VERSION} "${RELEASE_FLAG}""


ifeq ($(svn), yes)
SVN_VER = $(shell svn info ./|grep 'Last Changed Rev'|cut -c 19-)
endif

build:
	@echo "compile begin"
	@echo ${SVN_VER} > version
	@echo $(GOBIN)/go build ${BUILD_FLAGS} ${LDFLAGS} -o ${PROJECT_NAME} ${MAIN_PATH}
	@$(GOBIN)/go build ${BUILD_FLAGS} ${LDFLAGS} -o ${PROJECT_NAME}
	@if [ ! -d "${BUILD_PATH}" ]; then mkdir ${BUILD_PATH}; fi
	@cp -f ${PROJECT_NAME} ${BUILD_PATH}
	@echo "compile end"
    

generate:
	@$(GOBIN)/go generate

clean:
	@$(GOBIN)/go clean
	rm -rf ${PROJECT_NAME}

rpm:
	@echo "package begin"
	@rm -rf ${BUILD_PATH}/*.rpm
	@nfpm package -p rpm -t ${BUILD_PATH}
	@echo "package end"




