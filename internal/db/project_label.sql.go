// Code generated by sqlc. DO NOT EDIT.
// source: project_label.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createProjectLabel = `-- name: CreateProjectLabel :one
INSERT INTO project_label (project_id, label_color_id, created_date, name)
  VALUES ($1, $2, $3, $4) RETURNING project_label_id, project_id, label_color_id, created_date, name
`

type CreateProjectLabelParams struct {
	ProjectID    uuid.UUID      `json:"project_id"`
	LabelColorID uuid.UUID      `json:"label_color_id"`
	CreatedDate  time.Time      `json:"created_date"`
	Name         sql.NullString `json:"name"`
}

func (q *Queries) CreateProjectLabel(ctx context.Context, arg CreateProjectLabelParams) (ProjectLabel, error) {
	row := q.db.QueryRowContext(ctx, createProjectLabel,
		arg.ProjectID,
		arg.LabelColorID,
		arg.CreatedDate,
		arg.Name,
	)
	var i ProjectLabel
	err := row.Scan(
		&i.ProjectLabelID,
		&i.ProjectID,
		&i.LabelColorID,
		&i.CreatedDate,
		&i.Name,
	)
	return i, err
}

const deleteProjectLabelByID = `-- name: DeleteProjectLabelByID :exec
DELETE FROM project_label WHERE project_label_id = $1
`

func (q *Queries) DeleteProjectLabelByID(ctx context.Context, projectLabelID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProjectLabelByID, projectLabelID)
	return err
}

const getProjectLabelByID = `-- name: GetProjectLabelByID :one
SELECT project_label_id, project_id, label_color_id, created_date, name FROM project_label WHERE project_label_id = $1
`

func (q *Queries) GetProjectLabelByID(ctx context.Context, projectLabelID uuid.UUID) (ProjectLabel, error) {
	row := q.db.QueryRowContext(ctx, getProjectLabelByID, projectLabelID)
	var i ProjectLabel
	err := row.Scan(
		&i.ProjectLabelID,
		&i.ProjectID,
		&i.LabelColorID,
		&i.CreatedDate,
		&i.Name,
	)
	return i, err
}

const getProjectLabelsForProject = `-- name: GetProjectLabelsForProject :many
SELECT project_label_id, project_id, label_color_id, created_date, name FROM project_label WHERE project_id = $1
`

func (q *Queries) GetProjectLabelsForProject(ctx context.Context, projectID uuid.UUID) ([]ProjectLabel, error) {
	rows, err := q.db.QueryContext(ctx, getProjectLabelsForProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProjectLabel
	for rows.Next() {
		var i ProjectLabel
		if err := rows.Scan(
			&i.ProjectLabelID,
			&i.ProjectID,
			&i.LabelColorID,
			&i.CreatedDate,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProjectLabel = `-- name: UpdateProjectLabel :one
UPDATE project_label SET name = $2, label_color_id = $3 WHERE project_label_id = $1 RETURNING project_label_id, project_id, label_color_id, created_date, name
`

type UpdateProjectLabelParams struct {
	ProjectLabelID uuid.UUID      `json:"project_label_id"`
	Name           sql.NullString `json:"name"`
	LabelColorID   uuid.UUID      `json:"label_color_id"`
}

func (q *Queries) UpdateProjectLabel(ctx context.Context, arg UpdateProjectLabelParams) (ProjectLabel, error) {
	row := q.db.QueryRowContext(ctx, updateProjectLabel, arg.ProjectLabelID, arg.Name, arg.LabelColorID)
	var i ProjectLabel
	err := row.Scan(
		&i.ProjectLabelID,
		&i.ProjectID,
		&i.LabelColorID,
		&i.CreatedDate,
		&i.Name,
	)
	return i, err
}

const updateProjectLabelColor = `-- name: UpdateProjectLabelColor :one
UPDATE project_label SET label_color_id = $2 WHERE project_label_id = $1 RETURNING project_label_id, project_id, label_color_id, created_date, name
`

type UpdateProjectLabelColorParams struct {
	ProjectLabelID uuid.UUID `json:"project_label_id"`
	LabelColorID   uuid.UUID `json:"label_color_id"`
}

func (q *Queries) UpdateProjectLabelColor(ctx context.Context, arg UpdateProjectLabelColorParams) (ProjectLabel, error) {
	row := q.db.QueryRowContext(ctx, updateProjectLabelColor, arg.ProjectLabelID, arg.LabelColorID)
	var i ProjectLabel
	err := row.Scan(
		&i.ProjectLabelID,
		&i.ProjectID,
		&i.LabelColorID,
		&i.CreatedDate,
		&i.Name,
	)
	return i, err
}

const updateProjectLabelName = `-- name: UpdateProjectLabelName :one
UPDATE project_label SET name = $2 WHERE project_label_id = $1 RETURNING project_label_id, project_id, label_color_id, created_date, name
`

type UpdateProjectLabelNameParams struct {
	ProjectLabelID uuid.UUID      `json:"project_label_id"`
	Name           sql.NullString `json:"name"`
}

func (q *Queries) UpdateProjectLabelName(ctx context.Context, arg UpdateProjectLabelNameParams) (ProjectLabel, error) {
	row := q.db.QueryRowContext(ctx, updateProjectLabelName, arg.ProjectLabelID, arg.Name)
	var i ProjectLabel
	err := row.Scan(
		&i.ProjectLabelID,
		&i.ProjectID,
		&i.LabelColorID,
		&i.CreatedDate,
		&i.Name,
	)
	return i, err
}
