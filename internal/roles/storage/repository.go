package storage

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/entity"
)

type RolesRepository interface {
	AssignLeaderRole(ctx context.Context, userID int) error
	AssignPrimaryRole(ctx context.Context, userID int) error
	AssignPastorRole(ctx context.Context, userID int) error
	AssignMinistryLeaderRole(ctx context.Context, userID int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRolesRepository(db *sqlx.DB) RolesRepository {
	return &repository{
		db: db,
	}
}

// AssignMinistryLeaderRole assigns the Ministry Leader role to the given user.
func (r *repository) AssignMinistryLeaderRole(ctx context.Context, userID int) error {
	roleID, err := r.getRoleID(ctx, "Ministry Leader")
	if err != nil {
		return err
	}

	return r.assignRole(ctx, userID, roleID)
}

// AssignLeaderRole assigns the Leader role to the given user.
func (r *repository) AssignLeaderRole(ctx context.Context, userID int) error {
	roleID, err := r.getRoleID(ctx, "Leader")
	if err != nil {
		return err
	}

	return r.assignRole(ctx, userID, roleID)
}

// AssignPrimaryRole assigns the Primary role to the given user.
func (r *repository) AssignPrimaryRole(ctx context.Context, userID int) error {
	roleID, err := r.getRoleID(ctx, "Primary")
	if err != nil {
		return err
	}

	return r.assignRole(ctx, userID, roleID)
}

// AssignPastorRole assigns the Pastor role to the given user.
func (r *repository) AssignPastorRole(ctx context.Context, userID int) error {
	roleID, err := r.getRoleID(ctx, "Pastor")
	if err != nil {
		return err
	}

	return r.assignRole(ctx, userID, roleID)
}

// getRoleID retrieves the role ID for a given role name using the ORM.
func (r *repository) getRoleID(ctx context.Context, roleName string) (int, error) {
	role, err := entity.Roles(entity.RoleWhere.RoleName.EQ(roleName)).One(ctx, r.db)
	if err != nil {
		return 0, fmt.Errorf("failed to get role id for '%s': %w", roleName, err)
	}

	return role.ID, nil
}

// assignRole inserts a record into the user_roles join table using the ORM.
func (r *repository) assignRole(ctx context.Context, userID, roleID int) error {
	userRole := &entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	if err := userRole.Insert(ctx, r.db, boil.Infer()); err != nil {
		return fmt.Errorf("failed to assign role %d to user %d: %w", roleID, userID, err)
	}

	return nil
}
