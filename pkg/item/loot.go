package item

type LootItem struct {
	BaseItem
	Data LootData `json:"data"`
}

func (d *LootItem) ValidateAndFill() {
	d.BaseItem.Type = "loot"
	d.BaseItem.ValidateAndFill()
	d.Data.ValidateAndFill()
}

type LootData struct {
	BaseItemData
}

func (d *LootData) ValidateAndFill() {
	d.BaseItemData.ValidateAndFill()
}
