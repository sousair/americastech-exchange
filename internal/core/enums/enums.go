package enums

type (
	OrderDirection string
	OrderType      string
	OrderStatus    string
)

const (
	OrderDirectionBuy  OrderDirection = "BUY"
	OrderDirectionSell OrderDirection = "SELL"
)

const (
	OrderTypeLimit  OrderType = "LIMIT"
	OrderTypeMarket OrderType = "MARKET"
)

const (
	OrderStatusOpen            OrderStatus = "open"
	OrderStatusCanceled        OrderStatus = "cancelled"
	OrderStatusPartiallyFilled OrderStatus = "partially_filled"
	OrderStatusFilled          OrderStatus = "filled"
)
