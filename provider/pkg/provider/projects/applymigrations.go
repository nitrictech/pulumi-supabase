package projects

import (
	"bytes"
	"fmt"
	"os/exec"

	p "github.com/pulumi/pulumi-go-provider"
)

type ApplyMigrationsInput struct {
	DbUrl  string `pulumi:"db_host"`
	DbPass string `pulumi:"db_pass"`

	IncludeAll   bool `pulumi:"include_all,optional"`
	IncludeRoles bool `pulumi:"include_roles,optional"`
	IncludeSeed  bool `pulumi:"include_seed,optional"`
}

type ApplyMigrationsResult struct {
	Success bool `pulumi:"success"`
}

type ApplyMigrations struct {
}

func (ApplyMigrations) Call(ctx p.Context, input ApplyMigrationsInput) (output ApplyMigrationsResult, err error) {
	if _, err := exec.LookPath("supabase"); err != nil {
		return ApplyMigrationsResult{}, fmt.Errorf("supabase CLI not installed, required to push remote migrations")
	}

	extraFlags := []string{
		"db", "push", "--db-url", input.DbUrl, "-p", input.DbPass,
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

	err = migrateCmd.Run()
	if err != nil {
		return ApplyMigrationsResult{}, fmt.Errorf(
			"failed to apply migrations to project %s: %s, %s",
			input.DbUrl, outb.String(), errb.String(),
		)
	}

	return ApplyMigrationsResult{}, nil
}
