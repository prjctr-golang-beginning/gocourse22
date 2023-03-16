package messsage

type Message struct {
	DeletedID     int                    `json:"id"`
	RecordID      interface{}            `json:"record_id"`
	GroupName     string                 `json:"group_name"`
	GoodsID       interface{}            `json:"record_from_id"`
	RelationID    interface{}            `json:"record_to_id"`
	Name          string                 `json:"name"`
	Variant       string                 `json:"variant"`
	FieldsData    map[string]interface{} `json:"fields_data"`
	ChangedFields map[string]interface{} `json:"changed_fields"`
	SourceData    SourceData             `json:"source_data"`
}

type SourceData struct {
	Request string `json:"request"`
}
