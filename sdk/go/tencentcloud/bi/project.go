// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bi

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a bi project
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Bi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Bi.NewProject(ctx, "project", &Bi.ProjectArgs{
//				ColorCode: pulumi.String("#7BD936"),
//				Logo:      pulumi.String("TF-test"),
//				Mark:      pulumi.String("project mark"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// bi project can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Bi/project:Project project project_id
//
// ```
type Project struct {
	pulumi.CustomResourceState

	// Logo background color.
	ColorCode pulumi.StringOutput `pulumi:"colorCode"`
	// Project logo.
	Logo pulumi.StringPtrOutput `pulumi:"logo"`
	// Remark.
	Mark pulumi.StringPtrOutput `pulumi:"mark"`
	// Project name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ColorCode == nil {
		return nil, errors.New("invalid value for required argument 'ColorCode'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("tencentcloud:Bi/project:Project", name, args, &resource, opts...)
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
	err := ctx.ReadResource("tencentcloud:Bi/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// Logo background color.
	ColorCode *string `pulumi:"colorCode"`
	// Project logo.
	Logo *string `pulumi:"logo"`
	// Remark.
	Mark *string `pulumi:"mark"`
	// Project name.
	Name *string `pulumi:"name"`
}

type ProjectState struct {
	// Logo background color.
	ColorCode pulumi.StringPtrInput
	// Project logo.
	Logo pulumi.StringPtrInput
	// Remark.
	Mark pulumi.StringPtrInput
	// Project name.
	Name pulumi.StringPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Logo background color.
	ColorCode string `pulumi:"colorCode"`
	// Project logo.
	Logo *string `pulumi:"logo"`
	// Remark.
	Mark *string `pulumi:"mark"`
	// Project name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Logo background color.
	ColorCode pulumi.StringInput
	// Project logo.
	Logo pulumi.StringPtrInput
	// Remark.
	Mark pulumi.StringPtrInput
	// Project name.
	Name pulumi.StringPtrInput
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

// Logo background color.
func (o ProjectOutput) ColorCode() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ColorCode }).(pulumi.StringOutput)
}

// Project logo.
func (o ProjectOutput) Logo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Logo }).(pulumi.StringPtrOutput)
}

// Remark.
func (o ProjectOutput) Mark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Mark }).(pulumi.StringPtrOutput)
}

// Project name.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
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
