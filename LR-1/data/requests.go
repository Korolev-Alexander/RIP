package data

type DeviceRequest struct {
	DeviceID       int     // ID устройства
	UsedCount      int     // Количество устройств (пользовательское)
	CustomHoursDay float64 // Пользовательское число часов работы в сутки
	// уточнить что мы вообще расчитываем электро трафик или информационный?
}

type Request struct {
	ID             int
	DeviceRequests []DeviceRequest // устройства в заявке
}
