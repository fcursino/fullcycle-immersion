package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction { //criação de nova transação
	total := float64(shares) * price //:= é declaração e atribuição no go
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Price:        price,
		Total:        total,
	}
}

func (t *Transaction) CloseBuyOrder() { //marca ordem de compra como fechada
	if t.BuyingOrder.PendingShares == 0 {
		t.BuyingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) CloseSellOrder() { //marca ordem de venda como fechada
	if t.SellingOrder.PendingShares == 0 {
		t.SellingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) AddBuyOrderPendingShares(shares int) { //adiciona ordem de compra nas ordens pendentes
	t.BuyingOrder.PendingShares += shares
}

func (t *Transaction) AddSellOrderPendingShares(shares int) { //adiciona ordem de venda nas ordens pendentes
	t.SellingOrder.PendingShares += shares
}
