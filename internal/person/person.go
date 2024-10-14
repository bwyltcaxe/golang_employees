package person

import (
	"errors"
)

type Person struct {
	Name  string
	Gender string
	ID    int
}

var (
	ErrInvalidGender = errors.New("invalid gender")
	ErrInvalidName = errors.New("invalid name")
	ErrInvalidID = errors.New("invalid id")
)

// SetName sets the name of the person.
// 
// Parameters:
//  newName - The new name of the person.
// 
// Returns:
//  error - An error if the name is invalid (empty string).
func (p *Person) SetName(newName string) error {
	if len(newName) == 0 {
		return ErrInvalidName
	}
	p.Name = newName
	return nil
}

// SetGender sets the gender of the person.
// 
// Parameters:
//  newGender - The new gender of the person (either "man" or "woman").
// 
// Returns:
//  error - An error if the gender is invalid.
func (p *Person) SetGender(newGender string) error {
	if newGender != "man" && newGender != "woman" {
		return ErrInvalidGender
	}
	p.Gender = newGender
	return nil
}

// SetID sets the ID of the person.
// 
// Parameters:
//  newID - The new ID of the person (must be greater than 0).
// 
// Returns:
//  error - An error if the ID is invalid.
func (p *Person) SetID(newID int) error {
	if newID <= 0 {
		return ErrInvalidID
	}
	p.ID = newID
	return nil
}

