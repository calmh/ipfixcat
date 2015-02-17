package main

import (
	"code.google.com/p/gcfg"
	"github.com/calmh/ipfix"
)

type Field struct {
	Id         uint16
	Enterprise uint32
	Type       string
}

func (f Field) DictionaryEntry(name string) ipfix.DictionaryEntry {
	return ipfix.DictionaryEntry{
		Name:         name,
		EnterpriseId: f.Enterprise,
		FieldId:      f.Id,
		Type:         ipfix.FieldTypes[f.Type],
	}
}

type UserDictionary struct {
	Field map[string]*Field
}

func loadUserDictionary(fname string, i *ipfix.Interpreter) error {
	dict := UserDictionary{}
	err := gcfg.ReadFileInto(&dict, fname)
	if err != nil {
		return err
	}

	for name, entry := range dict.Field {
		i.AddDictionaryEntry(entry.DictionaryEntry(name))
	}

	return nil
}
