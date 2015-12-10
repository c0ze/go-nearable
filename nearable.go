package nearable

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type NearablePacket struct {
	ProtocolVersion string
	Uuid string
	Rssi int
}

func NewNearablePacket(line string) *NearablePacket {
	vals := strings.Split(line, " ")
	return &NearablePacket{
		ProtocolVersion: parseProtocolVersion(vals),
		Uuid: parseUuid(vals),
		Rssi: parseRssi(vals)}
}

func IsValid(str string) bool {
	r, err := regexp.Compile(`^04\ 3E\ .{2}\ 02\ 01\ .{41}\ 0F\ 18\ 17\ FF\ 5D\ 01`)
	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return false
	}

	return r.MatchString(str)
}

func parseProtocolVersion(vals []string) string {
	return vals[25]
}

func parseUuid(vals []string) string {
	return strings.Join(
		vals[26:34],
		"")
}

// func GetType(str string) string {
// 	vals := strings.Split(str, " ")
// 	return parseType(vals)
// }

// func parseType(vals []string) string {
// 	return vals[25]
// }

func parseRssi(vals []string) int {
	rssi, _ := strconv.ParseInt(vals[len(vals)-1], 16, 0)
	return int(rssi) - 256
}

// func parsePower(vals []string) int {
// 	power, _ := strconv.ParseInt(vals[26], 16, 0)
// 	return int(power) - 256
// }

// func parseMac(vals []string) string {
// 	return strings.Join(
// 		[]string{
// 			vals[12],
// 			vals[11],
// 			vals[10],
// 			vals[9],
// 			vals[8],
// 			vals[7]},
// 		"")
// }
