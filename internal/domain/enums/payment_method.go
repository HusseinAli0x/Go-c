package enums

type PaymentMethod string

const (
    PaymentCard     PaymentMethod = "card"
    PaymentCash     PaymentMethod = "cash"
    PaymentWallet   PaymentMethod = "wallet"
    PaymentPayPal   PaymentMethod = "paypal"
)