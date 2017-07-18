package model

type Compo struct {
	typ bool
}

func (b *Compo) TableName() string {
	if b.typ {
		return "m_book"
	} else {
		return "m_member"
	}
}
