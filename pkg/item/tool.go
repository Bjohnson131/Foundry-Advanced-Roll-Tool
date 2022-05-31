package item

import "FoundryGoRollTables/pkg/item/types"

type ToolItem struct {
	BaseItem
	Data ToolData `json:"data"`
}

func (d *ToolItem) ValidateAndFill() {
	d.BaseItem.Type = "tool"
	d.BaseItem.ValidateAndFill()
	d.Data.ValidateAndFill()
}

type ToolData struct {
	BaseItemData
	toolType   types.ToolType        `json:"toolType"`
	BaseItem   types.BaseTool        `json:"baseItem"`
	Ability    *types.AbilityType    `json:"ability,omitempty"`
	ChatFlavor *string               `json:"chatFlavor"`
	Proficient types.ProfeciencyType `json:"proficient"`
	Bonus      string                `json:"bonus"`
}

func (d *ToolData) ValidateAndFill() {
	switch d.BaseItem {
	case types.BaseDice, types.BaseChess, types.BasePlayingCards:
		d.toolType = types.ToolGamingSet
	case types.BaseAlchemist, types.BaseBrewer, types.BaseCalligrapher, types.BaseCartographer, types.BaseCobbler, types.BaseCook, types.BaseGlassblower,
		types.BaseJeweler, types.BaseLeatherworker, types.BaseMason, types.BasePainter, types.BasePotter, types.BaseSmith, types.BaseTinker, types.BaseWeaver:
		d.toolType = types.ToolArtisanTools
	case types.BaseBagpipes, types.BaseDrum, types.BaseDulcimer, types.BaseFlute, types.BaseHorn, types.BaseLute, types.BaseLyre, types.BasePanflute, types.BaseShawm,
		types.BaseViol:
		d.toolType = types.ToolInstrument
	default:
		d.toolType = ""
	}
	d.BaseItemData.ValidateAndFill()
}
