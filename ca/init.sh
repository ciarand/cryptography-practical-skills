#!/usr/bin/env bash

# source:
# https://help.github.com/enterprise/11.10.340/admin/articles/using-self-signed-ssl-certificates/

# generate root key
openssl genrsa -out rootCA.key 2048

# sign new cert with root key
openssl req -x509 -new -nodes -key rootCA.key -days 365 -out rootCA.crt

# these sections are unnecessary as we only need one key:

# generate a key for the website
# openssl genrsa -out host.key 2048

# generate a CSR from the key
# openssl req -new -key host.key -out host.csr

# generate a signed cert from the CSR
# openssl x509 -req -in host.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out host.crt -days 365

