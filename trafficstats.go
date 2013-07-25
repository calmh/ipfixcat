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

func accountTraffic(rec InterpretedRecord) {
	for _, f := range rec.Fields {
		if f.Name == "proceraIncomingOctets" {
			inOctets += f.Value.(uint64)
		} else if f.Name == "proceraOutgoingOctets" {
			outOctets += f.Value.(uint64)
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
