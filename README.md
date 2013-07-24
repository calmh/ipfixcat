ipfixcat
========

`ipfixcat` is a utility to parse and print an IPFIX stream, as defined by
RFC 5101. It's also the minimal demo of how to use the
github.com/calmh/ipfix package.

Installation
------------

Grab a binary release from https://github.com/calmh/ipfixcat/releases.

You can also build from source. Make sure you have Go 1.1 installed. See
http://golang.org/doc/install.

    $ go install github.com/calmh/ipfixcat

Output
------

The output format is JSON with one object per line. Each object has
fields `exportTime` (UNIX epoch seconds), `templateId` and `elements`.
The latter is a dict containing the information elements as `field: value`.

Standard fields are interpreted with name and value type. Vendor fields
display as `F[vendor,field]` with a byte array as value. A custom
dictionary can be loaded to support vendor fields; see
`procera-fields.ini` included.

Examples
--------

Parse a UDP IPFIX stream. Note that it might take a while to start
displaying datasets, because we need to receive the periodically sent
template sets first in order to be able to parse them.

    $ socat udp-recv:4739 stdout | ipfixcat
    {"elements":{"F[15397.12]":[0,0,0,0],"F[15397.18]":[],"F[15397.1]":[66,105,116,84,111,114,114,101,110,116,32,75,82,...
    {"elements":{"F[15397.12]":[0,0,0,0],"F[15397.18]":[],"F[15397.1]":[68,114,111,112,98,111,120,32,76,65,78,32,115,12...
    ...

Use a custom dictionary to interpret vendor fields.

    $ socat udp-recv:4739 stdout | ipfixcat -dict procera-fields.ini 
    {"elements":{"destinationIPv4Address":"172.16.32.15","proceraExternalRtt":47,"proceraIncomingOctets":146,"proeraOut...
    {"elements":{"destinationIPv4Address":"172.16.32.15","proceraExternalRtt":3,"proceraIncomingOctets":140,"proceraOut...
    {"elements":{"destinationIPv4Address":"172.16.32.15","proceraExternalRtt":4,"proceraIncomingOctets":642,"proceraOut...
    ...

Don't attempt to use netcat (`nc`) for reading UDP streams. Almost all
distributed versions are broken and truncate UDP packets at 1024 bytes.

License
-------

MIT
