// A direct table is a table where items have a weight heuristic.
// The higher the weight, the higher the chance to pick the item.
// As an example, the weight for this type of table could be the
// number of people that use this item on a daily basis.
// A direct table whose

package direct

import (
	"FoundryGoRollTables/pkg/table"
	"strconv"
)

type DirectTable struct {
	Name  string
	Items []DirectLI
}

type DirectLI struct {
	Name       string
	Weight     int
	BrtFormula string
	Collection table.CollectionType
}

func (i DirectTable) ToRollTable() table.RollTable {

	results := make([]table.RollTableResult, len(i.Items))
	totalRolls := 0
	for i, v := range i.Items {
		if string(v.Collection) == "" {
			v.Collection = table.Item
		}
		BrtFormula := "1"
		if v.BrtFormula != "" {
			BrtFormula = v.BrtFormula
		}
		results[i] = table.RollTableResult{
			Text:       v.Name,
			BrtFormula: BrtFormula,
			RangeL:     totalRolls + 1,
			RangeH:     totalRolls + v.Weight,
			Collection: v.Collection,
		}
		totalRolls += v.Weight
	}

	return table.RollTable{
		Name:       i.Name,
		ResultList: results,
		Formula:    "1d" + strconv.Itoa(totalRolls),
		TableType:  table.LootTableType,
	}
}
