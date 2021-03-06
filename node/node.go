package node

import "github.com/jackspirou/chip/typ"

// Node represents a Node descriptor.
type Node interface {
	Type() typ.Type
	String() string
}

type node struct {
	typ typ.Type
	val interface{}
}
