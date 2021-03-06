transcode
=========

transcode is a lib transcode between UTF-8 and GBK/GB2312/GB10830/HZGB2312 for golang.

[中文文档](https://github.com/piex/transcode/blob/master/README-CH.md)

## Install

The easiest way to install is to run `go get -u github.com/piex/transcode`.

You can also manually git clone the repository to $GOPATH/src/github.com/piex/transcode.

When you use vgo, you can run `vgo get github.com/piex/transcode@latest`.

## Package Files

[golang.org/x/text](https://github.com/golang/text/)

## Usage

To use the package, you'll need the appropriate import statement:

```go
import (
  "github.com/piex/transcode"
)
```

The libraries use chained calls.

* Transcode GBxxx to UTF-8
  ```go
  s := transcode.FromString(gbkString).Decode("GBK").ToString()
  b := transcode.FromByteArray(gbkByteArray).Decode("GBK").ToByteArray()
  ```

* Transcode UTF-8 to GBxxx
  ```go
  b := transcode.FromString(utf8String).Encode("GBK").ToByteArray()
  s := transcode.FromByteArray(utf8String).Encode("GBK").ToString()
  ```

### type Trans

An interface of Transcode, inclueds the func for transcode.

```go
type Trans interface {
  // Decode decode string to UTF-8
  Decode(string) Trans 
  // Encode encode UTF-8 to GBxxx
  Encode(string) Trans 
  // ToByteArray return result as byte
  ToByteArray() []byte  array
  // ToString return result as string
  ToString() string    
}
```

### func FromString

```go
func FromString(s string) Trans
```
Input a string type variable to transcode UTF-8 between GBxxx.

### func FromByteArray

```go
func FromByteArray(b []byte) Trans
```
Input a byte array type variable to transcode UTF-8 between GBxxx. 