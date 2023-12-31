// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package functions

import (
	"context"
	"reflect"

	"errors"
	"github.com/nitrictech/pulumi-supabase/sdk/v3/go/supabase/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DbFunction struct {
	pulumi.CustomResourceState

	Args             pulumi.StringArrayOutput `pulumi:"args"`
	Behaviour        pulumi.StringOutput      `pulumi:"behaviour"`
	Config_params    pulumi.MapOutput         `pulumi:"config_params"`
	Definition       pulumi.StringOutput      `pulumi:"definition"`
	Function_id      pulumi.IntOutput         `pulumi:"function_id"`
	Function_name    pulumi.StringOutput      `pulumi:"function_name"`
	Language         pulumi.StringOutput      `pulumi:"language"`
	Name             pulumi.StringPtrOutput   `pulumi:"name"`
	Project_ref      pulumi.StringOutput      `pulumi:"project_ref"`
	Return_type      pulumi.StringOutput      `pulumi:"return_type"`
	Schema           pulumi.StringOutput      `pulumi:"schema"`
	Security_definer pulumi.BoolOutput        `pulumi:"security_definer"`
}

// NewDbFunction registers a new resource with the given unique name, arguments, and options.
func NewDbFunction(ctx *pulumi.Context,
	name string, args *DbFunctionArgs, opts ...pulumi.ResourceOption) (*DbFunction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Behaviour == nil {
		return nil, errors.New("invalid value for required argument 'Behaviour'")
	}
	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.Language == nil {
		return nil, errors.New("invalid value for required argument 'Language'")
	}
	if args.Project_ref == nil {
		return nil, errors.New("invalid value for required argument 'Project_ref'")
	}
	if args.Return_type == nil {
		return nil, errors.New("invalid value for required argument 'Return_type'")
	}
	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	if args.Security_definer == nil {
		return nil, errors.New("invalid value for required argument 'Security_definer'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DbFunction
	err := ctx.RegisterResource("supabase:functions:DbFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbFunction gets an existing DbFunction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbFunctionState, opts ...pulumi.ResourceOption) (*DbFunction, error) {
	var resource DbFunction
	err := ctx.ReadResource("supabase:functions:DbFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbFunction resources.
type dbFunctionState struct {
}

type DbFunctionState struct {
}

func (DbFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbFunctionState)(nil)).Elem()
}

type dbFunctionArgs struct {
	Args             []string               `pulumi:"args"`
	Behaviour        string                 `pulumi:"behaviour"`
	Config_params    map[string]interface{} `pulumi:"config_params"`
	Definition       string                 `pulumi:"definition"`
	Language         string                 `pulumi:"language"`
	Name             *string                `pulumi:"name"`
	Project_ref      string                 `pulumi:"project_ref"`
	Return_type      string                 `pulumi:"return_type"`
	Schema           string                 `pulumi:"schema"`
	Security_definer bool                   `pulumi:"security_definer"`
}

// The set of arguments for constructing a DbFunction resource.
type DbFunctionArgs struct {
	Args             pulumi.StringArrayInput
	Behaviour        pulumi.StringInput
	Config_params    pulumi.MapInput
	Definition       pulumi.StringInput
	Language         pulumi.StringInput
	Name             pulumi.StringPtrInput
	Project_ref      pulumi.StringInput
	Return_type      pulumi.StringInput
	Schema           pulumi.StringInput
	Security_definer pulumi.BoolInput
}

func (DbFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbFunctionArgs)(nil)).Elem()
}

type DbFunctionInput interface {
	pulumi.Input

	ToDbFunctionOutput() DbFunctionOutput
	ToDbFunctionOutputWithContext(ctx context.Context) DbFunctionOutput
}

func (*DbFunction) ElementType() reflect.Type {
	return reflect.TypeOf((**DbFunction)(nil)).Elem()
}

func (i *DbFunction) ToDbFunctionOutput() DbFunctionOutput {
	return i.ToDbFunctionOutputWithContext(context.Background())
}

func (i *DbFunction) ToDbFunctionOutputWithContext(ctx context.Context) DbFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbFunctionOutput)
}

// DbFunctionArrayInput is an input type that accepts DbFunctionArray and DbFunctionArrayOutput values.
// You can construct a concrete instance of `DbFunctionArrayInput` via:
//
//	DbFunctionArray{ DbFunctionArgs{...} }
type DbFunctionArrayInput interface {
	pulumi.Input

	ToDbFunctionArrayOutput() DbFunctionArrayOutput
	ToDbFunctionArrayOutputWithContext(context.Context) DbFunctionArrayOutput
}

type DbFunctionArray []DbFunctionInput

func (DbFunctionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbFunction)(nil)).Elem()
}

func (i DbFunctionArray) ToDbFunctionArrayOutput() DbFunctionArrayOutput {
	return i.ToDbFunctionArrayOutputWithContext(context.Background())
}

func (i DbFunctionArray) ToDbFunctionArrayOutputWithContext(ctx context.Context) DbFunctionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbFunctionArrayOutput)
}

// DbFunctionMapInput is an input type that accepts DbFunctionMap and DbFunctionMapOutput values.
// You can construct a concrete instance of `DbFunctionMapInput` via:
//
//	DbFunctionMap{ "key": DbFunctionArgs{...} }
type DbFunctionMapInput interface {
	pulumi.Input

	ToDbFunctionMapOutput() DbFunctionMapOutput
	ToDbFunctionMapOutputWithContext(context.Context) DbFunctionMapOutput
}

type DbFunctionMap map[string]DbFunctionInput

func (DbFunctionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbFunction)(nil)).Elem()
}

func (i DbFunctionMap) ToDbFunctionMapOutput() DbFunctionMapOutput {
	return i.ToDbFunctionMapOutputWithContext(context.Background())
}

func (i DbFunctionMap) ToDbFunctionMapOutputWithContext(ctx context.Context) DbFunctionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbFunctionMapOutput)
}

type DbFunctionOutput struct{ *pulumi.OutputState }

func (DbFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbFunction)(nil)).Elem()
}

func (o DbFunctionOutput) ToDbFunctionOutput() DbFunctionOutput {
	return o
}

func (o DbFunctionOutput) ToDbFunctionOutputWithContext(ctx context.Context) DbFunctionOutput {
	return o
}

func (o DbFunctionOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.StringArrayOutput { return v.Args }).(pulumi.StringArrayOutput)
}

func (o DbFunctionOutput) Behaviour() pulumi.StringOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.StringOutput { return v.Behaviour }).(pulumi.StringOutput)
}

func (o DbFunctionOutput) Config_params() pulumi.MapOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.MapOutput { return v.Config_params }).(pulumi.MapOutput)
}

func (o DbFunctionOutput) Definition() pulumi.StringOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.StringOutput { return v.Definition }).(pulumi.StringOutput)
}

func (o DbFunctionOutput) Function_id() pulumi.IntOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.IntOutput { return v.Function_id }).(pulumi.IntOutput)
}

func (o DbFunctionOutput) Function_name() pulumi.StringOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.StringOutput { return v.Function_name }).(pulumi.StringOutput)
}

func (o DbFunctionOutput) Language() pulumi.StringOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.StringOutput { return v.Language }).(pulumi.StringOutput)
}

func (o DbFunctionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DbFunctionOutput) Project_ref() pulumi.StringOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.StringOutput { return v.Project_ref }).(pulumi.StringOutput)
}

func (o DbFunctionOutput) Return_type() pulumi.StringOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.StringOutput { return v.Return_type }).(pulumi.StringOutput)
}

func (o DbFunctionOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.StringOutput { return v.Schema }).(pulumi.StringOutput)
}

func (o DbFunctionOutput) Security_definer() pulumi.BoolOutput {
	return o.ApplyT(func(v *DbFunction) pulumi.BoolOutput { return v.Security_definer }).(pulumi.BoolOutput)
}

type DbFunctionArrayOutput struct{ *pulumi.OutputState }

func (DbFunctionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbFunction)(nil)).Elem()
}

func (o DbFunctionArrayOutput) ToDbFunctionArrayOutput() DbFunctionArrayOutput {
	return o
}

func (o DbFunctionArrayOutput) ToDbFunctionArrayOutputWithContext(ctx context.Context) DbFunctionArrayOutput {
	return o
}

func (o DbFunctionArrayOutput) Index(i pulumi.IntInput) DbFunctionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DbFunction {
		return vs[0].([]*DbFunction)[vs[1].(int)]
	}).(DbFunctionOutput)
}

type DbFunctionMapOutput struct{ *pulumi.OutputState }

func (DbFunctionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbFunction)(nil)).Elem()
}

func (o DbFunctionMapOutput) ToDbFunctionMapOutput() DbFunctionMapOutput {
	return o
}

func (o DbFunctionMapOutput) ToDbFunctionMapOutputWithContext(ctx context.Context) DbFunctionMapOutput {
	return o
}

func (o DbFunctionMapOutput) MapIndex(k pulumi.StringInput) DbFunctionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DbFunction {
		return vs[0].(map[string]*DbFunction)[vs[1].(string)]
	}).(DbFunctionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbFunctionInput)(nil)).Elem(), &DbFunction{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbFunctionArrayInput)(nil)).Elem(), DbFunctionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbFunctionMapInput)(nil)).Elem(), DbFunctionMap{})
	pulumi.RegisterOutputType(DbFunctionOutput{})
	pulumi.RegisterOutputType(DbFunctionArrayOutput{})
	pulumi.RegisterOutputType(DbFunctionMapOutput{})
}
