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

type InterpretedRecord struct {
	ExportTime uint32               `json:"exportTime"`
	TemplateId uint16               `json:"templateId"`
	Fields     []myInterpretedField `json:"fields"`
}

// Because we want to control JSON serialization
type myInterpretedField struct {
	Name         string      `json:"name"`
	EnterpriseId uint32      `json:"enterprise,omitempty"`
	FieldId      uint16      `json:"field"`
	Value        interface{} `json:"value,omitempty"`
	RawValue     []int       `json:"raw,omitempty"`
}

func messagesGenerator(s *ipfix.Session, i *ipfix.Interpreter) <-chan []InterpretedRecord {
	c := make(chan []InterpretedRecord)

	errors := 0
	go func() {
		for {
			msg, err := s.ReadMessage()
			if err != nil {
				errors++
				if errors > 3 {
					panic(err)
				} else {
					log.Println(err)
				}
				continue
			} else {
				errors = 0
			}

			irecs := make([]InterpretedRecord, len(msg.DataRecords))
			for j, record := range msg.DataRecords {
				ifs := i.Interpret(&record)
				mfs := make([]myInterpretedField, len(ifs))
				for k, iif := range ifs {
					mfs[k] = myInterpretedField{iif.Name, iif.EnterpriseId, iif.FieldId, iif.Value, integers(iif.RawValue)}
				}
				ir := InterpretedRecord{msg.Header.ExportTime, record.TemplateId, mfs}
				irecs[j] = ir
			}

			c <- irecs
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
	i := ipfix.NewInterpreter(s)

	if *dictFile != "" {
		loadUserDictionary(*dictFile, i)
	}

	msgs := messagesGenerator(s, i)
	tick := time.Tick(time.Duration(*statsIntv) * time.Second)
	for {
		select {
		case irecs := <-msgs:
			if *messageStats {
				accountMsgStats(irecs)
			}

			for _, rec := range irecs {
				if *trafficStats {
					accountTraffic(rec)
				}

				if *output {
					for i := range rec.Fields {
						f := &rec.Fields[i]
						switch v := f.Value.(type) {
						case []byte:
							f.RawValue = integers(v)
							f.Value = nil
						}
					}
					bs, _ := json.Marshal(rec)
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

func integers(s []byte) []int {
	if s == nil {
		return nil
	}

	r := make([]int, len(s))
	for i := range s {
		r[i] = int(s[i])
	}
	return r
}
