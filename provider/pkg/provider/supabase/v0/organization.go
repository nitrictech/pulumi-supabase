package supabase

import (
	"encoding/json"
	"fmt"
	"io"

	supabase "github.com/nitrictech/pulumi-supabase/provider/pkg/api/supabase/v0"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/samber/lo"
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
	OrganizationSlug string `pulumi:"organization_slug"`
	OrganizationName string `pulumi:"organization_name"`
}

// All resources must implement Create at a minumum.
func (Organization) Create(ctx p.Context, name string, input OrganizationArgs, preview bool) (string, OrganizationState, error) {
	state := OrganizationState{OrganizationArgs: input}
	if preview {
		return name, state, nil
	}

	config := infer.GetConfig[config.Config](ctx)
	supabaseClient := config.ExperimentalClient

	orgName := name
	if input.Name != "" {
		// TODO: Add random value
		orgName = input.Name
	}

	resp, err := supabaseClient.CreateOrganizationWithTier(ctx, supabase.CreateOrganizationBodyV2{
		Name: orgName,
		Kind: lo.ToPtr("PERSONAL"),
		// TODO: The v0 API doesn't accept billing information
		// So orgs will be created with legacy billing until updated
		// Tier: supabase.TierFree,
	})
	if err != nil {
		return name, state, err
	}

	if resp.StatusCode/100 != 2 {
		respBody, _ := io.ReadAll(resp.Body)
		// If not a 200 status
		return name, state, fmt.Errorf("received non 200 status: %d: %s", resp.StatusCode, string(respBody))
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
	state.OrganizationId = fmt.Sprintf("%2f", orgResp.Id)
	state.OrganizationSlug = orgResp.Slug

	return name, state, nil
}

func (Organization) Delete(ctx p.Context, name string, state OrganizationState) error {
	config := infer.GetConfig[config.Config](ctx)
	supabaseClient := config.ExperimentalClient

	resp, err := supabaseClient.DeleteOrganization(ctx, state.OrganizationSlug)
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
