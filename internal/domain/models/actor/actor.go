package actor

import "school21/task/RogueGame/internal/domain/models/item"

type Stats struct {
	Strength int
	Agility  int
	Armor    int
	Helth    int
	MaxHelth int
}

type Equipment struct {
	Weapon item.WeaponItem
	Armor  item.ArmorItem
}

type Actor struct {
	Stats
	Equipment
}
