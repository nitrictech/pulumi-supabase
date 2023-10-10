package projects

import p "github.com/pulumi/pulumi-go-provider"

type ApplyMigrationsInput struct {
}

type ApplyMigrationsResult struct {
}

type ApplyMigrations struct {
}

func (ApplyMigrations) Call(ctx p.Context, input ApplyMigrationsInput) (output ApplyMigrationsResult, err error) {
	return ApplyMigrationsResult{}, nil
}
