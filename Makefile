.PHONY: default
default: ca/server

.PHONY: rootca
rootca: ca/rootCA.crt

ca/rootCA.key:
	openssl genrsa -out $@ 2048

ca/rootCA.crt: ca/rootCA.key
	openssl req -x509 -new -nodes -key ca/rootCA.key -days 364 -out ca/rootCA.crt

ca/server: ca/main.go
	go build -o ca/server $<
