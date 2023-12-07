package puzzle19

type workflow struct {
	name  string
	rules []rule
	final string
}

type rule struct {
	attr   byte
	op     byte
	val    int
	target string
}

//nolint:gochecknoglobals // internal
var attrIdx = map[byte]int{
	'x': 0, 'm': 1, 'a': 2, 's': 3,
}
