package config

import (
	"fmt"
	"os"

	"github.com/nitrictech/pulumi-supabase/provider/pkg/api/supabase"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

const defaultApiHost = "https://api.supabase.com"

type Config struct {
	Token   string `pulumi:"token"`
	Client  *supabase.Client
	Version string `pulumi:"version"`
}

var _ = (infer.Annotated)((*Config)(nil))

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.Token, "Supbase Personal Access Token for account")

	// TODO: Look up ENV defaults for supabase management API if any
	a.SetDefault(&c.Token, "", "SUPABASE_ACCESS_TOKEN")
}

var _ infer.CustomConfigure = &Config{}

func (c *Config) Configure(ctx p.Context) error {
	var token = c.Token
	if c.Token == "" {
		token = os.Getenv("SUPABASE_ACCESS_TOKEN")
	}

	if token == "" {
		return fmt.Errorf("invalid access token provided")
	}

	client, err := supabase.NewClient(defaultApiHost, supabase.WithAuthToken(token))
	if err != nil {
		return fmt.Errorf("unable to create supabase API client: %w", err)
	}

	c.Client = client
	return nil
}
