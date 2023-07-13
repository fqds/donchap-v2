package dto

type PlayerParameterDto struct {
	ID          int `json:"-" db:"id"`
	Position    int `json:"position" db:"position"`
	ParameterID int `json:"parameter_id" db:"parameter_id"`
	Value       int `json:"value" db:"value"`
}
