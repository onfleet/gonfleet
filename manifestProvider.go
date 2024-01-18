package onfleet

type TurnByTurn struct {
	DrivingDistance string   `json:"driving_distance"`
	EndAddress      string   `json:"end_address"`
	ETA             int64    `json:"eta"`
	StartAddress    string   `json:"start_address"`
	Steps           []string `json:"steps"`
}

type Driver struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type ManifestGenerateParams struct {
	HubId    string `json:"hubId"`
	WorkerId string `json:"workerId"`
}

type DeliveryManifest struct {
	DepartureTime int64              `json:"departureTime"`
	Driver        Driver             `json:"driver"`
	HubAddress    string             `json:"hubAddress"`
	ManifestDate  int64              `json:"manifestDate"`
	Tasks         []Task             `json:"tasks"`
	TotalDistance string             `json:"totalDistance"`
	TurnByTurn    []TurnByTurn       `json:"turnByTurn"`
	Vehicle       WorkerVehicleParam `json:"vehicle"`
}
