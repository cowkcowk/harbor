#!/bin/bash
#docker version: 20.10.10+
#docker-compose version: 1.18.0+
#golang version: 1.12.0+

set +e

#
# Set Colors
#

reset=$(tput sgr0)

red=$(tput setaf 1)
green=$(tput setaf 76)

success() { printf "${green}✔ %s${reset}\n" "$@"
}