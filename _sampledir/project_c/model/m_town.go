package model

type MTown struct {
}

func (b *MTown) TableName() string {
	return "m_town"
}
