package entity

type Asset struct {
	ID           string
	Name         string
	MarketVolume int
}

func NewAsset(id string, name string, marketVolume int) *Asset { //cria uma carteira
	return &Asset{
		ID:           id,
		Name:         name,
		MarketVolume: marketVolume,
	}
}
