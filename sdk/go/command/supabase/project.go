// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package supabase

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Project struct {
	pulumi.CustomResourceState

	Database_host   pulumi.StringOutput    `pulumi:"database_host"`
	Db_pass         pulumi.StringOutput    `pulumi:"db_pass"`
	Kps_enabled     pulumi.BoolPtrOutput   `pulumi:"kps_enabled"`
	Name            pulumi.StringPtrOutput `pulumi:"name"`
	Organization_id pulumi.StringOutput    `pulumi:"organization_id"`
	Plan            pulumi.StringPtrOutput `pulumi:"plan"`
	Project_id      pulumi.StringOutput    `pulumi:"project_id"`
	Project_name    pulumi.StringOutput    `pulumi:"project_name"`
	Region          pulumi.StringOutput    `pulumi:"region"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Db_pass == nil {
		return nil, errors.New("invalid value for required argument 'Db_pass'")
	}
	if args.Organization_id == nil {
		return nil, errors.New("invalid value for required argument 'Organization_id'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Db_pass != nil {
		args.Db_pass = pulumi.ToSecret(args.Db_pass).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"db_pass",
	})
	opts = append(opts, secrets)
	var resource Project
	err := ctx.RegisterResource("supabase:supabase:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("supabase:supabase:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
}

type ProjectState struct {
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	Db_pass         string  `pulumi:"db_pass"`
	Kps_enabled     *bool   `pulumi:"kps_enabled"`
	Name            *string `pulumi:"name"`
	Organization_id string  `pulumi:"organization_id"`
	Plan            *string `pulumi:"plan"`
	Region          string  `pulumi:"region"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	Db_pass         pulumi.StringInput
	Kps_enabled     pulumi.BoolPtrInput
	Name            pulumi.StringPtrInput
	Organization_id pulumi.StringInput
	Plan            pulumi.StringPtrInput
	Region          pulumi.StringInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//	ProjectArray{ ProjectArgs{...} }
type ProjectArrayInput interface {
	pulumi.Input

	ToProjectArrayOutput() ProjectArrayOutput
	ToProjectArrayOutputWithContext(context.Context) ProjectArrayOutput
}

type ProjectArray []ProjectInput

func (ProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (i ProjectArray) ToProjectArrayOutput() ProjectArrayOutput {
	return i.ToProjectArrayOutputWithContext(context.Background())
}

func (i ProjectArray) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectArrayOutput)
}

// ProjectMapInput is an input type that accepts ProjectMap and ProjectMapOutput values.
// You can construct a concrete instance of `ProjectMapInput` via:
//
//	ProjectMap{ "key": ProjectArgs{...} }
type ProjectMapInput interface {
	pulumi.Input

	ToProjectMapOutput() ProjectMapOutput
	ToProjectMapOutputWithContext(context.Context) ProjectMapOutput
}

type ProjectMap map[string]ProjectInput

func (ProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (i ProjectMap) ToProjectMapOutput() ProjectMapOutput {
	return i.ToProjectMapOutputWithContext(context.Background())
}

func (i ProjectMap) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMapOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func (o ProjectOutput) Database_host() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Database_host }).(pulumi.StringOutput)
}

func (o ProjectOutput) Db_pass() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Db_pass }).(pulumi.StringOutput)
}

func (o ProjectOutput) Kps_enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolPtrOutput { return v.Kps_enabled }).(pulumi.BoolPtrOutput)
}

func (o ProjectOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ProjectOutput) Organization_id() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Organization_id }).(pulumi.StringOutput)
}

func (o ProjectOutput) Plan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Plan }).(pulumi.StringPtrOutput)
}

func (o ProjectOutput) Project_id() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Project_id }).(pulumi.StringOutput)
}

func (o ProjectOutput) Project_name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Project_name }).(pulumi.StringOutput)
}

func (o ProjectOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Project {
		return vs[0].([]*Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Project {
		return vs[0].(map[string]*Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
