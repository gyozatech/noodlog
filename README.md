# Noodlog

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
         Color: log.Enable,
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

### LogLevel

To set the logging level you can use the facility method:

```golang
log.LogLevel("warn")
```
or

```golang
log.SetConfigs(
      log.Configs{
         LogLevel: log.LevelWarn,
      },
    )

```
`log.LevelWarn` is a pre-built pointer to the string "warn".

**The default log level is info**.

### JSON Pretty Printing

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

**To disable** pretty printing you can use:

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

### Colors

**To enable** colors in JSON logs you can use:

```golang
log.EnableColors()
```
or

```golang
log.SetConfigs(
      log.Configs{
         Color: log.Enable,
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
         Color: log.Disable,
      },
    )

```
`log.Disable` is a pre-built pointer to the bool _false_.

**The default value is _false_**. 

** You can customize the single colors (for log level) by using:

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
         Color: log.Enable,
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

### Trace the caller

### Sensitive params

## Contribute to the project

## Benchmark

