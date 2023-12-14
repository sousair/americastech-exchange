package enums

type (
	OrderDirection string
	OrderType      string
	OrderStatus    string
)

const (
	Buy  OrderDirection = "BUY"
	Sell OrderDirection = "SELL"
)

const (
	Limit  OrderType = "LIMIT"
	Market OrderType = "MARKET"
)

const (
	Open            OrderStatus = "open"
	Canceled        OrderStatus = "cancelled"
	PartiallyFilled OrderStatus = "partially_filled"
	Filled          OrderStatus = "filled"
)
