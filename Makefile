SHELL := /bin/bash
BUILDPATH=$(CURDIR)
MAKEPATH=$(BUILDPATH)/make

# parameters

BUILDTARGET=build

# version prepare
# for docker image tag

BASEIMAGETAG=dev
BUILDBASETARGET=registry
BASEIMAGENAMESPACE=goharbor

#versions
REGISTRYVERSION=v2.8.0-patch-redis

# version of registry for pulling the source code
REGISTRY_SRC_TAG=v2.8.0

# dependency binaries


# docker parameters
DOCKERCMD=$(shell which docker)
DOCKERBUILD=$(DOCKERCMD) build

# go parameters

GOBUILDIMAGE=golang:1.19.4
GOBUILDPATHINCONTAINER=/harbor

GOBUILDMAKEPATH=make

# go build

# go build command

# makefile
MAKEFILEPATH_PHOTON=$(MAKEPATH)/photon

GOBUILDPATH_STANDALONE_DB_MIGRATOR=$(GOBUILDPATHINCONTAINER)/src/cmd/standalone-db-migrator

compile_standalone_db_migrator:
	@echo "compiling binary for standalone db migrator (golang image)..."
	@$(DOCKERCMD) run --rm -v $(BUILDPATH):$(GOBUILDPATHINCONTAINER) -w $(GOBUILDPATH_STANDALONE_DB_MIGRATOR) $(GOBUILDIMAGE)

build:
	make -f $(MAKEFILEPATH_PHOTON)/Makefile $(BUILDTARGET) \
	-e REGISTRYVERSION=$(REGISTRYVERSION) -e REGISTRY_SRC_TAG=$(REGISTRY_SRC_TAG)

build_base_docker:
	@for name in $(BUILDBASETARGET); do \
		echo $$name ; \
		sleep 1 ; \
		$(DOCKERBUILD) --pull --no-cache -f $(MAKEFILEPATH_PHOTON)/$$name/Dockerfile.base -t $(BASEIMAGENAMESPACE)/harbor-$$name-base:$(BASEIMAGETAG) --label base-build-date=$(date +"%Y%m%d") . ; \
	done
