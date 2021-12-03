package constraints

import "constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

type Addable interface {
	constraints.Integer | constraints.Float | constraints.Complex | ~string
}
