// Code generated by sqlc. DO NOT EDIT.

package seequell

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addUserDetailsStmt, err = db.PrepareContext(ctx, addUserDetails); err != nil {
		return nil, fmt.Errorf("error preparing query AddUserDetails: %w", err)
	}
	if q.assignUserToProjectStmt, err = db.PrepareContext(ctx, assignUserToProject); err != nil {
		return nil, fmt.Errorf("error preparing query AssignUserToProject: %w", err)
	}
	if q.createProjectStmt, err = db.PrepareContext(ctx, createProject); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProject: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getAllProjectsStmt, err = db.PrepareContext(ctx, getAllProjects); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllProjects: %w", err)
	}
	if q.getProjectByIdStmt, err = db.PrepareContext(ctx, getProjectById); err != nil {
		return nil, fmt.Errorf("error preparing query GetProjectById: %w", err)
	}
	if q.getProjectBySlugStmt, err = db.PrepareContext(ctx, getProjectBySlug); err != nil {
		return nil, fmt.Errorf("error preparing query GetProjectBySlug: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserDetailsStmt, err = db.PrepareContext(ctx, getUserDetails); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserDetails: %w", err)
	}
	if q.isEmailInUseStmt, err = db.PrepareContext(ctx, isEmailInUse); err != nil {
		return nil, fmt.Errorf("error preparing query IsEmailInUse: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.projectExistsByIdStmt, err = db.PrepareContext(ctx, projectExistsById); err != nil {
		return nil, fmt.Errorf("error preparing query ProjectExistsById: %w", err)
	}
	if q.projectExistsBySlugStmt, err = db.PrepareContext(ctx, projectExistsBySlug); err != nil {
		return nil, fmt.Errorf("error preparing query ProjectExistsBySlug: %w", err)
	}
	if q.updateProjectStmt, err = db.PrepareContext(ctx, updateProject); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProject: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	if q.userExistsStmt, err = db.PrepareContext(ctx, userExists); err != nil {
		return nil, fmt.Errorf("error preparing query UserExists: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addUserDetailsStmt != nil {
		if cerr := q.addUserDetailsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUserDetailsStmt: %w", cerr)
		}
	}
	if q.assignUserToProjectStmt != nil {
		if cerr := q.assignUserToProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing assignUserToProjectStmt: %w", cerr)
		}
	}
	if q.createProjectStmt != nil {
		if cerr := q.createProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProjectStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getAllProjectsStmt != nil {
		if cerr := q.getAllProjectsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllProjectsStmt: %w", cerr)
		}
	}
	if q.getProjectByIdStmt != nil {
		if cerr := q.getProjectByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProjectByIdStmt: %w", cerr)
		}
	}
	if q.getProjectBySlugStmt != nil {
		if cerr := q.getProjectBySlugStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProjectBySlugStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserDetailsStmt != nil {
		if cerr := q.getUserDetailsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserDetailsStmt: %w", cerr)
		}
	}
	if q.isEmailInUseStmt != nil {
		if cerr := q.isEmailInUseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing isEmailInUseStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.projectExistsByIdStmt != nil {
		if cerr := q.projectExistsByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing projectExistsByIdStmt: %w", cerr)
		}
	}
	if q.projectExistsBySlugStmt != nil {
		if cerr := q.projectExistsBySlugStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing projectExistsBySlugStmt: %w", cerr)
		}
	}
	if q.updateProjectStmt != nil {
		if cerr := q.updateProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProjectStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	if q.userExistsStmt != nil {
		if cerr := q.userExistsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userExistsStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                      DBTX
	tx                      *sql.Tx
	addUserDetailsStmt      *sql.Stmt
	assignUserToProjectStmt *sql.Stmt
	createProjectStmt       *sql.Stmt
	createUserStmt          *sql.Stmt
	deleteUserStmt          *sql.Stmt
	getAllProjectsStmt      *sql.Stmt
	getProjectByIdStmt      *sql.Stmt
	getProjectBySlugStmt    *sql.Stmt
	getUserStmt             *sql.Stmt
	getUserDetailsStmt      *sql.Stmt
	isEmailInUseStmt        *sql.Stmt
	listUsersStmt           *sql.Stmt
	projectExistsByIdStmt   *sql.Stmt
	projectExistsBySlugStmt *sql.Stmt
	updateProjectStmt       *sql.Stmt
	updateUserStmt          *sql.Stmt
	userExistsStmt          *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                      tx,
		tx:                      tx,
		addUserDetailsStmt:      q.addUserDetailsStmt,
		assignUserToProjectStmt: q.assignUserToProjectStmt,
		createProjectStmt:       q.createProjectStmt,
		createUserStmt:          q.createUserStmt,
		deleteUserStmt:          q.deleteUserStmt,
		getAllProjectsStmt:      q.getAllProjectsStmt,
		getProjectByIdStmt:      q.getProjectByIdStmt,
		getProjectBySlugStmt:    q.getProjectBySlugStmt,
		getUserStmt:             q.getUserStmt,
		getUserDetailsStmt:      q.getUserDetailsStmt,
		isEmailInUseStmt:        q.isEmailInUseStmt,
		listUsersStmt:           q.listUsersStmt,
		projectExistsByIdStmt:   q.projectExistsByIdStmt,
		projectExistsBySlugStmt: q.projectExistsBySlugStmt,
		updateProjectStmt:       q.updateProjectStmt,
		updateUserStmt:          q.updateUserStmt,
		userExistsStmt:          q.userExistsStmt,
	}
}
