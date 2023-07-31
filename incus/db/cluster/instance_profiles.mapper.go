//go:build linux && cgo && !agent

package cluster

// The code below was generated by lxd-generate - DO NOT EDIT!

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cyphar/incus/lxd/db/query"
	"github.com/cyphar/incus/shared/api"
)

var _ = api.ServerEnvironment{}

var instanceProfileObjects = RegisterStmt(`
SELECT instances_profiles.instance_id, instances_profiles.profile_id, instances_profiles.apply_order
  FROM instances_profiles
  ORDER BY instances_profiles.instance_id, instances_profiles.apply_order
`)

var instanceProfileObjectsByProfileID = RegisterStmt(`
SELECT instances_profiles.instance_id, instances_profiles.profile_id, instances_profiles.apply_order
  FROM instances_profiles
  WHERE ( instances_profiles.profile_id = ? )
  ORDER BY instances_profiles.instance_id, instances_profiles.apply_order
`)

var instanceProfileObjectsByInstanceID = RegisterStmt(`
SELECT instances_profiles.instance_id, instances_profiles.profile_id, instances_profiles.apply_order
  FROM instances_profiles
  WHERE ( instances_profiles.instance_id = ? )
  ORDER BY instances_profiles.instance_id, instances_profiles.apply_order
`)

var instanceProfileCreate = RegisterStmt(`
INSERT INTO instances_profiles (instance_id, profile_id, apply_order)
  VALUES (?, ?, ?)
`)

var instanceProfileDeleteByInstanceID = RegisterStmt(`
DELETE FROM instances_profiles WHERE instance_id = ?
`)

// GetProfileInstances returns all available Instances for the Profile.
// generator: instance_profile GetMany
func GetProfileInstances(ctx context.Context, tx *sql.Tx, profileID int) ([]Instance, error) {
	var err error

	// Result slice.
	objects := make([]InstanceProfile, 0)

	sqlStmt, err := Stmt(tx, instanceProfileObjectsByProfileID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get \"instanceProfileObjectsByProfileID\" prepared statement: %w", err)
	}

	args := []any{profileID}

	// Select.
	objects, err = getInstanceProfiles(ctx, sqlStmt, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instances_profiles\" table: %w", err)
	}

	result := make([]Instance, len(objects))
	for i, object := range objects {
		instance, err := GetInstances(ctx, tx, InstanceFilter{ID: &object.InstanceID})
		if err != nil {
			return nil, err
		}

		result[i] = instance[0]
	}

	return result, nil
}

// instanceProfileColumns returns a string of column names to be used with a SELECT statement for the entity.
// Use this function when building statements to retrieve database entries matching the InstanceProfile entity.
func instanceProfileColumns() string {
	return "instances_profiles.instance_id, instances_profiles.profile_id, instances_profiles.apply_order"
}

// getInstanceProfiles can be used to run handwritten sql.Stmts to return a slice of objects.
func getInstanceProfiles(ctx context.Context, stmt *sql.Stmt, args ...any) ([]InstanceProfile, error) {
	objects := make([]InstanceProfile, 0)

	dest := func(scan func(dest ...any) error) error {
		i := InstanceProfile{}
		err := scan(&i.InstanceID, &i.ProfileID, &i.ApplyOrder)
		if err != nil {
			return err
		}

		objects = append(objects, i)

		return nil
	}

	err := query.SelectObjects(ctx, stmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instances_profiles\" table: %w", err)
	}

	return objects, nil
}

// getInstanceProfilesRaw can be used to run handwritten query strings to return a slice of objects.
func getInstanceProfilesRaw(ctx context.Context, tx *sql.Tx, sql string, args ...any) ([]InstanceProfile, error) {
	objects := make([]InstanceProfile, 0)

	dest := func(scan func(dest ...any) error) error {
		i := InstanceProfile{}
		err := scan(&i.InstanceID, &i.ProfileID, &i.ApplyOrder)
		if err != nil {
			return err
		}

		objects = append(objects, i)

		return nil
	}

	err := query.Scan(ctx, tx, sql, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instances_profiles\" table: %w", err)
	}

	return objects, nil
}

// GetInstanceProfiles returns all available Profiles for the Instance.
// generator: instance_profile GetMany
func GetInstanceProfiles(ctx context.Context, tx *sql.Tx, instanceID int) ([]Profile, error) {
	var err error

	// Result slice.
	objects := make([]InstanceProfile, 0)

	sqlStmt, err := Stmt(tx, instanceProfileObjectsByInstanceID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get \"instanceProfileObjectsByInstanceID\" prepared statement: %w", err)
	}

	args := []any{instanceID}

	// Select.
	objects, err = getInstanceProfiles(ctx, sqlStmt, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instances_profiles\" table: %w", err)
	}

	result := make([]Profile, len(objects))
	for i, object := range objects {
		profile, err := GetProfiles(ctx, tx, ProfileFilter{ID: &object.ProfileID})
		if err != nil {
			return nil, err
		}

		result[i] = profile[0]
	}

	return result, nil
}

// CreateInstanceProfiles adds a new instance_profile to the database.
// generator: instance_profile Create
func CreateInstanceProfiles(ctx context.Context, tx *sql.Tx, objects []InstanceProfile) error {
	for _, object := range objects {
		args := make([]any, 3)

		// Populate the statement arguments.
		args[0] = object.InstanceID
		args[1] = object.ProfileID
		args[2] = object.ApplyOrder

		// Prepared statement to use.
		stmt, err := Stmt(tx, instanceProfileCreate)
		if err != nil {
			return fmt.Errorf("Failed to get \"instanceProfileCreate\" prepared statement: %w", err)
		}

		// Execute the statement.
		_, err = stmt.Exec(args...)
		if err != nil {
			return fmt.Errorf("Failed to create \"instances_profiles\" entry: %w", err)
		}

	}

	return nil
}

// DeleteInstanceProfiles deletes the instance_profile matching the given key parameters.
// generator: instance_profile DeleteMany
func DeleteInstanceProfiles(ctx context.Context, tx *sql.Tx, instanceID int) error {
	stmt, err := Stmt(tx, instanceProfileDeleteByInstanceID)
	if err != nil {
		return fmt.Errorf("Failed to get \"instanceProfileDeleteByInstanceID\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(int(instanceID))
	if err != nil {
		return fmt.Errorf("Delete \"instances_profiles\" entry failed: %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	return nil
}
