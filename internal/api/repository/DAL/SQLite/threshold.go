package SQLite

import (
	"context"
	"database/sql"
	"goapi/internal/api/repository/DAL"
	"goapi/internal/api/repository/models"
)

type ThresholdRepository struct {
	sqlDB        *sql.DB
	createStmt,
	readStmt,
	readManyStmt,
	updateStmt,
	deleteStmt *sql.Stmt
	ctx context.Context
}

func NewThresholdRepository(sqlDB DAL.SQLDatabase, ctx context.Context) (models.ThresholdRepository, error) {
	repo := &ThresholdRepository{
		sqlDB: sqlDB.Connection(),
		ctx:   ctx,
	}

	// Create the thresholds table if it doesn't exist
	if _, err := repo.sqlDB.Exec(`CREATE TABLE IF NOT EXISTS thresholds (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sensor_type VARCHAR(50) NOT NULL,
		min_value FLOAT NOT NULL,
		max_value FLOAT NOT NULL,
		updated_at TIMESTAMP
	);`); err != nil {
		repo.sqlDB.Close()
		return nil, err
	}

	// Prepare SQL statements
	createStmt, err := repo.sqlDB.Prepare(`INSERT INTO thresholds (sensor_type, min_value, max_value, updated_at) VALUES (?, ?, ?, ?)`)
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.createStmt = createStmt

	readStmt, err := repo.sqlDB.Prepare(`SELECT id, sensor_type, min_value, max_value, updated_at FROM thresholds WHERE id = ?`)
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readStmt = readStmt

	readManyStmt, err := repo.sqlDB.Prepare(`SELECT id, sensor_type, min_value, max_value, updated_at FROM thresholds LIMIT ? OFFSET ?`)
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.readManyStmt = readManyStmt

	updateStmt, err := repo.sqlDB.Prepare(`UPDATE thresholds SET sensor_type = ?, min_value = ?, max_value = ?, updated_at = ? WHERE id = ?`)
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.updateStmt = updateStmt

	deleteStmt, err := repo.sqlDB.Prepare(`DELETE FROM thresholds WHERE id = ?`)
	if err != nil {
		repo.sqlDB.Close()
		return nil, err
	}
	repo.deleteStmt = deleteStmt

	go CloseThreshold(ctx, repo)

	return repo, nil
}

func CloseThreshold(ctx context.Context, r *ThresholdRepository) {
	<-ctx.Done()
	r.createStmt.Close()
	r.readStmt.Close()
	r.readManyStmt.Close()
	r.updateStmt.Close()
	r.deleteStmt.Close()
	r.sqlDB.Close()
}

func (r *ThresholdRepository) Create(threshold *models.Threshold, ctx context.Context) error {
	res, err := r.createStmt.ExecContext(ctx, threshold.SensorType, threshold.MinValue, threshold.MaxValue, threshold.UpdatedAt)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	threshold.ID = int(id)
	return nil
}

func (r *ThresholdRepository) ReadOne(id int, ctx context.Context) (*models.Threshold, error) {
	row := r.readStmt.QueryRowContext(ctx, id)
	var threshold models.Threshold
	err := row.Scan(&threshold.ID, &threshold.SensorType, &threshold.MinValue, &threshold.MaxValue, &threshold.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &threshold, nil
}

func (r *ThresholdRepository) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Threshold, error) {
	offset := rowsPerPage * (page - 1)
	rows, err := r.readManyStmt.QueryContext(ctx, rowsPerPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var thresholds []*models.Threshold
	for rows.Next() {
		var threshold models.Threshold
		err := rows.Scan(&threshold.ID, &threshold.SensorType, &threshold.MinValue, &threshold.MaxValue, &threshold.UpdatedAt)
		if err != nil {
			return nil, err
		}
		thresholds = append(thresholds, &threshold)
	}
	return thresholds, nil
}

func (r *ThresholdRepository) Update(threshold *models.Threshold, ctx context.Context) (int64, error) {
	res, err := r.updateStmt.ExecContext(ctx, threshold.SensorType, threshold.MinValue, threshold.MaxValue, threshold.UpdatedAt, threshold.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *ThresholdRepository) Delete(threshold *models.Threshold, ctx context.Context) (int64, error) {
    res, err := r.deleteStmt.ExecContext(ctx, threshold.ID)
    if err != nil {
        return 0, err
    }
    return res.RowsAffected()
}

