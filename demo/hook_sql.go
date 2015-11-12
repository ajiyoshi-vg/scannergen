package demo

// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

import (
	"database/sql"
)

func ScanHook(row *sql.Row) (*Hook, error) {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 bool
	var v7 string
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 string
	var v15 string

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
		&v9,
		&v10,
		&v11,
		&v12,
		&v13,
		&v14,
		&v15,
	)
	if err != nil {
		return nil, err
	}

	v := &Hook{}
	v.ID = v0
	v.Sha = v1
	v.After = v2
	v.Before = v3
	v.Created = v4
	v.Deleted = v5
	v.Forced = v6
	v.HeadCommit = &Commit{}
	v.HeadCommit.ID = v7
	v.HeadCommit.Message = v8
	v.HeadCommit.Timestamp = v9
	v.HeadCommit.Author = &Author{}
	v.HeadCommit.Author.Name = v10
	v.HeadCommit.Author.Email = v11
	v.HeadCommit.Author.Username = v12
	v.HeadCommit.Committer = &Author{}
	v.HeadCommit.Committer.Name = v13
	v.HeadCommit.Committer.Email = v14
	v.HeadCommit.Committer.Username = v15

	return v, nil
}

func ScanHooks(rows *sql.Rows) ([]*Hook, error) {
	var err error
	var vv []*Hook

	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 bool
	var v7 string
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 string
	var v15 string

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,
			&v7,
			&v8,
			&v9,
			&v10,
			&v11,
			&v12,
			&v13,
			&v14,
			&v15,
		)
		if err != nil {
			return vv, err
		}

		v := &Hook{}
		v.ID = v0
		v.Sha = v1
		v.After = v2
		v.Before = v3
		v.Created = v4
		v.Deleted = v5
		v.Forced = v6
		v.HeadCommit = &Commit{}
		v.HeadCommit.ID = v7
		v.HeadCommit.Message = v8
		v.HeadCommit.Timestamp = v9
		v.HeadCommit.Author = &Author{}
		v.HeadCommit.Author.Name = v10
		v.HeadCommit.Author.Email = v11
		v.HeadCommit.Author.Username = v12
		v.HeadCommit.Committer = &Author{}
		v.HeadCommit.Committer.Name = v13
		v.HeadCommit.Committer.Email = v14
		v.HeadCommit.Committer.Username = v15

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceHook(v *Hook) []interface{} {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 bool
	var v5 bool
	var v6 bool
	var v7 string
	var v8 string
	var v9 string
	var v10 string
	var v11 string
	var v12 string
	var v13 string
	var v14 string
	var v15 string

	v0 = v.ID
	v1 = v.Sha
	v2 = v.After
	v3 = v.Before
	v4 = v.Created
	v5 = v.Deleted
	v6 = v.Forced
	if v.HeadCommit != nil {
		v7 = v.HeadCommit.ID
		v8 = v.HeadCommit.Message
		v9 = v.HeadCommit.Timestamp
		if v.HeadCommit.Author != nil {
			v10 = v.HeadCommit.Author.Name
			v11 = v.HeadCommit.Author.Email
			v12 = v.HeadCommit.Author.Username
		}

	}

	if v.HeadCommit.Committer != nil {
		v13 = v.HeadCommit.Committer.Name
		v14 = v.HeadCommit.Committer.Email
		v15 = v.HeadCommit.Committer.Username
	}

	return []interface{}{
		v0,
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,
		v7,
		v8,
		v9,
		v10,
		v11,
		v12,
		v13,
		v14,
		v15,
	}
}

func ScanCommit(row *sql.Row) (*Commit, error) {
	var v0 string
	var v1 string
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 string
	var v7 string
	var v8 string

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
	)
	if err != nil {
		return nil, err
	}

	v := &Commit{}
	v.ID = v0
	v.Message = v1
	v.Timestamp = v2
	v.Author = &Author{}
	v.Author.Name = v3
	v.Author.Email = v4
	v.Author.Username = v5
	v.Committer = &Author{}
	v.Committer.Name = v6
	v.Committer.Email = v7
	v.Committer.Username = v8

	return v, nil
}

func ScanCommits(rows *sql.Rows) ([]*Commit, error) {
	var err error
	var vv []*Commit

	var v0 string
	var v1 string
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 string
	var v7 string
	var v8 string

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,
			&v7,
			&v8,
		)
		if err != nil {
			return vv, err
		}

		v := &Commit{}
		v.ID = v0
		v.Message = v1
		v.Timestamp = v2
		v.Author = &Author{}
		v.Author.Name = v3
		v.Author.Email = v4
		v.Author.Username = v5
		v.Committer = &Author{}
		v.Committer.Name = v6
		v.Committer.Email = v7
		v.Committer.Username = v8

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceCommit(v *Commit) []interface{} {
	var v0 string
	var v1 string
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 string
	var v7 string
	var v8 string

	v0 = v.ID
	v1 = v.Message
	v2 = v.Timestamp
	if v.Author != nil {
		v3 = v.Author.Name
		v4 = v.Author.Email
		v5 = v.Author.Username
	}

	if v.Committer != nil {
		v6 = v.Committer.Name
		v7 = v.Committer.Email
		v8 = v.Committer.Username
	}

	return []interface{}{
		v0,
		v1,
		v2,
		v3,
		v4,
		v5,
		v6,
		v7,
		v8,
	}
}

func ScanAuthor(row *sql.Row) (*Author, error) {
	var v0 string
	var v1 string
	var v2 string

	err := row.Scan(
		&v0,
		&v1,
		&v2,
	)
	if err != nil {
		return nil, err
	}

	v := &Author{}
	v.Name = v0
	v.Email = v1
	v.Username = v2

	return v, nil
}

func ScanAuthors(rows *sql.Rows) ([]*Author, error) {
	var err error
	var vv []*Author

	var v0 string
	var v1 string
	var v2 string

	for rows.Next() {
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
		)
		if err != nil {
			return vv, err
		}

		v := &Author{}
		v.Name = v0
		v.Email = v1
		v.Username = v2

		vv = append(vv, v)
	}
	return vv, rows.Err()
}

func SliceAuthor(v *Author) []interface{} {
	var v0 string
	var v1 string
	var v2 string

	v0 = v.Name
	v1 = v.Email
	v2 = v.Username

	return []interface{}{
		v0,
		v1,
		v2,
	}
}
