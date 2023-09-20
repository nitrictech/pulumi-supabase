package provider

import (
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/config"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/functions"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/organizations"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/projects"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/secrets"
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/storage"
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
			infer.Resource[secrets.Secret, secrets.SecretArgs, secrets.SecretState](),
			// V0 API Resources
			infer.Resource[organizations.Organization, organizations.OrganizationArgs, organizations.OrganizationState](),
			infer.Resource[projects.Project, projects.ProjectArgs, projects.ProjectState](),
			infer.Resource[storage.Bucket, storage.BucketArgs, storage.BucketState](),
			infer.Resource[functions.DbFunction, functions.DbFunctionArgs, functions.DbFunctionState](),
		},
		Config: infer.Config[*config.Config](),
	})
}
