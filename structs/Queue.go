package structs

type Queue struct {
	Addresses []string `json:"addresses"`
	Email     Email    `json:"email"`
}
