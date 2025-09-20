package data

type DeviceRequest struct {
	DeviceID int
	Quantity int
}

type Request struct {
	ID             int
	Address        string
	DeviceRequests []DeviceRequest
}

// Предзаполненная заявка для демонстрации
var CurrentRequest = &Request{
	ID:      1,
	Address: "ул. Примерная, д. 1, кв. 5",
	DeviceRequests: []DeviceRequest{
		{DeviceID: 2, Quantity: 3}, // 3 лампочки
		{DeviceID: 4, Quantity: 2}, // 2 датчика
	},
}
