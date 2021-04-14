# Noodlog

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 
[![CodeFactor](https://www.codefactor.io/repository/github/alessandroargentieri/noodlog/badge)](https://www.codefactor.io/repository/github/alessandroargentieri/noodlog)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
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
    log "github.com/gyozatech/noodlog"
)

func init() {
   log.SetConfigs(
      log.Configs{
         LogLevel: log.LevelTrace,
         JSONPrettyPrint: log.Enable,
         TraceCaller: log.Enable,
         Colors: log.Enable,
         CustomColors: &log.CustomColors{ Trace: log.Cyan },
         ObscureSensitiveData: log.Enable,
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
    log "github.com/gyozatech/noodlog"
)
```
you can use the facility method:

```golang
log.LogLevel("warn")
```
or the `SetConfigs` function:

```golang
log.SetConfigs(
    log.Configs{
        LogLevel: log.LevelWarn,
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
    log "github.com/gyozatech/noodlog"
)
```

**To enable** pretty printing of the JSON logs you can use:

```golang
log.EnableJSONPrettyPrint()
```
or

```golang
log.SetConfigs(
    log.Configs{
    JSONPrettyPrint: log.Enable,
    },
)
```
`log.Enable` is a pre-built pointer to the bool _true_.

**to disable** pretty printing you can use:

```golang
log.DisableJSONPrettyPrint()
```
or

```golang
log.SetConfigs(
    log.Configs{
    JSONPrettyPrint: log.Disable,
    },
)
```
`log.Disable` is a pre-built pointer to the bool _false_.

**The default value is _false_**. 

----

### Colors

After importing the library with:

```golang
import (
    log "github.com/gyozatech/noodlog"
)
```

**to enable** colors in JSON logs you can use:

```golang
log.EnableColors()
```
or

```golang
log.SetConfigs(
    log.Configs{
        Colors: log.Enable,
    },
)
```
`log.Enable` is a pre-built pointer to the bool _true_.

**To disable** colors you can use:

```golang
log.DisableColors()
```
or

```golang
log.SetConfigs(
    log.Configs{
        Colors: log.Disable,
    },
)
```
`log.Disable` is a pre-built pointer to the bool _false_.

**The default value is _false_**. 

**You can customize the single colors** (for log level) by using:

```golang
log.SetTraceColor("cyan")
log.SetDebugColor("green")
log.SetInfoColor("default")
log.SetWarnColor("yellow")
log.SetErrorColor("purple")
```
or
```golang
log.SetConfigs(
    log.Configs{
        Colors: log.Enable,
        CustomColors: &log.CustomColors{ 
            Trace: log.Cyan, 
            Debug: log.Green,
            Info:  log.Default,
            Warn:  log.Yellow,
            Error: log.Purple,    
        },
    },
)
```
`log.Cyan`, `log.Green`, `log.Default`, `log.Yellow`, `log.Purple`, `log.Red`, `log.Blue` are pre-build pointers to the strings "cyan", "green", "default", "yellow", "purple", "red", "blue".

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
    log "github.com/gyozatech/noodlog"
)
```

**to enable** the trace caller you can use:
```golang
log.EnableTraceCaller()
```
or
```golang
log.SetConfigs(
    log.Configs{
        TraceCaller: log.Enable,
    },
)
```
`log.Enable` is a pre-built pointer to the bool _true_.

**To disable** it:
```golang
log.DisableTraceCaller()
```
or
```golang
log.SetConfigs(
    log.Configs{
        TraceCaller: log.Disable,
    },
)
```
`log.Disable` is a pre-built pointer to the bool _false_.

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

func init() {
    // configure logger once
    noodlog.SetConfig(
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
    noodlog.Info(message...)
}
```

----

### Sensitive params

Noodlog gives you the possibility to enable the **obscuration of sensitive params when recognized in the JSON structures** (not in the simple strings that you compose).

After importing the library with:

```golang
import (
    log "github.com/gyozatech/noodlog"
)
```

You can **enable** the sensitive params obscuration with the facility methods:

```golang
log.EnableObscureSensitiveData([]string{"param1", "param2", "param3"})
```
or with the `SetConfig` function:

```golang
log.SetConfigs(
    log.Configs{
        ObscureSensitiveData: log.Enable,
        SensitiveParams: []string{"param1", "param2", "param3"},
    },
)
```
Where `log.Enable` is a pre-built pointer to the bool _true_.

**To disable** the sensitive params obscuration you can set:

```golang
log.DisableObscureSensitiveData()
```
or 
```golang
log.SetConfigs(
    log.Configs{
        ObscureSensitiveData: log.Disable,
    },
)
```
Where `log.Disable` is a pre-built pointer to the bool _false_.

The *default* value for the obscuration is _false_.

----

## Contribute to the project

## Benchmark

