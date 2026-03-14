package enums

type UserRole string

const (
	RoleCustomer UserRole = "customer"
	RoleDriver   UserRole = "driver"
	RoleAdmin    UserRole = "admin"
)
