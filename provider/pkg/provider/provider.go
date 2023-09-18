package provider

import (
	"github.com/nitrictech/pulumi-supabase/provider/pkg/provider/config"
	supabaseExperimental "github.com/nitrictech/pulumi-supabase/provider/pkg/provider/supabase/v0"
	supabase "github.com/nitrictech/pulumi-supabase/provider/pkg/provider/supabase/v1"
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
			infer.Resource[supabase.Organization, supabase.OrganizationArgs, supabase.OrganizationState](),
			infer.Resource[supabase.Project, supabase.ProjectArgs, supabase.ProjectState](),
			infer.Resource[supabase.Secret, supabase.SecretArgs, supabase.SecretState](),
			infer.Resource[supabase.Migration, supabase.MigrationArgs, supabase.MigrationState](),
			infer.Resource[supabaseExperimental.Organization, supabaseExperimental.OrganizationArgs, supabaseExperimental.OrganizationState](),
		},
		Config: infer.Config[*config.Config](),
	})
}
