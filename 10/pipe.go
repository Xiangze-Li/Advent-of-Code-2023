package puzzle10

type connDirection int8

type pipe struct {
	direction connDirection
	i, j      int
	next      *pipe
}

func convertDirection(c rune) connDirection {
	switch c {
	case '.':
		return conn00
	case '|':
		return conn13
	case '-':
		return conn24
	case 'L':
		return conn12
	case 'F':
		return conn23
	case '7':
		return conn34
	case 'J':
		return conn41
	case 'S':
		return connSS
	default:
		panic("unknown direction" + string(c))
	}
}

const (
	conn00 connDirection = iota
	conn13
	conn24
	conn12
	conn23
	conn34
	conn41
	connSS
)
