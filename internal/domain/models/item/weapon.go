package item

type WeaponType int

// enum
const (
	Blade WeaponType = iota
	Axe
	Dagger
)

type WeaponItem struct {
	Item
	Type   WeaponType
	Damage int
}
