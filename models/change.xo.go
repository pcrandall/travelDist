// Package models contains the types for schema ''.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Change represents a row from 'CHANGE'.
type Change struct {
	ID        sql.NullInt64  `json:"id"`        // id
	Shuttle   string         `json:"shuttle"`   // shuttle
	Distance  int            `json:"distance"`  // distance
	Timestamp sql.NullString `json:"timestamp"` // timestamp
	Notes     sql.NullString `json:"notes"`     // notes
	UUID      sql.NullString `json:"uuid"`      // uuid

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Change exists in the database.
func (c *Change) Exists() bool {
	return c._exists
}

// Deleted provides information if the Change has been deleted from the database.
func (c *Change) Deleted() bool {
	return c._deleted
}

// Insert inserts the Change to the database.
func (c *Change) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO CHANGE (` +
		`shuttle, distance, timestamp, notes, uuid` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, c.Shuttle, c.Distance, c.Timestamp, c.Notes, c.UUID)
	res, err := db.Exec(sqlstr, c.Shuttle, c.Distance, c.Timestamp, c.Notes, c.UUID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	c.ID = sql.NullInt64(id)
	c._exists = true

	return nil
}

// Update updates the Change in the database.
func (c *Change) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE CHANGE SET ` +
		`shuttle = ?, distance = ?, timestamp = ?, notes = ?, uuid = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, c.Shuttle, c.Distance, c.Timestamp, c.Notes, c.UUID, c.ID)
	_, err = db.Exec(sqlstr, c.Shuttle, c.Distance, c.Timestamp, c.Notes, c.UUID, c.ID)
	return err
}

// Save saves the Change to the database.
func (c *Change) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Delete deletes the Change from the database.
func (c *Change) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM CHANGE WHERE id = ?`

	// run query
	XOLog(sqlstr, c.ID)
	_, err = db.Exec(sqlstr, c.ID)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// ChangeByID retrieves a row from 'CHANGE' as a Change.
//
// Generated from index 'CHANGE_id_pkey'.
func ChangeByID(db XODB, id sql.NullInt64) (*Change, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, shuttle, distance, timestamp, notes, uuid ` +
		`FROM CHANGE ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	c := Change{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&c.ID, &c.Shuttle, &c.Distance, &c.Timestamp, &c.Notes, &c.UUID)
	if err != nil {
		return nil, err
	}

	return &c, nil
}