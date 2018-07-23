package oop

import "fmt"

// type XXX interface {
// 	yyy()
// }

type Ironman struct {
	Name string
	Human
	Inhuman inhuman
}

func NewIronMan(name string, skill string, power uint, human Human) *Ironman {
	return &Ironman{
		Name:  name,
		Human: human,
		Inhuman: inhuman{
			power: power,
			skill: skill,
		},
	}
}

// func (i *Ironman) yyy() {
// }

func (i *Ironman) SetSkill(skill string) {
	i.Inhuman.skill = skill
}

func (i *Ironman) ShowSkill() {
	fmt.Printf("Skill: %s\n", i.Inhuman.skill)
}

func (i *Ironman) SetPower(power uint) {
	i.Inhuman.power = power
}

func (i *Ironman) ShowPower() {
	fmt.Printf("Power: %d\n", i.Inhuman.power)
}

type WarM struct {
	*Ironman
	NickName string
}
