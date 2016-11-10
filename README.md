# runss

Get DynamoDB Item.

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