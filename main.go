package main

import (
	"fmt"
	"github.com/calmh/ipfix"
	"os"
)

func main() {
	s := ipfix.NewSession(os.Stdin)

	for {
		msg, err := s.ReadMessage()
		if err != nil {
			panic(err)
		}

		for _, ts := range msg.TemplateSets {
			fmt.Printf("# Received Template Set %d\n", ts.TemplateHeader.TemplateId)
		}
		for _, ds := range msg.DataSets {
			fmt.Printf("--- %d %d\n", msg.Header.ExportTime, ds.TemplateId)
			if tpl := s.Templates[ds.TemplateId]; tpl != nil {
				for i, field := range tpl {
					fmt.Printf("%d.%d: %v\n", field.EnterpriseId, field.FieldId, ds.Records[i])
				}
			} else {
				fmt.Println("# Data Set with unknown template")
			}
		}
	}
}
