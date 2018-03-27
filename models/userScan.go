package models

import "database/sql"

// ScanUser db representatition
func ScanUser(r *sql.Row) (User, error) {
	var s User
	if err := r.Scan(
		&s.ID,
		&s.Username,
		&s.UsernameCanonical,
		&s.Email,
		&s.EmailCanonical,
		&s.Enabled,
		&s.Salt,
		&s.Password,
		&s.LastLogin,
		&s.ConfirmationToken,
		&s.PasswordRequestedAt,
		&s.Roles,
	); err != nil {
		return User{}, err
	}
	return s, nil
}

// ScanUsers db representatition
func ScanUsers(rs *sql.Rows) ([]User, error) {
	structs := make([]User, 0, 16)
	var err error
	for rs.Next() {
		var s User
		if err = rs.Scan(
			&s.ID,
			&s.Username,
			&s.UsernameCanonical,
			&s.Email,
			&s.EmailCanonical,
			&s.Enabled,
			&s.Salt,
			&s.Password,
			&s.LastLogin,
			&s.ConfirmationToken,
			&s.PasswordRequestedAt,
			&s.Roles,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}
