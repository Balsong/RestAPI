package macinfo

import (
	"API/internal/store/postgres/mac_info/model"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func buildSaveInfoQuery(info model.MacInfo) (string, []interface{}, error) {
	fields := map[string]interface{}{
		"id":   uuid.New(),
		"temp": info.Temperature,
		"cpu":  info.Cpu,
	}

	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert("macinfo").
		SetMap(fields).
		Suffix("returning id")

	q, args, err := query.ToSql()
	if err != nil {
		return "", nil, fmt.Errorf("sql.ToSql failed to save info into postgres: %w", err)
	}

	return q, args, nil
}
