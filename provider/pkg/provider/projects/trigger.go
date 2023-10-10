package projects

type Trigger struct{}

type TriggerArgs struct {
	Activation     string   `pulumi:"activation"`
	EnabledMode    string   `pulumi:"enabled_mode"`
	Events         []string `pulumi:"events"`
	FunctionArgs   []string `pulumi:"function_args"`
	FunctionName   string   `pulumi:"function_name"`
	FunctionSchema string   `pulumi:"function_schema"`

	// Name Name of your project, should not contain dots
	Name string `pulumi:"name,optional"`

	Orientation string `pulumi:"orientation"`

	Schema string `pulumi:"schema"`

	Table string `pulumi:"table"`

	ProjectRef int64 `pulumi:"project_ref"`
}
