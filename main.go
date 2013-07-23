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

type recordMap map[string]interface{}

func messagesGenerator(s *ipfix.Session) <-chan []recordMap {
	c := make(chan []recordMap)

	go func() {
		for {
			sets := make([]recordMap, 0, 32)
			msg, err := s.ReadMessage()
			if err != nil {
				panic(err)
			}

			for _, record := range msg.DataRecords {
				set := make(map[string]interface{})
				set["templateId"] = record.TemplateId
				set["exportTime"] = msg.Header.ExportTime
				set["elements"] = s.Interpret(&record)
				sets = append(sets, set)
			}

			c <- sets
		}
	}()

	return c
}

func main() {
	log.Println("ipfixcat", ipfixcatVersion)
	dictFile := flag.String("dict", "", "User dictionary file")
	messageStats := flag.Bool("mstats", false, "Log IPFIX message statistics")
	resourceStats := flag.Bool("rstats", false, "Log resource usage (CPU, memory)")
	trafficStats := flag.Bool("acc", false, "Log traffic rates (Procera)")
	output := flag.Bool("output", true, "Display received flow records in JSON format")
	statsIntv := flag.Int("statsintv", 60, "Statistics log interval (s)")
	flag.Parse()

	if *messageStats {
		log.Printf("Logging message statistics every %d seconds", *statsIntv)
	}

	if *resourceStats {
		log.Printf("Logging resource statistics every %d seconds", *statsIntv)
	}

	if *trafficStats {
		log.Printf("Logging traffic rates every %d seconds", *statsIntv)
	}

	if !*messageStats && !*resourceStats && !*trafficStats && !*output {
		log.Fatal("If you don't want me to do anything, don't run me at all.")
	}

	s := ipfix.NewSession(os.Stdin)

	if *dictFile != "" {
		loadUserDictionary(*dictFile, s)
	}

	msgs := messagesGenerator(s)
	tick := time.Tick(time.Duration(*statsIntv) * time.Second)
	for {
		select {
		case sets := <-msgs:
			if *messageStats {
				accountMsgStats(sets)
			}

			for _, set := range sets {
				if *trafficStats {
					accountTraffic(set)
				}

				if *output {
					bs, _ := json.Marshal(set)
					fmt.Println(string(bs))
				}
			}
		case <-tick:
			if *messageStats {
				logMsgStats()
			}

			if *trafficStats {
				logAccountedTraffic()
			}

			if *resourceStats {
				logResourceUsage()
			}
		}
	}
}
