package enums

type CarStatus string

const (
    CarAvailable   CarStatus = "available"
    CarBusy        CarStatus = "busy"
    CarOffline     CarStatus = "offline"
    CarMaintenance CarStatus = "maintenance"
)