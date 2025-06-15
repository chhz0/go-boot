package fields

import "github.com/chhz0/go-boot/pkg/meta/selection"

type Requirements []Requirement

type Requirement struct {
	Operator selection.Operator
	Field    string
	Value    string
}
