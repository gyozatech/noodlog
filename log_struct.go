package noodlog

import "io"

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
	LogLevel             *string
	LogWriter            io.Writer
	JSONPrettyPrint      *bool
	TraceCaller          *bool
	SinglePointTracing   *bool
	Colors               *bool
	CustomColors         *CustomColors
	ObscureSensitiveData *bool
	SensitiveParams      []string
}

// CustomColors struct is used to specify the custom colors for the various log levels
type CustomColors struct {
	Trace interface{}
	Debug interface{}
	Info  interface{}
	Warn  interface{}
	Error interface{}
}

// ~~~~~~~~~~~ Prebuilt pointers to be used in the SetConfigs ~~~~~~~~~~~~ //

// LevelTrace pointer for the Config struct
var LevelTrace = pointerOfString(traceLabel)

// LevelDebug pointer for the Config struct
var LevelDebug = pointerOfString(debugLabel)

// LevelInfo pointer for the Config struct
var LevelInfo = pointerOfString(infoLabel)

// LevelWarn pointer for the Config struct
var LevelWarn = pointerOfString(warnLabel)

// LevelError pointer for the Config struct
var LevelError = pointerOfString(errorLabel)

// Enable pointer for the Config struct
var Enable = pointerOfBool(true)

// Disable pointer for the Config struct
var Disable = pointerOfBool(false)

// Default pointer for the CustomColors struct
var Default = pointerOfString(defaultColor)

// Red pointer for the CustomColors struct
var Red = pointerOfString(redColor)

// Green pointer for the CustomColors struct
var Green = pointerOfString(greenColor)

// Yellow pointer for the CustomColors struct
var Yellow = pointerOfString(yellowColor)

// Blue pointer for the CustomColors struct
var Blue = pointerOfString(blueColor)

// Purple pointer for the CustomColors struct
var Purple = pointerOfString(purpleColor)

// Cyan pointer for the CustomColors struct
var Cyan = pointerOfString(cyanColor)
