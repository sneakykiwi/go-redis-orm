package reflector

import "reflect"

type ModelReflector struct {
	Fields []struct {
		Name string
		Tag  string
	}
}

func (r *ModelReflector) ParseModel(dst interface{}) error {
	val := reflect.Indirect(reflect.ValueOf(dst))
	var fieldNames []string

	for i := 0; i < val.NumField(); i++ {
		fieldNames = append(fieldNames, val.Type().Field(i).Name)
	}

	for _, fieldName := range fieldNames {
		field, ok := reflect.TypeOf(dst).Elem().FieldByName(fieldName)
		if !ok {
			return fieldAcquireError
		}

		r.Fields = append(r.Fields, struct {
			Name string
			Tag  string
		}{Name: fieldName, Tag: string(field.Tag)})

	}
	return nil
}
