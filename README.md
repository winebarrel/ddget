# ddget

Get DynamoDB Item.

[![Build Status](https://travis-ci.org/winebarrel/ddget.svg?branch=master)](https://travis-ci.org/winebarrel/ddget)

## Usage

```
Usage of ddget:
  -k string
      Item key
  -n  Do not print newline
  -t string
      Table name
  -v string
      Value attribute name
```

```
$ ddbcli
ap-northeast-1> select all * from my-table;
[
  {"key":"foo","value":"bar"},
  {"key":"zoo","value":"baz"}
]

$ ddget -t my-table -k foo
bar
```

```
$ ddbcli
ap-northeast-1> select all * from my-table;
[
  {"key":"foo","value":"bar","value1":"zoo"}
]

$ ddget -t my-table -k foo -v value1
zoo
```

## Installation

```sh
wget https://github.com/winebarrel/ddget/releases/download/vX.X.X/ddget-vX.X.X-linux-amd64.gz
gunzip -c ddget-vX.X.X-linux-amd64.gz > ddget
chmod +x ddget
```

```sh
brew install https://raw.githubusercontent.com/winebarrel/ddget/master/homebrew/ddget.rb
```
