// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Stock       int         `json:"stock"`
	Category    []*Category `json:"category"`
}

type Profile struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

type ShoppingCart struct {
	ID       string     `json:"id"`
	Status   int        `json:"status"`
	User     *User      `json:"user"`
	Products []*Product `json:"products"`
}

type User struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Status bool   `json:"status"`
	Type   int    `json:"type"`
}
