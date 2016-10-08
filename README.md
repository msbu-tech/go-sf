# go-sf

![](https://img.shields.io/badge/license-MIT-blue.svg)
[![](https://api.travis-ci.org/msbu-tech/go-sf.svg)](https://travis-ci.org/msbu-tech/go-sf)
[![Code Climate](https://codeclimate.com/github/msbu-tech/go-sf/badges/gpa.svg)](https://codeclimate.com/github/msbu-tech/go-sf)

Service Frame generator for Dpp in Go. <https://msbu-tech.github.io/go-sf/>

## Install

```
go get github.com/msbu-tech/go-sf/cmd/go-sf
```

You need to set `$GOPATH`:

```
$ mkdir $HOME/work
$ export GOPATH=$HOME/work
```

And then add `$GOPATH` to your `$PATH`:

```
$ export PATH=$PATH:$GOPATH/bin
```

See <https://golang.org/doc/code.html>

## Usage

```
go-sf -new -name=doctor \
      -author=msbu \
      -template=template/service \
      -output=/path/to/output
```

## License

MIT