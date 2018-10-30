package models

import "fmt"

func GetGPS(OrderVehicleId int, DriverId int) interface{} {
	var listGPS []GPS
	result, err := db.Query("SELECT latitude, longitude, created_at FROM gps_event_logs WHERE order_vehicle_id = ? AND driver_id = ?", OrderVehicleId, DriverId)
	if err != nil {
		panic(err.Error())
		fmt.Println("Co loi xay ra trong sql!")
		return nil
	}
	for result.Next() {
		var latitude, longitude float64
		var created_at int
		err = result.Scan(&latitude, &longitude, &created_at)
		if err != nil {
			panic(err.Error())
		}
		gps := GPS{Latitude: latitude, Longitude: longitude, CreateAt: created_at}
		listGPS = append(listGPS, gps)
	}
	if listGPS == nil {
		return APIResponseFail("Không có thông tin! Vui lòng kiểm tra lại!")
	}
	return APIResponseOk(listGPS)
}
func UpdateGPS(OrderVehicleId int, DriverId int, listGPS []GPS) interface{} {

	if len(listGPS) == 0 {
		return nil
	}
	transaction, err := db.Begin()
	if err != nil {
		return nil
	}
	for _, gps := range listGPS {
		_, err := transaction.Query("INSERT INTO gps_event_logs (order_vehicle_id,driver_id, latitude, longitude, created_at) VALUE (?,?,?,?,?)", OrderVehicleId, DriverId, gps.Latitude, gps.Longitude, gps.Time)
		if err != nil {
			transaction.Rollback()
			return nil
		}
	}
	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		return nil
	}

	return APIResponseOk(listGPS)
}
