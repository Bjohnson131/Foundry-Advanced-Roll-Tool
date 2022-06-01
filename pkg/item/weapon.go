package item

import (
	"FoundryGoRollTables/pkg/item/types"
	"FoundryGoRollTables/pkg/utils"
	"encoding/json"
)

type WeaponItem struct {
	BaseItem
	Data WeaponData `json:"data"`
}

func (d *WeaponItem) ValidateAndFill() {
	d.BaseItem.Type = "weapon"
	d.BaseItem.ValidateAndFill()
	d.Data.ValidateAndFill()
}

type ItemActivationCondition struct {
	ActivationCurrency types.ActivationType `json:"type"`
	Cost               int                  `json:"cost"`
	Condition          string               `json:"condition"`
}

func (d *ItemActivationCondition) ValidateAndFill() {
	switch d.ActivationCurrency {
	case types.ActivationNilTime:
		d.Cost = 0
		d.Condition = ""
	case types.ActivationNoneAction:
		d.Cost = 0
	case types.ActivationReaction, types.ActivationAction, types.ActivationBonusAction,
		types.ActivationCrewAction, types.ActivationLairAction, types.ActivationLegendaryAction:
		d.Cost = 1
	}
}

type ItemDuration struct {
	Value int                `json:"value"`
	Units types.DurationType `json:"units"`
}

func (d *ItemDuration) ValidateAndFill() {
	switch d.Units {
	case types.DurationPermanent:
		d.Value = 1
	case types.DurationNilTime, types.DurationInstantaneous:
		d.Value = 0
	}
}

type ItemTarget struct {
	Value *int                `json:"value"`
	Width *int                `json:"width"`
	Units types.DistanceUnits `json:"units"`
	Type  types.TargetType    `json:"type"`
}

func (d *ItemTarget) ValidateAndFill() {
	switch d.Units {
	case types.DistanceUnitSelf:
		d.Value = utils.Int(1)
		d.Width = nil
		d.Type = types.TargetSelf
	case types.DistanceUnitNil, types.DistanceUnitNone:
		d.Value = utils.Int(0)
		d.Width = nil
		d.Type = types.TargetNone
	case types.DistanceUnitTouch:
		d.Value = utils.Int(1)
	}
}

type ItemRange struct {
	ItemDuration
	Long int `json:"long"`
}

type ItemUses struct {
	CurrentUses *int                    `json:"value"`
	MaxUses     utils.DynamicInt        `json:"max"`
	Per         types.LimitedUsePeriods `json:"per"`
}

func (d *ItemUses) ValidateAndFill() {
	if d.MaxUses != "" && d.CurrentUses == nil {
		num, _ := d.MaxUses.Int64()
		d.CurrentUses = utils.Int(int(num))
	}
}

type ItemConsume struct {
	Type   types.AbilityConsumptionTypes `json:"type"`
	Target *string                       `json:"target"`
	Amount *int                          `json:"amount"`
}

func (d *ItemConsume) ValidateAndFill() {}

type ItemCritical struct {
	Threshold   *int   `json:"threshold"`
	ExtraDamage string `json:"damage"`
}

func (d *ItemCritical) ValidateAndFill() {}

type ItemDamage struct {
	Parts     [][2]string `json:"parts"`
	Versatile string      `json:"versatile"`
}

func (d *ItemDamage) ValidateAndFill() {}

type ItemSave struct {
	AbilityVs types.AbilityType             `json:"ability"`
	Dc        *int                          `json:"dc"`
	Scaling   types.SpellcastingAbilityType `json:"scaling"`
}

func (d *ItemSave) ValidateAndFill() {
	if d.Scaling == "" {
		d.Scaling = types.SAflat
	}
	if d.Dc != nil && *d.Dc < 0 {
		d.Dc = utils.Int(0)
	}
}

type ItemArmor struct {
	Value *int    `json:"value"`
	Type  *string `json:"type,omitempty"`
	Dex   *int    `json:"dex,omitempty"`
}

func (d *ItemArmor) ValidateAndFill() {
	if (d.Type != nil || d.Dex != nil) && d.Value == nil {
		d.Value = utils.Int(0)
	}
}

type ItemHP struct {
	Value      int    `json:"value"`
	Max        int    `json:"max"`
	Threshold  *int   `json:"dt"`
	Conditions string `json:"conditions"`
}

func (d *ItemHP) ValidateAndFill() {
	if d.Max == 0 && d.Value != 0 {
		d.Max = d.Value
	}
	if d.Max != 0 && d.Value == 0 {
		d.Value = d.Max
	}
	if d.Threshold != nil && *d.Threshold < 0 {
		d.Threshold = utils.Int(0)
	}
}

type WeaponProperties struct {
	Adamantine        bool `json:"ada"`
	Ammunition        bool `json:"amm"`
	Finesse           bool `json:"fin"`
	Firearm           bool `json:"fir"`
	SpellcastingFocus bool `json:"foc"`
	Heavy             bool `json:"hvy"`
	Light             bool `json:"lgt"`
	Loading           bool `json:"lod"`
	Magical           bool `json:"mgc"`
	Reach             bool `json:"rch"`
	FirearmReload     bool `json:"rel"`
	Returning         bool `json:"ret"`
	Silvered          bool `json:"sil"`
	Special           bool `json:"spc"`
	Thrown            bool `json:"thr"`
	TwoHanded         bool `json:"two"`
	Versatile         bool `json:"ver"`
}

func (d *WeaponProperties) ValidateAndFill() {}

type WeaponData struct {
	BaseItemData
	Activation  ItemActivationCondition `json:"activation"`
	Duration    ItemDuration            `json:"duration"`
	Target      ItemTarget              `json:"target"`
	Range       ItemRange               `json:"range"`
	Uses        ItemUses                `json:"uses"`
	Consume     ItemConsume             `json:"consume"`
	Ability     types.AbilityType       `json:"ability"`
	ActionType  string                  `json:"actionType"`
	AttackBonus json.Number             `json:"attackBonus"`
	ChatFlavor  string                  `json:"chatFlavor"`
	Critical    ItemCritical            `json:"critical"`
	Damage      ItemDamage              `json:"damage"`
	Formula     string                  `json:"formula"`
	Save        ItemSave                `json:"save"`
	Armor       *ItemArmor              `json:"armor,omitempty"`
	Hp          ItemHP                  `json:"hp"`
	WeaponType  types.WeaponType        `json:"weaponType"`
	BaseItem    string                  `json:"baseItem"`
	Properties  WeaponProperties        `json:"properties,omitempty"`
	Proficient  bool                    `json:"proficient,omitempty"`
	Cover       float32                 `json:"cover"`
}

func (d *WeaponData) ValidateAndFill() {
	d.BaseItemData.ValidateAndFill()
	d.Activation.ValidateAndFill()
	d.Target.ValidateAndFill()
	d.Range.ValidateAndFill()
	d.Uses.ValidateAndFill()
	d.Consume.ValidateAndFill()
	d.Critical.ValidateAndFill()
	d.Damage.ValidateAndFill()
	d.Save.ValidateAndFill()
	d.Properties.ValidateAndFill()

}
