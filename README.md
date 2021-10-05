# JSONLine container

This repository contains a command-line tool and a library that can handle JSONLine format.

## Features

- preserve order of keys
- format values to valid JSON types : string, numeric, boolean
- handle specific format that are not part of the JSON specification : binary, datetime, time, timestamp
- read JSON lines with specific underlying var type (e.g. store the binary string read from JSON inside a int64)
- validate format of input lines

## Command Line Usage

```text
Order keys and enforce format of JSON lines.

Usage:
  jl [flags]

Examples:
  jl -t '{"first":"string","second":"string"}' <dirty.jsonl

Flags:
  -t, --template string    row template definition (-t {"name":"format"} or -t {"name":"format(type)"}) or -t {"name":"format(type):format"})
                           possible formats : string, numeric, boolean, binary, datetime, time, timestamp, auto, hidden
                           possible types : int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool, byte, rune, string, []byte, time.Time, json.Number (default "{}")
  -f, --filename string    name of row template filename (default "./row.yml")
  -v, --verbosity string   set level of log verbosity : none (0), error (1), warn (2), info (3), debug (4), trace (5) (default "error")
      --debug              add debug information to logs (very slow)
      --log-json           output logs in JSON format
      --color string       use colors in log outputs : yes, no or auto (default "auto")
  -h, --help               help for jl
      --version            version for jl
```

### Example use case

Look at this file.

```json
{"title":"Jurassic Park", "year":1993, "release-date": 739828800}
{"year":1999, "release-date": "922910400", "title":"The Matrix", "running-time":136}
{"title":"Titanic", "running-time":"195", "release-date": "1997-12-19T08:00:00-04:00", "director":"James Cameron"}
```

Let's define a template  in a configuration file named `row.yml`, it will help organize columns.

```yaml
columns:
- name: "title"
- name: "director"
- name: "year"
  output: "numeric"
- name: "running-time"
  output: "numeric"
- name: "release-date"
  output: "datetime"
```

Use the `jl` command line to enforce line format.

```json
$ jl <movies.jsonl
{"title":"Jurassic Park","director":null,"year":1993,"running-time":null,"release-date":"1993-06-11T22:00:00+02:00"}
{"title":"The Matrix","director":null,"year":1999,"running-time":136,"release-date":"1999-03-31T22:00:00+02:00"}
{"title":"Titanic","director":"James Cameron","year":null,"running-time":195,"release-date":"1997-12-19T08:00:00-04:00"}
```

Finally, let's improve output display a bit with mlr.

```console
$ jl <movies.jsonl | mlr --j2p --barred cat
+---------------+---------------+------+--------------+---------------------------+
| title         | director      | year | running-time | release-date              |
+---------------+---------------+------+--------------+---------------------------+
| Jurassic Park | -             | 1993 | -            | 1993-06-11T22:00:00+02:00 |
| The Matrix    | -             | 1999 | 136          | 1999-03-31T22:00:00+02:00 |
| Titanic       | James Cameron | -    | 195          | 1997-12-19T08:00:00-04:00 |
+---------------+---------------+------+--------------+---------------------------+
```

Columns definition can also be defined by argument in command line, using the `-t` flag (or `--template`).

```bash
# give the same result as previous command
jl -t '{"title":"","director":"","year":"numeric","running-time":"numeric","release-date":"datetime"}' <movies.jsonl
```

### Sub rows use case

A row definition can contain sub rows.

```yaml
columns:
- name: "title"
- name: "director"

# this is a sub-row definition, it will be added if missing from the input
- name: "producer"
  columns:
    - name: "first-name"
    - name: "last-name"
```

```bash
# template version
jl -t '{"title":"string","director":"","producer":{"first_name":"","last_name":""}}' <movies.jsonl
```

### Specify the underlying struct

Check this file, it stores int64 integers in binary format.

```json
{"value":"AgAAAAAAAAA="}
{"value":"KgAAAAAAAAA="}
{"value":"aGVsbG8="}
```

But one of the lines is invalid (3rd line can't be an integer 64bit because it's only 5 bytes).

```bash
$ # this command doesn't catch the invalid value
$ jl -t '{"value":"binary"}' < file.jsonl
{"value":"AgAAAAAAAAA="}
{"value":"KgAAAAAAAAA="}
{"value":"aGVsbG8="}
```

```bash
$ # this command will catch the invalid value because the value will be cast to int64
$ jl -t '{"value":"binary(int64)"}' < file.jsonl
{"value":"AgAAAAAAAAA="}
{"value":"KgAAAAAAAAA="}
3:54PM ERR failed to process JSON line error="can't import type []uint8 to int64 format: unable to cast value to int64: []uint8([104 101 108 108 111])" line-number=2
```

```yaml
# same effect but with YAML configuration
columns:
  - name: "value"
    output: "binary(int64)"
```

Valid types are : `int`, `int64`, `int32`, `int16`, `int8`, `uint`, `uint64`, `uint32`, `uint16`, `uint8`, `float64`, `float32`, `bool`, `byte`, `rune`, `string`, `[]byte`, `time.Time`, `json.Number`

### Specify a different format between input and output

```yaml
columns:
  # this column will be read as a datetime, and written as a timestamp
  - name: "release-date"
    input: "datetime"
    output: "timestamp"
```

```json
{"release-date": 739828800}
{"release-date": "922910400"}
{"release-date": "1997-12-19T08:00:00-04:00"}
```

```console
$ jl <movies.jsonl
{"release-date":739828800}
{"release-date":922910400}
{"release-date":882532800}
```

## Library Usage

Check the [examples](examples/) folder.

### Import the package

```go
// Add in your go file
import "github.com/cgi-fr/jsonline/pkg/jsonline"
```

### Create rows

A row is like a map, it store key/value pairs, but you decide how it will output the values in JSON.
Different format options are available : String, Numeric, DateTime, Time, Timestamp, Binary, ...

```go
row := jsonline.NewRow()
row.Set("address", jsonline.NewValueString("123 Main Street, New York, NY 10030"))
row.Set("last-update", jsonline.NewValueDateTime(time.Now()))
```

But, unlike a map, the keys will always appear in the order of insertion.
Dates and times will be formatted to RFC3339.

```go
// Will print : {"address":"123 Main Street, New York, NY 10030","last-update":"2021-09-25T08:51:10+02:00"}
fmt.Println(row)
```

### Enforce format with templates

 A template defines a JSONLine structure.

```go
template := jsonline.NewTemplate().WithString("name").WithNumeric("age").WithDateTime("birthdate")
```

The template can create rows for you, give it either a map, or a slice.
Keys order and format will be enforced for every rows created from the template.

```go
person1, err := template.CreateRow([]interface{}{"Dorothy", 30, time.Date(1991, time.September, 24, 21, 21, 0, 0, time.UTC)})
if err != nil {
    fmt.Println(person1) // {"name":"Dorothy","age":30,"birthdate":"1991-09-24T21:21:00Z"}
}
```

Templates can contains sub templates, to nest JSON objects.

```go
template = template.WithRow("house", jsonline.NewTemplate().WithString("address").WithDateTime("last-update"))
person1.Set("house", row)
fmt.Println(person1) // {"name":"Dorothy","age":30,"birthdate":"1991-09-24T21:21:00Z","house":{"address":"123 Main Street, New York, NY 10030","last-update":"2021-09-25T09:22:54+02:00"}}
```

Standard Go interface Marshaler and Unmarshaler are supported.

```go
b, err := person1.MarshalJSON()
fmt.Println(string(b)) // same result as fmt.Println(person1)

person2 := jsonline.NewRow().UnmarshalJSON(b)
fmt.Println(person2) // same result as fmt.Println(person1)
```

Extra field that are not defined in the template will appear at the end of the JSONLine.

```go
person3, err := template.CreateRow(map[string]interface{}{"name":"Alice", "extra":true, "age":17, "birthdate":time.Date(2004, time.June, 15, 21, 8, 47, 0, time.UTC)})
if err != nil {
    fmt.Println(person3) // {"name":"Alice","age":17,"birthdate":"2004-06-15T21:08:47Z","extra":true}
} else {
    fmt.Println("ERROR:", err)
}
```

### Read and write JSONLine

An exporter will write objects as JSON lines into os.Writer.

```go
exporter := jsonline.NewExporter(os.Stdout).WithTemplate(template) // or template.GetExporter(os.Stdout)
exporter.Export([]interface{}{"Dorothy", 30, time.Date(1991, time.September, 24, 21, 21, 0, 0, time.UTC)})
exporter.Export([]interface{}{"Alice", 17, time.Date(2004, time.June, 15, 21, 8, 47, 0, time.UTC)})
```

An importer will read JSON lines from os.Reader.

```go
for importer := template.GetImporter(os.Stdin); importer.Import(); { // or importer := jsonline.NewImporter(os.Stdin).WithTemplate(template)
    row, err := importer.GetRow()
    if err != nil {
        fmt.Println("an error occurred!", err)
    } else {
        fmt.Println(row)
    }
}
```

A streamer will process JSON lines from os.Reader to os.Writer.

```go
importer := template.GetImporter(os.Stdin)
exporter := template.GetExporter(os.Stdout)

streamer := jsonline.NewStreamer(importer, exporter)
streamer.Stream()
```

## License

Copyright (C) 2021 CGI France

JL is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

JL is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

See the [LICENSE](LICENSE) file for more information.

Some files contains a [GPL linking exception](https://en.wikipedia.org/wiki/GPL_linking_exception) to allow linking of modules that are not derived from or based on this library (everything under pkg folder).
