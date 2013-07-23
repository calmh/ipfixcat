package main

import (
	"log"
	"time"
)

var (
	outOctets   uint64
	inOctets    uint64
	accountTime time.Time = time.Now()
)

type TrafStats struct {
	InKbps  int
	OutKbps int
}

func accountTraffic(set recordMap) {
	if elems, ok := set["elements"]; ok {
		// Account the in/out traffic rate if we have Procera Networks vendor fields available.
		elems := elems.(map[string]interface{})
		if in, ok := elems["proceraIncomingOctets"]; ok {
			inOctets += in.(uint64)
		}
		if out, ok := elems["proceraOutgoingOctets"]; ok {
			outOctets += out.(uint64)
		}
	}
}

func logAccountedTraffic() {
	now := time.Now()
	diff := now.Sub(accountTime).Seconds()

	ts := TrafStats{}
	ts.InKbps = int(float64(inOctets*8/1000) / diff)
	ts.OutKbps = int(float64(outOctets*8/1000) / diff)

	log.Printf("%#v\n", ts)

	outOctets = 0
	inOctets = 0
	accountTime = now
}
