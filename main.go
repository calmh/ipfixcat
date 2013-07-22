package main

import (
	"encoding/json"
	"fmt"
	"github.com/calmh/ipfix"
	"os"
	"code.google.com/p/gcfg"
	"flag"
)

type Field struct {
	Id uint16
	Enterprise uint32
	Type ipfix.FieldType
}

type UserDictionary struct {
	Field map[string]*Field
}

func main() {
	dictFile := flag.String("dict", "", "User dictionary file")
	flag.Parse()

	s := ipfix.NewSession(os.Stdin)

	if *dictFile != "" {
		dict := UserDictionary{}
		err := gcfg.ReadFileInto(&dict, *dictFile)
		if err != nil {
			panic(err)
		}
		for name, entry := range dict.Field {
			e := ipfix.DictionaryEntry{Name: name, FieldId: entry.Id, EnterpriseId: entry.Enterprise, Type: entry.Type}
			s.AddDictionaryEntry(e)
		}
	}


	for {
		msg, err := s.ReadMessage()
		if err != nil {
			panic(err)
		}

		for _, ds := range msg.DataSets {
			set := make(map[string]interface{})
			set["templateId"] = ds.TemplateId
			set["exportTime"] = msg.Header.ExportTime
			set["elements"] = s.Interpret(&ds)
			bs, err := json.Marshal(set)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(bs))
		}
	}
}
