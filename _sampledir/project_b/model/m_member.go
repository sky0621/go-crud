package model

type MMember struct {
}

func (b *MMember) TableName() string {
	return "m_member"
}
