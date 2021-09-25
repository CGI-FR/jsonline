package jsonline

type Template interface {
	WithString(string) Template
	WithNumeric(string) Template
	WithBoolean(string) Template
	WithBinary(string) Template
	WithDateTime(string) Template
	WithTime(string) Template
	WithTimestamp(string) Template
	WithAuto(string) Template
	WithHidden(string) Template

	Create(interface{}) Row

	UnmarshalJSON([]byte) (Row, error) //nolint:stdmethods
}

type template struct {
	empty Row
}

func NewTemplate() Template {
	return &template{NewRow()}
}

func (t *template) WithString(name string) Template {
	t.empty.Set(name, NewValueString(nil))

	return t
}

func (t *template) WithNumeric(name string) Template {
	t.empty.Set(name, NewValueNumeric(nil))

	return t
}

func (t *template) WithBoolean(name string) Template {
	t.empty.Set(name, NewValueBoolean(nil))

	return t
}

func (t *template) WithBinary(name string) Template {
	t.empty.Set(name, NewValueBinary(nil))

	return t
}

func (t *template) WithDateTime(name string) Template {
	t.empty.Set(name, NewValueDateTime(nil))

	return t
}

func (t *template) WithTime(name string) Template {
	t.empty.Set(name, NewValueTime(nil))

	return t
}

func (t *template) WithTimestamp(name string) Template {
	t.empty.Set(name, NewValueTimestamp(nil))

	return t
}

func (t *template) WithAuto(name string) Template {
	t.empty.Set(name, NewValueAuto(nil))

	return t
}

func (t *template) WithHidden(name string) Template {
	t.empty.Set(name, NewValueHidden(nil))

	return t
}

func (t *template) Create(v interface{}) Row {
	result := t.empty.Clone()

	switch values := v.(type) {
	case []interface{}:
		for i, val := range values {
			result.SetIndex(i, NewValueAuto(val))
		}
	case map[string]interface{}:
		for key, val := range values {
			result.Set(key, NewValueAuto(val))
		}
	}

	return result
}

//nolint:stdmethods
func (t *template) UnmarshalJSON([]byte) (Row, error) {
	return nil, nil
}
