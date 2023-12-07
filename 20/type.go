package puzzle20

type modType int8

const (
	dumb modType = iota
	broadcaster
	flipflop
	conjunction
)

type pulse struct {
	from, to string
	level    byte
}
