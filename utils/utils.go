package utils

import "database/sql"

func NullToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}
