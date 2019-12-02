// Package fixedwidth provides encoding and decoding for fixed-width formatted Data.
package fixedwidth

type Marshaler interface {
	MarshalFixedWidth(FieldSpec) ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalFixedWidth([]byte) error
}

type FieldSpec struct {
	StartPos, EndPos int
}
