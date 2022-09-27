package pretty

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

// Struct pretty the struct
func Struct(obj interface{}) string {
	typ := reflect.TypeOf(obj)
	if typ.Kind() != reflect.Ptr {
		panic(errors.New("object is not a pointer"))
	}

	data := make([][]string, 0)
	value := reflect.ValueOf(obj).Elem()
	for i := 0; i < value.NumField(); i++ {
		data = append(data, []string{value.Type().Field(i).Name, fmt.Sprintf("%v", value.Field(i).Interface())})
	}

	writer := bytes.NewBuffer(make([]byte, 0))
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"Field", "Value"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return writer.String()
}
