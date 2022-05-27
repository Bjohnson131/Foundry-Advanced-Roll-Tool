// simultaneous is a way of combining two lists such that all lists always
// occur together. An example may be a few roll tables which together form
// a few character traits.

package simultaneous

import (
	"FoundryGoRollTables/pkg/table"
)

type Table struct {
	Name   string
	Tables []*table.RollTable
}

func (e Table) ToRollTable() table.RollTable {
	tableItems := make([]table.RollTableResult, len(e.Tables))
	for i, v := range e.Tables {
		tableItems[i] = table.RollTableResult{
			Text:       v.Name,
			Collection: table.RollTableCollection,
			BrtFormula: "1",
			RangeL:     1,
			RangeH:     1,
		}
	}

	return table.RollTable{
		Name:       e.Name,
		Formula:    "1",
		TableType:  table.LootTableType,
		ResultList: tableItems,
	}
}
