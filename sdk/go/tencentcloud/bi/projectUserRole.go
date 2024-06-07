// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bi

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a bi projectUserRole
//
// > **NOTE:** You cannot use `Bi.UserRole` and `Bi.ProjectUserRole` at the same time to modify the `phoneNumber` and `email` of the same user.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := Bi.NewProjectUserRole(ctx, "projectUserRole", &Bi.ProjectUserRoleArgs{
//				AreaCode:    pulumi.String("+86"),
//				Email:       pulumi.String("123456@qq.com"),
//				PhoneNumber: pulumi.String("13130001000"),
//				ProjectId:   pulumi.Int(11015030),
//				RoleIdLists: pulumi.IntArray{
//					pulumi.Int(10629453),
//				},
//				UserId:   pulumi.String("100024664626"),
//				UserName: pulumi.String("keep-cam-user"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// bi project_user_role can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Bi/projectUserRole:ProjectUserRole project_user_role projectId#userId
// ```
type ProjectUserRole struct {
	pulumi.CustomResourceState

	// Mobile area code(Note: This field may return null, indicating that no valid value can be obtained).
	AreaCode pulumi.StringOutput `pulumi:"areaCode"`
	// E-mail(Note: This field may return null, indicating that no valid value can be obtained).
	Email pulumi.StringOutput `pulumi:"email"`
	// Phone number(Note: This field may return null, indicating that no valid value can be obtained).
	PhoneNumber pulumi.StringOutput `pulumi:"phoneNumber"`
	// Project id.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// Role id list.
	RoleIdLists pulumi.IntArrayOutput `pulumi:"roleIdLists"`
	// User id.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// Username.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewProjectUserRole registers a new resource with the given unique name, arguments, and options.
func NewProjectUserRole(ctx *pulumi.Context,
	name string, args *ProjectUserRoleArgs, opts ...pulumi.ResourceOption) (*ProjectUserRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AreaCode == nil {
		return nil, errors.New("invalid value for required argument 'AreaCode'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.PhoneNumber == nil {
		return nil, errors.New("invalid value for required argument 'PhoneNumber'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectUserRole
	err := ctx.RegisterResource("tencentcloud:Bi/projectUserRole:ProjectUserRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectUserRole gets an existing ProjectUserRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectUserRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectUserRoleState, opts ...pulumi.ResourceOption) (*ProjectUserRole, error) {
	var resource ProjectUserRole
	err := ctx.ReadResource("tencentcloud:Bi/projectUserRole:ProjectUserRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectUserRole resources.
type projectUserRoleState struct {
	// Mobile area code(Note: This field may return null, indicating that no valid value can be obtained).
	AreaCode *string `pulumi:"areaCode"`
	// E-mail(Note: This field may return null, indicating that no valid value can be obtained).
	Email *string `pulumi:"email"`
	// Phone number(Note: This field may return null, indicating that no valid value can be obtained).
	PhoneNumber *string `pulumi:"phoneNumber"`
	// Project id.
	ProjectId *int `pulumi:"projectId"`
	// Role id list.
	RoleIdLists []int `pulumi:"roleIdLists"`
	// User id.
	UserId *string `pulumi:"userId"`
	// Username.
	UserName *string `pulumi:"userName"`
}

type ProjectUserRoleState struct {
	// Mobile area code(Note: This field may return null, indicating that no valid value can be obtained).
	AreaCode pulumi.StringPtrInput
	// E-mail(Note: This field may return null, indicating that no valid value can be obtained).
	Email pulumi.StringPtrInput
	// Phone number(Note: This field may return null, indicating that no valid value can be obtained).
	PhoneNumber pulumi.StringPtrInput
	// Project id.
	ProjectId pulumi.IntPtrInput
	// Role id list.
	RoleIdLists pulumi.IntArrayInput
	// User id.
	UserId pulumi.StringPtrInput
	// Username.
	UserName pulumi.StringPtrInput
}

func (ProjectUserRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectUserRoleState)(nil)).Elem()
}

type projectUserRoleArgs struct {
	// Mobile area code(Note: This field may return null, indicating that no valid value can be obtained).
	AreaCode string `pulumi:"areaCode"`
	// E-mail(Note: This field may return null, indicating that no valid value can be obtained).
	Email string `pulumi:"email"`
	// Phone number(Note: This field may return null, indicating that no valid value can be obtained).
	PhoneNumber string `pulumi:"phoneNumber"`
	// Project id.
	ProjectId *int `pulumi:"projectId"`
	// Role id list.
	RoleIdLists []int `pulumi:"roleIdLists"`
	// User id.
	UserId string `pulumi:"userId"`
	// Username.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a ProjectUserRole resource.
type ProjectUserRoleArgs struct {
	// Mobile area code(Note: This field may return null, indicating that no valid value can be obtained).
	AreaCode pulumi.StringInput
	// E-mail(Note: This field may return null, indicating that no valid value can be obtained).
	Email pulumi.StringInput
	// Phone number(Note: This field may return null, indicating that no valid value can be obtained).
	PhoneNumber pulumi.StringInput
	// Project id.
	ProjectId pulumi.IntPtrInput
	// Role id list.
	RoleIdLists pulumi.IntArrayInput
	// User id.
	UserId pulumi.StringInput
	// Username.
	UserName pulumi.StringInput
}

func (ProjectUserRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectUserRoleArgs)(nil)).Elem()
}

type ProjectUserRoleInput interface {
	pulumi.Input

	ToProjectUserRoleOutput() ProjectUserRoleOutput
	ToProjectUserRoleOutputWithContext(ctx context.Context) ProjectUserRoleOutput
}

func (*ProjectUserRole) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectUserRole)(nil)).Elem()
}

func (i *ProjectUserRole) ToProjectUserRoleOutput() ProjectUserRoleOutput {
	return i.ToProjectUserRoleOutputWithContext(context.Background())
}

func (i *ProjectUserRole) ToProjectUserRoleOutputWithContext(ctx context.Context) ProjectUserRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectUserRoleOutput)
}

// ProjectUserRoleArrayInput is an input type that accepts ProjectUserRoleArray and ProjectUserRoleArrayOutput values.
// You can construct a concrete instance of `ProjectUserRoleArrayInput` via:
//
//	ProjectUserRoleArray{ ProjectUserRoleArgs{...} }
type ProjectUserRoleArrayInput interface {
	pulumi.Input

	ToProjectUserRoleArrayOutput() ProjectUserRoleArrayOutput
	ToProjectUserRoleArrayOutputWithContext(context.Context) ProjectUserRoleArrayOutput
}

type ProjectUserRoleArray []ProjectUserRoleInput

func (ProjectUserRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectUserRole)(nil)).Elem()
}

func (i ProjectUserRoleArray) ToProjectUserRoleArrayOutput() ProjectUserRoleArrayOutput {
	return i.ToProjectUserRoleArrayOutputWithContext(context.Background())
}

func (i ProjectUserRoleArray) ToProjectUserRoleArrayOutputWithContext(ctx context.Context) ProjectUserRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectUserRoleArrayOutput)
}

// ProjectUserRoleMapInput is an input type that accepts ProjectUserRoleMap and ProjectUserRoleMapOutput values.
// You can construct a concrete instance of `ProjectUserRoleMapInput` via:
//
//	ProjectUserRoleMap{ "key": ProjectUserRoleArgs{...} }
type ProjectUserRoleMapInput interface {
	pulumi.Input

	ToProjectUserRoleMapOutput() ProjectUserRoleMapOutput
	ToProjectUserRoleMapOutputWithContext(context.Context) ProjectUserRoleMapOutput
}

type ProjectUserRoleMap map[string]ProjectUserRoleInput

func (ProjectUserRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectUserRole)(nil)).Elem()
}

func (i ProjectUserRoleMap) ToProjectUserRoleMapOutput() ProjectUserRoleMapOutput {
	return i.ToProjectUserRoleMapOutputWithContext(context.Background())
}

func (i ProjectUserRoleMap) ToProjectUserRoleMapOutputWithContext(ctx context.Context) ProjectUserRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectUserRoleMapOutput)
}

type ProjectUserRoleOutput struct{ *pulumi.OutputState }

func (ProjectUserRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectUserRole)(nil)).Elem()
}

func (o ProjectUserRoleOutput) ToProjectUserRoleOutput() ProjectUserRoleOutput {
	return o
}

func (o ProjectUserRoleOutput) ToProjectUserRoleOutputWithContext(ctx context.Context) ProjectUserRoleOutput {
	return o
}

// Mobile area code(Note: This field may return null, indicating that no valid value can be obtained).
func (o ProjectUserRoleOutput) AreaCode() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectUserRole) pulumi.StringOutput { return v.AreaCode }).(pulumi.StringOutput)
}

// E-mail(Note: This field may return null, indicating that no valid value can be obtained).
func (o ProjectUserRoleOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectUserRole) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Phone number(Note: This field may return null, indicating that no valid value can be obtained).
func (o ProjectUserRoleOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectUserRole) pulumi.StringOutput { return v.PhoneNumber }).(pulumi.StringOutput)
}

// Project id.
func (o ProjectUserRoleOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectUserRole) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// Role id list.
func (o ProjectUserRoleOutput) RoleIdLists() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *ProjectUserRole) pulumi.IntArrayOutput { return v.RoleIdLists }).(pulumi.IntArrayOutput)
}

// User id.
func (o ProjectUserRoleOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectUserRole) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// Username.
func (o ProjectUserRoleOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectUserRole) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type ProjectUserRoleArrayOutput struct{ *pulumi.OutputState }

func (ProjectUserRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectUserRole)(nil)).Elem()
}

func (o ProjectUserRoleArrayOutput) ToProjectUserRoleArrayOutput() ProjectUserRoleArrayOutput {
	return o
}

func (o ProjectUserRoleArrayOutput) ToProjectUserRoleArrayOutputWithContext(ctx context.Context) ProjectUserRoleArrayOutput {
	return o
}

func (o ProjectUserRoleArrayOutput) Index(i pulumi.IntInput) ProjectUserRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectUserRole {
		return vs[0].([]*ProjectUserRole)[vs[1].(int)]
	}).(ProjectUserRoleOutput)
}

type ProjectUserRoleMapOutput struct{ *pulumi.OutputState }

func (ProjectUserRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectUserRole)(nil)).Elem()
}

func (o ProjectUserRoleMapOutput) ToProjectUserRoleMapOutput() ProjectUserRoleMapOutput {
	return o
}

func (o ProjectUserRoleMapOutput) ToProjectUserRoleMapOutputWithContext(ctx context.Context) ProjectUserRoleMapOutput {
	return o
}

func (o ProjectUserRoleMapOutput) MapIndex(k pulumi.StringInput) ProjectUserRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectUserRole {
		return vs[0].(map[string]*ProjectUserRole)[vs[1].(string)]
	}).(ProjectUserRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectUserRoleInput)(nil)).Elem(), &ProjectUserRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectUserRoleArrayInput)(nil)).Elem(), ProjectUserRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectUserRoleMapInput)(nil)).Elem(), ProjectUserRoleMap{})
	pulumi.RegisterOutputType(ProjectUserRoleOutput{})
	pulumi.RegisterOutputType(ProjectUserRoleArrayOutput{})
	pulumi.RegisterOutputType(ProjectUserRoleMapOutput{})
}
