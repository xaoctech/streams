package streams

import "constraints"

type Number interface {
	constraints.Integer | constraints.Float
}
