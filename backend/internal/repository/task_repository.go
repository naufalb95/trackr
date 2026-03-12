package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/naufalb95/trackr/internal/model"
)

type TaskRepository interface {
	FindAll(ctx context.Context) ([]model.Task, error)
	FindById(ctx context.Context, taskId string) (*model.Task, error)
	Create(ctx context.Context, task *model.Task) error
	UpdateStatus(ctx context.Context, taskId string, updatedFields map[string]any) error
	Delete(ctx context.Context, taskId string) error
}

type PostgresTaskRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresTaskRepository(pool *pgxpool.Pool) TaskRepository {
	return &PostgresTaskRepository{pool: pool}
}

func (r *PostgresTaskRepository) FindAll(ctx context.Context) ([]model.Task, error) {
	query := `
		SELECT
			"id",
			"title",
			"description",
			"status",
			"created_at",
			"updated_at"
		FROM "tasks"
		WHERE "status" != 'deleted'
		ORDER BY "created_at" DESC;
	`
	rows, err := r.pool.Query(ctx, query)

	if err != nil {
		fmt.Printf("Error occurred when trying to retrieve all tasks: %s", err)
		return nil, err
	}

	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var task model.Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("Error when scanning field: %w", err)
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error when retrieving rows: %w", err)
	}

	return tasks, nil
}

func (r *PostgresTaskRepository) FindById(ctx context.Context, taskId string) (*model.Task, error) {
	task := &model.Task{}

	query := `
		SELECT
			"id",
			"title",
			"description",
			"status",
			"created_at",
			"updated_at"
		FROM "tasks"
		WHERE "id" = $1 AND "status" != 'deleted';
	`

	err := r.pool.QueryRow(
		ctx,
		query,
		taskId,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return task, fmt.Errorf("Error when trying to find specific task: %w", err)
	}

	return task, nil
}

func (r *PostgresTaskRepository) Create(ctx context.Context, task *model.Task) error {
	query := `
		INSERT INTO "tasks" ("title", "description", "status")
		VALUES ($1, $2, $3)
		RETURNING "id", "created_at", "updated_at";
	`

	err := r.pool.QueryRow(ctx, query, task.Title, task.Description, task.Status).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return fmt.Errorf("Error when trying to create task: %w", err)
	}

	return nil
}

func (r *PostgresTaskRepository) UpdateStatus(ctx context.Context, taskId string, updatedFields map[string]any) error {
	if len(updatedFields) == 0 {
		return errors.New("No fields to update.")
	}

	setClauses := []string{}
	args := []any{}
	argIdx := 2 // first index for ID where clause

	args = append(args, taskId)

	for key, val := range updatedFields {
		clause := fmt.Sprintf(`"%s" = $%d`, key, argIdx)
		args = append(args, val)
		setClauses = append(setClauses, clause)
		argIdx++
	}

	query := fmt.Sprintf(
		`
			UPDATE "tasks"
			SET "updated_at" = NOW(), %s
			WHERE %s AND "status" != 'deleted';
		`,
		strings.Join(setClauses, ", "),
		`"id" = $1`,
	)

	fmt.Println(query)

	_, err := r.pool.Exec(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("Error when trying to update task status: %w", err)
	}

	return nil
}

func (r *PostgresTaskRepository) Delete(ctx context.Context, taskId string) error {
	query := `
		UPDATE "tasks"
		SET status = 'deleted'
		WHERE "id" = $1 AND "status" != 'deleted';
	`

	_, err := r.pool.Exec(ctx, query, taskId)

	if err != nil {
		return fmt.Errorf("Error when trying to delete task: %w", err)
	}

	return nil
}
