package database

import "strings"

func IsUniqueViolation(err error) bool {
	return strings.Contains(err.Error(), "violates unique constraint")
}
