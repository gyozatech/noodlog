# Noodlog

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![CodeFactor](https://www.codefactor.io/repository/github/alessandroargentieri/noodlog/badge)](https://www.codefactor.io/repository/github/alessandroargentieri/noodlog)
[![Go Report Card](https://goreportcard.com/badge/github.com/gyozatech/noodlog)](https://goreportcard.com/report/github.com/gyozatech/noodlog)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![codecov](https://codecov.io/gh/gyozatech/noodlog/branch/main/graph/badge.svg?token=7V3XRVR58L)](https://codecov.io/gh/gyozatech/noodlog)
[![Open Source Love svg1](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

![alt text](assets/noodlogo.png?raw=true)

## Summary

**Noodlog** is a Golang JSON parametrized  and highly configurable logging library.

It allows you to:
- print **go structs** as JSON messages;
- print JSON strings and raw strings messages as **pure JSONs**;
- obscure some **sensitive params** from your logging;
- **chain objects** or strings in your logs;
- apply string **templates** to your logs;
- choose to **trace the caller** file and function and fine tune the settings;
- apply **pretty printing** or not;
- apply **colors** to your logging;
- **customize colors** per log level.

## Import 

``` go get github.com/gyozatech/noodlog ```

## Usage

Let's assume you have Go 1.16+ istalled on your computer.
Execute the following:

```shell
$ mkdir example && cd example
$ go mod init example
$ go get github.com/gyozatech/noodlog
$ touch main.go
```
Open `main.go` and paste the following code:

```golang
package main

import (
    "github.com/gyozatech/noodlog"
)

var log *noodlog.Logger

func init() {
   log = noodlog.NewLogger().SetConfigs(
      noodlog.Configs{
         LogLevel: noodlog.LevelTrace,
         JSONPrettyPrint: noodlog.Enable,
         TraceCaller: noodlog.Enable,
         Colors: noodlog.Enable,
         CustomColors: &noodlog.CustomColors{ Trace: noodlog.Cyan },
         ObscureSensitiveData: noodlog.Enable,
         SensitiveParams: []string{"password"},
      },
    )
}

func main() {
    // simple string message (with custom color)
    log.Trace("Hello world!")
    
    // chaining elements
    log.Info("You've reached", 3, "login attemps")
    
    // using string formatting
    log.Warn("You have %d attempts left", 2)
    
    // logging a struct with a JSON
    log.Error(struct{Code int; Error string}{500, "Generic Error"})
    
    // logging a raw JSON string with a JSON (with obscuring "password")
    log.Info(`{"username": "gyozatech", "password": "Gy0zApAssw0rd"}`)
    
    // logging a JSON string with a JSON (with obscuring "password")
    log.Info("{\"username\": \"nooduser\", \"password\": \"N0oDPasSw0rD\"}")
}
```

Running this example with:
```shell
$ go run main.go
```
You'll get the following output:

![alt text](assets/example.png?raw=true)

## Settings

**Noodlog** allows you to customize the logs through various settings.
You can use various facility functions or the `SetConfigs` function which wraps all the configs together.

----

### LogLevel

To set the logging level, after importing the library with:

```golang
import (
    "github.com/gyozatech/noodlog"
)

var log *noodlog.Logger

func init() {
   log = noodlog.NewLogger()
}
```
you can use the facility method:

```golang
log.LogLevel("warn")
```
or the `SetConfigs` function:

```golang
log.SetConfigs(
    noodlog.Configs{
        LogLevel: noodlog.LevelWarn,
    },
)
```
`log.LevelWarn` is a pre-built pointer to the string "warn".

**The default log level is info**.

----

### JSON Pretty Printing

After importing the library with:

```golang
import (
    "github.com/gyozatech/noodlog"
)

var log *noodlog.Logger

func init() {
   log = noodlog.NewLogger()
}
```

**To enable** pretty printing of the JSON logs you can use:

```golang
log.EnableJSONPrettyPrint()
```
or

```golang
log.SetConfigs(
    noodlog.Configs{
       JSONPrettyPrint: noodlog.Enable,
    },
)
```
`noodlog.Enable` is a pre-built pointer to the bool _true_.

**to disable** pretty printing you can use:

```golang
log.DisableJSONPrettyPrint()
```
or

```golang
log.SetConfigs(
    noodlog.Configs{
       JSONPrettyPrint: noodlog.Disable,
    },
)
```
`noodlog.Disable` is a pre-built pointer to the bool _false_.

**The default value is _false_**. 

----

### Colors

After importing the library with:

```golang
import (
    "github.com/gyozatech/noodlog"
)

var log *noodlog.Logger

func init() {
   log = noodlog.NewLogger()
}
```

**to enable** colors in JSON logs you can use:

```golang
log.EnableColors()
```
or

```golang
log.SetConfigs(
    noodlog.Configs{
        Colors: noodlog.Enable,
    },
)
```
`noodlog.Enable` is a pre-built pointer to the bool _true_.

**To disable** colors you can use:

```golang
log.DisableColors()
```
or

```golang
log.SetConfigs(
    noodlog.Configs{
        Colors: noodlog.Disable,
    },
)
```
`noodlog.Disable` is a pre-built pointer to the bool _false_.

**The default value is _false_**. 

#### Color

The **basic** way to use a custom color is declaring using a pointer of a string representing the color. <br />
`log.Cyan`, `log.Green`, `log.Default`, `log.Yellow`, `log.Purple`, `log.Red`, `log.Blue` are pre-build pointers to the strings "cyan", "green", "default", "yellow", "purple", "red", "blue".

For instance, you can customize trace color by:

```golang
log.SetTraceColor(noodlog.Cyan)
```

A more detailed explanation of each log level is available later into this section.

##### Composition of a color

Color can be composed with text color and background color.
For each level it can be composed using a string or a true color notation.

**Trivial** usage is creating a new color like: 
```golang
log.NewColor(noodlog.Red)
```
It results a red text on default background

Adding a background color can be done through:
```golang
log.NewColor(noodlog.Red).Background(noodlog.Cyan)
```
In this scenario it prints red text on cyan background

A third option is to edit just background color using default text color:
```golang
log.Background(noodlog.Cyan)
```
A list of pre-built pointer of a string is [here](#Composition of a color).

Library provides also more customization through the usage of true color notation (RGB value).
Before the usage of this notation, please consider if your terminal supports truecolor.
For instance if you execute (printf required):

```bash
printf '\033[38;2;255;0;0mHello World\033[0m'
```
a red text "Hello World" should be displayed on the screen

In this way a wider set of color is available for logging, besides of the previous way it can be created a color as:
```golang
log.NewColorRGB(255,0,0).BackgroundRGB(0,0,255)
```
Where a red text (255 for red, 0 the others) is showed on blue background (255 for blue, 0 for others).

As in the previous scenario, ``NewColorRGB`` and ``BackgroundRGB`` hasn't to be executed combined.

Color can be used to set color of Trace log, by typing:
```golang
log.SetTraceColor(noodlog.NewColorRGB(255,0,0).BackgroundRGB(0,0,255))
```

**You can customize the single colors** (for log level) by using:

```golang
log.SetTraceColor(noodlog.Cyan)
log.SetDebugColor(noodlog.NewColorRGB(255,255,0))
log.SetInfoColor(noodlog.NewColor(noodlog.Red).Background(noodlog.Cyan))
log.SetWarnColor(noodlog.NewColor(noodlog.Green).BackgroundRGB(0,255,255))
log.SetErrorColor(noodlog.NewColorRGB(128,255,0).Background(noodlog.Purple))
```
or
```golang
log.SetConfigs(
    noodlog.Configs{
        Colors: noodlog.Enable,
        CustomColors: &noodlog.CustomColors{ 
            Trace: noodlog.Cyan, 
            Debug: noodlog.NewColorRGB(255,255,0),
            Info:  noodlog.NewColor(noodlog.Red).Background(noodlog.Cyan),
            Warn:  noodlog.NewColor(noodlog.Green).BackgroundRGB(0,255,255),
            Error: noodlog.NewColorRGB(128,255,0).Background(noodlog.Purple),    
        },
    },
)
```

Here we highlight all the different combination available to customize colors.

When enabled, the **default colors** are:
- _trace_: "default"
- _info_:  "default"
- _debug_: "green"
- _warn_:  "yellow"
- _error_: "red"

----

### Trace the caller

Noodles allows you to print the file and the function which are calling the log functions.

After importing the library with:

```golang
import (
    "github.com/gyozatech/noodlog"
)

var log *noodlog.Logger

func init() {
   log = noodlog.NewLogger()
}
```

**to enable** the trace caller you can use:
```golang
log.EnableTraceCaller()
```
or
```golang
log.SetConfigs(
    noodlog.Configs{
        TraceCaller: noodlog.Enable,
    },
)
```
`noodlog.Enable` is a pre-built pointer to the bool _true_.

**To disable** it:
```golang
log.DisableTraceCaller()
```
or
```golang
log.SetConfigs(
    noodlog.Configs{
        TraceCaller: noodlog.Disable,
    },
)
```
`noodlog.Disable` is a pre-built pointer to the bool _false_.

The **default value** is _false_.

**Important**: if you want to import **noodlog** only in one package of your project (in order to configure it once) and wraps the logging functions you can use the `EnableSinglePointTracing` to trace file and function the real caller and not of your logging package.

For example:

`main.go`
```golang
package main

import (
   log "example/logging"
)

func main() {
   // main.main real caller we want to track 
   log.Info("Hello folks!")
}
```

`logging/logger.go`
```golang
package logging

import (
    "github.com/gyozatech/noodlog"
)

var l *noodlog.Logger

func init() {
    l = noodlog.NewLogger()
    // configure logger once
    l.SetConfig(
        noodlog.Configs{
         TraceCaller: noodlog.Enable,
         SinglePointTracing: noodlog.Enable,
      },
    )
}

// wrapper function
func Info(message ...interface{}) {
    // if we wouldn't enable SinglePointTracing
    // logger.Info would have been considered the caller to be tracked
    l.Info(message...)
}
```

----

### Sensitive params

Noodlog gives you the possibility to enable the **obscuration of sensitive params when recognized in the JSON structures** (not in the simple strings that you compose).

After importing the library with:

```golang
import (
    "github.com/gyozatech/noodlog"
)

var log *noodlog.Logger

func init() {
   log = noodlog.NewLogger()
}
```

You can **enable** the sensitive params obscuration with the facility methods:

```golang
log.EnableObscureSensitiveData([]string{"param1", "param2", "param3"})
```
or with the `SetConfig` function:

```golang
log.SetConfigs(
    noodlog.Configs{
        ObscureSensitiveData: noodlog.Enable,
        SensitiveParams: []string{"param1", "param2", "param3"},
    },
)
```
Where `noodlog.Enable` is a pre-built pointer to the bool _true_.

**To disable** the sensitive params obscuration you can set:

```golang
log.DisableObscureSensitiveData()
```
or 
```golang
log.SetConfigs(
    noodlog.Configs{
        ObscureSensitiveData: noodlog.Disable,
    },
)
```
Where `noodlog.Disable` is a pre-built pointer to the bool _false_.

The *default* value for the obscuration is _false_.

----

## Contribute to the project

If you want to contribute to the project follow the following [guidelines](https://github.com/gyozatech/noodlog/blob/main/CONTRIBUTING.md).
Any form of contribution is encouraged!
