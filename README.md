transcode
=========

transcode is a lib transcode between UTF-8 and GBK/GB2312/GB10830 for golang.

## Install

he easiest way to install is to run `go get -u github.com/foreverzmy/transcode`.
You can also manually git clone the repository to $GOPATH/src/github.com/foreverzmy/transcode.

When you use vgo, you can run `vgo get github.com/foreverzmy/transcode@latest`.

## Package Files

[golang.org/x/text](https://github.com/golang/text/)

## Usage

To use the package, you'll need the appropriate import statement:

```go
import (
  "github.com/foreverzmy/transcode"
)
```

The libraries use chained calls.

* Transcode GBxxx to UTF-8
  ```go
  s := transcode.FromString(gbkString).Encode("GBK").ToString()
  ```

* Transcode UTF-8 to GBxxx
  ```go
  s := transcode.FromString(utf8String).Decode("GBK").ToString()
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