ipfixcat
========

`ipfixcat` is a utility to parse and print an IPFIX stream, as defined by
RFC 5101.

Installation
------------

    $ go install github.com/calmh/ipfixcat

Examples
--------

Parse a UDP IPFIX stream.

    $ socat udp-recv:4739 stdout | ipfixcat
    --- 1374477175 49836
    # Data Set with unknown template
    --- 1374477175 49836
    # Data Set with unknown template
    # Received Template Set 10299
    # Received Template Set 49836
    --- 1374477180 49836
    0.12: [172 16 32 15]
    15397.32780: [0 0 0 0]
    15397.32771: [0 0 0 0 0 0 0 145]
    15397.32772: [0 0 0 0 0 0 1 54]
    15397.32786: []
    15397.32769: [66 105 116 84 111 114 114 101 110 116 32 75 82 80 67]
    15397.32796: []
    0.8: [14 39 135 171]
    --- 1374477180 49836
    0.12: [172 16 32 15]
    15397.32780: [0 0 0 0]
    15397.32771: [0 0 0 0 0 0 0 166]
    15397.32772: [0 0 0 0 0 0 0 0]
    15397.32786: []
    15397.32769: [66 105 116 84 111 114 114 101 110 116 32 75 82 80 67]
    15397.32796: []
    0.8: [195 24 238 68]
    ...

Don't attempt to use netcat (`nc`). Almost all distributed versions are
broken and truncate UDP packets at 1024 bytes.
