# Instruction for running Lazy-Bot

## Build executable file

### MacOS

if you don't have an executable file (with name `pizzabot`), then run the next command in command line:
```shell
env GOOS=darwin GOARCH=amd64 go build -o pizzabot
```

### Linux

if you don't have an executable file (with name `pizzabot`), then run the next command in command line:
```shell
env GOOS=linux GOARCH=amd64 go build -o pizzabot
```

### Windows

if you don't have an executable file (with name `pizzabot.exe`), then run the next command in command line:
```shell
env GOOS=windows GOARCH=amd64 go build -o pizzabot.exe
```

## Running bot

to run the bot use the next command for macos or linux:
```shell
./pizzabot "5x5 (1, 3) (4, 4)"
```
or this for windows
```shell
pizzabot.exe "5x5 (1, 3) (4, 4)"
```

where `5x5` is a size of pizzabot's filed, and `(1, 3) (4, 4)` is a coordinates of delivery points.

it is an example. you can use your own values for the size and delivery points.

## What is in the result

in the result of work of `pizzabot`script you will have a string alike this:
```
ENNNDEEEND
```

an output depends on from set of delivery points.