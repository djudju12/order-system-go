// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package repository

import ()

type Product struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}
