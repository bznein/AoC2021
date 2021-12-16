package parser

import (
	"fmt"

	"github.com/bznein/AoC2021/pkg/math"
)

const (
	sum = iota
	product
	min
	max
	lit
	gt
	lt
	eq
)

type parser struct {
	bs     []byte
	pos    int
	verSum int
}

func (p parser) VersionSum() int {
	return p.verSum
}

func NewParser(b []byte) parser {
	return parser{
		bs: b,
	}
}

func (p *parser) ParsePacket() int {
	id := p.parseHeader()
	if id == lit {
		return p.parseLiteral()
	}
	vals := p.parseOperator()
	return eval(id, vals)
}

func (p *parser) parseVal(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res <<= 1
		res += int(p.bs[p.pos+i] % 2)
	}
	p.pos += n
	return res
}

func (p *parser) parseHeader() int {
	p.verSum += p.parseVal(3)
	return p.parseVal(3)
}

func (p *parser) parseLiteral() int {
	var res int
	for {
		res <<= 4
		keepGoing := p.parseVal(1) == 1
		res += p.parseVal(4)
		if !keepGoing {
			return res
		}
	}
}

func (p *parser) parseOperator() []int {
	var res []int
	if p.parseVal(1) == 0 {
		width := p.parseVal(15)
		start := p.pos
		for p.pos-start < width {
			res = append(res, p.ParsePacket())
		}
	} else {
		n := p.parseVal(11)
		for i := 0; i < n; i++ {
			res = append(res, p.ParsePacket())
		}
	}
	return res
}

func eval(id int, vals []int) int {
	switch id {
	case sum:
		for i := 1; i < len(vals); i++ {
			vals[0] += vals[i]
		}
	case product:
		for i := 1; i < len(vals); i++ {
			vals[0] *= vals[i]
		}
	case min:
		for i := 1; i < len(vals); i++ {
			vals[0] = math.Min(vals[0], vals[i])
		}
	case max:
		for i := 1; i < len(vals); i++ {
			vals[0] = math.Max(vals[0], vals[i])
		}
	case lt:
		if vals[0] < vals[1] {
			return 1
		}
		return 0
	case gt:
		if vals[0] > vals[1] {
			return 1
		}
		return 0
	case eq:
		if vals[0] == vals[1] {
			return 1
		}
		return 0
	default:
		panic(fmt.Sprintf("Invalid id received: %d", id))
	}
	return vals[0]
}
