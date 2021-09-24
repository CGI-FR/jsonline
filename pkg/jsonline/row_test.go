package jsonline_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/adrienaury/go-template/pkg/jsonline"
)

func TestRow(t *testing.T) {
	r := jsonline.NewRow()
	r.Set("name", jsonline.NewString("Dorothy"))
	r.Set("age", jsonline.NewNumeric(30))
	r.Set("car", jsonline.NewNull())
	r.Set("pet", jsonline.NewRow())
	r.Set("house", jsonline.NewRow().
		Set("address", jsonline.NewString("123 Main Street, New York, NY 10030")).
		Set("buy-date", jsonline.NewDateTime(time.Now())))
	fmt.Println(r.String())
}
