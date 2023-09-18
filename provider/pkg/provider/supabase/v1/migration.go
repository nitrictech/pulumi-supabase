package v1

import (
	"bytes"
	"fmt"
	"os/exec"

	p "github.com/pulumi/pulumi-go-provider"
)

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type Migration struct{}

// Each resource has in input struct, defining what arguments it accepts.
type MigrationArgs struct {
	// ProjectId - The remote host to apply the migration to
	ProjectId string `pulumi:"project_id,optional"`
	// DbPassword - The database password
	DbPassword string `pulumi:"db_password"`
	// IncludeAll - Include all migrations not found on remote history table.
	IncludeAll bool `pulumi:"include_all,optional"`
	// IncludeRoles - Include custom roles from supabase/roles.sql.
	IncludeRoles bool `pulumi:"include_roles,optional"`
	// IncludeSeed - Include seed data from supabase/seed.sql.
	IncludeSeed bool `pulumi:"include_seed,optional"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type MigrationState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	MigrationArgs
}

// All resources must implement Create at a minumum.
func (Migration) Create(ctx p.Context, name string, input MigrationArgs, preview bool) (string, MigrationState, error) {
	state := MigrationState{MigrationArgs: input}
	if preview {
		return name, state, nil
	}

	if _, err := exec.LookPath("supabase"); err != nil {
		return name, state, fmt.Errorf("supabase CLI not installed, required to push remote migrations")
	}

	// linkCmd := exec.Command(
	// 	"supabase",
	// 	"link",
	// 	"--project-ref",
	// 	input.ProjectId,
	// )

	// err := linkCmd.Run()
	// if err != nil {
	// 	return name, state, fmt.Errorf("failed to link supabase project %s: %w", input.ProjectId, err)
	// }

	extraFlags := []string{
		"db", "push", "-p", input.DbPassword,
	}

	if input.ProjectId != "" {
		extraFlags = append(extraFlags, "--db-url")
		extraFlags = append(extraFlags, fmt.Sprintf("postgresql://postgres:%s@db.%s.supabase.co:5432/postgres", input.DbPassword, input.ProjectId))
	}

	if input.IncludeAll {
		extraFlags = append(extraFlags, "--include-all")
	}

	if input.IncludeRoles {
		extraFlags = append(extraFlags, "--include-roles")
	}

	if input.IncludeSeed {
		extraFlags = append(extraFlags, "--include-seed")
	}

	migrateCmd := exec.Command(
		"supabase",
		extraFlags...,
	)

	var outb, errb bytes.Buffer

	migrateCmd.Stderr = &errb
	migrateCmd.Stdout = &outb

	err := migrateCmd.Run()
	if err != nil {
		return name, state, fmt.Errorf("failed to apply migrations to project %s: %s, %s", input.ProjectId, outb.String(), errb.String())
	}

	return name, state, nil
}
