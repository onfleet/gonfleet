package onfleet

type TurnByTurn struct {
	StartAddress    string   `json:"startAddress"`
	EndAddress      string   `json:"endAddress"`
	ETA             int64    `json:"eta"`
	DrivingDistance string   `json:"drivingDistance"`
	Steps           []string `json:"steps"`
}

type Driver struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type ManifestGenerateParams struct {
	WorkerId string `json:"workerId"`
	HubId    string `json:"hubId"`
}

type DeliveryManifest struct {
	ManifestDate  int64              `json:"manifestDate"`
	DepartureTime int64              `json:"departureTime"`
	Driver        Driver             `json:"driver"`
	Vehicle       WorkerVehicleParam `json:"vehicle"`
	HubAddress    string             `json:"hubAddress"`
	Tasks         []Task             `json:"tasks"`
	TurnByTurn    []TurnByTurn       `json:"turnByTurn"`
	TotalDistance string             `json:"totalDistance"`
}
