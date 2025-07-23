package subscriptions

import (
	"context"
	"database/sql"
)

func (r *repo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM subscriptions WHERE user_id = $1`

	result, err := r.postgres.Pool.Exec(ctx, query, id)
	if err != nil {
		r.logger.Error("failed to delete subscription: " + err.Error())
		return err
	}

	if result.RowsAffected() == 0 {
		r.logger.Warn("no subscription found with id: " + id)
		return sql.ErrNoRows
	}

	r.logger.Info("deleted subscription with id: " + id)
	return nil
}
