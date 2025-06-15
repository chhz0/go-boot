package errcodegen

type ErrMeta struct {
	Name  string
	Value int
	Desc  string
	// extra fields can be added as needed
	Category string
	Template string
}
