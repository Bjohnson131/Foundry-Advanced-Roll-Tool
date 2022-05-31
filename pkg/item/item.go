package item

// EXPERIMENTAL. this will change in the future.

import (
	"crypto/sha1"
	"encoding/base64"
)

type Validatable interface {
	ValidateAndFill()
}

type ItemSpeed struct {
	Value      *int   `json:"value"`
	Conditions string `json:"conditions"`
}

type Data struct {
	BaseItemData
	Speed    ItemSpeed `json:"speed,omitempty"`
	Strength int       `json:"strength,omitempty"`
	Stealth  bool      `json:"stealth,omitempty"`
}
type EffectDuration struct {
	StartTime  *int `json:"startTime"`
	Rounds     int  `json:"rounds,omitempty"`
	Seconds    int  `json:"seconds,omitempty"`
	Turns      int  `json:"turns,omitempty"`
	StartRound int  `json:"startRound,omitempty"`
	StartTurn  int  `json:"startTurn,omitempty"`
}

func (e EffectDuration) ValidateAndFill() {
}

type ItemEffect struct {
	Id       string         `json:"_id"`
	Changes  []interface{}  `json:"changes"`
	Disabled bool           `json:"disabled"`
	Duration EffectDuration `json:"duration"`
	Icon     string         `json:"icon"`
	Label    string         `json:"label"`
	Transfer bool           `json:"transfer"`
	Flags    struct{}       `json:"flags"`
	Tint     string         `json:"tint"`
}

func (i *ItemEffect) ValidateAndFill() {
	i.Duration.ValidateAndFill()
	if i.Icon == "" {
		i.Icon = "icons/svg/aura.svg"
	}
	if i.Label == "" {
		i.Label = "Default_Name"
	}
	if i.Id == "" {
		id := i.GetID()
		i.Id = string(id[:])
	}
}

func (i *ItemEffect) GetID() [16]byte {
	hasher := sha1.New()
	hasher.Write([]byte(i.Label))
	hasher.Write([]byte(i.Icon))
	hasher.Write([]byte(i.Tint))
	sha := []byte(base64.URLEncoding.EncodeToString(hasher.Sum(nil)))
	id := *(*[16]byte)(sha[0:16])
	return id
}
