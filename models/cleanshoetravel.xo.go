// Package models contains the types for schema ''.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// CleanShoeTravel represents a row from 'clean_shoe_travel'.
type CleanShoeTravel struct {
	Shuttle             sql.NullString `json:"Shuttle"`               // Shuttle
	LastUpdated         sql.NullString `json:"Last_Updated"`          // Last_Updated
	ShoeTravel          sql.NullString `json:"Shoe_Travel"`           // Shoe_Travel
	DaysInstalled       sql.NullString `json:"Days_Installed"`        // Days_Installed
	ShoesLastDistance   sql.NullInt64  `json:"Shoes_Last_Distance"`   // Shoes_Last_Distance
	ShoesChangeDistance sql.NullString `json:"Shoes_Change_Distance"` // Shoes_Change_Distance
	ShoesLastChanged    sql.NullString `json:"Shoes_Last_Changed"`    // Shoes_Last_Changed
	Notes               sql.NullString `json:"Notes"`                 // Notes
	UUID                sql.NullString `json:"UUID"`                  // UUID
}