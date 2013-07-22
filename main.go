package main

import (
	"encoding/json"
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

		for _, ds := range msg.DataSets {
			set := make(map[string]interface{})
			set["templateId"] = ds.TemplateId
			set["exportTime"] = msg.Header.ExportTime
			elements := make(map[string]interface{})
			set["elements"] = elements

			if tpl := s.Templates[ds.TemplateId]; tpl != nil {
				for i, field := range tpl {
					var fieldName string
					if field.EnterpriseId > 0 {
						fieldName = fmt.Sprintf("V%d.%d", field.EnterpriseId, field.FieldId)
					} else {
						fieldName = fmt.Sprintf("F%d", field.FieldId)
					}

					// json.Marshal will format a []byte as base64 string, but []int as integer array.
					elements[fieldName] = integers(ds.Records[i])
				}
			}
			bs, err := json.Marshal(set)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(bs))
		}
	}
}

func integers(s []byte) []int {
	r := make([]int, len(s))
	for i := range s {
		r[i] = int(s[i])
	}
	return r
}
