package structf

import (
	"log"
	"reflect"
)

func PrintStruct(in interface{}) {
	if in == nil {
		return
	}
	val := reflect.ValueOf(in)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

		log.Printf("\tname=%s, type=%s, value=%v, tag=`%s`\n",
			typeField.Name,
			typeField.Type,
			val.Field(i),
			typeField.Tag,
		)
	}
}

func UpdateStruct(in interface{}, values map[string]interface{}) {
	if in == nil {
		return
	}
	val := reflect.ValueOf(in)

	if val.Kind() != reflect.Ptr {
		return
	}
	if val.Kind() != reflect.Struct {
		return
	}

	val = val.Elem()

	for i := 0; i < val.NumField(); i++ {
		valField := val.Field(i)
		typeField := val.Type().Field(i)

		v, ok := values[typeField.Name]
		if !ok {
			continue
		}

		if reflect.TypeOf(v) == valField.Type() {
			if valField.CanSet() {
				valField.Set(reflect.ValueOf(v))
			}
		}
	}
}
