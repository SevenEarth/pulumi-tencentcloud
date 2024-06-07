// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ckafka

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a Ckafka Acl.
//
// ## Example Usage
//
// ### Ckafka Acl
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ckafka"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ckafka.NewAcl(ctx, "foo", &Ckafka.AclArgs{
//				InstanceId:     pulumi.String("ckafka-f9ife4zz"),
//				ResourceType:   pulumi.String("TOPIC"),
//				ResourceName:   pulumi.String("topic-tf-test"),
//				OperationType:  pulumi.String("WRITE"),
//				PermissionType: pulumi.String("ALLOW"),
//				Host:           pulumi.String("*"),
//				Principal:      pulumi.Any(tencentcloud_ckafka_user.Foo.Account_name),
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
// Ckafka acl can be imported using the instance_id#permission_type#principal#host#operation_type#resource_type#resource_name, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Ckafka/acl:Acl foo ckafka-f9ife4zz#ALLOW#test#*#WRITE#TOPIC#topic-tf-test
// ```
type Acl struct {
	pulumi.CustomResourceState

	// The default is *, which means that any host can access it. Support filling in IP or network segment, and support `;`separation.
	Host pulumi.StringPtrOutput `pulumi:"host"`
	// ID of the ckafka instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
	OperationType pulumi.StringOutput `pulumi:"operationType"`
	// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
	PermissionType pulumi.StringPtrOutput `pulumi:"permissionType"`
	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list. For example: `root` meaning user root can access.
	Principal pulumi.StringPtrOutput `pulumi:"principal"`
	// ACL resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name.
	ResourceName pulumi.StringOutput `pulumi:"resourceName"`
	// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
	ResourceType pulumi.StringPtrOutput `pulumi:"resourceType"`
}

// NewAcl registers a new resource with the given unique name, arguments, and options.
func NewAcl(ctx *pulumi.Context,
	name string, args *AclArgs, opts ...pulumi.ResourceOption) (*Acl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.OperationType == nil {
		return nil, errors.New("invalid value for required argument 'OperationType'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Acl
	err := ctx.RegisterResource("tencentcloud:Ckafka/acl:Acl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAcl gets an existing Acl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclState, opts ...pulumi.ResourceOption) (*Acl, error) {
	var resource Acl
	err := ctx.ReadResource("tencentcloud:Ckafka/acl:Acl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Acl resources.
type aclState struct {
	// The default is *, which means that any host can access it. Support filling in IP or network segment, and support `;`separation.
	Host *string `pulumi:"host"`
	// ID of the ckafka instance.
	InstanceId *string `pulumi:"instanceId"`
	// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
	OperationType *string `pulumi:"operationType"`
	// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
	PermissionType *string `pulumi:"permissionType"`
	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list. For example: `root` meaning user root can access.
	Principal *string `pulumi:"principal"`
	// ACL resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name.
	ResourceName *string `pulumi:"resourceName"`
	// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
	ResourceType *string `pulumi:"resourceType"`
}

type AclState struct {
	// The default is *, which means that any host can access it. Support filling in IP or network segment, and support `;`separation.
	Host pulumi.StringPtrInput
	// ID of the ckafka instance.
	InstanceId pulumi.StringPtrInput
	// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
	OperationType pulumi.StringPtrInput
	// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
	PermissionType pulumi.StringPtrInput
	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list. For example: `root` meaning user root can access.
	Principal pulumi.StringPtrInput
	// ACL resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name.
	ResourceName pulumi.StringPtrInput
	// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
	ResourceType pulumi.StringPtrInput
}

func (AclState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclState)(nil)).Elem()
}

type aclArgs struct {
	// The default is *, which means that any host can access it. Support filling in IP or network segment, and support `;`separation.
	Host *string `pulumi:"host"`
	// ID of the ckafka instance.
	InstanceId string `pulumi:"instanceId"`
	// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
	OperationType string `pulumi:"operationType"`
	// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
	PermissionType *string `pulumi:"permissionType"`
	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list. For example: `root` meaning user root can access.
	Principal *string `pulumi:"principal"`
	// ACL resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name.
	ResourceName string `pulumi:"resourceName"`
	// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
	ResourceType *string `pulumi:"resourceType"`
}

// The set of arguments for constructing a Acl resource.
type AclArgs struct {
	// The default is *, which means that any host can access it. Support filling in IP or network segment, and support `;`separation.
	Host pulumi.StringPtrInput
	// ID of the ckafka instance.
	InstanceId pulumi.StringInput
	// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
	OperationType pulumi.StringInput
	// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
	PermissionType pulumi.StringPtrInput
	// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list. For example: `root` meaning user root can access.
	Principal pulumi.StringPtrInput
	// ACL resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name.
	ResourceName pulumi.StringInput
	// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
	ResourceType pulumi.StringPtrInput
}

func (AclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclArgs)(nil)).Elem()
}

type AclInput interface {
	pulumi.Input

	ToAclOutput() AclOutput
	ToAclOutputWithContext(ctx context.Context) AclOutput
}

func (*Acl) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (i *Acl) ToAclOutput() AclOutput {
	return i.ToAclOutputWithContext(context.Background())
}

func (i *Acl) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclOutput)
}

// AclArrayInput is an input type that accepts AclArray and AclArrayOutput values.
// You can construct a concrete instance of `AclArrayInput` via:
//
//	AclArray{ AclArgs{...} }
type AclArrayInput interface {
	pulumi.Input

	ToAclArrayOutput() AclArrayOutput
	ToAclArrayOutputWithContext(context.Context) AclArrayOutput
}

type AclArray []AclInput

func (AclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acl)(nil)).Elem()
}

func (i AclArray) ToAclArrayOutput() AclArrayOutput {
	return i.ToAclArrayOutputWithContext(context.Background())
}

func (i AclArray) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclArrayOutput)
}

// AclMapInput is an input type that accepts AclMap and AclMapOutput values.
// You can construct a concrete instance of `AclMapInput` via:
//
//	AclMap{ "key": AclArgs{...} }
type AclMapInput interface {
	pulumi.Input

	ToAclMapOutput() AclMapOutput
	ToAclMapOutputWithContext(context.Context) AclMapOutput
}

type AclMap map[string]AclInput

func (AclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acl)(nil)).Elem()
}

func (i AclMap) ToAclMapOutput() AclMapOutput {
	return i.ToAclMapOutputWithContext(context.Background())
}

func (i AclMap) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclMapOutput)
}

type AclOutput struct{ *pulumi.OutputState }

func (AclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (o AclOutput) ToAclOutput() AclOutput {
	return o
}

func (o AclOutput) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return o
}

// The default is *, which means that any host can access it. Support filling in IP or network segment, and support `;`separation.
func (o AclOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringPtrOutput { return v.Host }).(pulumi.StringPtrOutput)
}

// ID of the ckafka instance.
func (o AclOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.
func (o AclOutput) OperationType() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.OperationType }).(pulumi.StringOutput)
}

// ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.
func (o AclOutput) PermissionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringPtrOutput { return v.PermissionType }).(pulumi.StringPtrOutput)
}

// User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list. For example: `root` meaning user root can access.
func (o AclOutput) Principal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringPtrOutput { return v.Principal }).(pulumi.StringPtrOutput)
}

// ACL resource name, which is related to `resourceType`. For example, if `resourceType` is `TOPIC`, this field indicates the topic name; if `resourceType` is `GROUP`, this field indicates the group name.
func (o AclOutput) ResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.ResourceName }).(pulumi.StringOutput)
}

// ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.
func (o AclOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringPtrOutput { return v.ResourceType }).(pulumi.StringPtrOutput)
}

type AclArrayOutput struct{ *pulumi.OutputState }

func (AclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acl)(nil)).Elem()
}

func (o AclArrayOutput) ToAclArrayOutput() AclArrayOutput {
	return o
}

func (o AclArrayOutput) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return o
}

func (o AclArrayOutput) Index(i pulumi.IntInput) AclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Acl {
		return vs[0].([]*Acl)[vs[1].(int)]
	}).(AclOutput)
}

type AclMapOutput struct{ *pulumi.OutputState }

func (AclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acl)(nil)).Elem()
}

func (o AclMapOutput) ToAclMapOutput() AclMapOutput {
	return o
}

func (o AclMapOutput) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return o
}

func (o AclMapOutput) MapIndex(k pulumi.StringInput) AclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Acl {
		return vs[0].(map[string]*Acl)[vs[1].(string)]
	}).(AclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclInput)(nil)).Elem(), &Acl{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclArrayInput)(nil)).Elem(), AclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclMapInput)(nil)).Elem(), AclMap{})
	pulumi.RegisterOutputType(AclOutput{})
	pulumi.RegisterOutputType(AclArrayOutput{})
	pulumi.RegisterOutputType(AclMapOutput{})
}
