package main

import (
	"FoundryGoRollTables/pkg/item"
	"FoundryGoRollTables/pkg/table/abstract/fair"
	"FoundryGoRollTables/pkg/table/combinatorial/exclusive"
	"encoding/json"
	"fmt"
)

func main() {
	myWeapon := item.WeaponItem{}
	err := json.Unmarshal([]byte(myJSON), &myWeapon)

	fmt.Println(err)

	SunnyRollTable := SunnyWeatherTable.ToRollTable()
	RainyRollTable := RainyWeatherTable.ToRollTable()
	StormyRollTable := StormyWeatherTable.ToRollTable()
	WeatherTable := exclusive.Table{
		Name: "Weather and activities",
		Tables: []exclusive.WeightedTables{
			{
				Weight: 80,
				Table:  &SunnyRollTable,
			},

			{
				Weight: 12,
				Table:  &RainyRollTable,
			},

			{
				Weight: 8,
				Table:  &StormyRollTable,
			},
		},
	}
	completeTable := WeatherTable.ToRollTable().ToJSONTable()
	output, _ := json.Marshal(completeTable)
	SunnyWeatherTableOut, _ := json.Marshal(SunnyRollTable.ToJSONTable())
	RainyRollTableOut, _ := json.Marshal(RainyRollTable.ToJSONTable())
	StormyRollTableOut, _ := json.Marshal(StormyRollTable.ToJSONTable())

	fmt.Println()
	fmt.Println(string(output))
	fmt.Println()
	fmt.Println(string(SunnyWeatherTableOut))
	fmt.Println()
	fmt.Println(string(RainyRollTableOut))
	fmt.Println()
	fmt.Println(string(StormyRollTableOut))
	fmt.Println()

}

// every wed is pennance day!

var SunnyWeatherTable = fair.FairTable{
	Name:        "Sunny Weather activities",
	Items:       SunnyWeatherElements,
	Replacement: false,
}
var SunnyWeatherElements = []fair.FairLI{
	fair.AsFairLIText("A pair of children are flying a kite, when one of them is swept off their feet and into the air.."),
	fair.AsFairLIText("An old man who is gardening in his front yard accidentally spills a potion on his lawn..."),
	fair.AsFairLIText("A paper boy is suddenly attacked by a group of vicious pigeons"),
	fair.AsFairLIText("A sailor's parrot has been kidnapped by a rival"),
	fair.AsFairLIText("An orphanage starts on fire"),
	fair.AsFairLIText("A triad of children have decided to take revenge on their bullies"),
	fair.AsFairLIText("A kid is stuck in the well"),
	fair.AsFairLIText("A scholar's homework is due in 10 minutes"),
}

var RainyWeatherTable = fair.FairTable{
	Name:        "Rainy Weather activities",
	Items:       RainyWeatherElements,
	Replacement: false,
}
var RainyWeatherElements = []fair.FairLI{
	fair.AsFairLIText("A home seems to have an obscene amount of water pouring into it..."),
}

var StormyWeatherTable = fair.FairTable{
	Name:        "Stormy Weather activities",
	Items:       StormyWeatherElements,
	Replacement: false,
}
var StormyWeatherElements = []fair.FairLI{
	fair.AsFairLIText("The local pet shop was struck by lightning"),
}

var AnyWeatherElements = []fair.FairLI{
	fair.AsFairLIText(""),
	fair.AsFairLIText("The sea gate has been clogged by the cargo barge, the ElvenGiven."),
	fair.AsFairLIText("The Brimstone triumvirate is holding their annual forged in hell competition, and are in need of soul coal"),
	fair.AsFairLIText("The Brimstone triumvirate is running desperately low on mountainous deepslate"),
	fair.AsFairLIText("The old coin fort has announced they are going to mint new bonds, and lower backings"),
	fair.AsFairLIText("The old coin fort has announced their monthly bounties"),
	fair.AsFairLIText("The old coin fort has an ominous yellow glow, espicially at night."),
	fair.AsFairLIText("Mercy o' the sea needs to send out an emergency stork, but they dont have the shrimp that the stork likes to eat"),
	fair.AsFairLIText("Mercy o' the sea is dangerously low on yuan-ti antivenom"),
	fair.AsFairLIText("The Alaion observatory has observed a new star in the sky"),
	fair.AsFairLIText("Kaya's orientery has gone haywire- the centering no longer point to the loadstone."),
	fair.AsFairLIText("Umberhulks from the nearby mountains have begun their breeding season."),
}

var myJSON = `{
  "name": "Rapier",
  "type": "weapon",
  "img": "icons/weapons/swords/sword-guard-brown.webp",
  "data": {
    "description": {
      "value": "<p>A thin tensile metal blade, light in weight but long in reach designed for quick darting attacks to target weak points in enemy defenses with lightning swiftness.</p>",
      "chat": "",
      "unidentified": ""
    },
    "source": "PHB pg. 149",
    "quantity": 1,
    "weight": 2,
    "price": "25",
    "attunement": 0,
    "equipped": false,
    "rarity": "common",
    "identified": true,
    "activation": {
      "type": "action",
      "cost": 1,
      "condition": ""
    },
    "duration": {
      "value": null,
      "units": ""
    },
    "target": {
      "value": null,
      "width": null,
      "units": "",
      "type": ""
    },
    "range": {
      "value": 5,
      "long": null,
      "units": "ft"
    },
    "uses": {
      "value": null,
      "max": "",
      "per": null
    },
    "consume": {
      "type": "",
      "target": null,
      "amount": null
    },
    "ability": null,
    "actionType": "mwak",
    "attackBonus": 0,
    "chatFlavor": "",
    "critical": {
      "threshold": null,
      "damage": ""
    },
    "damage": {
      "parts": [
        [
          "1d8 + @mod",
          "piercing"
        ]
      ],
      "versatile": ""
    },
    "formula": "",
    "save": {
      "ability": "",
      "dc": null,
      "scaling": "spell"
    },
    "armor": {
      "value": 10
    },
    "hp": {
      "value": 0,
      "max": 0,
      "dt": null,
      "conditions": ""
    },
    "weaponType": "martialM",
    "baseItem": "rapier",
    "properties": {
      "fin": true
    },
    "proficient": false
  },
  "effects": [],
  "flags": {
    "core": {
      "sourceId": "Compendium.dnd5e.items.Tobce1hexTnDk4sV"
    },
    "exportSource": {
      "world": "zendikar",
      "system": "dnd5e",
      "coreVersion": "9.269",
      "systemVersion": "1.6.0"
    }
  }
}`
