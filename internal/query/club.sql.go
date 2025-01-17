// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: club.sql

package query

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/lib/pq"
)

const createClubAdmin = `-- name: CreateClubAdmin :exec
INSERT INTO club_admin (
    club_name, position, roll_number
)
VALUES (
    (SELECT name from club WHERE name = $1 or alias = $1),
    $2,
    $3
)
`

type CreateClubAdminParams struct {
	Name       string `json:"name"`
	Position   string `json:"position"`
	RollNumber string `json:"roll_number"`
}

func (q *Queries) CreateClubAdmin(ctx context.Context, arg CreateClubAdminParams) error {
	_, err := q.db.ExecContext(ctx, createClubAdmin, arg.Name, arg.Position, arg.RollNumber)
	return err
}

const createClubFaculty = `-- name: CreateClubFaculty :exec
INSERT INTO club_faculty (
    club_name, emp_id
)
VALUES (
    (SELECT c.name from club c WHERE c.name = $1 or c.alias = $1),
    $2
)
`

type CreateClubFacultyParams struct {
	Name  string `json:"name"`
	EmpID int32  `json:"emp_id"`
}

func (q *Queries) CreateClubFaculty(ctx context.Context, arg CreateClubFacultyParams) error {
	_, err := q.db.ExecContext(ctx, createClubFaculty, arg.Name, arg.EmpID)
	return err
}

const createClubMember = `-- name: CreateClubMember :exec
INSERT INTO club_member (
    club_name, roll_number
)
VALUES (
    (SELECT name from club WHERE name = $1 or alias = $1),
    $2
)
`

type CreateClubMemberParams struct {
	Name       string `json:"name"`
	RollNumber string `json:"roll_number"`
}

func (q *Queries) CreateClubMember(ctx context.Context, arg CreateClubMemberParams) error {
	_, err := q.db.ExecContext(ctx, createClubMember, arg.Name, arg.RollNumber)
	return err
}

const createClubSocial = `-- name: CreateClubSocial :exec
INSERT INTO club_social (
    name, platform_type, link
)
VALUES (
    (SELECT c.name from club c WHERE c.name = $1 or c.alias = $1),
    $2,
    $3
)
`

type CreateClubSocialParams struct {
	Name         string `json:"name"`
	PlatformType string `json:"platform_type"`
	Link         string `json:"link"`
}

func (q *Queries) CreateClubSocial(ctx context.Context, arg CreateClubSocialParams) error {
	_, err := q.db.ExecContext(ctx, createClubSocial, arg.Name, arg.PlatformType, arg.Link)
	return err
}

const deleteClubAdmin = `-- name: DeleteClubAdmin :exec
DELETE FROM club_admin
WHERE
    club_name = (SELECT name FROM club WHERE name = $1 OR alias = $1)
    AND roll_number = $2
`

type DeleteClubAdminParams struct {
	Name       string `json:"name"`
	RollNumber string `json:"roll_number"`
}

func (q *Queries) DeleteClubAdmin(ctx context.Context, arg DeleteClubAdminParams) error {
	_, err := q.db.ExecContext(ctx, deleteClubAdmin, arg.Name, arg.RollNumber)
	return err
}

const deleteClubFaculty = `-- name: DeleteClubFaculty :exec
DELETE FROM club_faculty cf
WHERE
    cf.club_name = (SELECT c.name FROM club c WHERE c.name = $1 OR c.alias = $1)
    AND cf.emp_id = $2
`

type DeleteClubFacultyParams struct {
	Name  string `json:"name"`
	EmpID int32  `json:"emp_id"`
}

func (q *Queries) DeleteClubFaculty(ctx context.Context, arg DeleteClubFacultyParams) error {
	_, err := q.db.ExecContext(ctx, deleteClubFaculty, arg.Name, arg.EmpID)
	return err
}

const deleteClubMember = `-- name: DeleteClubMember :exec
DELETE FROM club_member
WHERE
    club_name = (SELECT name FROM club WHERE name = $1 OR alias = $1)
    AND roll_number = $2
`

type DeleteClubMemberParams struct {
	Name       string `json:"name"`
	RollNumber string `json:"roll_number"`
}

func (q *Queries) DeleteClubMember(ctx context.Context, arg DeleteClubMemberParams) error {
	_, err := q.db.ExecContext(ctx, deleteClubMember, arg.Name, arg.RollNumber)
	return err
}

const deleteClubSocial = `-- name: DeleteClubSocial :exec
DELETE FROM club_social
WHERE
    club_name = (SELECT name FROM club WHERE name = $1 OR alias = $1)
    AND platform_type = $2
`

type DeleteClubSocialParams struct {
	Name         string `json:"name"`
	PlatformType string `json:"platform_type"`
}

func (q *Queries) DeleteClubSocial(ctx context.Context, arg DeleteClubSocialParams) error {
	_, err := q.db.ExecContext(ctx, deleteClubSocial, arg.Name, arg.PlatformType)
	return err
}

const getClub = `-- name: GetClub :one
SELECT
    name, alias, branch, kind, description,
    (
        SELECT
            COALESCE(JSON_AGG(JSON_BUILD_OBJECT('name', f.name, 'phone', f.mobile) ORDER BY f.name), '[]')::JSON
        FROM
            faculty AS f
        JOIN club_faculty AS cf ON f.emp_id = cf.emp_id
        WHERE
            cf.club_name = club.name
    ) AS faculties,
    (
        SELECT
            JSON_AGG(JSON_BUILD_OBJECT('platform', cs.platform_type, 'link', cs.link) ORDER BY cs.platform_type)
        FROM
            club_social AS cs
        WHERE
            cs.club_name = club.name
    ) AS socials,
    (
        SELECT
            COALESCE(JSON_AGG(JSON_BUILD_OBJECT('position', ca.position, 'roll', ca.roll_number)), '[]')::JSON
        FROM
            club_admin AS ca
        WHERE
            ca.club_name = club.name
    ) AS admins
FROM
    club
WHERE
    club.name = $1
    OR club.alias = $1
`

type GetClubRow struct {
	Name        string          `json:"name"`
	Alias       sql.NullString  `json:"alias"`
	Branch      []string        `json:"branch"`
	Kind        string          `json:"kind"`
	Description string          `json:"description"`
	Faculties   json.RawMessage `json:"faculties"`
	Socials     json.RawMessage `json:"socials"`
	Admins      json.RawMessage `json:"admins"`
}

func (q *Queries) GetClub(ctx context.Context, name string) (GetClubRow, error) {
	row := q.db.QueryRowContext(ctx, getClub, name)
	var i GetClubRow
	err := row.Scan(
		&i.Name,
		&i.Alias,
		pq.Array(&i.Branch),
		&i.Kind,
		&i.Description,
		&i.Faculties,
		&i.Socials,
		&i.Admins,
	)
	return i, err
}

const getClubAdmins = `-- name: GetClubAdmins :many
SELECT
    s.roll_number, s.section, s.name, s.gender, s.mobile, s.birth_date, s.email, s.batch, s.hostel_id, s.room_id, s.discord_id, s.is_verified, admin.position
FROM
    student s
    JOIN club_admin admin ON s.roll_number = admin.roll_number
WHERE
    admin.club_name = $1
    OR $1 = (SELECT alias FROM club WHERE name = admin.club_name)
`

type GetClubAdminsRow struct {
	RollNumber string         `json:"roll_number"`
	Section    string         `json:"section"`
	Name       string         `json:"name"`
	Gender     sql.NullString `json:"gender"`
	Mobile     sql.NullString `json:"mobile"`
	BirthDate  sql.NullTime   `json:"birth_date"`
	Email      string         `json:"email"`
	Batch      int16          `json:"batch"`
	HostelID   string         `json:"hostel_id"`
	RoomID     sql.NullString `json:"room_id"`
	DiscordID  sql.NullInt64  `json:"discord_id"`
	IsVerified bool           `json:"is_verified"`
	Position   string         `json:"position"`
}

func (q *Queries) GetClubAdmins(ctx context.Context, clubName string) ([]GetClubAdminsRow, error) {
	rows, err := q.db.QueryContext(ctx, getClubAdmins, clubName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetClubAdminsRow
	for rows.Next() {
		var i GetClubAdminsRow
		if err := rows.Scan(
			&i.RollNumber,
			&i.Section,
			&i.Name,
			&i.Gender,
			&i.Mobile,
			&i.BirthDate,
			&i.Email,
			&i.Batch,
			&i.HostelID,
			&i.RoomID,
			&i.DiscordID,
			&i.IsVerified,
			&i.Position,
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

const getClubFaculty = `-- name: GetClubFaculty :many
SELECT
    f.name, f.mobile
FROM
    faculty AS f
    JOIN club_faculty AS cf ON f.emp_id = cf.emp_id
WHERE
    cf.club_name = $1
    OR $1 = (SELECT alias FROM club WHERE name = cf.club_name)
`

type GetClubFacultyRow struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

func (q *Queries) GetClubFaculty(ctx context.Context, clubName string) ([]GetClubFacultyRow, error) {
	rows, err := q.db.QueryContext(ctx, getClubFaculty, clubName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetClubFacultyRow
	for rows.Next() {
		var i GetClubFacultyRow
		if err := rows.Scan(&i.Name, &i.Mobile); err != nil {
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

const getClubMembers = `-- name: GetClubMembers :many
SELECT
    s.roll_number, s.section, s.name, s.gender, s.mobile, s.birth_date, s.email, s.batch, s.hostel_id, s.room_id, s.discord_id, s.is_verified
FROM
    student s
    JOIN club_member member ON s.roll_number = member.roll_number
WHERE
    member.club_name = $1
    OR $1 = (SELECT alias FROM club WHERE name = member.club_name)
`

func (q *Queries) GetClubMembers(ctx context.Context, clubName string) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, getClubMembers, clubName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.RollNumber,
			&i.Section,
			&i.Name,
			&i.Gender,
			&i.Mobile,
			&i.BirthDate,
			&i.Email,
			&i.Batch,
			&i.HostelID,
			&i.RoomID,
			&i.DiscordID,
			&i.IsVerified,
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

const getClubSocials = `-- name: GetClubSocials :many
SELECT
    platform_type,
    link
FROM
    club_social
WHERE
    club_name = $1
    OR $1 = (SELECT alias FROM club WHERE name = club_name)
`

type GetClubSocialsRow struct {
	PlatformType string `json:"platform_type"`
	Link         string `json:"link"`
}

func (q *Queries) GetClubSocials(ctx context.Context, clubName string) ([]GetClubSocialsRow, error) {
	rows, err := q.db.QueryContext(ctx, getClubSocials, clubName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetClubSocialsRow
	for rows.Next() {
		var i GetClubSocialsRow
		if err := rows.Scan(&i.PlatformType, &i.Link); err != nil {
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

const getClubs = `-- name: GetClubs :many
SELECT
    name, alias, branch, kind, description,
    (
        SELECT
            COALESCE(JSON_AGG(JSON_BUILD_OBJECT('name', f.name, 'phone', f.mobile) ORDER BY f.name), '[]')::JSON
        FROM
            faculty AS f
        JOIN club_faculty AS cf ON f.emp_id = cf.emp_id
        WHERE
            cf.club_name = club.name
    ) AS faculties,
    (
        SELECT
            JSON_AGG(JSON_BUILD_OBJECT('platform', cs.platform_type, 'link', cs.link) ORDER BY cs.platform_type)
        FROM
            club_social AS cs
        WHERE
            cs.club_name = club.name
    ) AS socials,
    (
        SELECT
            COALESCE(JSON_AGG(JSON_BUILD_OBJECT('position', ca.position, 'roll', ca.roll_number)), '[]')::JSON
        FROM
            club_admin AS ca
        WHERE
            ca.club_name = club.name
    ) AS admins
FROM
    club
ORDER BY
    club.name
`

type GetClubsRow struct {
	Name        string          `json:"name"`
	Alias       sql.NullString  `json:"alias"`
	Branch      []string        `json:"branch"`
	Kind        string          `json:"kind"`
	Description string          `json:"description"`
	Faculties   json.RawMessage `json:"faculties"`
	Socials     json.RawMessage `json:"socials"`
	Admins      json.RawMessage `json:"admins"`
}

func (q *Queries) GetClubs(ctx context.Context) ([]GetClubsRow, error) {
	rows, err := q.db.QueryContext(ctx, getClubs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetClubsRow
	for rows.Next() {
		var i GetClubsRow
		if err := rows.Scan(
			&i.Name,
			&i.Alias,
			pq.Array(&i.Branch),
			&i.Kind,
			&i.Description,
			&i.Faculties,
			&i.Socials,
			&i.Admins,
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

const updateClubSocials = `-- name: UpdateClubSocials :exec
UPDATE
    club_social
SET
    link = $2
WHERE
    platform_type = $1
    AND club_name = $3
    OR $3 = (SELECT alias FROM club WHERE name = club_name)
`

type UpdateClubSocialsParams struct {
	PlatformType string `json:"platform_type"`
	Link         string `json:"link"`
	ClubName     string `json:"club_name"`
}

func (q *Queries) UpdateClubSocials(ctx context.Context, arg UpdateClubSocialsParams) error {
	_, err := q.db.ExecContext(ctx, updateClubSocials, arg.PlatformType, arg.Link, arg.ClubName)
	return err
}
