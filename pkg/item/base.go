package item

import "FoundryGoRollTables/pkg/item/types"

type BaseItem struct {
	Name    string         `json:"name"`
	Type    types.ItemType `json:"type"`
	Img     string         `json:"img"`
	Effects []ItemEffect   `json:"effects"`
	Flags   struct{}       `json:"flags"`
}

func (b *BaseItem) ValidateAndFill() {
	if b.Name == "" {
		b.Name = "Default Item"
	}
	if b.Img == "" {
		b.Img = "icons/svg/item-bag.svg"
	}
	if b.Effects == nil {
		b.Effects = make([]ItemEffect, 0)
	}
	if string(b.Type) == "" {
		b.Type = types.TypeLoot
	}
	b.Flags = struct{}{}
}

type ItemDescription struct {
	Value        string `json:"value"`
	Chat         string `json:"chat"`
	Unidentified string `json:"unidentified"`
}

type BaseItemData struct {
	Description ItemDescription      `json:"description"`
	Source      string               `json:"source"`
	Quantity    int                  `json:"quantity"`
	Weight      float64              `json:"weight"`
	Price       string               `json:"price"`
	Attunement  types.AttunementType `json:"attunement"`
	Equipped    bool                 `json:"equipped"`
	Rarity      types.Rarity         `json:"rarity"`
	Identified  bool                 `json:"identified"`
}

func (d *BaseItemData) ValidateAndFill() {
	if d.Quantity == 0 {
		d.Quantity = 1
	}
	if d.Weight == 0.0 {
		d.Weight = 1.0
	}
}
