// An exclusive list takes two or more lists, and combines them together in a way
// such that the options from both lists never occur together. An example
// of this may be to roll on a table for the weather that day, as well as
// the activities of the town. The tables input would be the towns activities
// for a certain weather, and the weight would be how often the weather occurs

package exclusive

import (
	"FoundryGoRollTables/pkg/table"
	"strconv"
)

type Table struct {
	Name   string
	Tables []WeightedTables
}

type WeightedTables struct {
	Weight int
	Table  *table.RollTable
}

// Wt1Item is a shorthand way to make a table item with
// only one weight.
func Wt1Item(table *table.RollTable) WeightedTables {
	return WeightedTables{
		Weight: 1,
		Table:  table,
	}
}

func (e Table) ToRollTable() table.RollTable {
	tableItems := make([]table.RollTableResult, len(e.Tables))
	cumulativeCount := 1
	for i, v := range e.Tables {
		tableItems[i] = table.RollTableResult{
			Text:       v.Table.Name,
			Collection: table.RollTableCollection,
			BrtFormula: "1",
			RangeL:     cumulativeCount,
			RangeH:     cumulativeCount + v.Weight,
		}
		cumulativeCount += v.Weight + 1
	}

	dieSides := strconv.Itoa(cumulativeCount - 1)

	return table.RollTable{
		Name:       e.Name,
		Formula:    "1d" + dieSides,
		TableType:  table.LootTableType,
		ResultList: tableItems,
	}
}
