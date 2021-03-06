package typ

import "github.com/jackspirou/chip/token"

// Param describes a parameter type.
type Param struct {
	typ  Type
	name string
	next *Param
}

// newParam returns a new parameter of the provided type.
func newParam(typ Type) *Param {
	return &Param{typ: typ, name: typ.String()}
}

// Value returns the parameter's token type.
func (p Param) Value() Type {
	return p.typ
}

// Token returns the token.Type.
func (p Param) Token() token.Type {
	return p.typ.Token()
}

// String impliments the fmt.Stringer interface and returns the name of the
// parameter.
func (p Param) String() string {
	return p.name
}

// Next returns the next parameter or nil.
func (p *Param) Next() *Param {
	return p.next
}
