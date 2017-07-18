package model

type MMovie struct {
}

func (b *MMovie) TableName() string {
	return "m_movie"
}
