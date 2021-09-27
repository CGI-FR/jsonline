# JSONLine container

This repository contains a library that can handle JSONLine format.

## Features

- order of insertion is preserved
- configure format values to JSON types

## Library Usage

### Import the package

```go
// Add in your go file
import "github.com/adrienaury/go-template/pkg/jsonline"
```

### Create rows

```go
// A row is like a map, it store key/value pairs, but you decide how it will output the values in JSON.
// Different format options are available : String, Numeric, DateTime, Time, Timestamp, Binary, ...
row := jsonline.NewRow()
row.Set("address", jsonline.NewValueString("123 Main Street, New York, NY 10030"))
row.Set("last-update", jsonline.NewValueDateTime(time.Now()))

// Unlike a map, the keys will always appear in the order of insertion.
// Dates will be formatted to RFC3339.
// {"address":"123 Main Street, New York, NY 10030","last-update":"2021-09-25T08:51:10+02:00"}
fmt.Println(row)
```

### Enforce format with templates

```go
// A template defines a JSONLine, order and format will be enforced.
template := jsonline.NewTemplate().WithString("name").WithNumeric("age").WithDateTime("birthdate")

// The template can create a row for you, give it either a map, or a slice (use the same order for the values).
person1 := template.Create([]interface{}{"Dorothy", 30, time.Date(1991, time.September, 24, 21, 21, 0, 0, time.UTC)})

// Template can contains sub templates, to nest JSON objects.
template = template.WithRow("house", jsonline.NewTemplate().WithString("address").WithDateTime("last-update"))
person1.Set("house", row)

// Standard Go interface Marshaler and Unmarshaler are supported.
b, err := person1.MarshalJSON()
 // {"name":"Dorothy","age":30,"birthdate":"1991-09-24T21:21:00Z","house":{"address":"123 Main Street, New York, NY 10030","last-update":"2021-09-25T09:22:54+02:00"}
fmt.Println(string(b))
person2 := NewRow().UnmarshalJSON(b)

// Extra field that are not defined in the template will appear at the end of the JSONLine.
person3 := template.Create(map[string]interface{}{"name":"Alice", "extra":true, "age":17, "birthdate":time.Date(2004, time.June, 15, 21, 8, 47, 0, time.UTC)})
fmt.Println(person3) // {"name":"Alice","age":17,"birthdate":"2004-06-15T21:08:47Z","extra":true}
```

### Read and write JSONLine

```go
// An exporter will write objects as JSON lines into os.Writer.
exporter := jsonline.NewExporter(os.Stdout).WithTemplate(template) // or template.GetExporter(os.Stdout)
exporter.Export([]interface{}{"Dorothy", 30, time.Date(1991, time.September, 24, 21, 21, 0, 0, time.UTC)})
exporter.Export([]interface{}{"Alice", 17, time.Date(2004, time.June, 15, 21, 8, 47, 0, time.UTC)})

// An importer will read JSON lines from os.Reader.
importer := jsonline.NewImporter(os.Stdin).WithTemplate(template) // or template.GetImporter(os.Stdin)
for row := importer.Import(); row != null; row = importer.Import() {
    fmt.Println(row)
}

// A streamer will process JSON lines from os.Reader to os.Writer.
streamer := jsonline.NewStreamer(importer, exporter)
streamer.Stream()
```

## Command Line Usage

```text
Order keys and enforce format of JSON lines.

Usage:
  jl [flags]

Examples:
  jl -c '{name: first, type: string}' -c '{name: second, type: string}' <dirty.jsonl
  jl -t '{"first":"string","second":"string"}' <dirty.jsonl

Flags:
  -c, --column stringArray   inline column definition in minified YAML (-c {name: title, type: string})
                             use this flag multiple times, one for each column
                             possible types : string, numeric, boolean, binary, datetime, time, timestamp, row, auto, hidden
  -t, --template string      row template definition in JSON (-t {"title":"string"})
                             possible types : string, numeric, boolean, binary, datetime, time, timestamp, auto, hidden
  -f, --filename string      name of row template filename (default "./row.yml")
  -v, --verbosity string     set level of log verbosity : none (0), error (1), warn (2), info (3), debug (4), trace (5) (default "error")
      --debug                add debug information to logs (very slow)
      --log-json             output logs in JSON format
      --color string         use colors in log outputs : yes, no or auto (default "auto")
  -h, --help                 help for jl
      --version              version for jl
```

### Example use case

Look at this file.

```json
{"title":"Jurassic Park", "year":1993}
{"year":1999, "title":"The Matrix", "running-time":136}
{"title":"Titanic", "running-time":"195", "director":"James Cameron"}
```

Let's define a template, in a configuration file named `row.yml`.

```yaml
columns:
- name: "title"
  type: "string"
- name: "year"
  type: "numeric"
- name: "director"
  type: "string"
- name: "running-time"
  type: "numeric"
```

Use the `jl` command line to enforce line format.

```json
$ jl <movies.jsonl
{"title":"Jurassic Park","year":1993,"director":null,"running-time":null}
{"title":"The Matrix","year":1999,"director":null,"running-time":136}
{"title":"Titanic","year":1999,"director":"James Cameron","running-time":195}
```

Columns definition can also be inlined in the command.

```bash
# give the same result as previous command
jl -c '{name: title, type: string}' -c '{name: year, type: numeric}' -c '{name: director, type: string}' -c '{name: running-time, type: numeric}' <movies.jsonl
```

Or you can use the more compact template version.

```bash
# give the same result as previous command
jl -t '{"title":"string","year":"numeric","director":"string","running-time":"numeric"}' <movies.jsonl
```

### Sub rows

A row definition can contain sub rows.

```yaml
columns:
- name: "title"
  type: "string"
- name: "year"
  type: "numeric"
- name: "director"
  type: "row"
  columns:
    - name: "first_name"
      type: "string"
    - name: "last_name"
      type: "string"
```

```bash
# inline columns version
jl -c '{name: title, type: string}' -c '{name: year, type: numeric}' -c '{name: director, type: row, columns: [{name: first_name, type: string}, {name: last_name, type: string}]}' <movies.jsonl
```

```bash
# template
jl -t '{"title":"string","year":"numeric","director":{"first_name":"string","last_name":"string"}}' <movies.jsonl
```

## License

Copyright (C) 2021 CGI France

LINO is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

LINO is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

See the [LICENSE](LICENSE) file for more information.
