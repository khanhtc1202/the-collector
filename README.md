# Boogeyman

[![][goreportcard-svg]][goreportcard] 
[![][CodeFactor]](https://www.codefactor.io/repository/github/khanhtc1202/boogeyman)
[![][Build Status]](https://travis-ci.org/khanhtc1202/boogeyman)

[Build Status]: https://travis-ci.org/khanhtc1202/boogeyman.svg?branch=master
[CodeFactor]: https://www.codefactor.io/repository/github/khanhtc1202/boogeyman/badge
[goreportcard]: https://goreportcard.com/report/github.com/khanhtc1202/boogeyman
[goreportcard-svg]: https://goreportcard.com/badge/github.com/khanhtc1202/boogeyman

A simple program that help you get search results from multi search engines instead of just from google. Try it [here](https://search.khanhtc.me/).

## What can it does

![showcase](asserts/sample.gif)

This program searches through multi search engines and returns search results under some of strategies:

> Top

Return top result of each search engines. 

> Cross Matching

Return matched results cross through multi search engines. (Appeared in 2 or more search engines)

> All (with limit 20)

Return all :)

On ver `1.2.6`, search engines list:

1. Ask
2. Bing
3. Google
4. Yahoo

This list will be updated by far ~

## The design

The program's architecture implemented under `the clean architecture` design. More info go [here](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html).

![boogeymain design](asserts/boogeyman_design.jpg)

## Usage

In case exec file you downloaded's name is `boogeyman`.

Sample full params command

```bash
$ ./boogeyman -e=bing -s=top -k="some anything"
```

Type `-h` to get help. Return value be like

```$xslt
Usage of ./boogeyman:
  -e string
        search engine(s): google | bing | ask | yahoo | all (default "all")
  -k string
        search (query) string (default "bar")
  -s string
        result show strategy: top | cross | all (default "top")
  -v    show application version
  -version
        show application version
```

## Run on local

Boogeyman development environment requires: 

1. Golang (1.9.2 or higher). Install go [here](https://golang.org/doc/install).
2. dep (Go dependency management tool). Install go [here](https://github.com/golang/dep).
3. go-bindata (Go generate bindata from template file). Install go [here](https://github.com/jteeuwen/go-bindata).

Run by `go`

```bash
$ go run main.go
```

or check [Makefile](https://github.com/khanhtc1202/boogeyman/blob/master/Makefile) for building bin on your local.

## Contribution

All contributions are welcomed in this project.

## Download

For linux x64 : [download](bin/boogeyman-linux-64)

For MacOS x64 : [download](bin/boogeyman-darwin-64)

## License
The MIT License (MIT). Please see [LICENSE](LICENSE) for more information.
