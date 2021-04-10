package noodlog

// record struct represents the schema for every log record
type record struct {
	Level    string      `json:"level,omitempty"`
	File     *string     `json:"file,omitempty"`
	Function *string     `json:"function,omitempty"`
	Message  interface{} `json:"message,omitempty"`
	Time     string      `json:"time,omitempty"`
}
