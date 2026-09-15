package models

type Spaceship struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Fuel       int    `json:"fuel"`
	IsReady    bool   `json:"is_ready"`
	EngineType string `json:"engine_type"`
	CrewCount  int    `json:"crew_count"`
}
