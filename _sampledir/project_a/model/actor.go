package model

type Actor struct {
}

func (b *Actor) TableName() string {
	return "actor"
}
