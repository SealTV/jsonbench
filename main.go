package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	gojson "github.com/goccy/go-json"
	"github.com/mailru/easyjson"
)

func main() {
	fmt.Println("hello world")
}

type Data struct {
	Ival int    `json:"ival"`
	Sval string `json:"sval"`
	Arr  []Str  `json:"arr"`
}

type Str struct {
	I int    `json:"i"`
	S string `json:"s"`
}

func StdJSONMarshal(data *Data) []byte {
	b, _ := json.Marshal(data)
	return b
}

func CustromJSONMarshal(data *Data) []byte {
	result := []byte{}
	result = append(result, '{')
	result = append(result, `"ival":`...)
	result = append(result, []byte(strconv.Itoa(data.Ival))...)
	result = append(result, `,"sval":"`...)
	result = append(result, data.Sval...)
	result = append(result, `"`...)
	result = append(result, `,"arr":[`...)
	for i, v := range data.Arr {
		if i > 0 {
			result = append(result, ',')
		}
		result = append(result, `{"i":`...)
		result = append(result, strconv.Itoa(v.I)...)
		result = append(result, `,"s":"`...)
		result = append(result, v.S...)
		result = append(result, `"}`...)
	}
	result = append(result, `]`...)
	result = append(result, '}')
	return result
}
func BufferJsonMarshal(data *Data) []byte {
	builder := strings.Builder{}

	builder.WriteRune('{')
	builder.WriteString(`"ival":`)
	builder.WriteString(strconv.Itoa(data.Ival))

	builder.WriteString(`,"sval":"`)
	builder.WriteString(data.Sval)
	builder.WriteString(`",[`)
	for i, v := range data.Arr {
		if i > 0 {
			builder.WriteRune(',')
		}
		builder.WriteString(`{"i":`)
		builder.WriteString(strconv.Itoa(v.I))
		builder.WriteString(`,"s":"`)
		builder.WriteString(v.S)
		builder.WriteString(`"}`)
	}

	builder.WriteRune(']')
	builder.WriteRune('}')

	return []byte(builder.String())
}

func GoJSONMarshal(data *Data) []byte {
	b, _ := gojson.Marshal(data)
	return b
}

func EasyJSONMarshal(data *Data) []byte {
	b, _ := easyjson.Marshal(data)
	return b
}
