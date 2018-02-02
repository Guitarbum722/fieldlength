# qualified

A simple library to deal with delimited strings that _could_ contain the delimiter in the data itself.


### Features
* Split strings by a delimiter and escape the delimiter character if contained in a qualified text field.
    * ex. `one,two,"three,four",five`

```go
fields := qualified.SplitWithQual("Ohhi, mark", ",", "")
fmt.Println(fields) // [Ohhi  mark]
```

* Get a map that contains the length of each field

```go
lens := qualified.FieldLengths("\"Ohhi, mark\",\"Hey bud\"", ",", "\"")
fmt.Println(lens) // map[0:12 1:9]
```