# Makefile for Harbor project
#
# Targets:
#
# all:			prepare env, compile binaries, build images and install images
# prepare: 		prepare env
# compile: 		compile core and jobservice code
#
# compile_golangimage:
#			compile from golang image
#			for example: make compile_golangimage -e GOBUILDIMAGE= \
#							golang:1.18.5
# compile_core, compile_jobservice: compile specific binary
#
# build:	build Harbor docker images from photon baseimage
#
# install:		include compile binaries, build images, prepare specific \
#				version composefile and startup Harbor instance
#
# start:		startup Harbor instance
#
# down:			shutdown Harbor instance
#
# package_online:
#				prepare online install package
#			for example: make package_online -e DEVFLAG=false\
#							REGISTRYSERVER=reg-bj.goharbor.io \
#							REGISTRYPROJECTNAME=harborrelease
#
# package_offline:
#				prepare offline install package
#
# pushimage:	push Harbor images to specific registry server
#			for example: make pushimage -e DEVFLAG=false REGISTRYUSER=admin \
#							REGISTRYPASSWORD=***** \
#							REGISTRYSERVER=reg-bj.goharbor.io/ \
#							REGISTRYPROJECTNAME=harborrelease
#				note**: need add "/" on end of REGISTRYSERVER. If not setting \
#						this value will push images directly to dockerhub.
#						 make pushimage -e DEVFLAG=false REGISTRYUSER=goharbor \
#							REGISTRYPASSWORD=***** \
#							REGISTRYPROJECTNAME=goharbor
#
# clean:        remove binary, Harbor images, specific version docker-compose \
#               file, specific version tag and online/offline install package
# cleanbinary:	remove core and jobservice binary
# cleanbaseimage:
#               remove the base images of Harbor images
# cleanimage: 	remove Harbor images
# cleandockercomposefile:
#				remove specific version docker-compose
# cleanversiontag:
#				cleanpackageremove specific version tag
# cleanpackage: remove online/offline install package
#
# other example:
#	clean specific version binaries and images:
#				make clean -e VERSIONTAG=[TAG]
#				note**: If commit new code to github, the git commit TAG will \
#				change. Better use this command clean previous images and \
#				files with specific TAG.
#   By default DEVFLAG=true, if you want to release new version of Harbor, \
#		should setting the flag to false.
#				make XXXX -e DEVFLAG=false

SHELL := /bin/bash
BUILDPATH=$(CURDIR)
MAKEPATH=$(BUILDPATH)/make
MAKE_PREPARE_PATH=$(MAKEPATH)/photon/prepare
SRCPATH=./src
TOOLSPATH=$(BUILDPATH)/tools
CORE_PATH=$(BUILDPATH)/src/core
PORTAL_PATH=$(BUILDPATH)/src/portal
CHECKENVCMD=checkenv.sh

# parameters
REGISTRYSERVER=
REGISTRYPROJECTNAME=goharbor
DEVFLAG=true
TRIVYFLAG=false
HTTPPROXY=
BUILDBIN=true
NPM_REGISTRY=https://registry.npmjs.org
BUILDTARGET=build
GEN_TLS=

# version prepare
# for docker image tag
VERSIONTAG=dev
# for base docker image tag
BUILD_BASE=true
PUSHBASEIMAGE=false
BASEIMAGETAG=dev
# BUILDBASETARGET=trivy-adapter core db jobservice log nginx portal prepare redis registry registryctl exporter
BUILDBASETARGET=registry
IMAGENAMESPACE=goharbor
BASEIMAGENAMESPACE=goharbor
# #input true/false only
PULL_BASE_FROM_DOCKERHUB=true

# for harbor package name
PKGVERSIONTAG=dev

PREPARE_VERSION_NAME=versions

#versions
REGISTRYVERSION=v2.8.3-patch-redis
TRIVYVERSION=v0.47.0
TRIVYADAPTERVERSION=v0.30.19

# version of registry for pulling the source code
REGISTRY_SRC_TAG=v2.8.3

# dependency binaries
REGISTRYURL=https://storage.googleapis.com/harbor-builds/bin/registry/release-${REGISTRYVERSION}/registry
TRIVY_DOWNLOAD_URL=https://github.com/aquasecurity/trivy/releases/download/$(TRIVYVERSION)/trivy_$(TRIVYVERSION:v%=%)_Linux-64bit.tar.gz
TRIVY_ADAPTER_DOWNLOAD_URL=https://github.com/aquasecurity/harbor-scanner-trivy/releases/download/$(TRIVYADAPTERVERSION)/harbor-scanner-trivy_$(TRIVYADAPTERVERSION:v%=%)_Linux_x86_64.tar.gz

define VERSIONS_FOR_PREPARE
VERSION_TAG: $(VERSIONTAG)
REGISTRY_VERSION: $(REGISTRYVERSION)
TRIVY_VERSION: $(TRIVYVERSION)
TRIVY_ADAPTER_VERSION: $(TRIVYADAPTERVERSION)
endef

# docker parameters
DOCKERCMD=$(shell which docker)
DOCKERBUILD=$(DOCKERCMD) build
DOCKERRMIMAGE=$(DOCKERCMD) rmi
DOCKERPULL=$(DOCKERCMD) pull
DOCKERIMAGES=$(DOCKERCMD) images
DOCKERSAVE=$(DOCKERCMD) save
DOCKERCOMPOSECMD=$(shell which docker-compose)
DOCKERTAG=$(DOCKERCMD) tag

# go parameters
GOCMD=$(shell which go)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w
GOBUILDIMAGE=golang:1.21.5
GOBUILDPATHINCONTAINER=/harbor

# go build
PKG_PATH=github.com/goharbor/harbor/src/pkg
GITCOMMIT := $(shell git rev-parse --short=8 HEAD)
RELEASEVERSION := $(shell cat VERSION)
GOFLAGS="-buildvcs=false"
GOTAGS=$(if $(GOBUILDTAGS),-tags "$(GOBUILDTAGS)",)
GOLDFLAGS=$(if $(GOBUILDLDFLAGS),--ldflags "-w -s $(GOBUILDLDFLAGS)",)
CORE_LDFLAGS=-X $(PKG_PATH)/version.GitCommit=$(GITCOMMIT) -X $(PKG_PATH)/version.ReleaseVersion=$(RELEASEVERSION)
ifneq ($(GOBUILDLDFLAGS),)
	CORE_LDFLAGS += $(GOBUILDLDFLAGS)
endif

# go build command
GOIMAGEBUILDCMD=/usr/local/go/bin/go build
GOIMAGEBUILD_COMMON=$(GOIMAGEBUILDCMD) $(GOFLAGS) ${GOTAGS} ${GOLDFLAGS}
GOIMAGEBUILD_CORE=$(GOIMAGEBUILDCMD) $(GOFLAGS) ${GOTAGS} --ldflags "-w -s $(CORE_LDFLAGS)"

GOBUILDPATH_CORE=$(GOBUILDPATHINCONTAINER)/src/core
GOBUILDPATH_JOBSERVICE=$(GOBUILDPATHINCONTAINER)/src/jobservice
GOBUILDPATH_REGISTRYCTL=$(GOBUILDPATHINCONTAINER)/src/registryctl
GOBUILDPATH_MIGRATEPATCH=$(GOBUILDPATHINCONTAINER)/src/cmd/migrate-patch
GOBUILDPATH_STANDALONE_DB_MIGRATOR=$(GOBUILDPATHINCONTAINER)/src/cmd/standalone-db-migrator
GOBUILDPATH_EXPORTER=$(GOBUILDPATHINCONTAINER)/src/cmd/exporter
GOBUILDMAKEPATH=make
GOBUILDMAKEPATH_CORE=$(GOBUILDMAKEPATH)/photon/core
GOBUILDMAKEPATH_JOBSERVICE=$(GOBUILDMAKEPATH)/photon/jobservice
GOBUILDMAKEPATH_REGISTRYCTL=$(GOBUILDMAKEPATH)/photon/registryctl
GOBUILDMAKEPATH_STANDALONE_DB_MIGRATOR=$(GOBUILDMAKEPATH)/photon/standalone-db-migrator
GOBUILDMAKEPATH_EXPORTER=$(GOBUILDMAKEPATH)/photon/exporter

# binary
CORE_BINARYPATH=$(BUILDPATH)/$(GOBUILDMAKEPATH_CORE)
CORE_BINARYNAME=harbor_core
JOBSERVICEBINARYPATH=$(BUILDPATH)/$(GOBUILDMAKEPATH_JOBSERVICE)
JOBSERVICEBINARYNAME=harbor_jobservice
REGISTRYCTLBINARYPATH=$(BUILDPATH)/$(GOBUILDMAKEPATH_REGISTRYCTL)
REGISTRYCTLBINARYNAME=harbor_registryctl
MIGRATEPATCHBINARYNAME=migrate-patch
STANDALONE_DB_MIGRATOR_BINARYPATH=$(BUILDPATH)/$(GOBUILDMAKEPATH_STANDALONE_DB_MIGRATOR)
STANDALONE_DB_MIGRATOR_BINARYNAME=migrate

# configfile
CONFIGPATH=$(MAKEPATH)
INSIDE_CONFIGPATH=/compose_location
CONFIGFILE=harbor.yml

# prepare parameters
PREPAREPATH=$(TOOLSPATH)
PREPARECMD=prepare
PREPARECMD_PARA=--conf $(INSIDE_CONFIGPATH)/$(CONFIGFILE)
ifeq ($(TRIVYFLAG), true)
	PREPARECMD_PARA+= --with-trivy
endif

# makefile
MAKEFILEPATH_PHOTON=$(MAKEPATH)/photon

# common dockerfile
DOCKERFILEPATH_COMMON=$(MAKEPATH)/common

# docker image name
DOCKER_IMAGE_NAME_PREPARE=$(IMAGENAMESPACE)/prepare
DOCKERIMAGENAME_PORTAL=$(IMAGENAMESPACE)/harbor-portal
DOCKERIMAGENAME_CORE=$(IMAGENAMESPACE)/harbor-core
DOCKERIMAGENAME_JOBSERVICE=$(IMAGENAMESPACE)/harbor-jobservice
DOCKERIMAGENAME_LOG=$(IMAGENAMESPACE)/harbor-log
DOCKERIMAGENAME_DB=$(IMAGENAMESPACE)/harbor-db
DOCKERIMAGENAME_REGCTL=$(IMAGENAMESPACE)/harbor-registryctl
DOCKERIMAGENAME_EXPORTER=$(IMAGENAMESPACE)/harbor-exporter

# docker-compose files
DOCKERCOMPOSEFILEPATH=$(MAKEPATH)
DOCKERCOMPOSEFILENAME=docker-compose.yml

SEDCMD=$(shell which sed)
SEDCMDI=$(SEDCMD) -i
ifeq ($(shell uname),Darwin)
    SEDCMDI=$(SEDCMD) -i ''
endif

# package
TARCMD=$(shell which tar)
ZIPCMD=$(shell which gzip)
DOCKERIMGFILE=harbor
HARBORPKG=harbor

# pull/push image
PUSHSCRIPTPATH=$(MAKEPATH)
PUSHSCRIPTNAME=pushimage.sh
REGISTRYUSER=
REGISTRYPASSWORD=

# cmds
DOCKERSAVE_PARA=$(DOCKER_IMAGE_NAME_PREPARE):$(VERSIONTAG) \
		$(DOCKERIMAGENAME_PORTAL):$(VERSIONTAG) \
		$(DOCKERIMAGENAME_CORE):$(VERSIONTAG) \
		$(DOCKERIMAGENAME_LOG):$(VERSIONTAG) \
		$(DOCKERIMAGENAME_DB):$(VERSIONTAG) \
		$(DOCKERIMAGENAME_JOBSERVICE):$(VERSIONTAG) \
		$(DOCKERIMAGENAME_REGCTL):$(VERSIONTAG) \
		$(DOCKERIMAGENAME_EXPORTER):$(VERSIONTAG) \
		$(IMAGENAMESPACE)/redis-photon:$(VERSIONTAG) \
		$(IMAGENAMESPACE)/nginx-photon:$(VERSIONTAG) \
		$(IMAGENAMESPACE)/registry-photon:$(VERSIONTAG)

PACKAGE_OFFLINE_PARA=-zcvf harbor-offline-installer-$(PKGVERSIONTAG).tgz \
					$(HARBORPKG)/$(DOCKERIMGFILE).$(VERSIONTAG).tar.gz \
					$(HARBORPKG)/prepare \
					$(HARBORPKG)/LICENSE $(HARBORPKG)/install.sh \
					$(HARBORPKG)/common.sh \
					$(HARBORPKG)/harbor.yml.tmpl

PACKAGE_ONLINE_PARA=-zcvf harbor-online-installer-$(PKGVERSIONTAG).tgz \
					$(HARBORPKG)/prepare \
					$(HARBORPKG)/LICENSE \
					$(HARBORPKG)/install.sh \
					$(HARBORPKG)/common.sh \
					$(HARBORPKG)/harbor.yml.tmpl

DOCKERCOMPOSE_FILE_OPT=-f $(DOCKERCOMPOSEFILEPATH)/$(DOCKERCOMPOSEFILENAME)

ifeq ($(TRIVYFLAG), true)
	DOCKERSAVE_PARA+= $(IMAGENAMESPACE)/trivy-adapter-photon:$(VERSIONTAG)
endif


RUNCONTAINER=$(DOCKERCMD) run --rm -u $(shell id -u):$(shell id -g) -v $(BUILDPATH):$(BUILDPATH) -w $(BUILDPATH)

# $1 the name of the docker image
# $2 the tag of the docker image
# $3 the command to build the docker image
define prepare_docker_image
	@if [ "$(shell ${DOCKERIMAGES} -q $(1):$(2) 2> /dev/null)" == "" ]; then \
		$(3) && echo "build $(1):$(2) successfully" || (echo "build $(1):$(2) failed" && exit 1) ; \
	fi
endef

SWAGGER_IMAGENAME=$(IMAGENAMESPACE)/swagger
SWAGGER_VERSION=v0.25.0
SWAGGER=$(RUNCONTAINER) ${SWAGGER_IMAGENAME}:${SWAGGER_VERSION}
SWAGGER_GENERATE_SERVER=${SWAGGER} generate server --template-dir=$(TOOLSPATH)/swagger/templates --exclude-main --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC
SWAGGER_IMAGE_BUILD_CMD=${DOCKERBUILD} -f ${TOOLSPATH}/swagger/Dockerfile --build-arg GOLANG=${GOBUILDIMAGE} --build-arg SWAGGER_VERSION=${SWAGGER_VERSION} -t ${SWAGGER_IMAGENAME}:$(SWAGGER_VERSION) .

# $1 the path of swagger spec
# $2 the path of base directory for generating the files
# $3 the name of the application
define swagger_generate_server
	@echo "generate all the files for API from $(1)"
	@rm -rf $(2)/{models,restapi}
	@mkdir -p $(2)
	@$(SWAGGER_GENERATE_SERVER) -f $(1) -A $(3) --target $(2)
endef

gen_apis:
	$(call prepare_docker_image,${SWAGGER_IMAGENAME},${SWAGGER_VERSION},${SWAGGER_IMAGE_BUILD_CMD})
	$(call swagger_generate_server,api/v2.0/swagger.yaml,src/server/v2.0,harbor)

compile_standalone_db_migrator:
	@echo "compiling binary for standalone db migrator (golang image)..."
	@$(DOCKERCMD) run --rm -v $(BUILDPATH):$(GOBUILDPATHINCONTAINER) -w $(GOBUILDPATH_STANDALONE_DB_MIGRATOR) $(GOBUILDIMAGE)

build:
	make -f $(MAKEFILEPATH_PHOTON)/Makefile $(BUILDTARGET) \
	-e REGISTRYVERSION=$(REGISTRYVERSION) -e REGISTRY_SRC_TAG=$(REGISTRY_SRC_TAG)

build_base_docker:
	@for name in $(BUILDBASETARGET); do \
		echo $$name ; \
		sleep 30 ; \
		$(DOCKERBUILD) --pull --no-cache -f $(MAKEFILEPATH_PHOTON)/$$name/Dockerfile.base -t $(BASEIMAGENAMESPACE)/harbor-$$name-base:$(BASEIMAGETAG) --label base-build-date=$(date +"%Y%m%d") . ; \
	done

swagger_client:
	@echo "Generate swagger client"
	wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/4.3.1/openapi-generator-cli-4.3.1.jar -O openapi-generator-cli.jar
	rm -rf harborclient
	mkdir

clean:
	@echo "  make cleanall:		remove binary, Harbor images, specific version docker-compose"
