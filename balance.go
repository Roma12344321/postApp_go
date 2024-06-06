package postApp

type Balance struct {
	PersonId int     `json:"person_id" db:"person_id"`
	Sum      float64 `json:"sum" db:"sum"`
}
