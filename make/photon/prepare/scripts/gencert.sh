#! /bin/bash
set -e

if [ -z "$1" ]; then
    echo "No argument supplied set days to 365"
    DAYS=365
else
    echo "No argument supplied set days to $1"
    DAYS=$1
fi

CA_KEY="harbor_internal_ca.key"
CA_CRT="harbor_internal_ca.crt"

# CA key and certificate
if [[ ! -f $CA_KEY && ! -f $CA_CRT ]]; then
openssl req -x509 -nodes -days $DAYS -newkey rsa:4096 \
        -keyout $CA_KEY -out $CA_CRT \
        -subj "/C=CN/ST=Beijing/L=Beijing/O=VMware"
else
    echo "$CA_KEY and $CA_CRT exist, use them to generate certs"
fi

# generate proxy key and csr
openssl req -new -newkey rsa:4096 -nodes -sha256 \
        -keyout proxy.key \
        -out proxy.csr \
        -subj "/C=CN/ST=Beijing/L=Beijing/O=VMware/CN=proxy"

echo subjectAltName = DNS.1:proxy > extfile.cnf