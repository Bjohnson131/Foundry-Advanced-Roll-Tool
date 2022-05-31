package types

type AbilityType string

var (
	AbilityInt AbilityType = "int"
	AbilityWis AbilityType = "wis"
	AbilityCha AbilityType = "cha"
	AbilityStr AbilityType = "str"
	AbilityDex AbilityType = "dex"
	AbilityCon AbilityType = "con"
)

type SpellcastingAbilityType string

var (
	SAInt  SpellcastingAbilityType = "int"
	SAWis  SpellcastingAbilityType = "wis"
	SACha  SpellcastingAbilityType = "cha"
	SAStr  SpellcastingAbilityType = "str"
	SADex  SpellcastingAbilityType = "dex"
	SACon  SpellcastingAbilityType = "con"
	SAflat SpellcastingAbilityType = "flat"
	SAscm  SpellcastingAbilityType = "spellcasting"
)

type TargetType string

var (
	TargetAlly     TargetType = "ally"
	TargetCone     TargetType = "cone"
	TargetCreature TargetType = "creature"
	TargetCube     TargetType = "cube"
	TargetCylinder TargetType = "cylinder"
	TargetEnemy    TargetType = "enemy"
	TargetLine     TargetType = "line"
	TargetNone     TargetType = "none"
	TargetObject   TargetType = "object"
	TargetRadius   TargetType = "radius"
	TargetSelf     TargetType = "self"
	TargetSpace    TargetType = "space"
	TargetSphere   TargetType = "sphere"
	TargetSquare   TargetType = "square"
	TargetWall     TargetType = "wall"
)

type ProfeciencyType float32

var (
	ProficientNone      ProfeciencyType = 0.0
	ProficientHalf      ProfeciencyType = 0.5
	Proficient          ProfeciencyType = 1.0
	ProficientExpertise ProfeciencyType = 2.0
)

type ToolType string

var (
	ToolArtisanTools ToolType = "art"
	ToolGamingSet    ToolType = "game"
	ToolInstrument   ToolType = "music"
)

type BaseTool string

var (
	BaseBagpipes BaseTool = "bagpipes"
	BaseDrum     BaseTool = "drum"
	BaseDulcimer BaseTool = "dulcimer"
	BaseFlute    BaseTool = "flute"
	BaseHorn     BaseTool = "horn"
	BaseLute     BaseTool = "lute"
	BaseLyre     BaseTool = "lyre"
	BasePanflute BaseTool = "panflute"
	BaseShawm    BaseTool = "shawm"
	BaseViol     BaseTool = "viol"

	BaseAlchemist     BaseTool = "alchemist"
	BaseBrewer        BaseTool = "brewer"
	BaseCalligrapher  BaseTool = "calligrapher"
	BaseCartographer  BaseTool = "cartographer"
	BaseCobbler       BaseTool = "cobbler"
	BaseCook          BaseTool = "cook"
	BaseGlassblower   BaseTool = "glassblower"
	BaseJeweler       BaseTool = "jeweler"
	BaseLeatherworker BaseTool = "leatherworker"
	BaseMason         BaseTool = "mason"
	BasePainter       BaseTool = "painter"
	BasePotter        BaseTool = "potter"
	BaseSmith         BaseTool = "smith"
	BaseTinker        BaseTool = "tinker"
	BaseWeaver        BaseTool = "weaver"
	BaseWoodcarver    BaseTool = "woodcarver"

	BaseDice         BaseTool = "dice"
	BaseChess        BaseTool = "chess"
	BasePlayingCards BaseTool = "cards"
)

type ItemType string

var (
	TypeWeapon     ItemType = "weapon"
	TypeEquipment  ItemType = "equipment"
	TypeConsumable ItemType = "consumable"
	TypeTool       ItemType = "tool"
	TypeLoot       ItemType = "loot"
	TypeBackground ItemType = "background"
)

type AttunementType int

var (
	NoAttunement       AttunementType = 0
	AttunementRequired AttunementType = 1
	Attuned            AttunementType = 2
)

type Rarity string

var (
	None      Rarity = ""
	Common    Rarity = "common"
	Uncommon  Rarity = "uncommon"
	Rare      Rarity = "rare"
	VeryRare  Rarity = "very rare"
	Legendary Rarity = "legendary"
	Artifact  Rarity = "artifact"
)

type WeaponType string

var (
	MartialMelee  WeaponType = "martialM"
	SimpleMelee   WeaponType = "simpleM"
	MartialRanged WeaponType = "martialR"
	SimpleRanged  WeaponType = "simpleR"
	NaturalWeapon WeaponType = "natural"
	Siege         WeaponType = "siege"
)

type ActivationType string

var (
	ActivationAction          ActivationType = "action"
	ActivationBonusAction     ActivationType = "bonus"
	ActivationCrewAction      ActivationType = "crew"
	ActivationDays            ActivationType = "day"
	ActivationHours           ActivationType = "hour"
	ActivationMinutes         ActivationType = "minute"
	ActivationLairAction      ActivationType = "lair"
	ActivationLegendaryAction ActivationType = "legendary"
	ActivationReaction        ActivationType = "reaction"
	ActivationSpecial         ActivationType = "special"
	ActivationNoneAction      ActivationType = "none"
	ActivationNilTime         ActivationType = ""
)

type DurationType string

var (
	DurationSpecial       DurationType = "spec"
	DurationPermanent     DurationType = "perm"
	DurationYears         DurationType = "year"
	DurationMonths        DurationType = "month"
	DurationDays          DurationType = "day"
	DurationHours         DurationType = "hour"
	DurationMinutes       DurationType = "minute"
	DurationInstantaneous DurationType = "inst"
	DurationTurns         DurationType = "turn"
	DurationRounds        DurationType = "round"
	DurationNilTime       DurationType = ""
)

type DistanceUnits string

var (
	DistanceUnitNil   DistanceUnits = ""
	DistanceUnitNone  DistanceUnits = "none"
	DistanceUnitSelf  DistanceUnits = "self"
	DistanceUnitTouch DistanceUnits = "touch"
	DistanceUnitSpec  DistanceUnits = "spec"
	DistanceUnitAny   DistanceUnits = "any"
	DistanceUnitFt    DistanceUnits = "ft"
	DistanceUnitMi    DistanceUnits = "mi"
	DistanceUnitM     DistanceUnits = "m"
	DistanceUnitKm    DistanceUnits = "km"
)

type LimitedUsePeriods string

var (
	UsePeriodShortRest LimitedUsePeriods = "sr"
	UsePeriodLongRest  LimitedUsePeriods = "lr"
	UsePeriodDay       LimitedUsePeriods = "day"
	UsePeriodCharges   LimitedUsePeriods = "charges"
)

type AbilityConsumptionTypes string

var (
	AbConsAmmo     AbilityConsumptionTypes = "ammo"
	AbConsAtt      AbilityConsumptionTypes = "attribute"
	AbConsHitDice  AbilityConsumptionTypes = "hitDice"
	AbConsCharges  AbilityConsumptionTypes = "charges"
	AbConsMaterial AbilityConsumptionTypes = "material"
)
