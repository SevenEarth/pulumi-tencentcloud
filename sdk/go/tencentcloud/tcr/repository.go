// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create tcr repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tcr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcr"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		test, err := Tcr.GetInstances(ctx, &tcr.GetInstancesArgs{
// 			Name: pulumi.StringRef("test"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Tcr.NewRepository(ctx, "foo", &Tcr.RepositoryArgs{
// 			InstanceId:    pulumi.String(test.InstanceLists[0].Id),
// 			NamespaceName: pulumi.String("exampleNamespace"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// tcr repository can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Tcr/repository:Repository foo cls-cda1iex1#namespace#repository
// ```
type Repository struct {
	pulumi.CustomResourceState

	// Brief description of the repository. Valid length is [1~100].
	BriefDesc pulumi.StringPtrOutput `pulumi:"briefDesc"`
	// Create time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the repository. Valid length is [1~1000].
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ID of the TCR instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Indicate the repository is public or not.
	IsPublic pulumi.BoolOutput `pulumi:"isPublic"`
	// Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the TCR namespace.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Last updated time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// URL of the repository.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Repository
	err := ctx.RegisterResource("tencentcloud:Tcr/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("tencentcloud:Tcr/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// Brief description of the repository. Valid length is [1~100].
	BriefDesc *string `pulumi:"briefDesc"`
	// Create time.
	CreateTime *string `pulumi:"createTime"`
	// Description of the repository. Valid length is [1~1000].
	Description *string `pulumi:"description"`
	// ID of the TCR instance.
	InstanceId *string `pulumi:"instanceId"`
	// Indicate the repository is public or not.
	IsPublic *bool `pulumi:"isPublic"`
	// Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
	Name *string `pulumi:"name"`
	// Name of the TCR namespace.
	NamespaceName *string `pulumi:"namespaceName"`
	// Last updated time.
	UpdateTime *string `pulumi:"updateTime"`
	// URL of the repository.
	Url *string `pulumi:"url"`
}

type RepositoryState struct {
	// Brief description of the repository. Valid length is [1~100].
	BriefDesc pulumi.StringPtrInput
	// Create time.
	CreateTime pulumi.StringPtrInput
	// Description of the repository. Valid length is [1~1000].
	Description pulumi.StringPtrInput
	// ID of the TCR instance.
	InstanceId pulumi.StringPtrInput
	// Indicate the repository is public or not.
	IsPublic pulumi.BoolPtrInput
	// Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
	Name pulumi.StringPtrInput
	// Name of the TCR namespace.
	NamespaceName pulumi.StringPtrInput
	// Last updated time.
	UpdateTime pulumi.StringPtrInput
	// URL of the repository.
	Url pulumi.StringPtrInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// Brief description of the repository. Valid length is [1~100].
	BriefDesc *string `pulumi:"briefDesc"`
	// Description of the repository. Valid length is [1~1000].
	Description *string `pulumi:"description"`
	// ID of the TCR instance.
	InstanceId string `pulumi:"instanceId"`
	// Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
	Name *string `pulumi:"name"`
	// Name of the TCR namespace.
	NamespaceName string `pulumi:"namespaceName"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// Brief description of the repository. Valid length is [1~100].
	BriefDesc pulumi.StringPtrInput
	// Description of the repository. Valid length is [1~1000].
	Description pulumi.StringPtrInput
	// ID of the TCR instance.
	InstanceId pulumi.StringInput
	// Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
	Name pulumi.StringPtrInput
	// Name of the TCR namespace.
	NamespaceName pulumi.StringInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

// RepositoryArrayInput is an input type that accepts RepositoryArray and RepositoryArrayOutput values.
// You can construct a concrete instance of `RepositoryArrayInput` via:
//
//          RepositoryArray{ RepositoryArgs{...} }
type RepositoryArrayInput interface {
	pulumi.Input

	ToRepositoryArrayOutput() RepositoryArrayOutput
	ToRepositoryArrayOutputWithContext(context.Context) RepositoryArrayOutput
}

type RepositoryArray []RepositoryInput

func (RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (i RepositoryArray) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return i.ToRepositoryArrayOutputWithContext(context.Background())
}

func (i RepositoryArray) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryArrayOutput)
}

// RepositoryMapInput is an input type that accepts RepositoryMap and RepositoryMapOutput values.
// You can construct a concrete instance of `RepositoryMapInput` via:
//
//          RepositoryMap{ "key": RepositoryArgs{...} }
type RepositoryMapInput interface {
	pulumi.Input

	ToRepositoryMapOutput() RepositoryMapOutput
	ToRepositoryMapOutputWithContext(context.Context) RepositoryMapOutput
}

type RepositoryMap map[string]RepositoryInput

func (RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (i RepositoryMap) ToRepositoryMapOutput() RepositoryMapOutput {
	return i.ToRepositoryMapOutputWithContext(context.Background())
}

func (i RepositoryMap) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMapOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

// Brief description of the repository. Valid length is [1~100].
func (o RepositoryOutput) BriefDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.BriefDesc }).(pulumi.StringPtrOutput)
}

// Create time.
func (o RepositoryOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the repository. Valid length is [1~1000].
func (o RepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// ID of the TCR instance.
func (o RepositoryOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Indicate the repository is public or not.
func (o RepositoryOutput) IsPublic() pulumi.BoolOutput {
	return o.ApplyT(func(v *Repository) pulumi.BoolOutput { return v.IsPublic }).(pulumi.BoolOutput)
}

// Name of the TCR repository. Valid length is [2~200]. It can only contain lowercase letters, numbers and separators (`.`, `_`, `-`, `/`), and cannot start, end or continue with separators. Support the use of multi-level address formats, such as `sub1/sub2/repo`.
func (o RepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name of the TCR namespace.
func (o RepositoryOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.NamespaceName }).(pulumi.StringOutput)
}

// Last updated time.
func (o RepositoryOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// URL of the repository.
func (o RepositoryOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type RepositoryArrayOutput struct{ *pulumi.OutputState }

func (RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) Index(i pulumi.IntInput) RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].([]*Repository)[vs[1].(int)]
	}).(RepositoryOutput)
}

type RepositoryMapOutput struct{ *pulumi.OutputState }

func (RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (o RepositoryMapOutput) ToRepositoryMapOutput() RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) MapIndex(k pulumi.StringInput) RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].(map[string]*Repository)[vs[1].(string)]
	}).(RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryArrayInput)(nil)).Elem(), RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryMapInput)(nil)).Elem(), RepositoryMap{})
	pulumi.RegisterOutputType(RepositoryOutput{})
	pulumi.RegisterOutputType(RepositoryArrayOutput{})
	pulumi.RegisterOutputType(RepositoryMapOutput{})
}
