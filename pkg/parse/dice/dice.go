package dice

import (
	"regexp"
	"strconv"
)

type DieRollExpression struct {
	Rolls        int
	DieSides     int
	Multiplicand int
	Addend       int

	ShouldMultiply bool
	ShouldAdd      bool
}

func (d DieRollExpression) ToString() string {
	countString := strconv.Itoa(d.Rolls)
	dieIndicator := "d"
	sidesString := strconv.Itoa(d.DieSides)
	if d.DieSides == 0 {
		dieIndicator = ""
		sidesString = ""
	}
	multiplicationSymbol := "/"
	if d.ShouldMultiply {
		multiplicationSymbol = "*"
	}
	multiplier := strconv.Itoa(d.Multiplicand)
	if d.Multiplicand == 0 {
		multiplier = ""
		multiplicationSymbol = ""
		dieIndicator = ""
	}
	if d.Multiplicand == 1 {
		multiplier = ""
		multiplicationSymbol = ""
	}
	addSymbol := "-"
	if d.ShouldAdd {
		addSymbol = "+"
	}
	adder := strconv.Itoa(d.Addend)
	if d.Addend == 0 {
		addSymbol = ""
		adder = ""
	}
	if d.Addend < 0 {
		addSymbol = "-"
		if !d.ShouldAdd {
			addSymbol = "+"
		}
		adder = strconv.Itoa(-d.Addend)
	}
	return countString + dieIndicator + sidesString + multiplicationSymbol + multiplier + addSymbol + adder
}

var DieParseRegex = regexp.MustCompile(`^(?:[\t\f\v ]*)([0-9]+)(?:[\t\f\v ])*d?(?:[\t\f\v ]*)([\d]*)(?:[\t\f\v ]*)([*\/]?)(?:[\t\f\v ]*)([+-]?[\t\f\v ]?[\d]*)(?:[\t\f\v ]*)([+-]?)(?:[\t\f\v ]*)([+-]?[\t\f\v ]?[\d]*)(?:[\t\f\v ]*)$`)

func ParseRolCt(rollFormula string) DieRollExpression {
	segs := DieParseRegex.FindStringSubmatch(rollFormula)

	rollCount, _ := strconv.Atoi(segs[1])
	diceSides, _ := strconv.Atoi(segs[2])
	multiplicand, _ := strconv.Atoi(segs[4])
	if segs[3] == "" {
		multiplicand = 1
	}
	addend, _ := strconv.Atoi(segs[6])
	if segs[5] == "" {
		addend = 0
	}

	shouldMultiply := segs[3] == "*"
	shouldAdd := segs[5] == "+"

	return DieRollExpression{
		Rolls:          rollCount,
		DieSides:       diceSides,
		Multiplicand:   multiplicand,
		Addend:         addend,
		ShouldMultiply: shouldMultiply,
		ShouldAdd:      shouldAdd,
	}

}
