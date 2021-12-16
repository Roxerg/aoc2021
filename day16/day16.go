package day16

import (
	"fmt"
	"math"
	"roxerg_aoc/utils"
	"strconv"
	"strings"
)

type Packet struct {
	version      uint8
	typeID       uint8
	rawContent   string
	literalValue uint64
	totalLen     int
	innerPackets []*Packet
}

func stringToPacket(in string) Packet {
	v, _ := strconv.ParseUint(in[:3], 2, 3)
	id, _ := strconv.ParseUint(in[3:6], 2, 3)
	c := in[6:]
	newPacket := Packet{uint8(v), uint8(id), c, 0, 3 + 3, []*Packet{}}

	if newPacket.typeID == 4 {
		newPacket.literalValue, newPacket.totalLen = newPacket.readLiteral()
	} else {
		newPacket.innerPackets, newPacket.totalLen = newPacket.readOperator()
	}

	return newPacket
}

func (p Packet) readLiteral() (uint64, int) {
	res := ""
	consumable := p.rawContent
	for len(consumable) > 0 {
		chunk := consumable[:5]
		p.totalLen += len(chunk)
		consumable = consumable[5:]
		res = res + chunk[1:]
		if chunk[0] == '0' {
			break
		}
	}
	p.literalValue, _ = strconv.ParseUint(res, 2, len(res))
	return p.literalValue, p.totalLen
}

func (p Packet) readOperator() ([]*Packet, int) {
	method := p.rawContent[0] == '1'
	// method true   - next 15 bits are a number that represents the total length in
	//	               bits of the sub-packets contained by this packet.
	// method false  - next 11 bits are a number that represents the number of
	//				   sub-packets immediately contained by this packet.

	if !method {
		infoBits := p.rawContent[1:16]
		totalChildrenLen, _ := strconv.ParseUint(infoBits, 2, len(infoBits))

		childrenBits := p.rawContent[16 : 16+totalChildrenLen]

		p.totalLen += 16 + int(totalChildrenLen)
		for len(childrenBits) > 0 {
			newBoi := stringToPacket(childrenBits)
			p.innerPackets = append(p.innerPackets, &newBoi)

			childrenBits = childrenBits[newBoi.totalLen:]
		}

	} else {
		infoBits := p.rawContent[1:12]
		numberOfChildren, _ := strconv.ParseUint(infoBits, 2, len(infoBits))

		childrenBits := p.rawContent[12:]

		p.totalLen += 12
		for n := 0; n < int(numberOfChildren); n++ {
			newBoi := stringToPacket(childrenBits)
			p.innerPackets = append(p.innerPackets, &newBoi)
			childrenBits = childrenBits[newBoi.totalLen:]
			p.totalLen += newBoi.totalLen
		}
	}

	return p.innerPackets, p.totalLen

	// 00111000000000000110111101000101001010010001001000000000
	//       00000000000110111101000101001010010001001000000000
	// VVVTTTILLLLLLLLLLLLLLLAAAAAAAAAAABBBBBBBBBBBBBBBB
	// 010 100 1000100100
}

func (p Packet) countVersion() uint64 {
	verSum := uint64(p.version)
	for _, pp := range p.innerPackets {
		verSum += uint64(pp.countVersion())
	}
	return verSum
}

func (p Packet) Eval() uint64 {
	res := uint64(0)

	switch p.typeID {
	case 1:
		res = 1
	case 2:
		res = math.MaxInt
	case 4:
		res = p.literalValue
	}

	if p.typeID < 4 {
		for _, pp := range p.innerPackets {
			switch p.typeID {
			case 0:
				res += pp.Eval()
			case 1:
				res *= pp.Eval()
			case 2: // minimum
				childVal := pp.Eval()
				if res > childVal {
					res = childVal
				}
			case 3: // maximum
				childVal := pp.Eval()
				if res < childVal {
					res = childVal
				}
			}
		}
	} else if p.typeID > 4 {
		switch p.typeID {
		case 5:
			if p.innerPackets[0].Eval() > p.innerPackets[1].Eval() {
				res = 1
			} else {
				res = 0
			}
		case 6:
			if p.innerPackets[0].Eval() < p.innerPackets[1].Eval() {
				res = 1
			} else {
				res = 0
			}
		case 7:
			if p.innerPackets[0].Eval() == p.innerPackets[1].Eval() {
				res = 1
			} else {
				res = 0
			}
		}
	}

	return res
}

func Run() {
	_, inarr := utils.LoadFile("day16", "\n")
	hex := inarr[0]
	allBits := ""
	for _, h := range strings.Split(hex, "") {
		hh, _ := strconv.ParseUint(h, 16, 32)
		allBits = allBits + asBits(hh)
	}

	initialPacket := stringToPacket(allBits)
	// part 1
	fmt.Println(initialPacket.countVersion())
	// part 2
	fmt.Println(initialPacket.Eval())
}

func asBits(val uint64) string {
	bits := ""
	for i := 0; i < 4; i++ {
		bits = strconv.Itoa(int(val&0x1)) + bits
		val = val >> 1
	}
	return bits
}
