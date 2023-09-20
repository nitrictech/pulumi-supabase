package supabase

import (
	"encoding/json"
	"fmt"
	"io"

	supabase "github.com/nitrictech/pulumi-supabase/provider/pkg/api/supabase/v0"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type PgFunction struct{}

// Each resource has in input struct, defining what arguments it accepts.
type PgFunctionArgs struct {
	// The name of the organization to create, if not provided one will be generated from the resource name
	Name *string `pulumi:"name,optional"`

	ProjectRef string `pulumi:"project_ref"`

	Args []string `pulumi:"args,optional"`

	Behavior supabase.CreateFunctionBodyBehavior `pulumi:"behaviour"`

	ConfigParams *map[string]any `pulumi:"config_params,optional"`

	Definition string `pulumi:"definition"`

	Language string `pulumi:"language"`

	ReturnType string `pulumi:"return_type"`

	Schema string `pulumi:"schema"`

	SecurityDefiner bool `pulumi:"verify_jwt"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type PgFunctionState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	PgFunctionArgs

	FunctionName string `pulumi:"function_name"`
}

// All resources must implement Create at a minumum.
func (PgFunction) Create(ctx p.Context, name string, input PgFunctionArgs, preview bool) (string, PgFunctionState, error) {
	functionName := name
	if input.Name != nil && *input.Name != "" {
		// TODO: Add random value
		functionName = *input.Name
	}

	state := PgFunctionState{
		PgFunctionArgs: input,
		FunctionName:   functionName,
	}
	if preview {
		return name, state, nil
	}

	args := input.Args
	if args == nil {
		// Quirk of API is that it expects an empty array
		args = []string{}
	}

	config := infer.GetConfig[config.Config](ctx)
	supabaseClient := config.ExperimentalClient

	projResp, err := supabaseClient.GetProject(ctx, input.ProjectRef)
	if err != nil {
		return name, state, err
	}

	projBody, err := io.ReadAll(projResp.Body)
	if err != nil {
		return name, state, err
	}

	project := &supabase.ProjectDetailResponse{}
	if err := json.Unmarshal(projBody, project); err != nil {
		return name, state, err
	}

	resp, err := supabaseClient.CreateFunction(ctx, input.ProjectRef, &supabase.CreateFunctionParams{
		XConnectionEncrypted: project.ConnectionString,
	}, supabase.CreateFunctionBody{
		Args:            args,
		Behavior:        supabase.CreateFunctionBodyBehavior(input.Behavior),
		ConfigParams:    input.ConfigParams,
		Definition:      input.Definition,
		Language:        input.Language,
		Name:            functionName,
		ReturnType:      input.ReturnType,
		Schema:          input.Schema,
		SecurityDefiner: input.SecurityDefiner,
	})
	if err != nil {
		return name, state, err
	}

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode/100 != 2 {
		// If not a 200 status
		return name, state, fmt.Errorf("received non 200 status: %d: %s", resp.StatusCode, string(respBody))
	}

	return name, state, nil
}
