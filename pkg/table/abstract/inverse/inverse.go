// An inverse table is a table where items have a weight heuristic.
// The higher the weight, the lower the chance to pick the item.
// As an example, the weight for this type of table could be the value in
// GP. The more GP the item is worth, the less of a chance it'll be picked.

package inverse

import (
	"FoundryGoRollTables/pkg/table"
	"strconv"
)

// PrecisionMultiplier determines how precise the roll table is. The higher
// the setting, the more precise, but the more calculation heavy.
var PrecisionMultiplier = 1.5

type InverseTable struct {
	Name  string
	Items []InverseLI
}

type InverseLI struct {
	Name       string
	Weight     int
	BrtFormula string
}

func (i InverseTable) ToRollTable() table.RollTable {
	maxWt := 0
	for _, v := range i.Items {
		if v.Weight > maxWt {
			maxWt = v.Weight
		}
	}
	results := make([]table.RollTableResult, len(i.Items))

	totalRolls := 0
	for i, v := range i.Items {
		rangeSize := int(float64(maxWt)*PrecisionMultiplier) / v.Weight
		BrtFormula := "1"
		if v.BrtFormula != "" {
			BrtFormula = v.BrtFormula
		}
		results[i] = table.RollTableResult{
			Text:       v.Name,
			BrtFormula: BrtFormula,
			RangeL:     totalRolls + 1,
			RangeH:     totalRolls + rangeSize,
			Collection: table.Item,
		}
		totalRolls += rangeSize
	}

	return table.RollTable{
		Name:       i.Name,
		ResultList: results,
		Formula:    "1d" + strconv.Itoa(totalRolls),
		TableType:  table.LootTableType,
	}
}
