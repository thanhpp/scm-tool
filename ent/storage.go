// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/thanhpp/scm/ent/storage"
)

// Storage is the model entity for the Storage schema.
type Storage struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StorageQuery when eager-loading is set.
	Edges StorageEdges `json:"edges"`
}

// StorageEdges holds the relations/edges for other nodes in the graph.
type StorageEdges struct {
	// StorageSerial holds the value of the storage_serial edge.
	StorageSerial []*Serial `json:"storage_serial,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// StorageSerialOrErr returns the StorageSerial value or an error if the edge
// was not loaded in eager-loading.
func (e StorageEdges) StorageSerialOrErr() ([]*Serial, error) {
	if e.loadedTypes[0] {
		return e.StorageSerial, nil
	}
	return nil, &NotLoadedError{edge: "storage_serial"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Storage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case storage.FieldName, storage.FieldLocation:
			values[i] = new(sql.NullString)
		case storage.FieldCreateTime, storage.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case storage.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Storage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Storage fields.
func (s *Storage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case storage.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case storage.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case storage.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case storage.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case storage.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				s.Location = value.String
			}
		}
	}
	return nil
}

// QueryStorageSerial queries the "storage_serial" edge of the Storage entity.
func (s *Storage) QueryStorageSerial() *SerialQuery {
	return (&StorageClient{config: s.config}).QueryStorageSerial(s)
}

// Update returns a builder for updating this Storage.
// Note that you need to call Storage.Unwrap() before calling this method if this Storage
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Storage) Update() *StorageUpdateOne {
	return (&StorageClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Storage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Storage) Unwrap() *Storage {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Storage is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Storage) String() string {
	var builder strings.Builder
	builder.WriteString("Storage(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteString(", location=")
	builder.WriteString(s.Location)
	builder.WriteByte(')')
	return builder.String()
}

// Storages is a parsable slice of Storage.
type Storages []*Storage

func (s Storages) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}