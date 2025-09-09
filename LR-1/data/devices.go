package data

type Device struct {
	ID          int
	Name        string
	Model       string
	PowerMin    float64 // минимальная мощность, Вт
	PowerMax    float64 // максимальная мощность, Вт
	AvgHours    float64 // по умолчанию: часов в сутки работы
	Image       string  // имя картинки
	Description string
	Protocol    string
}

var Devices = []Device{
	{
		ID:          1,
		Name:        "Хаб",
		Model:       "Яндекс Хаб",
		PowerMin:    10,
		PowerMax:    12,
		AvgHours:    24,
		Image:       "hub.png",
		Description: "Умный пульт Яндекс Хаб для устройств",
		Protocol:    "Wi-Fi",
	},
	{
		ID:          2,
		Name:        "Лампочка",
		Model:       "Яндекс, E27",
		PowerMin:    8,
		PowerMax:    10,
		AvgHours:    5,
		Image:       "lamp.png",
		Description: "Умная лампочка Яндекс, E27",
		Protocol:    "Wi-Fi",
	},
	{
		ID:          3,
		Name:        "Розетка",
		Model:       "YNDX-00540",
		PowerMin:    1,
		PowerMax:    2,
		AvgHours:    24,
		Image:       "socket.png",
		Description: "Умная розетка Яндекс",
		Protocol:    "Wi-Fi",
	},
	// ...потом дополню
}
