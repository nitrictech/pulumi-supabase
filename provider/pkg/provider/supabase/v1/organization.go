package v1

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/nitrictech/pulumi-supabase/provider/pkg/api/supabase"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
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
type Organization struct{}

// Each resource has in input struct, defining what arguments it accepts.
type OrganizationArgs struct {
	// The name of the organization to create, if not provided one will be generated from the resource name
	Name string `pulumi:"name,optional"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type OrganizationState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	OrganizationArgs

	OrganizationId   string `pulumi:"organization_id"`
	OrganizationName string `pulumi:"organization_name"`
}

// All resources must implement Create at a minumum.
func (Organization) Create(ctx p.Context, name string, input OrganizationArgs, preview bool) (string, OrganizationState, error) {
	state := OrganizationState{OrganizationArgs: input}
	if preview {
		return name, state, nil
	}

	config := infer.GetConfig[config.Config](ctx)
	supabaseClient := config.Client

	orgName := name
	if input.Name != "" {
		// TODO: Add random value
		orgName = input.Name
	}

	resp, err := supabaseClient.CreateOrganization(ctx, supabase.CreateOrganizationBody{
		Name: orgName,
	})
	if err != nil {
		return name, state, err
	}

	if resp.StatusCode/100 != 2 {
		// If not a 200 status
		return name, state, fmt.Errorf("received non 200 status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return name, state, err
	}

	orgResp := supabase.OrganizationResponse{}

	err = json.Unmarshal(body, &orgResp)
	if err != nil {
		return name, state, err
	}

	state.OrganizationName = orgResp.Name
	state.OrganizationId = orgResp.Id

	return name, state, nil
}
