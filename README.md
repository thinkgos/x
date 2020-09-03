# go-core-package

[![GoDoc](https://godoc.org/github.com/thinkgos/go-core-package?status.svg)](https://godoc.org/github.com/thinkgos/go-core-package)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/thinkgos/go-core-package?tab=doc)
[![Build Status](https://www.travis-ci.org/thinkgos/go-core-package.svg?branch=master)](https://www.travis-ci.org/thinkgos/go-core-package)
[![codecov](https://codecov.io/gh/thinkgos/go-core-package/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkgos/go-core-package)
![Action Status](https://github.com/thinkgos/go-core-package/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/go-core-package)](https://goreportcard.com/report/github.com/thinkgos/go-core-package)
[![License](https://img.shields.io/github/license/thinkgos/go-core-package)](https://github.com/thinkgos/go-core-package/raw/master/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/thinkgos/go-core-package)](https://github.com/thinkgos/go-core-package/tags)


## Feature 
- [extbase64](#extbase64) base64封装
- [extcert](#extcert) 简易cert封装
- [extio](#extio) 常用io
- [extmath](#extmath) 数学常用
- [extnet](extnet) 常用net方法和自定义net.conn
- [extos](#extos) 常用os封装
- [extrand](#extrand) 常用rand方法
- [extrsa](#extrsa) rsa简易封装
- [extssh](#extssh) ssh简易封装
- [gopool](#gopool) 协程池接口
- [normalize](#normalize) 标准化
- [lib](#lib) 基础库包
    - [algo](#algo) 常用算法加密
    - [bpool](#bpool) 切片缓存池
    - [encrypt](#encrypt) 加密流
    - [password](#password) 密码加密和检查
    - [regex](#regex) 常用正则
    - [ternary](#ternary) 常用三目运算

## Reference package

- [strext](https://github.com/thinkgos/strext) string extend package
- [digital](https://github.com/thinkgos/digital) numerical treatment
- [meter](https://github.com/thinkgos/meter) storage metering,like B,KB,MB,GB,TB,PB,EB
- [container](https://github.com/thinkgos/container) collection container
- [gpool](https://github.com/thinkgos/gpool) gpool is a high-performance and low-cost goroutine pool in Go, use [ants](https://github.com/panjf2000/ants) instead.
- [wheel](https://github.com/thinkgos/wheel) time wheel library, which similar linux time wheel
- [timing](https://github.com/thinkgos/timing) time scheduler

## Third party base package
- [go-humanize](https://github.com/dustin/go-humanize) Go Humans! (formatters for units to human friendly sizes)
- [atomic](https://github.com/uber-go/atomic) Wrapper types for sync/atomic which enforce atomic access
- [multierr](https://github.com/uber-go/multierr) Combine one or more Go errors together
- [archiver](https://github.com/mholt/archiver) Easily create & extract archives, and compress & decompress files of various formats
- [compress](https://github.com/klauspost/compress) Optimized compression packages
## Reference web
- [render](https://github.com/thinkgos/render)  render extract from gin,but with optional build tags,useful for net/http or embedded systems,reduce program size
- [binding](https://github.com/thinkgos/binding)  binding extract from gin,but with optional build tags,useful for net/http or embedded linux systems,reduce program size
- [gin-middlewares](https://github.com/thinkgos/gin-middlewares) middleware for Gin
- [http-middlewares](https://github.com/thinkgos/http-middlewares)  middleware for net/http

