package fields

import "github.com/chhz0/gpkg/pkg/meta/selection"

type Requirements []Requirement

type Requirement struct {
	Operator selection.Operator
	Field    string
	Value    string
}
