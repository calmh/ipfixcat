package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/calmh/ipfix"
	"log"
	"os"
	"time"
)

const (
	statReportInterval = 60
)

func messagesGenerator(s *ipfix.Session) <-chan map[string]interface{} {
	c := make(chan map[string]interface{})

	go func() {
		for {
			msg, err := s.ReadMessage()
			if err != nil {
				panic(err)
			}

			for _, record := range msg.DataRecords {
				set := make(map[string]interface{})
				set["templateId"] = record.TemplateId
				set["exportTime"] = msg.Header.ExportTime
				set["elements"] = s.Interpret(&record)

				c <- set
			}
		}
	}()

	return c
}

func main() {
	log.Println("ipfixcat", ipfixcatVersion)
	dictFile := flag.String("dict", "", "User dictionary file")
	flag.Parse()

	s := ipfix.NewSession(os.Stdin)

	if *dictFile != "" {
		loadUserDictionary(*dictFile, s)
	}

	msgC := messagesGenerator(s)
	tick := time.Tick(time.Duration(statReportInterval) * time.Second)
	t0 := time.Now()
	records := 0
	for {
		select {
		case msg := <-msgC:
			records++
			bs, err := json.Marshal(msg)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(bs))
		case <-tick:
			now := time.Now()
			diff := now.Sub(t0).Seconds()
			recsPs := float64(records) / diff
			log.Printf("%d records (%3.2f/s)\n", records, recsPs)
			t0 = now
			records = 0
		}
	}
}
