package noodlog

// record struct represents the schema for every log record
type record struct {
	Level     string      `json:"level"`
	File      string      `json:"file"`
	Function  string      `json:"function"`
	Message   interface{} `json:"message"`
	Time      string      `json:"time"`
}