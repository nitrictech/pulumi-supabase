package supabase

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"time"

	_ "github.com/lib/pq"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/api/supabase"
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
	OrganizationId string `pulumi:"organization_id"`

	// Plan Subscription plan
	Plan string `pulumi:"plan,optional"`

	// Region Region you want your server to reside in
	Region string `pulumi:"region"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type ProjectState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	ProjectArgs

	ProjectId   string `pulumi:"project_id"`
	ProjectName string `pulumi:"project_name"`

	DatabaseHost string `pulumi:"database_host"`
}

// All resources must implement Create at a minumum.
func (Project) Create(ctx p.Context, name string, input ProjectArgs, preview bool) (string, ProjectState, error) {
	state := ProjectState{ProjectArgs: input}
	if preview {
		return name, state, nil
	}

	supabaseClient := infer.GetConfig[config.Config](ctx).Client

	projectName := name
	if input.Name != "" {
		// TODO: Add random value
		projectName = input.Name
	}

	plan := input.Plan
	if plan == "" {
		plan = "free"
	}

	resp, err := supabaseClient.CreateProject(ctx, supabase.CreateProjectBody{
		OrganizationId: input.OrganizationId,
		DbPass:         input.DbPass,
		Name:           projectName,
		KpsEnabled:     &input.KpsEnabled,
		Plan:           supabase.CreateProjectBodyPlan(input.Plan),
		Region:         supabase.CreateProjectBodyRegion(input.Region),
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

	projectResp := supabase.ProjectResponse{}

	err = json.Unmarshal(body, &projectResp)
	if err != nil {
		return name, state, err
	}

	state.ProjectId = projectResp.Id
	state.ProjectName = projectResp.Name
	state.DatabaseHost = fmt.Sprintf("db.%s.supabase.co", projectResp.Id)

	_, _, err = lo.AttemptWithDelay(15, 10*time.Second, func(index int, duration time.Duration) error {
		postgresqlDbInfo := fmt.Sprintf("host=%s port=5432 user=postgres "+
			"password=%s dbname=postgres",
			state.DatabaseHost, input.DbPass)
		db, err := sql.Open("postgres", postgresqlDbInfo)
		if err != nil {
			return err
		}
		defer db.Close()
		err = db.Ping()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return name, state, fmt.Errorf("timed out on project creation, database may still be initialising: %w", err)
	}

	return name, state, nil
}

// The Update method will be run on every update.
func (c *Project) Update(ctx p.Context, id string, old ProjectState, new ProjectArgs, preview bool) (ProjectState, error) {
	state := ProjectState{ProjectArgs: new}

	// If in preview, don't run the command.
	if preview {
		return state, nil
	}

	// TODO: determine properties that can actually be updated

	return state, nil
}
