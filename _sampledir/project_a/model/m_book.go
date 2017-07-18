package model

type MBook struct {
}

func (b *MBook) TableName() string {
	return "m_book"
}
