package messsage

import (
	"encoding/json"
	"testing"
)

var example = `{
	"id": 921,
	"record_id": "IUOgiuyt(&^gi76gI76f76",
	"group_name": "Products",
	"record_from_id": 217,
	"record_to_id": 12,
	"name": "Tourizm",
	"variant": "tourism.jpg",
	"fields_data": "",
	"changed_fields": "id, name",
	"source_data": {
		"request": "BO^*gf87^F*75TVU^r5Dc6rTVu^T5vu^5IU&"
	}
}`

func BenchmarkMessageUnmarshalInMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res1 := map[string]any{}
		_ = json.Unmarshal([]byte(example), &res1)
	}
}

func BenchmarkMessageUnmarshalInStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res2 := Message{}
		_ = json.Unmarshal([]byte(example), &res2)
	}
}

func BenchmarkMessageEasyJSONUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res3 := Message{}
		_ = res3.UnmarshalJSON([]byte(example))
	}
}
