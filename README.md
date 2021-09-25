# JSONLine container

This repository contains a library that can handle JSONLine format.

## Features

- order of insertion is preserved
- configure format values to JSON types

## Usage

```go
// row
row :=
    jsonline.NewRow().
        Set("address", jsonline.NewValueString("123 Main Street, New York, NY 10030")).
        Set("last-update", jsonline.NewValueDateTime(time.Now()))
fmt.Println(row) // {"address":"123 Main Street, New York, NY 10030","last-update":"2021-09-25T08:51:10+02:00"}

// template
template := jsonline.NewTemplate().WithString("name").WithNumeric("age").WithDateTime("birthdate")
person1 := template.Create([]interface{}{"Dorothy", 30, time.Date(1991, time.September, 24, 21, 21, 0, 0, time.UTC)})
person1.Set("house", row)
b, err := person1.MarshalJSON()
fmt.Println(string(b)) // {"name":"Dorothy","age":30,"birthdate":"1991-09-24T21:21:00Z","house":{"address":"123 Main Street, New York, NY 10030","last-update":"2021-09-25T09:22:54+02:00"}
person2 := template.UnmarshalJSON(b)
person3 := template.Create(map[string]interface{}{"name":"Alice", "age":17, "birthdate":time.Date(2004, time.June, 15, 21, 8, 47, 0, time.UTC), "extra":true})
fmt.Println(person3) // {"name":"Alice","age":17,"birthdate":"2004-06-15T21:08:47Z","extra":true}

// exporter
exporter := jsonline.NewExporter(os.Stdout).WithTemplate(template) // or template.GetExporter(os.Stdout)
exporter.Export([]interface{}{"Dorothy", 30, time.Date(1991, time.September, 24, 21, 21, 0, 0, time.UTC)})
exporter.Export([]interface{}{"Alice", 17, time.Date(2004, time.June, 15, 21, 8, 47, 0, time.UTC)})
```

## License

This work is not licensed yet
