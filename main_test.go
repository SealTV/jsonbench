package main

import "testing"

var testData = &Data{
	Ival: 123,
	Sval: "some text string here",
	Arr: []Str{
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
		{I: 123, S: "asdasdasdasd asda"},
	},
}

func BenchmarkStdJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdJSONMarshal(testData)
	}
}

func BenchmarkCustromJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CustromJSONMarshal(testData)
	}
}

func BenchmarkBufferJsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BufferJsonMarshal(testData)
	}
}

func BenchmarkGoJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoJSONMarshal(testData)
	}
}

func BenchmarkEasyJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EasyJSONMarshal(testData)
	}
}
