# x
 go stand library extend api

[![GoDoc](https://godoc.org/github.com/thinkgos/x?status.svg)](https://godoc.org/github.com/thinkgos/x)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/thinkgos/x?tab=doc)
[![Build Status](https://www.travis-ci.org/thinkgos/x.svg?branch=master)](https://www.travis-ci.org/thinkgos/x)
[![codecov](https://codecov.io/gh/thinkgos/x/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkgos/x)
![Action Status](https://github.com/thinkgos/x/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/x)](https://goreportcard.com/report/github.com/thinkgos/x)
[![License](https://img.shields.io/github/license/thinkgos/x)](https://github.com/thinkgos/x/raw/master/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/thinkgos/x)](https://github.com/thinkgos/x/tags)


## Feature 
- [extbase64](#extbase64) base64封装
- [extcert](#extcert) 简易cert封装
- [extime](#extime) 常用time封装
- [extimg](#extimg) 常用图片封装
- [extio](#extio) 常用io
- [extjson](#extjson) 常用json
- [extmath](#extmath) 数学常用
- [extnet](extnet) 常用net方法和自定义net.conn
- [extos](#extos) 常用os封装
- [extrand](#extrand) 常用rand方法
- [extrsa](#extrsa) rsa简易封装
- [extssh](#extssh) ssh简易封装
- [extstr](#extstr) 常用字符串封装
- [gopool](#gopool) 协程池接口
- [normalize](#normalize) 标准化
- [numeric](#numeric) 数值类型封装
- [lib](#lib) 基础库包
    - [algo](#algo) 常用算法加密
    - [bpool](#bpool) 切片缓存池
    - [encrypt](#encrypt) 加密流
    - [habit](#habit) 常用无法分类的常用库
    - [logger](#logger) 标准日志及日志接口
    - [parallel](#parallel) 有限制的并行控制库
    - [password](#password) 密码加密和检查
    - [regex](#regex) 常用正则
    - [ternary](#ternary) 常用三目运算
    - [textcolor](#textcolor) 字符串文字简易加色
    - [track](#track) 跟综goroutine的生存,死亡.
    - [univ](#univ) 综合库

## Reference package

- [meter](https://github.com/thinkgos/meter) storage metering,like B,KB,MB,GB,TB,PB,EB
- [container](https://github.com/thinkgos/container) collection container
- [wheel](https://github.com/thinkgos/wheel) time wheel library, which similar linux time wheel
- [timing](https://github.com/thinkgos/timing) time scheduler

## Third party base package
- [go-internal](https://github.com/rogpeppe/go-internal) Selected Go-internal packages factored out from the standard library
- [go-humanize](https://github.com/dustin/go-humanize) Go Humans! (formatters for units to human friendly sizes)
- [atomic](https://github.com/uber-go/atomic) Wrapper types for sync/atomic which enforce atomic access
- [multierr](https://github.com/uber-go/multierr) Combine one or more Go errors together
- [cast](https://github.com/spf13/cast) safe and easy casting from one type to another in Go
- [archiver](https://github.com/mholt/archiver) Easily create & extract archives, and compress & decompress files of various formats
- [inflection](https://github.com/jinzhu/inflection) Pluralizes and singularizes English nouns
- [compress](https://github.com/klauspost/compress) Optimized compression packages
- [strcase](https://github.com/iancoleman/strcase) A golang package for converting to snake_case or CamelCase
- [color](https://github.com/fatih/color) Color package for Go
- [runewidth](https://github.com/mattn/go-runewidth) Provides functions to get fixed width of the character or string.

- [ntp](https://github.com/beevik/ntp) a simple ntp client package for go
## Reference web
- [render](https://github.com/thinkgos/render)  render extract from gin,but with optional build tags,useful for net/http or embedded systems,reduce program size
- [binding](https://github.com/thinkgos/binding)  binding extract from gin,but with optional build tags,useful for net/http or embedded linux systems,reduce program size
- [schema](https://github.com/gorilla/schema) Package gorilla/schema fills a struct with form values.
- [gin-middlewares](https://github.com/thinkgos/gin-middlewares) middleware for Gin
- [http-middlewares](https://github.com/thinkgos/http-middlewares)  middleware for net/http
- [user_agent](https://github.com/mssola/user_agent) HTTP User Agent parser for the Go programming language.
- [resty](https://github.com/go-resty/resty) Simple HTTP and REST client library for Go
- [gout](https://github.com/guonaihong/gout) gout to become the Swiss Army Knife of the http client 
## other 
- [ants](https://github.com/panjf2000/ants) ants is a high-performance and low-cost goroutine pool in Go
- [treeprint](https://github.com/xlab/treeprint) Package treeprint provides a simple ASCII tree composing tool.
- [tablewriter](https://github.com/olekukonko/tablewriter) Generate ASCII table on the fly
- [promptui](https://github.com/manifoldco/promptui) Interactive prompt for command-line applications
- [go-prompt](https://github.com/c-bata/go-prompt) Building powerful interactive prompts in Go, inspired by python-prompt-toolkit.
- [cidranger](https://github.com/yl2chen/cidranger) Fast IP to CIDR lookup in Golang
- [go-version](https://github.com/hashicorp/go-version) A Go (golang) library for parsing and verifying versions and version constraints.

## id
- [nuid](https://github.com/nats-io/nuid) NATS Unique Identifiers
- [xid](https://github.com/rs/xid) xid is a globally unique id generator thought for the web
- [snowflake](github.com/bwmarrin/snowflake) A simple to use Go (golang) package to generate or parse Twitter snowflake IDs
## image
- [barcode](github.com/boombuler/barcode) This is a package for GO which can be used to create different types of barcodes.
- [qrcode](github.com/skip2/go-qrcode) QR Code encoder (Go)
- [gozxing](https://github.com/makiuchi-d/gozxing) ZXing is an open-source, multi-format 1D/2D barcode image processing library for Java. This project is a port of ZXing core library to pure Go.

## Donation

if package help you a lot,you can support us by:

**Alipay**

![alipay](https://github.com/thinkgos/thinkgos/blob/master/asserts/alipay.jpg)

**WeChat Pay**

![wxpay](https://github.com/thinkgos/thinkgos/blob/master/asserts/wxpay.jpg)