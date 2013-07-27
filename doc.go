/*
`ipfixcat` is a utility to parse and print an IPFIX stream, as defined
by RFC 5101. It's also the minimal demo of how to use the
github.com/calmh/ipfix package.

Installation

Grab a binary release from https://github.com/calmh/ipfixcat/releases.

You can also build from source. Make sure you have Go 1.1 installed. See
http://golang.org/doc/install.

    $ go install github.com/calmh/ipfixcat

Output

The output format is JSON with one object per line. Each object has
fields `exportTime` (UNIX epoch seconds), `templateId` and `elements`.
The latter is an array containing the information elements in the same
order as received by the exporter.

Each information element has the fields `name`, `enterprise`, `field`,
`value` and `rawvalue`. For vendor fields that are not described by a
user dictionary, `name` and `value` will be empty and `rawvalue`
contains a byte array. For fully understood fields, `value` contains the
parsed value and `rawvalue` is empty.

There are some statistics that can be enabled as well, see
`ipfixcat -help` for more information.

Examples

Parse a UDP IPFIX stream, using a custom dictionary to interpret vendor
fields. Note that it might take a while to start displaying datasets,
because we need to receive the periodically sent template sets first in
order to be able to parse them.

    $ socat udp-recv:4739 stdout | ipfixcat -dict procera-fields.ini 
    {"exportTime":1374745620,"templateId":49836,"fields":[{"name":"destinationIPv4Address","field":12,"value":"194.153....
    {"exportTime":1374745620,"templateId":10299,"fields":[{"name":"destinationIPv6Address","field":28,"value":"2001:470...
    {"exportTime":1374745620,"templateId":10299,"fields":[{"name":"destinationIPv6Address","field":28,"value":"2001:470...
    ...

Don't attempt to use netcat (`nc`) for reading UDP streams. Almost all
distributed versions are broken and truncate UDP packets at 1024 bytes.

License

The MIT License.
*/
package main
