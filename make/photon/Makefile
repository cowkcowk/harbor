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

# photon dockerfile
DOCKERFILEPATH=$(MAKEPATH)/photon

REGISTRY=registry
DOCKERFILEPATH_REG=$(DOCKERFILEPATH)/$(REGISTRY)
DOCKERFILENAME_REG=Dockerfile
DOCKERIMAGENAME_REG=$(IMAGENAMESPACE)/$(REGISTRY)-photon

_build_registry:
	cd $(DOCKERFILEPATH_REG) && $(DOCKERFILEPATH_REG)/builder $(REGISTRY_SRC_TAG) && cd - ;
	

define _build_base

endef

build: _build_registry
