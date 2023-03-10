#!/bin/sh

set -e

if ! grep -q "Photon" /etc/lsb-release; then
    echo "Current OS is not Photon, skip appending ca bundle"
    exit 0
fi

ORIGINAL_LOCATION=$(dirname "$0")

if [ ! -f $ORIGINAL_LOCATION/ca-bundle.crt.original ]; then
    cp /etc/pki/tls/certs/ca-bundle.crt $ORIGINAL_LOCATION/ca-bundle.crt.original
fi

cp $ORIGINAL_LOCATION/ca-bundle.crt.original /etc/pki/tls/certs/ca-bundle.crt

# Install /etc/harbor/ssl/{component}/ca.crt to trust CA.
echo "Appending internal tls trust CA to ca-bundle ..."
for caFile in `find /etc/harbor/ssl -maxdepth 2 -name ca.crt`; do
    cat $caFile >> /etc/pki/tls/certs/ca-bundle.crt
    echo "Internal tls trust CA $caFile appended ..."
done
echo "Internal tls trust CA appending is Done."

