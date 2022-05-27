package table

var DefaultIcon = "icons/svg/blood.svg"
var ItemDefaultIcon = ""

type CollectionType string

var Item CollectionType = "Item"
var RollTableCollection CollectionType = "RollTable"
var Compendium5e CollectionType = "dnd5e.items"
var Text CollectionType = "Text"

type TableType string

var LootTableType TableType = "loot"

type RollTable struct {
	Name       string
	ResultList []RollTableResult
	Formula    string
	TableType  TableType
	DisableReplacement bool
}

type RollTableResult struct {
	Text       string         `json:"text"`
	Collection CollectionType `json:"collection"`
	BrtFormula string         `json:"formula"`
	RangeL     int            `json:"rangeL"`
	RangeH     int            `json:"rangeH"`
}

func (r RollTable) ToJSONTable() Table {
	Tresults := make([]TResult, len(r.ResultList))
	for i, v := range r.ResultList {
		Tresults[i] = v.ToTResult()
	}
	return Table{
		Name:        r.Name,
		Img:         DefaultIcon,
		Formula:     r.Formula,
		Replacement: !r.DisableReplacement,
		DisplayRoll: true,
		Results:     Tresults,
		Flags: struct {
			BetterRolltables struct {
				Tags                string `json:"tags"`
				TableType           string `json:"table-type"`
				TableCurrencyString string `json:"table-currency-string"`
				LootAmountKey       string `json:"loot-amount-key"`
				LootActorName       string `json:"loot-actor-name"`
			} `json:"better-rolltables"`
		}(struct {
			BetterRolltables struct {
				Tags                string `json:"tags"`
				TableType           string `json:"table-type"`
				TableCurrencyString string `json:"table-currency-string"`
				LootAmountKey       string `json:"loot-amount-key"`
				LootActorName       string `json:"loot-actor-name"`
			}
		}{BetterRolltables: struct {
			Tags                string `json:"tags"`
			TableType           string `json:"table-type"`
			TableCurrencyString string `json:"table-currency-string"`
			LootAmountKey       string `json:"loot-amount-key"`
			LootActorName       string `json:"loot-actor-name"`
		}(struct {
			Tags                string
			TableType           string
			TableCurrencyString string
			LootAmountKey       string
			LootActorName       string
		}{Tags: "", TableType: string(r.TableType), TableCurrencyString: "", LootAmountKey: "1", LootActorName: ""})}),
	}
}

func (r RollTableResult) ToTResult() TResult {
	colType := 1

	switch r.Collection {
	case Compendium5e:
		colType = 2
	case Text:
		colType = 0
	}
	return TResult{
		Type:       colType,
		Text:       r.Text,
		Img:        ItemDefaultIcon,
		Collection: string(r.Collection),
		Weight:     1,
		Range:      []int{r.RangeL, r.RangeH},
		Drawn:      false,
		RangeH:     r.RangeH,
		RangeL:     r.RangeL,
		Flags: struct {
			BetterRolltables struct {
				BrtResultFormula struct {
					Formula string `json:"formula"`
				} `json:"brt-result-formula"`
			} `json:"better-rolltables"`
		}(struct {
			BetterRolltables struct {
				BrtResultFormula struct {
					Formula string `json:"formula"`
				} `json:"brt-result-formula"`
			}
		}{BetterRolltables: struct {
			BrtResultFormula struct {
				Formula string `json:"formula"`
			} `json:"brt-result-formula"`
		}(struct {
			BrtResultFormula struct {
				Formula string `json:"formula"`
			}
		}{BrtResultFormula: struct {
			Formula string `json:"formula"`
		}(struct{ Formula string }{Formula: r.BrtFormula})})}),
	}
}

type Table struct {
	Name        string    `json:"name"`
	Img         string    `json:"img"`
	Formula     string    `json:"formula"`
	Replacement bool      `json:"replacement"`
	DisplayRoll bool      `json:"displayRoll"`
	Results     []TResult `json:"results"`
	Flags       struct {
		BetterRolltables struct {
			Tags                string `json:"tags"`
			TableType           string `json:"table-type"`
			TableCurrencyString string `json:"table-currency-string"`
			LootAmountKey       string `json:"loot-amount-key"`
			LootActorName       string `json:"loot-actor-name"`
		} `json:"better-rolltables"`
	} `json:"flags"`
}

type TResult struct {
	Type       int    `json:"type"`
	Text       string `json:"text"`
	Img        string `json:"img"`
	Collection string `json:"collection"`
	Weight     int    `json:"weight"`
	Range      []int  `json:"range"`
	Drawn      bool   `json:"drawn"`
	RangeL     int    `json:"rangeL"`
	RangeH     int    `json:"rangeH"`
	Flags      struct {
		BetterRolltables struct {
			BrtResultFormula struct {
				Formula string `json:"formula"`
			} `json:"brt-result-formula"`
		} `json:"better-rolltables"`
	} `json:"flags,omitempty"`
}
