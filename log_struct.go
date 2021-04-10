package noodlog

// record struct represents the schema for every log record
type record struct {
	Level    string      `json:"level,omitempty"`
	File     *string     `json:"file,omitempty"`
	Function *string     `json:"function,omitempty"`
	Message  interface{} `json:"message,omitempty"`
	Time     string      `json:"time,omitempty"`
}

// Configs struct contains all possible configs for noodlog
type Configs struct {
	LogLevel        *string
	JSONPrettyPrint *bool
	Color           *bool
	CustomColors    *CustomColors
	SensitiveParams []string
}

// CustomColors struct is used to specify the custom colors for the various log levels
type CustomColors struct {
	Trace *string
	Debug *string
	Info  *string
	Warn  *string
	Error *string
}
