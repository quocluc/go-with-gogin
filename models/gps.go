package models

type GPS struct {
	Id             int     `json:"id"`
	AccountId      int     `json:"account_id"`
	Timestamp      string  `json:"timestamp"`
	StatusCode     int     `json:"status_code"`
	Latitude       float64 `json:"lat"`
	Longitude      float64 `json:"long"`
	Altitude       float64 `json:"altitude"`
	OrderId        int     `json:"order_id"`
	OrderVehicleId int     `json:"order_vehicle_id"`
	DeviceId       int     `json:"device_id"`
	TransportId    int     `json:"transport_id"`
	DriverId       int     `json:"driver_id"`
	Speed          float64 `json:"speed"`
	FuelLevel      float64 `json:"fuel_level"`
	FuelTotal      float64 `json:"fuel_total"`
	Course         float64 `json:"course"`
	Runtime        float64 `json:"runtime"`
	DayRuntime     float64 `json:"day_runtime"`
	Duration       float64 `json:"duration"`
	SpeedLimit     float64 `json:"speed_limit"`
	Type           int     `json:"type"`
	Platform       int     `json:"platform"`
	CreateAt       int     `json:"create_at"`
	UpdatedAt      int     `json:"updated_at"`
	Time           int     `json:"time"`
}
