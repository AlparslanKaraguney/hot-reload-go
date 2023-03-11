# This is a working example of cosmtrek air for go hot reload.

This example shows how to properly setup cosmtrek air library for your go project. It also shows how to use .env file in this setup. Currently i managed to get env variables using a third part library so that .env file autmatically set on project start. I hope it will help some one...

First you must get the library:
```
go get -u github.com/cosmtrek/air
```

Then type to create .air.toml file: 
```
air init
```

Edit .air.toml file for your project. For example i have to edit following lines to run correctly for this project setup:

First i have to edit `cmd`:
##
From:
##
```
"cmd = "go build -o ./tmp/main.exe ."
```

To:
##
```
cmd = "go build -o ./tmp/main.exe ./cmd/"
```

Since i run this example in windows i did not changed `bin` command in `.air.toml` file but in different operating systems you my also have to change this one. Mine as follows for this project setup: (care for `\\` this might change)
```
bin = "tmp\\main.exe"
```

