package models

type UpdateJSON struct {
	OrderVehicleID int   `json:"order_vehicle_id"`
	DriverID       int   `json:"driver_id"`
	BodyGps        []GPS `json:"body_gps"`
}
