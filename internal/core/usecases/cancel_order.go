package usecases

type (
	CancelOrderParams struct {
		OrderID string
		UserID  string
	}

	CancelOrderUseCase interface {
		Cancel(params CancelOrderParams) error
	}
)
