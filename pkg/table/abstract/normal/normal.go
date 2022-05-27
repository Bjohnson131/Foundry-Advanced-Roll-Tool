// A normal table is a table with multiple binomial distributions that
// can be picked independently of one another. As an example, say that we're
// creating a loot table for a thief. Thieves can have crowbars, lockpicks, and
// ball bearings. We can set an average and standard deviation for all of these
// to generate all of these items separately.

package normal

import (
	"FoundryGoRollTables/pkg/parse/dice"
	"FoundryGoRollTables/pkg/table"
	"math"
)

// DefaultRollCt determines how normal-like the distribution curve is
// The higher the number, the more normal the curve is, however,
// it will take longer to roll. Despite converting to float64, it should
// be an integer.
var DefaultRollCt = float64(4)

// DefaultDivisor is the number that the roll is divided by. Higher numbers
// mean less quality in the results. Do not set to below 1.0.
var DefaultDivisor = float64(1)

type PolynomialTable struct {
	Name  string
	Items []PolynomialLI
}

type PolynomialLI struct {
	Name      string
	AverageCt float64 // assumed 1 if not specified.
	StdDev    float64 // assumed 1 if not specified. 0 indicates a constant result.
}

func (i PolynomialTable) ToRollTable() table.RollTable {
	results := make([]table.RollTableResult, len(i.Items))

	for i, v := range i.Items {
		results[i] = v.ToRollTableWithDivisor(DefaultDivisor)
	}

	return table.RollTable{
		Name:       i.Name,
		ResultList: results,
		Formula:    "1",
		TableType:  table.LootTableType,
	}

}

func (i PolynomialLI) ToRollTableWithDivisor(divisor float64) table.RollTableResult {
	stdDev := i.StdDev
	if stdDev == 0.0 {
		stdDev = 1.0
	}
	avg := i.AverageCt
	if avg == 0.0 {
		avg = 1
	}
	rollCt := 0.0
	DiceSides := 0.0
	sideErrorMin := 100.0
	for testRollCt := DefaultRollCt; testRollCt < DefaultRollCt+100; testRollCt++ {
		variance := stdDev * stdDev
		numerator := 12*variance*divisor + testRollCt
		TestDiceSides := math.Sqrt(numerator) / math.Sqrt(testRollCt)

		// In this case, we can determine the error directly by dividing the
		// int by the float to get a fraction representing the decimal portion
		denominator := float64(int64(TestDiceSides))
		sidesError := float64(TestDiceSides / denominator)

		if sidesError < sideErrorMin {
			sideErrorMin = sidesError
			DiceSides = TestDiceSides
			rollCt = testRollCt
			// if we find an exact solution, break
			if sidesError == 1.0 {
				break
			}
		}

	}
	if int(DiceSides) < 10 {
		return i.ToRollTableWithDivisor(divisor+1)
	}

	subtractor := (-2.0*avg*divisor + rollCt*DiceSides + rollCt) / (2.0 * divisor)

	dieRoll := dice.DieRollExpression{
		Rolls:        int(rollCt),
		DieSides:     int(DiceSides),
		Multiplicand: int(divisor),
		Addend:       int(subtractor),
	}

	return table.RollTableResult{
		Text:       i.Name,
		BrtFormula: dieRoll.ToString(),
		RangeL:     1,
		RangeH:     1,
		Collection: table.Item,
	}
}
