package main

import (
	"testing"

	"structConvert/dry"
)

const (
	field1      = "test"
	field2      = 999
	field3      = true
	innerField1 = "test2"
	innerField2 = 4234
	innerField3 = false
)

func TestConvertWithMarshal(t *testing.T) {
	field5 := []string{"testtest1", "testtest2", "testtest3"}

	in := income{
		Field1: field1,
		Field2: field2,
		Field3: field3,
		Field4: struct {
			InnerField1 string `json:"inner_field_1"`
			InnerField2 int64  `json:"inner_field_2"`
			InnerField3 bool   `json:"inner_field_3"`
		}{
			InnerField1: innerField1,
			InnerField2: innerField2,
			InnerField3: innerField3,
		},
		Field5: field5,
	}

	out := convertWithMarshal(in)

	dry.TestCheckEqual(t, field1, out.Field1)
	dry.TestCheckEqual(t, int64(field2), out.Field2)
	dry.TestCheckEqual(t, field3, out.Field3)
	dry.TestCheckEqual(t, innerField1, out.Field4.InnerField1)
	dry.TestCheckEqual(t, int64(innerField2), out.Field4.InnerField2)
	dry.TestCheckEqual(t, innerField3, out.Field4.InnerField3)
	dry.TestCheckEqual(t, field5, out.Field5)
}

func TestConvertWithMapping(t *testing.T) {
	field5 := []string{"testtest1", "testtest2", "testtest3"}

	in := income{
		Field1: field1,
		Field2: field2,
		Field3: field3,
		Field4: struct {
			InnerField1 string `json:"inner_field_1"`
			InnerField2 int64  `json:"inner_field_2"`
			InnerField3 bool   `json:"inner_field_3"`
		}{
			InnerField1: innerField1,
			InnerField2: innerField2,
			InnerField3: innerField3,
		},
		Field5: field5,
	}

	out := convertWithMapping(in)

	dry.TestCheckEqual(t, field1, out.Field1)
	dry.TestCheckEqual(t, int64(field2), out.Field2)
	dry.TestCheckEqual(t, field3, out.Field3)
	dry.TestCheckEqual(t, innerField1, out.Field4.InnerField1)
	dry.TestCheckEqual(t, int64(innerField2), out.Field4.InnerField2)
	dry.TestCheckEqual(t, innerField3, out.Field4.InnerField3)
	dry.TestCheckEqual(t, field5, out.Field5)
}
