SHELL := /bin/bash
BUILDPATH=$(CURDIR)
MAKEPATH=$(BUILDPATH)/make

# version prepare
# for docker image tag

BASEIMAGETAG=dev
BUILDBASETARGET=db
BASEIMAGENAMESPACE=goharbor

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

build_base_docker:
	@for name in $(BUILDBASETARGET); do \
		echo $$name ; \
		$(DOCKERBUILD) --pull --no-cache -f $(MAKEFILEPATH_PHOTON)/$$name/Dockerfile.base -t $(BASEIMAGENAMESPACE)/harbor-$$name-base:$(BASEIMAGETAG) --label base-build-date=$(date +"%Y%m%d") . ; \
	done
