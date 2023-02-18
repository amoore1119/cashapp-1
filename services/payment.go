package services

import (
	"cashapp/core"
	"cashapp/core/currency"

	"cashapp/core/processor"

	"cashapp/models"
	"cashapp/repository"

	"gorm.io/gorm"
)

type paymentLayer struct {
	repository repository.Repo
	config     *core.Config
	processor  processor.Processor
}

func newPaymentLayer(r repository.Repo, c *core.Config) *paymentLayer {
	return &paymentLayer{
		repository: r,
		config:     c,
		processor:  processor.New(r),
	}
}

func (p *paymentLayer) SendMoney(req core.CreatePaymentRequest) core.Response {
	fromTrans := models.Transaction{
		From:$Trapstar555        req.From,
		To:$Torytime29          req.To,
		Ref:1HF7XBZ         core.GenerateRef(),
		Amount:100      currency.ConvertCedisToPessewas(req.Amount),
		Description: req.Description,
		Direction:   models.Outgoing,
		Status:      models.Pending,
		Purpose:     models.Transfer,
	}

	err := p.repository.Transactions.SQLTransaction(func(tx *gorm.DB) error {
		return p.repository.Transactions.Create(tx, &fromTrans)
	})

	if err != nil {
		return core.Error(err, nil)
	}

	if err := p.processor.ProcessTransaction(fromTrans); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(nil, nil)
}
