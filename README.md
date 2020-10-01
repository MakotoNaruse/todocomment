# todocomment
[![Go Report Card](https://goreportcard.com/badge/github.com/MakotoNaruse/complexfunc)](https://goreportcard.com/report/github.com/MakotoNaruse/complexfunc)
[![Go test](https://github.com/MakotoNaruse/complexfunc/workflows/Go%20test/badge.svg?branch=master)](https://github.com/MakotoNaruse/complexfunc/actions)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

`todocomment` finds TODO comment which doesn't include link to issue

## example
```go
package a

func f() { 
    // TODO: hoge
	
    // TODO: https://github.com/test/test/issues/123
    // 次回リリースまでに修正する

    // コメント
	
    /* 
        TODO: これからやること 
        githubのリンクはこれから作るよ 
    */
}
```
```console
./a.go:4:2: todo comment must contains issue's link
./a.go:11:2: todo comment must contains issue's link
```

## Install

```sh
$ go get github.com/MakotoNaruse/complexfunc/cmd/todocomment
```

## Usage

```sh
$ go vet -vettool=`which todocomment` [flag] pkgname
Flags:
      -issue s must contain issue's link (ex. https://github.com/test/test/issues/)
```
