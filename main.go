package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/calmh/ipfix"
	"os"
)

func main() {
	dictFile := flag.String("dict", "", "User dictionary file")
	flag.Parse()

	s := ipfix.NewSession(os.Stdin)

	if *dictFile != "" {
		loadUserDictionary(*dictFile, s)
	}

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

			bs, err := json.Marshal(set)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(bs))
		}
	}
}
