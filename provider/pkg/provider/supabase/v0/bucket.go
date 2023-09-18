package v0

import (
	"fmt"
	"io"
	"time"

	supabase "github.com/nitrictech/pulumi-supabase/provider/pkg/api/supabase/v0"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/samber/lo"
)

type Bucket struct{}

// Each resource has in input struct, defining what arguments it accepts.
type BucketArgs struct {
	// The name of the organization to create, if not provided one will be generated from the resource name
	Name string `pulumi:"name,optional"`

	ProjectRef string `pulumi:"project_ref"`

	AllowedMimeTypes []string `pulumi:"allowed_mime_types,optional"`
	FileSizeLimit    int64    `pulumi:"file_size_limit,optional"`
	Public           bool     `pulumi:"public,optional"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type BucketState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	BucketArgs
}

// All resources must implement Create at a minumum.
func (Bucket) Create(ctx p.Context, name string, input BucketArgs, preview bool) (string, BucketState, error) {
	state := BucketState{BucketArgs: input}
	if preview {
		return name, state, nil
	}

	config := infer.GetConfig[config.Config](ctx)
	supabaseClient := config.ExperimentalClient

	state.Name = name
	if input.Name != "" {
		// TODO: Add random value
		state.Name = input.Name
	}

	allowedMimeTypes := input.AllowedMimeTypes
	if len(allowedMimeTypes) < 1 {
		allowedMimeTypes = nil
	}

	_, _, err := lo.AttemptWithDelay(3, 5*time.Second, func(index int, duration time.Duration) error {
		resp, err := supabaseClient.CreateBucket(ctx, input.ProjectRef, supabase.CreateStorageBucketBody{
			Id:               state.Name,
			Public:           input.Public,
			FileSizeLimit:    float32(input.FileSizeLimit),
			AllowedMimeTypes: allowedMimeTypes,
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
	})
	if err != nil {
		return name, state, err
	}

	return name, state, nil
}

func (Bucket) Delete(ctx p.Context, name string, state BucketState) error {
	config := infer.GetConfig[config.Config](ctx)
	supabaseClient := config.ExperimentalClient

	resp, err := supabaseClient.DeleteBucket(ctx, state.ProjectRef, state.Name)
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
