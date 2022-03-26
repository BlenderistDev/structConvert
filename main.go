package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type income struct {
	Field1 string `json:"field_1"`
	Field2 int64  `json:"field_2"`
	Field3 bool   `json:"field_3"`
	Field4 struct {
		InnerField1 string `json:"inner_field_1"`
		InnerField2 int64  `json:"inner_field_2"`
		InnerField3 bool   `json:"inner_field_3"`
	}
	Field5 []string `json:"field_5"`
}

type outcome struct {
	Field1 string `json:"field_1"`
	Field2 int64  `json:"field_2"`
	Field3 bool   `json:"field_3"`
	Field4 struct {
		InnerField1 string `json:"inner_field_1"`
		InnerField2 int64  `json:"inner_field_2"`
		InnerField3 bool   `json:"inner_field_3"`
	}
	Field5 []string `json:"field_5"`
}

func main() {
	in := income{
		Field1: "test",
		Field2: 999,
		Field3: true,
		Field4: struct {
			InnerField1 string `json:"inner_field_1"`
			InnerField2 int64  `json:"inner_field_2"`
			InnerField3 bool   `json:"inner_field_3"`
		}{
			InnerField1: "test2",
			InnerField2: 4234,
			InnerField3: false,
		},
		Field5: []string{"testtest1", "testtest2", "testtest3"},
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		start := time.Now()

		for i := 0; i < 1000000; i++ {
			_ = convertWithMarshal(in)
		}

		t := time.Since(start)
		fmt.Printf("convertWithMarshal took %s \n", t)
		wg.Done()
	}()

	go func() {
		start := time.Now()

		for i := 0; i < 1000000; i++ {
			_ = convertWithMapping(in)
		}

		t := time.Since(start)
		fmt.Printf("convertWithMapping took %s \n", t)
		wg.Done()
	}()

	wg.Wait()
}

func convertWithMarshal(i income) outcome {
	s, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	var o outcome
	err = json.Unmarshal(s, &o)
	if err != nil {
		panic(err)
	}

	return o
}

func convertWithMapping(i income) outcome {
	return outcome{
		Field1: i.Field1,
		Field2: i.Field2,
		Field3: i.Field3,
		Field4: struct {
			InnerField1 string `json:"inner_field_1"`
			InnerField2 int64  `json:"inner_field_2"`
			InnerField3 bool   `json:"inner_field_3"`
		}{
			InnerField1: i.Field4.InnerField1,
			InnerField2: i.Field4.InnerField2,
			InnerField3: i.Field4.InnerField3,
		},
		Field5: i.Field5,
	}
}
