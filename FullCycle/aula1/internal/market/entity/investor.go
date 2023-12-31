package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition //ponteiro para deixar o registro único na memória (parece C)
}

func NewInvestor(id string) *Investor { //registra novo investidor, sem nenhuma posição inicial na carteira
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func (i Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) { //adicionando posição de carteira para um investidor
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtdShares int) { //atualiza posição de carteira
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, qtdShares))
	} else {
		assetPosition.Shares += qtdShares
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition { //procura por posição de carteira
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil //nulo
}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition { //adiciona posição de carteira
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}
