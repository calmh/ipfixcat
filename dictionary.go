package main

import (
	"code.google.com/p/gcfg"
	"github.com/calmh/ipfix"
)

type Field struct {
	Id         uint16
	Enterprise uint32
	Type       ipfix.FieldType
}

type UserDictionary struct {
	Field map[string]*Field
}

func loadUserDictionary(fname string, s *ipfix.Session) {
	dict := UserDictionary{}
	err := gcfg.ReadFileInto(&dict, fname)
	if err != nil {
		panic(err)
	}

	for name, entry := range dict.Field {
		e := ipfix.DictionaryEntry{Name: name, FieldId: entry.Id, EnterpriseId: entry.Enterprise, Type: entry.Type}
		s.AddDictionaryEntry(e)
	}
}
