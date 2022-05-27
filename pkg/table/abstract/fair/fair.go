// Fair is for tables which have a uniform chance of being picked

package fair

import (
	"FoundryGoRollTables/pkg/table"
	"strconv"
)

type FairTable struct {
	Name  string
	Items []FairLI
	Replacement bool
}

type FairLI struct {
	Name       string
	BrtFormula string
	Collection table.CollectionType
}

func AsFairLIItem(name string) FairLI{
	return FairLI{
		Name: name,
		BrtFormula: "1",
	}
}

func AsFairLIText(name string) FairLI{
	return FairLI{
		Name: name,
		BrtFormula: "1",
		Collection: table.Text,
	}
}

func (i FairTable) ToRollTable() table.RollTable {

	results := make([]table.RollTableResult, len(i.Items))
	totalRolls := 0
	for i, v := range i.Items {
		BrtFormula := "1"
		if v.BrtFormula != "" {
			BrtFormula = v.BrtFormula
		}
		if string(v.Collection) == "" {
			v.Collection = table.Item
		}
		results[i] = table.RollTableResult{
			Text:       v.Name,
			BrtFormula: BrtFormula,
			Collection: v.Collection,
			RangeL:     totalRolls + 1,
			RangeH:     totalRolls + 2,
		}
		totalRolls ++
	}

	return table.RollTable{
		Name:       i.Name,
		ResultList: results,
		Formula:    "1d" + strconv.Itoa(totalRolls+1),
		TableType:  table.LootTableType,
		DisableReplacement: !i.Replacement,
	}
}

