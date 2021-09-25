package jsonline_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/adrienaury/go-template/pkg/jsonline"
)

func TestRow(t *testing.T) {
	r := jsonline.NewRow()
	r.Set("name", jsonline.NewValueString("Dorothy"))
	r.Set("age", jsonline.NewValueNumeric(30))
	r.Set("car", jsonline.NewValueNil())
	r.Set("pet", jsonline.NewRow())
	r.Set("house", jsonline.NewRow().
		Set("address", jsonline.NewValueString("123 Main Street, New York, NY 10030")).
		Set("buy-date", jsonline.NewValueDateTime(time.Now())))
	fmt.Println(r.String())

	row :=
		jsonline.NewRow().
			Set("address", jsonline.NewValueString("123 Main Street, New York, NY 10030")).
			Set("last-update", jsonline.NewValueDateTime(time.Now()))
	fmt.Println(row)
}
