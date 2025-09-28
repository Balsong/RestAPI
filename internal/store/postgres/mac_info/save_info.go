package macinfo

import (
	"API/internal/store/postgres/mac_info/model"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *repository) SaveInfo(ctx context.Context, info model.MacInfo) error {
	query, args, err := buildSaveInfoQuery(info)
	if err != nil {
		return err
	}

	var id uuid.UUID

	err = r.db.GetContext(ctx, id, query, args...)
	if err != nil {
		return err
	}

	if id == uuid.Nil {
		return fmt.Errorf("uuid is nil")
	}

	return nil
}
