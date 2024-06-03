// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a bi embedToken
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
//			_, err := Bi.NewEmbedTokenApply(ctx, "embedToken", &Bi.EmbedTokenApplyArgs{
//				ExpireTime: pulumi.String("240"),
//				PageId:     pulumi.Int(10520483),
//				ProjectId:  pulumi.Int(11015030),
//				Scope:      pulumi.String("page"),
//				UserCorpId: pulumi.String("100022975249"),
//				UserId:     pulumi.String("100024664626"),
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
type EmbedTokenApply struct {
	pulumi.CustomResourceState

	// Create the generated token.
	BiToken pulumi.StringOutput `pulumi:"biToken"`
	// Create time.
	CreateAt pulumi.StringOutput `pulumi:"createAt"`
	// Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
	ExpireTime pulumi.StringPtrOutput `pulumi:"expireTime"`
	// Sharing page id, this is empty value 0 when embedding the board.
	PageId pulumi.IntPtrOutput `pulumi:"pageId"`
	// Share project id.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// Page means embedding the page, and panel means embedding the entire board.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
	TicketNum pulumi.IntPtrOutput `pulumi:"ticketNum"`
	// Upadte time.
	UdpateAt pulumi.StringOutput `pulumi:"udpateAt"`
	// User enterprise ID (for multi-user only).
	UserCorpId pulumi.StringPtrOutput `pulumi:"userCorpId"`
	// UserId (for multi-user only).
	UserId pulumi.StringPtrOutput `pulumi:"userId"`
}

// NewEmbedTokenApply registers a new resource with the given unique name, arguments, and options.
func NewEmbedTokenApply(ctx *pulumi.Context,
	name string, args *EmbedTokenApplyArgs, opts ...pulumi.ResourceOption) (*EmbedTokenApply, error) {
	if args == nil {
		args = &EmbedTokenApplyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EmbedTokenApply
	err := ctx.RegisterResource("tencentcloud:Bi/embedTokenApply:EmbedTokenApply", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEmbedTokenApply gets an existing EmbedTokenApply resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEmbedTokenApply(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EmbedTokenApplyState, opts ...pulumi.ResourceOption) (*EmbedTokenApply, error) {
	var resource EmbedTokenApply
	err := ctx.ReadResource("tencentcloud:Bi/embedTokenApply:EmbedTokenApply", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EmbedTokenApply resources.
type embedTokenApplyState struct {
	// Create the generated token.
	BiToken *string `pulumi:"biToken"`
	// Create time.
	CreateAt *string `pulumi:"createAt"`
	// Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
	ExpireTime *string `pulumi:"expireTime"`
	// Sharing page id, this is empty value 0 when embedding the board.
	PageId *int `pulumi:"pageId"`
	// Share project id.
	ProjectId *int `pulumi:"projectId"`
	// Page means embedding the page, and panel means embedding the entire board.
	Scope *string `pulumi:"scope"`
	// Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
	TicketNum *int `pulumi:"ticketNum"`
	// Upadte time.
	UdpateAt *string `pulumi:"udpateAt"`
	// User enterprise ID (for multi-user only).
	UserCorpId *string `pulumi:"userCorpId"`
	// UserId (for multi-user only).
	UserId *string `pulumi:"userId"`
}

type EmbedTokenApplyState struct {
	// Create the generated token.
	BiToken pulumi.StringPtrInput
	// Create time.
	CreateAt pulumi.StringPtrInput
	// Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
	ExpireTime pulumi.StringPtrInput
	// Sharing page id, this is empty value 0 when embedding the board.
	PageId pulumi.IntPtrInput
	// Share project id.
	ProjectId pulumi.IntPtrInput
	// Page means embedding the page, and panel means embedding the entire board.
	Scope pulumi.StringPtrInput
	// Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
	TicketNum pulumi.IntPtrInput
	// Upadte time.
	UdpateAt pulumi.StringPtrInput
	// User enterprise ID (for multi-user only).
	UserCorpId pulumi.StringPtrInput
	// UserId (for multi-user only).
	UserId pulumi.StringPtrInput
}

func (EmbedTokenApplyState) ElementType() reflect.Type {
	return reflect.TypeOf((*embedTokenApplyState)(nil)).Elem()
}

type embedTokenApplyArgs struct {
	// Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
	ExpireTime *string `pulumi:"expireTime"`
	// Sharing page id, this is empty value 0 when embedding the board.
	PageId *int `pulumi:"pageId"`
	// Share project id.
	ProjectId *int `pulumi:"projectId"`
	// Page means embedding the page, and panel means embedding the entire board.
	Scope *string `pulumi:"scope"`
	// Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
	TicketNum *int `pulumi:"ticketNum"`
	// User enterprise ID (for multi-user only).
	UserCorpId *string `pulumi:"userCorpId"`
	// UserId (for multi-user only).
	UserId *string `pulumi:"userId"`
}

// The set of arguments for constructing a EmbedTokenApply resource.
type EmbedTokenApplyArgs struct {
	// Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
	ExpireTime pulumi.StringPtrInput
	// Sharing page id, this is empty value 0 when embedding the board.
	PageId pulumi.IntPtrInput
	// Share project id.
	ProjectId pulumi.IntPtrInput
	// Page means embedding the page, and panel means embedding the entire board.
	Scope pulumi.StringPtrInput
	// Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
	TicketNum pulumi.IntPtrInput
	// User enterprise ID (for multi-user only).
	UserCorpId pulumi.StringPtrInput
	// UserId (for multi-user only).
	UserId pulumi.StringPtrInput
}

func (EmbedTokenApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*embedTokenApplyArgs)(nil)).Elem()
}

type EmbedTokenApplyInput interface {
	pulumi.Input

	ToEmbedTokenApplyOutput() EmbedTokenApplyOutput
	ToEmbedTokenApplyOutputWithContext(ctx context.Context) EmbedTokenApplyOutput
}

func (*EmbedTokenApply) ElementType() reflect.Type {
	return reflect.TypeOf((**EmbedTokenApply)(nil)).Elem()
}

func (i *EmbedTokenApply) ToEmbedTokenApplyOutput() EmbedTokenApplyOutput {
	return i.ToEmbedTokenApplyOutputWithContext(context.Background())
}

func (i *EmbedTokenApply) ToEmbedTokenApplyOutputWithContext(ctx context.Context) EmbedTokenApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmbedTokenApplyOutput)
}

// EmbedTokenApplyArrayInput is an input type that accepts EmbedTokenApplyArray and EmbedTokenApplyArrayOutput values.
// You can construct a concrete instance of `EmbedTokenApplyArrayInput` via:
//
//	EmbedTokenApplyArray{ EmbedTokenApplyArgs{...} }
type EmbedTokenApplyArrayInput interface {
	pulumi.Input

	ToEmbedTokenApplyArrayOutput() EmbedTokenApplyArrayOutput
	ToEmbedTokenApplyArrayOutputWithContext(context.Context) EmbedTokenApplyArrayOutput
}

type EmbedTokenApplyArray []EmbedTokenApplyInput

func (EmbedTokenApplyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmbedTokenApply)(nil)).Elem()
}

func (i EmbedTokenApplyArray) ToEmbedTokenApplyArrayOutput() EmbedTokenApplyArrayOutput {
	return i.ToEmbedTokenApplyArrayOutputWithContext(context.Background())
}

func (i EmbedTokenApplyArray) ToEmbedTokenApplyArrayOutputWithContext(ctx context.Context) EmbedTokenApplyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmbedTokenApplyArrayOutput)
}

// EmbedTokenApplyMapInput is an input type that accepts EmbedTokenApplyMap and EmbedTokenApplyMapOutput values.
// You can construct a concrete instance of `EmbedTokenApplyMapInput` via:
//
//	EmbedTokenApplyMap{ "key": EmbedTokenApplyArgs{...} }
type EmbedTokenApplyMapInput interface {
	pulumi.Input

	ToEmbedTokenApplyMapOutput() EmbedTokenApplyMapOutput
	ToEmbedTokenApplyMapOutputWithContext(context.Context) EmbedTokenApplyMapOutput
}

type EmbedTokenApplyMap map[string]EmbedTokenApplyInput

func (EmbedTokenApplyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmbedTokenApply)(nil)).Elem()
}

func (i EmbedTokenApplyMap) ToEmbedTokenApplyMapOutput() EmbedTokenApplyMapOutput {
	return i.ToEmbedTokenApplyMapOutputWithContext(context.Background())
}

func (i EmbedTokenApplyMap) ToEmbedTokenApplyMapOutputWithContext(ctx context.Context) EmbedTokenApplyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmbedTokenApplyMapOutput)
}

type EmbedTokenApplyOutput struct{ *pulumi.OutputState }

func (EmbedTokenApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmbedTokenApply)(nil)).Elem()
}

func (o EmbedTokenApplyOutput) ToEmbedTokenApplyOutput() EmbedTokenApplyOutput {
	return o
}

func (o EmbedTokenApplyOutput) ToEmbedTokenApplyOutputWithContext(ctx context.Context) EmbedTokenApplyOutput {
	return o
}

// Create the generated token.
func (o EmbedTokenApplyOutput) BiToken() pulumi.StringOutput {
	return o.ApplyT(func(v *EmbedTokenApply) pulumi.StringOutput { return v.BiToken }).(pulumi.StringOutput)
}

// Create time.
func (o EmbedTokenApplyOutput) CreateAt() pulumi.StringOutput {
	return o.ApplyT(func(v *EmbedTokenApply) pulumi.StringOutput { return v.CreateAt }).(pulumi.StringOutput)
}

// Expiration. Unit: Minutes Maximum value: 240. i.e. 4 hours Default: 240.
func (o EmbedTokenApplyOutput) ExpireTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmbedTokenApply) pulumi.StringPtrOutput { return v.ExpireTime }).(pulumi.StringPtrOutput)
}

// Sharing page id, this is empty value 0 when embedding the board.
func (o EmbedTokenApplyOutput) PageId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EmbedTokenApply) pulumi.IntPtrOutput { return v.PageId }).(pulumi.IntPtrOutput)
}

// Share project id.
func (o EmbedTokenApplyOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EmbedTokenApply) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// Page means embedding the page, and panel means embedding the entire board.
func (o EmbedTokenApplyOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmbedTokenApply) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// Access limit, the limit range is 1-99999, if it is empty, no access limit will be set.
func (o EmbedTokenApplyOutput) TicketNum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EmbedTokenApply) pulumi.IntPtrOutput { return v.TicketNum }).(pulumi.IntPtrOutput)
}

// Upadte time.
func (o EmbedTokenApplyOutput) UdpateAt() pulumi.StringOutput {
	return o.ApplyT(func(v *EmbedTokenApply) pulumi.StringOutput { return v.UdpateAt }).(pulumi.StringOutput)
}

// User enterprise ID (for multi-user only).
func (o EmbedTokenApplyOutput) UserCorpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmbedTokenApply) pulumi.StringPtrOutput { return v.UserCorpId }).(pulumi.StringPtrOutput)
}

// UserId (for multi-user only).
func (o EmbedTokenApplyOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmbedTokenApply) pulumi.StringPtrOutput { return v.UserId }).(pulumi.StringPtrOutput)
}

type EmbedTokenApplyArrayOutput struct{ *pulumi.OutputState }

func (EmbedTokenApplyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EmbedTokenApply)(nil)).Elem()
}

func (o EmbedTokenApplyArrayOutput) ToEmbedTokenApplyArrayOutput() EmbedTokenApplyArrayOutput {
	return o
}

func (o EmbedTokenApplyArrayOutput) ToEmbedTokenApplyArrayOutputWithContext(ctx context.Context) EmbedTokenApplyArrayOutput {
	return o
}

func (o EmbedTokenApplyArrayOutput) Index(i pulumi.IntInput) EmbedTokenApplyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EmbedTokenApply {
		return vs[0].([]*EmbedTokenApply)[vs[1].(int)]
	}).(EmbedTokenApplyOutput)
}

type EmbedTokenApplyMapOutput struct{ *pulumi.OutputState }

func (EmbedTokenApplyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EmbedTokenApply)(nil)).Elem()
}

func (o EmbedTokenApplyMapOutput) ToEmbedTokenApplyMapOutput() EmbedTokenApplyMapOutput {
	return o
}

func (o EmbedTokenApplyMapOutput) ToEmbedTokenApplyMapOutputWithContext(ctx context.Context) EmbedTokenApplyMapOutput {
	return o
}

func (o EmbedTokenApplyMapOutput) MapIndex(k pulumi.StringInput) EmbedTokenApplyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EmbedTokenApply {
		return vs[0].(map[string]*EmbedTokenApply)[vs[1].(string)]
	}).(EmbedTokenApplyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EmbedTokenApplyInput)(nil)).Elem(), &EmbedTokenApply{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmbedTokenApplyArrayInput)(nil)).Elem(), EmbedTokenApplyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmbedTokenApplyMapInput)(nil)).Elem(), EmbedTokenApplyMap{})
	pulumi.RegisterOutputType(EmbedTokenApplyOutput{})
	pulumi.RegisterOutputType(EmbedTokenApplyArrayOutput{})
	pulumi.RegisterOutputType(EmbedTokenApplyMapOutput{})
}
