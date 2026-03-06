package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	model "github.com/naufalb95/trackr/internal/models"
)

type TaskRepository interface {
	FindAll(ctx context.Context) ([]model.Task, error)
	FindById(ctx context.Context, taskId string) (*model.Task, error)
	Create(ctx context.Context, task *model.Task) error
	Update(ctx context.Context, taskId string, task *model.Task) error
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
			id,
			title,
			description,
			status,
			created_at,
			updated_at
		FROM tasks
		ORDER BY created_at DESC;
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
		WHERE "id" = $1;
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

// TODO: To Be Implemented
func (r *PostgresTaskRepository) Update(ctx context.Context, taskId string, task *model.Task) error {
	return nil
}

// TODO: To Be Implemented
func (r *PostgresTaskRepository) Delete(ctx context.Context, taskId string) error {
	return nil
}
