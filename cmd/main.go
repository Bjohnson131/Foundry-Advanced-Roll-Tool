package main

import (
	"FoundryGoRollTables/pkg/parse/dice"
	"FoundryGoRollTables/pkg/table/abstract/inverse"
	"FoundryGoRollTables/pkg/table/abstract/normal"
	"encoding/json"
	"fmt"
)

func main2() {
	/*myInverseTable := inverse.InverseTable{
		Name:  "Potions By Rarity",
		Items: potions,
	}
	//completeTable := myInverseTable.ToRollTable().ToJSONTable()
	//output, err := json.Marshal(completeTable)*/

	myPolynomialTable := normal.PolynomialTable{
		Name:  "Polynomial table",
		Items: potionsPolynomial,
	}
	completeTable2 := myPolynomialTable.ToRollTable().ToJSONTable()
	output2, err := json.Marshal(completeTable2)

	fmt.Print(err)
	fmt.Print(string(output2))
	//fmt.Print(string(output))

	for _, v := range parseTest {
		fmt.Printf("%s : %s\n", v, dice.ParseRolCt(v).ToString())
	}

}

var parseTest = []string{
	"1 d 4 * 3 - 1",
	"20d4/1 - -50",
	"1 d 4 * 3 + 1",
	"1d4*3-1",
	"1d4/3-1",
	"10 d 4 * 3 + 1",
	"10 d 4 / 3 + 1",
	"10d4*3-1",
	"10d400*3000-10000",
	"2",
	"10",
	"1d4+1"}

var potionsPolynomial = []normal.PolynomialLI{
	{
		Name:      "potion1",
		AverageCt: 101,
		StdDev:    4,
	},
	{
		Name:      "potion2",
		AverageCt: 100,
		StdDev:    100,
	},
	{
		Name:      "potion3",
		AverageCt: 100,
		StdDev:    20,
	},
	{
		Name:      "potion4",
		AverageCt: 200,
		StdDev:    100,
	},
}

var potions = []inverse.InverseLI{
	{
		Name:   "Potion of Supreme Healing",
		Weight: 1350,
	},
	{
		Name:   "Potion of Invisibility",
		Weight: 180,
	},
	{
		Name:   "Potion of Speed",
		Weight: 400,
	},
	{
		Name:   "Potion of Flying",
		Weight: 500,
	},
	{
		Name:   "Potion of Vitality",
		Weight: 960,
	},
	{
		Name:   "Potion of Invulnerability",
		Weight: 3840,
	},
	{
		Name:   "Oil of Sharpness",
		Weight: 3200,
	},
	{
		Name:   "Potion of Longevity",
		Weight: 9000,
	},
	{
		Name:   "Potion of Cloud Giant Strength",
		Weight: 1990,
	},
	{
		Name:   "Potion of Mental Fortitude",
		Weight: 950,
	},
	{
		Name:   "Potion of Greater Healing",
		Weight: 150,
	},
	{
		Name:   "Potion of Poison",
		Weight: 100,
	},
	{
		Name:   "Philter of Love",
		Weight: 90,
	},
	{
		Name:   "Potion of Fire Breath",
		Weight: 150,
	},
	{
		Name:   "Potion of Water Breathing",
		Weight: 180,
	},
	{
		Name:   "Potion of Animal Friendship",
		Weight: 200,
	},
	{
		Name:   "Potion of Growth",
		Weight: 270,
	},
	{
		Name:   "Potion of Resistance",
		Weight: 300,
	},
	{
		Name:   "Oil of Slipperiness",
		Weight: 480,
	},
	{
		Name:   "Potion of Hill Giant Strength",
		Weight: 400,
	},
	{
		Name:   "Potion of Berserker Rage",
		Weight: 1000,
	},
	{
		Name:   "Potion of Spell Restoration",
		Weight: 500,
	},
	{
		Name:   "Potion of Vigor",
		Weight: 450,
	},
	{
		Name:   "Potion of Wild Magic",
		Weight: 100,
	},
	{
		Name:   "Oil of Chain Reaction",
		Weight: 300,
	},
	{
		Name:   "Potion of Darkvision",
		Weight: 200,
	},
	{
		Name:   "Potion of sleep",
		Weight: 200,
	},
	{
		Name:   "Potion of Hallucinations",
		Weight: 400,
	},
	{
		Name:   "Potion of Rest",
		Weight: 200,
	},
	{
		Name:   "Potion of Superior Healing",
		Weight: 450,
	},
	{
		Name:   "Elixir of Health",
		Weight: 120,
	},
	{
		Name:   "Potion of Heroism",
		Weight: 180,
	},
	{
		Name:   "Potion of Mind Reading",
		Weight: 180,
	},
	{
		Name:   "Potion of Diminution",
		Weight: 270,
	},
	{
		Name:   "Potion of Gaseous Form",
		Weight: 300,
	},
	{
		Name:   "Potion of Clairvoyance",
		Weight: 960,
	},
	{
		Name:   "Oil of Etherealness",
		Weight: 1920,
	},
	{
		Name:   "Potion of Frost Giant Strength",
		Weight: 750,
	},
	{
		Name:   "Potion of Fire Giant Strength",
		Weight: 1300,
	},
	{
		Name:   "Potion of Major Spell Restoration",
		Weight: 900,
	},
	{
		Name:   "Potion of Polymorph",
		Weight: 700,
	},
	{
		Name:   "Potion of Second Chance",
		Weight: 2000,
	},
	{
		Name:   "Potion of Storm Giant Strength",
		Weight: 3000,
	},
	{
		Name:   "Potion of Deep Slumber",
		Weight: 5000,
	},
	{
		Name:   "Potion of Luck",
		Weight: 20000,
	},
	{
		Name:   "Potion of Healing",
		Weight: 50,
	},
	{
		Name:   "Potion of Climbing",
		Weight: 180,
	},
	{
		Name:   "Potion of Minor Spell Restoration",
		Weight: 200,
	},
	{
		Name:   "Unstable Healing Potion",
		Weight: 50,
	},
	{
		Name:   "Oil of Firefly",
		Weight: 100,
	},
	{
		Name:   "Potion of Message",
		Weight: 25,
	},
}
