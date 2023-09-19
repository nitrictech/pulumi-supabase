package provider

import (
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/config"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/supabase"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
)

func NewProvider() p.Provider {
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			Description: "Supabase",
			DisplayName: "The Pulumi Supabase provider enables you to manage supabase platform resources",
			Keywords: []string{
				"pulumi",
				"supabase",
				"kind/native",
			},
			Homepage:          "https://github.com/nitrictech/pulumi-supabase",
			Repository:        "https://github.com/nitrictech/pulumi-supabase",
			Publisher:         "Nitric",
			LogoURL:           "",
			License:           "MIT",
			PluginDownloadURL: "",
			LanguageMap: map[string]any{
				"nodejs": map[string]any{
					"packageName":        "@nitric/pulumi-supabase",
					"packageDescription": "A pulumi provider that manages supabase platform resources",
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
				},
				"go": map[string]any{
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/pulumi/pulumi-command/sdk/go/command",
				},
			},
		},
		Resources: []infer.InferredResource{
			// V1 API Resources
			// Adding into the v0 package for now to save on confusion
			infer.Resource[supabase.Secret, supabase.SecretArgs, supabase.SecretState](),
			// V0 API Resources
			infer.Resource[supabase.Migration, supabase.MigrationArgs, supabase.MigrationState](),
			infer.Resource[supabase.Organization, supabase.OrganizationArgs, supabase.OrganizationState](),
			infer.Resource[supabase.Project, supabase.ProjectArgs, supabase.ProjectState](),
			infer.Resource[supabase.Bucket, supabase.BucketArgs, supabase.BucketState](),
		},
		Config: infer.Config[*config.Config](),
	})
}
