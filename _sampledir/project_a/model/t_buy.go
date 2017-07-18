package model

type TBuy struct {
}

func (b *TBuy) TableName() string {
	return "t_buy"
}
