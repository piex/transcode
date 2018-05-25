transcode
=========

transcode 是一个为 go 语言提供的在 UTF-8 和 GBK/GB2312/GB10830/HZGB2312 之间转码的库。

## 安装

最简单的安装方式是直接运行 `go get -u github.com/piex/transcode`。

你也可以克隆下来这个仓库到你的 `$GOPATH/src/github.com/piex/transcode` 目录。

如果你用 vgo 管理包，你可以运行 `vgo get github.com/piex/transcode@latest`。

## 依赖

[golang.org/x/text](https://github.com/golang/text/)

## 使用

要使用该软件包，需要先导入：

```go
import (
  "github.com/piex/transcode"
)
```

该包使用了链式调用。

* 把 GBxxx 转为 UTF-8
  ```go
  s := transcode.FromString(gbkString).Decode("GBK").ToString()
  b := transcode.FromByteArray(gbkByteArray).Decode("GBK").ToByteArray()
  ```

* 把 UTF-8 转为 GBxxx
  ```go
  b := transcode.FromString(utf8String).Encode("GBK").ToByteArray()
  s := transcode.FromByteArray(utf8String).Encode("GBK").ToString()
  ```

### type Trans

Transcode 的接口, 包含 transcode 的所以方法。

```go
type Trans interface {
  // Decode 将 string 转为 UTF-8
  Decode(string) Trans 
  // Encode 将 UTF-8 转为 GBxxx
  Encode(string) Trans 
  // ToByteArray 返回一个比特数组类型的值
  ToByteArray() []byte  array
  // ToString 返回一个字符串类型的值
  ToString() string    
}
```

### func FromString

```go
func FromString(s string) Trans
```
输入一个字符串类型的变量，返回 Trans 的实例。

### func FromByteArray

```go
func FromByteArray(b []byte) Trans
```
输入一个比特数组的类型的变量，返回 Trans 的实例。