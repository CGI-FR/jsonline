package jsonline_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/adrienaury/go-template/pkg/jsonline"
)

func TestTemplate(t *testing.T) {
	row :=
		jsonline.NewRow().
			Set("address", jsonline.NewValueString("123 Main Street, New York, NY 10030")).
			Set("last-update", jsonline.NewValueDateTime(time.Now()))

	template := jsonline.NewTemplate().WithString("name").WithNumeric("age").WithDateTime("birthdate")
	person1 := template.Create([]interface{}{"Dorothy", 30, time.Date(1991, time.September, 24, 21, 21, 0, 0, time.UTC)})
	person1.Set("house", row)
	fmt.Println(person1)

	person3 := template.Create(
		map[string]interface{}{
			"name":      "Alice",
			"age":       17,
			"birthdate": time.Date(2004, time.June, 15, 21, 8, 47, 0, time.UTC),
			"extra":     true,
		})
	fmt.Println(person3)
}
