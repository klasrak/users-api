package utils

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/klasrak/users-api/rerrors"
)

func SanitizeUpdateParams(u interface{}) (map[string]interface{}, error) {
	var r map[string]interface{}

	us, err := json.Marshal(u)

	if err != nil {
		return nil, rerrors.NewInternal()
	}

	if err := json.Unmarshal(us, &r); err != nil {
		return nil, rerrors.NewInternal()
	}

	for i, v := range r {
		switch t := v.(type) {
		case string:
			if v == "" {
				r[i] = sql.NullString{}
			} else {
				data, err := time.Parse(time.RFC3339, v.(string))

				if err == nil && data.IsZero() {
					r[i] = sql.NullString{}
				} else if !data.IsZero() {
					r[i] = data
				} else {
					r[i] = t
				}
			}
		default:
			r[i] = t
		}
	}

	return r, nil
}
