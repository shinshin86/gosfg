# gosfg
[Simple favicon generator](https://github.com/shinshin86/simple-favicon-generator) implemented in Go.

## Install
```sh
go install github.com/shinshin86/gosfg@latest
```

## Usage

```sh
gosfg -i logo.png -n "Your site name"
```

You can check the detailed options with the help command.

```
Usage of gosfg:
  -d string
    	Specify output directory. If the directory does not exist, create it. (default "public")
  -displayMode string
    	Specify display mode. (default "standalone")
  -i string
    	[Required] Specify target image.
  -n string
    	Specify your site name.
  -themeColor string
    	Specify theme color. (default "#ffffff")
  -tileColor string
    	Specify tile color. (default "#da532c")
```

## Why do you output favicon as png instead of ico?
Go does not seem to natively support the ico format in stdlib at this time.  
https://pkg.go.dev/image#section-directories
