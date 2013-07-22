ipfixcat
========

`ipfixcat` is a utility to parse and print an IPFIX stream, as defined by
RFC 5101.

Installation
------------

    $ go install github.com/calmh/ipfixcat

Output
------

The output format is JSON with one object per line. Each object has
fields `exportTime` (UNIX epoch seconds), `templateId` and `elements`.
The latter is a dict containing the information elements as `field: value`.

Unless a dictionary is loaded, the field name will be `F<fieldId>` for
standard fields or `V<vendorId>.<fieldId>` for vendor specific fields.
Values are byte arrays unless converted by a dictionary (the IPFIX
stream contains no type information).

Examples
--------

Parse a UDP IPFIX stream.

    $ socat udp-recv:4739 stdout | ipfixcat
    {"elements":{},"exportTime":1374483650,"templateId":49836}
    {"elements":{},"exportTime":1374483650,"templateId":49836}
    {"elements":{},"exportTime":1374483655,"templateId":49836}
    {"elements":{"F12":[172,16,32,15],"F8":[59,91,233,213],"V15397.1":[66,105,116,84,111,114,114,101,110,116,32,75,82,80,67],"V15397.12":[0,0,0,0],"V15397.18":[],"V15397.28":[],"V15397.3":[0,0,0,0,0,0,0,145],"V15397.4":[0,0,0,0,0,0,1,54]},"exportTime":1374483660,"templateId":49836}
    {"elements":{"F12":[172,16,32,15],"F8":[223,204,77,117],"V15397.1":[66,105,116,84,111,114,114,101,110,116,32,75,82,80,67],"V15397.12":[0,0,0,0],"V15397.18":[],"V15397.28":[],"V15397.3":[0,0,0,0,0,0,0,145],"V15397.4":[0,0,0,0,0,0,1,54]},"exportTime":1374483660,"templateId":49836}
    ...

Don't attempt to use netcat (`nc`). Almost all distributed versions are
broken and truncate UDP packets at 1024 bytes.
