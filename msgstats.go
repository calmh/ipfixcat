package main

import (
	"log"
	"time"
)

var (
	msgT0   time.Time = time.Now()
	records int
	msgs    int
)

func accountMsgStats(sets []recordMap) {
	msgs++
	records += len(sets)
}

type MsgStats struct {
	Msgs       int
	Msgsps     int
	Records    int
	Recordsps  int
	AvgMsgRecs int
}

func logMsgStats() {
	now := time.Now()
	diff := now.Sub(msgT0).Seconds()

	ms := MsgStats{}
	ms.Msgs = msgs
	ms.Msgsps = int(float64(msgs) / diff)
	ms.Records = records
	ms.Recordsps = int(float64(records) / diff)
	if records > 0 {
		ms.AvgMsgRecs = records / msgs
	}

	log.Printf("%#v\n", ms)
	msgT0 = now
	records = 0
	msgs = 0
}
