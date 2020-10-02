# todocomment
[![Go Report Card](https://goreportcard.com/badge/github.com/MakotoNaruse/todocomment)](https://goreportcard.com/report/github.com/MakotoNaruse/todocomment)
[![Go test](https://github.com/MakotoNaruse/todocomment/workflows/Go%20test/badge.svg)](https://github.com/MakotoNaruse/todocomment/actions)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

`todocomment` finds TODO comment which doesn't include link to issue

## example
```
package a

func f() { 
    // TODO: hoge <- NG
	
    // TODO: hoge 
    // nolint: todocomment
	
    // TODO: comment
    // https://github.com/test/test/issues/ 

    // comment
	
    /* TODO github issue will be created */ <- NG
	
    /*
        TODO github issue will be created
        nolint: todocomment
    */
}

// TODO is a stuct name
// nolint: todocomment
type TODO struct {
}

```
```console
./a.go:4:2: todo comment must contains issue's link
./a.go:14:2: todo comment must contains issue's link
```

## Install

```sh
% go get github.com/MakotoNaruse/todocomment/cmd/todocomment
```

## Usage

```sh
% go vet -vettool=`which todocomment` [flag] pkgname
Flags:
      --issue s must contain issue's link (ex. https://github.com/test/test/issues/)
```

## Example
```sh
% go vet -vettool=$(which todocomment) --issue "github.com/MakotoNaruse/todocomment/issues" ./...
% go vet -vettool=$(which todocomment) --issue "trello.com/b/aaaaa/" main.go
% go vet -vettool=$(which todocomment) --issue "issue" ./src
```

