# goodystone

A Go library for parsing Eddystone packets. It is designed to parse hex strings dumped by [hcidump](http://linux.die.net/man/8/hcidump).

All three packet types supported. I tried to cross check with [Sandeep Mistry's Eddystone Scanner](https://github.com/sandeepmistry/node-eddystone-beacon-scanner).

[Eddystone Specifications](https://github.com/google/eddystone) :

# Eddystone-UID

The Eddystone-UID frame broadcasts an opaque, unique 16-byte Beacon ID composed of a 10-byte namespace and a 6-byte instance. The Beacon ID may be useful in mapping a device to a record in external storage. The namespace portion of the ID may be used to group a particular set of beacons, while the instance portion of the ID identifies individual devices in the group. The division of the ID into namespace and instance components may also be used to optimize BLE scanning strategies, e.g. by filtering only on the namespace.

## Frame Specification

The UID frame is encoded in the advertisement as a Service Data block associated with the Eddystone service UUID. The layout is:

Byte offset | Field | Description
------------|-------|------------
0 | Frame Type | Value = `0x00`
1 | Ranging Data | Calibrated Tx power at 0 m
2 | NID[0] | 10-byte Namespace
3 | NID[1]
4 | NID[2]
5 | NID[3]
6 | NID[4]
7 | NID[5]
8 | NID[6]
9 | NID[7]
10 | NID[8]
11 | NID[9]
12 | BID[0] | 6-byte Instance
13 | BID[1]
14 | BID[2]
15 | BID[3]
16 | BID[4]
17 | BID[5]
18 | RFU | Reserved for future use, must be`0x00`
19 | RFU | Reserved for future use, must be`0x00`

All multi-byte values are big-endian.

# Eddystone-URL

The Eddystone-URL frame broadcasts a URL using a compressed encoding format in order to fit more within the limited advertisement packet.

Once decoded, the URL can be used by any client with access to the internet.  For example, if an Eddystone-URL beacon were to broadcast the URL `https://goo.gl/Aq18zF`, then any client that received this packet could choose to [visit that url](https://goo.gl/Aq18zF).

The Eddystone-URL frame forms the backbone of the [Physical Web](http://physical-web.org), an effort to enable frictionless discovery of web content relating to one’s surroundings. Eddystone-URL incorporates all of the learnings from the [UriBeacon](http://uribeacon.org) format from which it evolved.

## Frame Specification

Byte offset | Field | Description
------------|-------|------------
0 | Frame Type | Value = `0x10`
1 | TX Power | Calibrated Tx power at 0 m
2 | URL Scheme | Encoded Scheme Prefix
3+ | Encoded URL | Length 0-17

# Eddystone-TLM

Eddystone beacons may transmit data about their own operation to clients. This data is called _telemetry_ and is useful for monitoring the health and operation of a fleet of beacons. Since the Eddystone-TLM frame does not contain a beacon ID, it must be paired with an identifying frame which provides the ID, either of type Eddystone-UID or Eddystone-URL. See Interleaving Telemetry for details.

Like the Eddystone-UID and Eddystone-URL frame types, Eddystone-TLM is broadcast in the clear, without message integrity validation. You should design your application to be tolerant of the open nature of such a broadcast.

## Frame Specification

The TLM frame is encoded in the advertisement as a Service Data block associated with the Eddystone service UUID. The layout is:

Byte offset | Field | Description
------------|-------|------------
0 | Frame Type | Value = `0x20`
1 | Version | TLM version, value = `0x00`
2 | VBATT[0] | Battery voltage, 1 mV/bit
3 | VBATT[1] |
4 | TEMP[0] | Beacon temperature
5 | TEMP[1] |
6 | ADV_CNT[0] | Advertising PDU count
7 | ADV_CNT[1] |
8 | ADV_CNT[2] |
9 | ADV_CNT[3] |
10 | SEC_CNT[0] | Time since power-on or reboot
11 | SEC_CNT[1] |
12 | SEC_CNT[2] |
13 | SEC_CNT[3] |

All multi-byte values are big-endian.
