package supabase

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

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
type Project struct{}

// Each resource has in input struct, defining what arguments it accepts.
type ProjectArgs struct {
	DbPass     string `pulumi:"db_pass" provider:"secret"`
	KpsEnabled bool   `pulumi:"kps_enabled,optional"`

	// Name Name of your project, should not contain dots
	Name string `pulumi:"name,optional"`

	// OrganizationId Slug of your organization
	OrganizationId int64 `pulumi:"organization_id"`

	// Plan Subscription plan
	Plan string `pulumi:"plan,optional"`

	// Region you want your server to reside in
	Region string `pulumi:"region,optional"`

	// Cloud you want your server to reside in ("AWS" currently supported)
	Cloud string `pulumi:"cloud,optional"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type ProjectState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	ProjectArgs

	ProjectId       int64  `pulumi:"project_id"`
	ProjectName     string `pulumi:"project_name"`
	ProjectRef      string `pulumi:"project_ref"`
	ProjectEndpoint string `pulumi:"project_endpoint"`

	DatabaseHost string `pulumi:"database_host"`
}

// All resources must implement Create at a minumum.
func (Project) Create(ctx p.Context, name string, input ProjectArgs, preview bool) (string, ProjectState, error) {
	state := ProjectState{ProjectArgs: input}
	if preview {
		return name, state, nil
	}

	supabaseClient := infer.GetConfig[config.Config](ctx).ExperimentalClient

	projectName := name
	if input.Name != "" {
		// TODO: Add random value
		projectName = input.Name
	}

	plan := input.Plan
	if plan == "" {
		plan = "tier_free"
	}

	cloud := input.Cloud
	if cloud == "" {
		cloud = "AWS"
	}

	region := input.Region
	if region == "" {
		region = "East US (North Virginia)"
	}

	resp, err := supabaseClient.CreateProject(ctx, supabase.CreateProjectBody{
		OrgId:           float32(input.OrganizationId),
		DbPass:          input.DbPass,
		Name:            projectName,
		KpsEnabled:      &input.KpsEnabled,
		CloudProvider:   cloud,
		DbRegion:        region,
		DbPricingTierId: plan,
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

	projectResp := supabase.CreateProjectReply{}

	err = json.Unmarshal(body, &projectResp)
	if err != nil {
		return name, state, err
	}

	state.ProjectId = int64(projectResp.Id)
	state.ProjectName = projectResp.Name
	state.ProjectRef = projectResp.Ref
	state.DatabaseHost = fmt.Sprintf("db.%s.supabase.co", projectResp.Ref)
	state.ProjectEndpoint = projectResp.Endpoint

	_, _, err = lo.AttemptWithDelay(15, 10*time.Second, func(index int, duration time.Duration) error {
		resp, err := supabaseClient.GetProject(ctx, state.ProjectRef)
		if err != nil {
			return err
		}

		projectInfo := supabase.ProjectInfo{}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading response")
		}

		err = json.Unmarshal(body, &projectInfo)
		if err != nil {
			return fmt.Errorf("error unmarshalling response")
		}

		if projectInfo.Status != "ACTIVE_HEALTHY" {
			return fmt.Errorf("project not healthy")
		}

		return nil
	})
	if err != nil {
		return name, state, fmt.Errorf("timed out on project creation, database may still be initialising: %w", err)
	}

	return name, state, nil
}

// The Delete method will be run on deletion.
func (Project) Delete(ctx p.Context, name string, state ProjectState) error {
	config := infer.GetConfig[config.Config](ctx)
	supabaseClient := config.ExperimentalClient

	resp, err := supabaseClient.DeleteProject(ctx, state.ProjectRef)
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
