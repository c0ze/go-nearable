package nearable

import (
	"strings"
	"testing"
)

// onDiscover: manufacturerData = 5d0101100550295ef974540401aa01b0fcfcc24a0356, rssi = -82
// { id: '100550295ef97454',
// 	uuid: 'd0d3fa86ca7645ec9bd96af4100550295ef97454',
// 	type: 'SB0',
// 	firmware: 'unknown',
// 	bootloader: 'SB1.0.0',
// 	temperature: 26.625,
// 	moving: false,
// 	batteryLevel: 'unknown',
// 	acceleration: { x: -62.5, y: -62.5, z: -968.75 },
// 	currentMotionStateDuration: 600,
// 	previousMotionStateDuration: 3,
// 	power: 4,
// 	firmwareState: 'app' }
const StickerPacket1 = `04 3E 2B 02 01 03 01 A1 82 D6 1B 03 C1 1F 02 01 04 03 03 0F 18 17 FF 5D 01 01 10 05 50 29 5E F9 74 54 04 01 B9 01 B2 F7 FD C3 21 01 56 BF`

// onDiscover: manufacturerData = 5d01018f6d0823c95c1b130401a681b1fbffc3c20066, rssi = -84
// { id: '8f6d0823c95c1b13',
// 	uuid: 'd0d3fa86ca7645ec9bd96af48f6d0823c95c1b13',
// 	type: 'SB0',
// 	firmware: 'unknown',
// 	bootloader: 'SB1.0.0',
// 	temperature: 26.375,
// 	moving: false,
// 	batteryLevel: 'unknown',
// 	acceleration: { x: -78.125, y: -15.625, z: -953.125 },
// 	currentMotionStateDuration: 2,
// 	previousMotionStateDuration: 0,
// 	power: 4,
// 	firmwareState: 'app' }
const StickerPacket2 = `04 3E 2B 02 01 03 01 70 5C A0 20 33 E4 1F 02 01 04 03 03 0F 18 17 FF 5D 01 01 8F 6D 08 23 C9 5C 1B 13 04 82 BB 61 35 FA FF C3 41 00 76 B2`

func TestIsValid(t *testing.T) {
	if !IsValid(StickerPacket1) {
		t.Errorf("Validation failed for Packet1")
	}

	if !IsValid(StickerPacket2) {
		t.Errorf("Validation failed for Packet2")
	}
}

func TestParseProtocolVersion(t *testing.T) {
	protocol := parseProtocolVersion(strings.Split(StickerPacket1, " "))
	expectedProtocol := "01"
	if protocol != expectedProtocol {
		t.Errorf("Parsing protocol {%v} failed for packet1: %v", expectedProtocol, protocol)
	}

	protocol = parseProtocolVersion(strings.Split(StickerPacket2, " "))
	if protocol != expectedProtocol {
		t.Errorf("Parsing protocol {%v} failed for packet2: %v", expectedProtocol, protocol)
	}
}

func TestParseUuid(t *testing.T) {
	uuid := parseUuid(strings.Split(StickerPacket1, " "))
	expectedUuid := strings.ToUpper("100550295ef97454")
	if uuid != expectedUuid {
		t.Errorf("Parsing uuid {%v} failed for packet1: %v", expectedUuid, uuid)
	}

	uuid = parseUuid(strings.Split(StickerPacket2, " "))
	expectedUuid = strings.ToUpper("8f6d0823c95c1b13")
	if uuid != expectedUuid {
		t.Errorf("Parsing uuid {%v} failed for packet2: %v", expectedUuid, uuid)
	}
}

func TestParseRssi(t *testing.T) {
	rssi := parseRssi(strings.Split(StickerPacket1, " "))
	expectedRssi := -65
	if rssi != expectedRssi {
		t.Errorf("Parsing rssi {%v} failed for packet1: %v", expectedRssi, rssi)
	}

	rssi = parseRssi(strings.Split(StickerPacket2, " "))
	expectedRssi = -78
	if rssi != expectedRssi {
		t.Errorf("Parsing rssi {%v} failed for packet2: %v", expectedRssi, rssi)
	}

}

// func TestParseType(t *testing.T) {
// 	Type := parseType(strings.Split(UidPacket1, " "))
// 	expectedType := "00"
// 	if Type != expectedType {
// 		t.Errorf("Parsing type {%v} failed for packet1: %v", expectedType, Type)
// 	}

// 	Type = parseType(strings.Split(UidPacket2, " "))
// 	if Type != expectedType {
// 		t.Errorf("Parsing type {%v} failed for packet2: %v", expectedType, Type)
// 	}

// 	Type = parseType(strings.Split(TlmPacket, " "))
// 	expectedType = "20"
// 	if Type != expectedType {
// 		t.Errorf("Parsing type {%v} failed for tlm packet: %v", expectedType, Type)
// 	}

// 	Type = parseType(strings.Split(UrlPacket1, " "))
// 	expectedType = "10"
// 	if Type != expectedType {
// 		t.Errorf("Parsing type {%v} failed for url packet: %v", expectedType, Type)
// 	}
// }

// func TestParsePower(t *testing.T) {
// 	power := parsePower(strings.Split(UidPacket1, " "))
// 	expectedPower := -21
// 	if power != expectedPower {
// 		t.Errorf("Parsing power {%v} failed for packet1: %v", expectedPower, power)
// 	}

// 	power = parsePower(strings.Split(UidPacket2, " "))
// 	if power != expectedPower {
// 		t.Errorf("Parsing power {%v} failed for packet2: %v", expectedPower, power)
// 	}

// 	power = parsePower(strings.Split(UrlPacket2, " "))
// 	expectedPower = -27
// 	if power != expectedPower {
// 		t.Errorf("Parsing power {%v} failed for url packet2: %v", expectedPower, power)
// 	}

// 	power = parsePower(strings.Split(UrlPacket3, " "))
// 	expectedPower = -21
// 	if power != expectedPower {
// 		t.Errorf("Parsing power {%v} failed for url packet3: %v", expectedPower, power)
// 	}
// }

// func TestParseMac(t *testing.T) {
// 	mac := parseMac(strings.Split(UidPacket1, " "))
// 	expectedMac := "EBB1CB1C7CC1"
// 	if mac != expectedMac {
// 		t.Errorf("Parsing mac {%v} failed for packet1: %v", expectedMac, mac)
// 	}

// 	mac = parseMac(strings.Split(UidPacket2, " "))
// 	if mac != expectedMac {
// 		t.Errorf("Parsing mac {%v} failed for packet2: %v", expectedMac, mac)
// 	}

// 	mac = parseMac(strings.Split(TlmPacket, " "))
// 	if mac != expectedMac {
// 		t.Errorf("Parsing mac {%v} failed for tlm packet: %v", expectedMac, mac)
// 	}

// 	mac = parseMac(strings.Split(UrlPacket1, " "))
// 	if mac != expectedMac {
// 		t.Errorf("Parsing mac {%v} failed for url packet: %v", expectedMac, mac)
// 	}
// }

// func TestToString(t *testing.T) {
// 	uidP := NewUIDPacket(UidPacket1)
// 	expectedStr := "INT  UID EDD1EBEAC04E5DEFA017 INSTANCE EBB1CB1C7CC1 RSSI -48"
// 	if uidP.ToString() != expectedStr {
// 		t.Errorf("Parsing packet {%v} failed for packet1", uidP.ToString())
// 	}

// 	uidP = NewUIDPacket(UidPacket2)
// 	expectedStr = "INT  UID EDD1EBEAC04E5DEFA017 INSTANCE EBB1CB1C7CC1 RSSI -49"
// 	if uidP.ToString() != expectedStr {
// 		t.Errorf("Parsing packet {%v} failed for packet2", uidP.ToString())
// 	}
// }

// func TestMapKey(t *testing.T) {
// 	uidP := NewUIDPacket(UidPacket1)
// 	expectedStr := "EDD1EBEAC04E5DEFA017|EBB1CB1C7CC1"
// 	if uidP.MapKey() != expectedStr {
// 		t.Errorf("Parsing packet {%v} failed for packet1", uidP.MapKey())
// 	}

// 	uidP = NewUIDPacket(UidPacket2)
// 	if uidP.MapKey() != expectedStr {
// 		t.Errorf("Parsing packet {%v} failed for packet2", uidP.MapKey())
// 	}

// 	urlP := NewURLPacket(UrlPacket2)
// 	expectedStr = "EBB1CB1C7CC1"
// 	if urlP.MapKey() != expectedStr {
// 		t.Errorf("Parsing packet {%v} failed for url packet2", urlP.MapKey())
// 	}
// }

// func TestParseVersion(t *testing.T) {
// 	version := parseVersion(strings.Split(TlmPacket, " "))
// 	expectedVersion := "00"
// 	if version != expectedVersion {
// 		t.Errorf("Parsing version {%v} failed for packet: %v", expectedVersion, version)
// 	}
// }

// func TestParseBattery(t *testing.T) {
// 	battery := parseBattery(strings.Split(TlmPacket, " "))
// 	expectedBattery := 2925
// 	if battery != expectedBattery {
// 		t.Errorf("Parsing battery {%v} failed for packet: %v", expectedBattery, battery)
// 	}
// }

// func TestParseTemperature(t *testing.T) {
// 	temperature := parseTemperature(strings.Split(TlmPacket, " "))
// 	expectedTemperature := 32.75
// 	if temperature != expectedTemperature {
// 		t.Errorf("Parsing temperature {%v} failed for packet: %v", expectedTemperature, temperature)
// 	}
// }

// func TestParsePacketCount(t *testing.T) {
// 	packetCount := parsePacketCount(strings.Split(TlmPacket, " "))
// 	expectedPacketCount := 400151
// 	if packetCount != expectedPacketCount {
// 		t.Errorf("Parsing packetCount {%v} failed for packet: %v", expectedPacketCount, packetCount)
// 	}
// }

// func TestParseTimeCount(t *testing.T) {
// 	timeCount := parseTimeCount(strings.Split(TlmPacket, " "))
// 	expectedTimeCount := 4422840
// 	if timeCount != expectedTimeCount {
// 		t.Errorf("Parsing timeCount {%v} failed for packet: %v", expectedTimeCount, timeCount)
// 	}
// }

// func TestParseScheme(t *testing.T) {
// 	scheme := parseScheme(strings.Split(UrlPacket1, " "))
// 	expectedScheme := "http://"
// 	if scheme != expectedScheme {
// 		t.Errorf("Parsing scheme {%v} failed for packet: %v", expectedScheme, scheme)
// 	}
// }

// func TestParseUrl(t *testing.T) {
// 	url := parseUrl(strings.Split(UrlPacket1, " "))
// 	expectedUrl := "go.esti.be"
// 	if url != expectedUrl {
// 		t.Errorf("Parsing url {%v} failed for url packet1: %v", expectedUrl, url)
// 	}

// 	url = parseUrl(strings.Split(UrlPacket2, " "))
// 	expectedUrl = "go.esti.ben"
// 	if url != expectedUrl {
// 		t.Errorf("Parsing url {%v} failed for url packet2: %v", expectedUrl, url)
// 	}

// 	url = parseUrl(strings.Split(UrlPacket3, " "))
// 	expectedUrl = "go.esti.ben"
// 	if url != expectedUrl {
// 		t.Errorf("Parsing url {%v} failed for url packet3: %v", expectedUrl, url)
// 	}

// 	url = parseUrl(strings.Split(UrlPacket4, " "))
// 	expectedUrl = "go.esti.be"
// 	if url != expectedUrl {
// 		t.Errorf("Parsing url {%v} failed for url packet4: %v", expectedUrl, url)
// 	}
// }
