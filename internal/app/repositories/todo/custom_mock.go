package repositories

import (
	"database/sql/driver"
	"time"
)

// Match satisfies sqlmock.Argument interface
type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

// UUIDArg is a custom argument type for matching UUIDs in SQL mock expectations.
type UUIDArg struct {
	Value string
}

// Match satisfies the sqlmock.Argument interface for custom UUID matching.
func (ua UUIDArg) Match(v driver.Value) bool {
	uuidValue, ok := v.(string)
	if !ok {
		return false
	}
	return ua.Value == uuidValue
}
