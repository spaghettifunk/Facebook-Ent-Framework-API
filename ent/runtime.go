// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/GhvstCode/LR-ENT/ent/notes"
	"github.com/GhvstCode/LR-ENT/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	notesFields := schema.Notes{}.Fields()
	_ = notesFields
	// notesDescPrivate is the schema descriptor for Private field.
	notesDescPrivate := notesFields[2].Descriptor()
	// notes.DefaultPrivate holds the default value on creation for the Private field.
	notes.DefaultPrivate = notesDescPrivate.Default.(bool)
	// notesDescCreatedAt is the schema descriptor for created_at field.
	notesDescCreatedAt := notesFields[3].Descriptor()
	// notes.DefaultCreatedAt holds the default value on creation for the created_at field.
	notes.DefaultCreatedAt = notesDescCreatedAt.Default.(func() time.Time)
}
