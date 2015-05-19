= Crypography final (CPSC 433)

== Group name
The Fighting Mongooses of Cryptography

== Document
The document is in PDF form in the `document.pdf` file.

== Running the Go code
Assuming you have a working Go toolchain in your `$PATH` (tested on Ubuntu
14.04 after running `sudo apt-get install -y golang`) you can run:

----
go build -o server main.go
----

And then `./server -help` for further instruction.

== Authors

Ciaran Downey <ciarand@csu.fullerton.edu>;
Kourun Sok <sokkourun@gmail.com>;
Stratton Aguilar <strataguilar@csu.fullerton.edu>;
Brian Brick <brianbrick@csu.fullerton.edu>;
Andrew Tsay <atsay714@gmail.com>

== Contributions

Kourun: SSH, OpenVPN, IPsec, documentation +
Ciaran: SSH settings, root CA, HTTPS server, documentation +
Strat:  Testing on OpenVPN and IPsec, documentation +
Brian:  IPsec writeup, documentation +
Andrew: VPN writeup

== Files

- client_config.ovpn:          OpenVPN client config file
- crypto_ssh_example_rsa:      Example RSA private key
- crypto_ssh_example_rsa.pub:  Example RSA public key
- host.crt:                    certificate for the HTTPS site
- host.csr:                    certificate request for the HTTPS site
- host.key:                    private key for the HTTPS site
- ipsec.psk:                   Private Shared Key for IPsec
- main.go:                     HTTPS server
- rootCA.crt:                  certificate for our root CA
- rootCA.key:                  key for our root CA
- rootCA.srl:                  serial number file for our root CA
- server.conf:                 OpenVPN server config file
- sshd_config:                 OpenSSH config file
- vars:                        vars file for easy-rsa

