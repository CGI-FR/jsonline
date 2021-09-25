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
person2 := template.UnmarshalJSON(b)

// Extra field that are not defined in the template will appear at the end of the JSONLine.
person3 := template.Create(map[string]interface{}{"name":"Alice", "extra":true, "age":17, "birthdate":time.Date(2004, time.June, 15, 21, 8, 47, 0, time.UTC)})
fmt.Println(person3) // {"name":"Alice","age":17,"birthdate":"2004-06-15T21:08:47Z","extra":true}
```

### Read and write JSONLine

```go
// exporter
exporter := jsonline.NewExporter(os.Stdout).WithTemplate(template) // or template.GetExporter(os.Stdout)
exporter.Export([]interface{}{"Dorothy", 30, time.Date(1991, time.September, 24, 21, 21, 0, 0, time.UTC)})
exporter.Export([]interface{}{"Alice", 17, time.Date(2004, time.June, 15, 21, 8, 47, 0, time.UTC)})
```

## Command Line Usage

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

## License

This work is not licensed yet
