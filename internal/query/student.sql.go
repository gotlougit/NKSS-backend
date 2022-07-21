// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: student.sql

package query

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const getStudent = `-- name: GetStudent :one
SELECT
    roll_number, section, name, gender, mobile, birth_date, email, batch, hostel_id, room_id, discord_id, is_verified,
    CAST(ARRAY(SELECT cm.club_name FROM club_member AS cm WHERE cm.roll_number = $1) AS VARCHAR[]) AS clubs
FROM
    student
WHERE roll_number = $1
`

type GetStudentRow struct {
	RollNumber string         `json:"roll_number"`
	Section    string         `json:"section"`
	Name       string         `json:"name"`
	Gender     sql.NullString `json:"gender"`
	Mobile     sql.NullString `json:"mobile"`
	BirthDate  sql.NullTime   `json:"birth_date"`
	Email      string         `json:"email"`
	Batch      int16          `json:"batch"`
	HostelID   sql.NullString `json:"hostel_id"`
	RoomID     sql.NullString `json:"room_id"`
	DiscordID  sql.NullInt64  `json:"discord_id"`
	IsVerified bool           `json:"is_verified"`
	Clubs      []string       `json:"clubs"`
}

func (q *Queries) GetStudent(ctx context.Context, rollNumber string) (GetStudentRow, error) {
	row := q.db.QueryRowContext(ctx, getStudent, rollNumber)
	var i GetStudentRow
	err := row.Scan(
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
		pq.Array(&i.Clubs),
	)
	return i, err
}