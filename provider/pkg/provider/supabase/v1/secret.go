package v1

import (
	"fmt"
	"io"

	"github.com/nitrictech/pulumi-supabase/provider/pkg/api/supabase"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Secret struct{}

// Each resource has in input struct, defining what arguments it accepts.
type SecretArgs struct {
	// ProjectId - Id of the project to create the secret in
	ProjectId string `pulumi:"project_id"`
	// Name - Name of the secret
	Name string `pulumi:"name"`
	// Value - Value of the secret
	Value string `pulumi:"value" provider:"secret"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type SecretState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	SecretArgs
}

// All resources must implement Create at a minumum.
func (Secret) Create(ctx p.Context, name string, input SecretArgs, preview bool) (string, SecretState, error) {
	state := SecretState{SecretArgs: input}
	if preview {
		return name, state, nil
	}

	supabaseClient := infer.GetConfig[config.Config](ctx).Client

	resp, err := supabaseClient.CreateSecrets(ctx, input.ProjectId, []supabase.CreateSecretBody{
		{
			Name:  input.Name,
			Value: input.Value,
		},
	})

	if err != nil {
		return name, state, err
	}

	if resp.StatusCode/100 != 2 {
		// If not a 200 status
		return name, state, fmt.Errorf("received non 200 status: %d", resp.StatusCode)
	}

	return name, state, nil
}

// The Delete method will be run on deletion.
func (Secret) Delete(ctx p.Context, name string, state SecretState) error {
	config := infer.GetConfig[config.Config](ctx)
	supabaseClient := config.Client

	resp, err := supabaseClient.DeleteSecrets(ctx, state.ProjectId, []string{state.Name})
	if err != nil {
		return err
	}

	if resp.StatusCode/100 != 2 {
		respBody, _ := io.ReadAll(resp.Body)
		// If not a 200 status
		return fmt.Errorf("received non 200 status: %d: %s", resp.StatusCode, string(respBody))
	}

	return nil
}
