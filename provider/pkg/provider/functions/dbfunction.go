package functions

import (
	"encoding/json"
	"fmt"
	"io"

	supabase "github.com/nitrictech/pulumi-supabase/provider/pkg/api/supabase/v0"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/config"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/projects"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type DbFunction struct{}

// Each resource has in input struct, defining what arguments it accepts.
type DbFunctionArgs struct {
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

	SecurityDefiner bool `pulumi:"security_definer"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type DbFunctionState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	DbFunctionArgs

	FunctionName string `pulumi:"function_name"`

	Id int64 `pulumi:"function_id"`
}

// All resources must implement Create at a minumum.
func (DbFunction) Create(ctx p.Context, name string, input DbFunctionArgs, preview bool) (string, DbFunctionState, error) {
	functionName := name
	if input.Name != nil && *input.Name != "" {
		// TODO: Add random value
		functionName = *input.Name
	}

	state := DbFunctionState{
		DbFunctionArgs: input,
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

	projectDetails, err := projects.GetProjectDetails(ctx, input.ProjectRef)
	if err != nil {
		return name, state, err
	}

	resp, err := supabaseClient.CreateFunction(ctx, input.ProjectRef, &supabase.CreateFunctionParams{
		XConnectionEncrypted: projectDetails.ConnectionString,
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

	function := &supabase.PostgresFunction{}

	if err := json.Unmarshal(respBody, function); err != nil {
		return name, state, err
	}
	state.Id = int64(function.Id)

	return name, state, nil
}

func (DbFunction) Delete(ctx p.Context, name string, state DbFunctionState) error {
	config := infer.GetConfig[config.Config](ctx)
	supabaseClient := config.ExperimentalClient

	projectDetails, err := projects.GetProjectDetails(ctx, state.ProjectRef)
	if err != nil {
		return err
	}

	resp, err := supabaseClient.DeleteFunction(ctx, state.ProjectRef, &supabase.DeleteFunctionParams{
		Id:                   float32(state.Id),
		XConnectionEncrypted: projectDetails.ConnectionString,
	})
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
