---
layout: go-sf
---
## 介绍

go-sf 是一个 Dpp Service 框架生成工具。

go-sf 是一个合成词，Go 代表这个工具是用 Go 编写，sf 是『Service Frame』（微服务框架代码）的缩写。

这是一个简单的提升开发效率的工具，由 MSBU Tech 成员利用业余时间完成。使用 Go 编写主要是为了练手，并没有什么特别的原因。如果有其它的实现，也欢迎分享出来。

## 安装

```shell
go get github.com/msbu-tech/go-sf
```

## 使用方法

```shell
go-sf -new -name=doctor \
      -author=msbu \
      -template=template/service \
      -output=/path/to/output
```

## 问题反馈

如果遇到 bug 或疑问，可以[新建 Issue](https://github.com/msbu-tech/go-sf/issues/new) 进行反馈。

## 贡献代码

<a href="https://github.com/msbu-tech/go-sf">msbu-tech/go-sf</a>
