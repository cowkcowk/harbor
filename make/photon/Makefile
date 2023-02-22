# Makefile for a harbor project
#
# Targets:
#
# build: 	build harbor photon images
# clean:	clean core and jobservice harbor images

# common
SHELL := /bin/bash
BUILDPATH=$(CURDIR)
MAKEPATH=$(BUILDPATH)/make
SRCPATH=./src
TOOLSPATH=$(CURDIR)/tools
SEDCMD=$(shell which sed)
WGET=$(shell which wget)
CURL=$(shell which curl)
TIMESTAMP=$(shell date +"%Y%m%d")

# docker parameters
DOCKERCMD=$(shell which docker)
DOCKERBUILD=$(DOCKERCMD) build --no-cache
DOCKERBUILD_WITH_PULL_PARA=$(DOCKERBUILD) --pull=$(PULL_BASE_FROM_DOCKERHUB)
DOCKERRMIMAGE=$(DOCKERCMD) rmi
DOCKERIMAGES=$(DOCKERCMD) images
IMAGENAMESPACE=goharbor
BASEIMAGENAMESPACE=goharbor

# photon dockerfile
DOCKERFILEPATH=$(MAKEPATH)/photon

PREPARE=prepare
DOCKERFILEPATH_PREPARE=$(DOCKERFILEPATH)/$(PREPARE)
DOCKERFILENAME_PREPARE=Dockerfile
DOCKERIMAGENAME_PREPARE=$(IMAGENAMESPACE)/$(PREPARE)

REGISTRY=registry
DOCKERFILEPATH_REG=$(DOCKERFILEPATH)/$(REGISTRY)
DOCKERFILENAME_REG=Dockerfile
DOCKERIMAGENAME_REG=$(IMAGENAMESPACE)/$(REGISTRY)-photon

_build_prepare:
	@$(call _build_base,$(PREPARE),$(DOCKERFILEPATH_PREPARE))
	@echo "building prepare container for photon..."
	@$(DOCKERBUILD_WITH_PULL_PARA) --build-arg harbor_base_image_version=$(BASEIMAGETAG) --build-arg harbor_base_namespace=$(BASEIMAGENAMESPACE) -f $(DOCKERFILEPATH_PREPARE)/$(DOCKERFILENAME_PREPARE) -t $(DOCKERIMAGENAME_PREPARE):$(VERSIONTAG) .
	@echo "Done."

_build_registry:
	cd $(DOCKERFILEPATH_REG) && $(DOCKERFILEPATH_REG)/builder $(REGISTRY_SRC_TAG) && cd - ;
	

define _build_base

endef

build: _build_prepare

cleanimage:
	@echo "cleaning image for photon..."
	
.PHONY: clean
clean: cleanimage	
