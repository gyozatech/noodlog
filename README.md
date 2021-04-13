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
         CustomColors: &log.CustomColors{ Trace: log.Purple },
         ObscureSensitiveData: log.Enable,
         SensitiveParams: []string{"password"},
      },
    )
}

func main() {
    log.Trace("Hello world!")
    log.Info("You've reached", 3, "login attemps")
    log.Warn("You have %d attempts left", 2)

    log.Error(struct{Code int; Error string}{500, "Generic Error"})
    log.Info(`{"username": "gyozatech", "password": "Gy0zApAssw0rd"}`)
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

### JSON Pretty Printing

### Colors

### Trace the caller

### Sensitive params

## Contribute to the project

## Benchmark

