// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create a KubernetesClusterEndpoint. This resource allows you to create an empty cluster first without any workers. Only all attached node depends create complete, cluster endpoint will finally be enabled.
//
// > **NOTE:** Recommend using `dependsOn` to make sure endpoint create after node pools or workers does.
//
// ## Import
//
// KubernetesClusterEndpoint instance can be imported by passing cluster id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Kubernetes/clusterEndpoint:ClusterEndpoint test cluster-id
// ```
type ClusterEndpoint struct {
	pulumi.CustomResourceState

	// The certificate used for access.
	CertificationAuthority pulumi.StringOutput `pulumi:"certificationAuthority"`
	// Cluster deploy type of `MANAGED_CLUSTER` or `INDEPENDENT_CLUSTER`.
	ClusterDeployType pulumi.StringOutput `pulumi:"clusterDeployType"`
	// External network address to access.
	ClusterExternalEndpoint pulumi.StringOutput `pulumi:"clusterExternalEndpoint"`
	// Specify cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Open internet access or not.
	ClusterInternet pulumi.BoolPtrOutput `pulumi:"clusterInternet"`
	// Domain name for cluster Kube-apiserver internet access.  Be careful if you modify value of this parameter, the clusterExternalEndpoint value may be changed automatically too.
	ClusterInternetDomain pulumi.StringPtrOutput `pulumi:"clusterInternetDomain"`
	// Specify security group, NOTE: This argument must not be empty if cluster internet enabled.
	ClusterInternetSecurityGroup pulumi.StringPtrOutput `pulumi:"clusterInternetSecurityGroup"`
	// Open intranet access or not.
	ClusterIntranet pulumi.BoolPtrOutput `pulumi:"clusterIntranet"`
	// Domain name for cluster Kube-apiserver intranet access. Be careful if you modify value of this parameter, the pgwEndpoint value may be changed automatically too.
	ClusterIntranetDomain pulumi.StringPtrOutput `pulumi:"clusterIntranetDomain"`
	// Subnet id who can access this independent cluster, this field must and can only set  when `clusterIntranet` is true. `clusterIntranetSubnetId` can not modify once be set.
	ClusterIntranetSubnetId pulumi.StringPtrOutput `pulumi:"clusterIntranetSubnetId"`
	// Domain name for access.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The LB parameter. Only used for public network access.
	ExtensiveParameters pulumi.StringPtrOutput `pulumi:"extensiveParameters"`
	// this argument was deprecated, use `clusterInternetSecurityGroup` instead. Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all. This field can only set when field `clusterDeployType` is 'MANAGED_CLUSTER' and `clusterInternet` is true. `managedClusterInternetSecurityPolicies` can not delete or empty once be set.
	//
	// Deprecated: this argument was deprecated, use `cluster_internet_security_group` instead.
	ManagedClusterInternetSecurityPolicies pulumi.StringArrayOutput `pulumi:"managedClusterInternetSecurityPolicies"`
	// Password of account.
	Password pulumi.StringOutput `pulumi:"password"`
	// The Intranet address used for access.
	PgwEndpoint pulumi.StringOutput `pulumi:"pgwEndpoint"`
	// User name of account.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewClusterEndpoint registers a new resource with the given unique name, arguments, and options.
func NewClusterEndpoint(ctx *pulumi.Context,
	name string, args *ClusterEndpointArgs, opts ...pulumi.ResourceOption) (*ClusterEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ClusterEndpoint
	err := ctx.RegisterResource("tencentcloud:Kubernetes/clusterEndpoint:ClusterEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterEndpoint gets an existing ClusterEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterEndpointState, opts ...pulumi.ResourceOption) (*ClusterEndpoint, error) {
	var resource ClusterEndpoint
	err := ctx.ReadResource("tencentcloud:Kubernetes/clusterEndpoint:ClusterEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterEndpoint resources.
type clusterEndpointState struct {
	// The certificate used for access.
	CertificationAuthority *string `pulumi:"certificationAuthority"`
	// Cluster deploy type of `MANAGED_CLUSTER` or `INDEPENDENT_CLUSTER`.
	ClusterDeployType *string `pulumi:"clusterDeployType"`
	// External network address to access.
	ClusterExternalEndpoint *string `pulumi:"clusterExternalEndpoint"`
	// Specify cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Open internet access or not.
	ClusterInternet *bool `pulumi:"clusterInternet"`
	// Domain name for cluster Kube-apiserver internet access.  Be careful if you modify value of this parameter, the clusterExternalEndpoint value may be changed automatically too.
	ClusterInternetDomain *string `pulumi:"clusterInternetDomain"`
	// Specify security group, NOTE: This argument must not be empty if cluster internet enabled.
	ClusterInternetSecurityGroup *string `pulumi:"clusterInternetSecurityGroup"`
	// Open intranet access or not.
	ClusterIntranet *bool `pulumi:"clusterIntranet"`
	// Domain name for cluster Kube-apiserver intranet access. Be careful if you modify value of this parameter, the pgwEndpoint value may be changed automatically too.
	ClusterIntranetDomain *string `pulumi:"clusterIntranetDomain"`
	// Subnet id who can access this independent cluster, this field must and can only set  when `clusterIntranet` is true. `clusterIntranetSubnetId` can not modify once be set.
	ClusterIntranetSubnetId *string `pulumi:"clusterIntranetSubnetId"`
	// Domain name for access.
	Domain *string `pulumi:"domain"`
	// The LB parameter. Only used for public network access.
	ExtensiveParameters *string `pulumi:"extensiveParameters"`
	// this argument was deprecated, use `clusterInternetSecurityGroup` instead. Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all. This field can only set when field `clusterDeployType` is 'MANAGED_CLUSTER' and `clusterInternet` is true. `managedClusterInternetSecurityPolicies` can not delete or empty once be set.
	//
	// Deprecated: this argument was deprecated, use `cluster_internet_security_group` instead.
	ManagedClusterInternetSecurityPolicies []string `pulumi:"managedClusterInternetSecurityPolicies"`
	// Password of account.
	Password *string `pulumi:"password"`
	// The Intranet address used for access.
	PgwEndpoint *string `pulumi:"pgwEndpoint"`
	// User name of account.
	UserName *string `pulumi:"userName"`
}

type ClusterEndpointState struct {
	// The certificate used for access.
	CertificationAuthority pulumi.StringPtrInput
	// Cluster deploy type of `MANAGED_CLUSTER` or `INDEPENDENT_CLUSTER`.
	ClusterDeployType pulumi.StringPtrInput
	// External network address to access.
	ClusterExternalEndpoint pulumi.StringPtrInput
	// Specify cluster ID.
	ClusterId pulumi.StringPtrInput
	// Open internet access or not.
	ClusterInternet pulumi.BoolPtrInput
	// Domain name for cluster Kube-apiserver internet access.  Be careful if you modify value of this parameter, the clusterExternalEndpoint value may be changed automatically too.
	ClusterInternetDomain pulumi.StringPtrInput
	// Specify security group, NOTE: This argument must not be empty if cluster internet enabled.
	ClusterInternetSecurityGroup pulumi.StringPtrInput
	// Open intranet access or not.
	ClusterIntranet pulumi.BoolPtrInput
	// Domain name for cluster Kube-apiserver intranet access. Be careful if you modify value of this parameter, the pgwEndpoint value may be changed automatically too.
	ClusterIntranetDomain pulumi.StringPtrInput
	// Subnet id who can access this independent cluster, this field must and can only set  when `clusterIntranet` is true. `clusterIntranetSubnetId` can not modify once be set.
	ClusterIntranetSubnetId pulumi.StringPtrInput
	// Domain name for access.
	Domain pulumi.StringPtrInput
	// The LB parameter. Only used for public network access.
	ExtensiveParameters pulumi.StringPtrInput
	// this argument was deprecated, use `clusterInternetSecurityGroup` instead. Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all. This field can only set when field `clusterDeployType` is 'MANAGED_CLUSTER' and `clusterInternet` is true. `managedClusterInternetSecurityPolicies` can not delete or empty once be set.
	//
	// Deprecated: this argument was deprecated, use `cluster_internet_security_group` instead.
	ManagedClusterInternetSecurityPolicies pulumi.StringArrayInput
	// Password of account.
	Password pulumi.StringPtrInput
	// The Intranet address used for access.
	PgwEndpoint pulumi.StringPtrInput
	// User name of account.
	UserName pulumi.StringPtrInput
}

func (ClusterEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterEndpointState)(nil)).Elem()
}

type clusterEndpointArgs struct {
	// Specify cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Open internet access or not.
	ClusterInternet *bool `pulumi:"clusterInternet"`
	// Domain name for cluster Kube-apiserver internet access.  Be careful if you modify value of this parameter, the clusterExternalEndpoint value may be changed automatically too.
	ClusterInternetDomain *string `pulumi:"clusterInternetDomain"`
	// Specify security group, NOTE: This argument must not be empty if cluster internet enabled.
	ClusterInternetSecurityGroup *string `pulumi:"clusterInternetSecurityGroup"`
	// Open intranet access or not.
	ClusterIntranet *bool `pulumi:"clusterIntranet"`
	// Domain name for cluster Kube-apiserver intranet access. Be careful if you modify value of this parameter, the pgwEndpoint value may be changed automatically too.
	ClusterIntranetDomain *string `pulumi:"clusterIntranetDomain"`
	// Subnet id who can access this independent cluster, this field must and can only set  when `clusterIntranet` is true. `clusterIntranetSubnetId` can not modify once be set.
	ClusterIntranetSubnetId *string `pulumi:"clusterIntranetSubnetId"`
	// The LB parameter. Only used for public network access.
	ExtensiveParameters *string `pulumi:"extensiveParameters"`
	// this argument was deprecated, use `clusterInternetSecurityGroup` instead. Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all. This field can only set when field `clusterDeployType` is 'MANAGED_CLUSTER' and `clusterInternet` is true. `managedClusterInternetSecurityPolicies` can not delete or empty once be set.
	//
	// Deprecated: this argument was deprecated, use `cluster_internet_security_group` instead.
	ManagedClusterInternetSecurityPolicies []string `pulumi:"managedClusterInternetSecurityPolicies"`
}

// The set of arguments for constructing a ClusterEndpoint resource.
type ClusterEndpointArgs struct {
	// Specify cluster ID.
	ClusterId pulumi.StringInput
	// Open internet access or not.
	ClusterInternet pulumi.BoolPtrInput
	// Domain name for cluster Kube-apiserver internet access.  Be careful if you modify value of this parameter, the clusterExternalEndpoint value may be changed automatically too.
	ClusterInternetDomain pulumi.StringPtrInput
	// Specify security group, NOTE: This argument must not be empty if cluster internet enabled.
	ClusterInternetSecurityGroup pulumi.StringPtrInput
	// Open intranet access or not.
	ClusterIntranet pulumi.BoolPtrInput
	// Domain name for cluster Kube-apiserver intranet access. Be careful if you modify value of this parameter, the pgwEndpoint value may be changed automatically too.
	ClusterIntranetDomain pulumi.StringPtrInput
	// Subnet id who can access this independent cluster, this field must and can only set  when `clusterIntranet` is true. `clusterIntranetSubnetId` can not modify once be set.
	ClusterIntranetSubnetId pulumi.StringPtrInput
	// The LB parameter. Only used for public network access.
	ExtensiveParameters pulumi.StringPtrInput
	// this argument was deprecated, use `clusterInternetSecurityGroup` instead. Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all. This field can only set when field `clusterDeployType` is 'MANAGED_CLUSTER' and `clusterInternet` is true. `managedClusterInternetSecurityPolicies` can not delete or empty once be set.
	//
	// Deprecated: this argument was deprecated, use `cluster_internet_security_group` instead.
	ManagedClusterInternetSecurityPolicies pulumi.StringArrayInput
}

func (ClusterEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterEndpointArgs)(nil)).Elem()
}

type ClusterEndpointInput interface {
	pulumi.Input

	ToClusterEndpointOutput() ClusterEndpointOutput
	ToClusterEndpointOutputWithContext(ctx context.Context) ClusterEndpointOutput
}

func (*ClusterEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterEndpoint)(nil)).Elem()
}

func (i *ClusterEndpoint) ToClusterEndpointOutput() ClusterEndpointOutput {
	return i.ToClusterEndpointOutputWithContext(context.Background())
}

func (i *ClusterEndpoint) ToClusterEndpointOutputWithContext(ctx context.Context) ClusterEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointOutput)
}

// ClusterEndpointArrayInput is an input type that accepts ClusterEndpointArray and ClusterEndpointArrayOutput values.
// You can construct a concrete instance of `ClusterEndpointArrayInput` via:
//
//          ClusterEndpointArray{ ClusterEndpointArgs{...} }
type ClusterEndpointArrayInput interface {
	pulumi.Input

	ToClusterEndpointArrayOutput() ClusterEndpointArrayOutput
	ToClusterEndpointArrayOutputWithContext(context.Context) ClusterEndpointArrayOutput
}

type ClusterEndpointArray []ClusterEndpointInput

func (ClusterEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterEndpoint)(nil)).Elem()
}

func (i ClusterEndpointArray) ToClusterEndpointArrayOutput() ClusterEndpointArrayOutput {
	return i.ToClusterEndpointArrayOutputWithContext(context.Background())
}

func (i ClusterEndpointArray) ToClusterEndpointArrayOutputWithContext(ctx context.Context) ClusterEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointArrayOutput)
}

// ClusterEndpointMapInput is an input type that accepts ClusterEndpointMap and ClusterEndpointMapOutput values.
// You can construct a concrete instance of `ClusterEndpointMapInput` via:
//
//          ClusterEndpointMap{ "key": ClusterEndpointArgs{...} }
type ClusterEndpointMapInput interface {
	pulumi.Input

	ToClusterEndpointMapOutput() ClusterEndpointMapOutput
	ToClusterEndpointMapOutputWithContext(context.Context) ClusterEndpointMapOutput
}

type ClusterEndpointMap map[string]ClusterEndpointInput

func (ClusterEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterEndpoint)(nil)).Elem()
}

func (i ClusterEndpointMap) ToClusterEndpointMapOutput() ClusterEndpointMapOutput {
	return i.ToClusterEndpointMapOutputWithContext(context.Background())
}

func (i ClusterEndpointMap) ToClusterEndpointMapOutputWithContext(ctx context.Context) ClusterEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointMapOutput)
}

type ClusterEndpointOutput struct{ *pulumi.OutputState }

func (ClusterEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterEndpoint)(nil)).Elem()
}

func (o ClusterEndpointOutput) ToClusterEndpointOutput() ClusterEndpointOutput {
	return o
}

func (o ClusterEndpointOutput) ToClusterEndpointOutputWithContext(ctx context.Context) ClusterEndpointOutput {
	return o
}

// The certificate used for access.
func (o ClusterEndpointOutput) CertificationAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.CertificationAuthority }).(pulumi.StringOutput)
}

// Cluster deploy type of `MANAGED_CLUSTER` or `INDEPENDENT_CLUSTER`.
func (o ClusterEndpointOutput) ClusterDeployType() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.ClusterDeployType }).(pulumi.StringOutput)
}

// External network address to access.
func (o ClusterEndpointOutput) ClusterExternalEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.ClusterExternalEndpoint }).(pulumi.StringOutput)
}

// Specify cluster ID.
func (o ClusterEndpointOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Open internet access or not.
func (o ClusterEndpointOutput) ClusterInternet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.BoolPtrOutput { return v.ClusterInternet }).(pulumi.BoolPtrOutput)
}

// Domain name for cluster Kube-apiserver internet access.  Be careful if you modify value of this parameter, the clusterExternalEndpoint value may be changed automatically too.
func (o ClusterEndpointOutput) ClusterInternetDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringPtrOutput { return v.ClusterInternetDomain }).(pulumi.StringPtrOutput)
}

// Specify security group, NOTE: This argument must not be empty if cluster internet enabled.
func (o ClusterEndpointOutput) ClusterInternetSecurityGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringPtrOutput { return v.ClusterInternetSecurityGroup }).(pulumi.StringPtrOutput)
}

// Open intranet access or not.
func (o ClusterEndpointOutput) ClusterIntranet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.BoolPtrOutput { return v.ClusterIntranet }).(pulumi.BoolPtrOutput)
}

// Domain name for cluster Kube-apiserver intranet access. Be careful if you modify value of this parameter, the pgwEndpoint value may be changed automatically too.
func (o ClusterEndpointOutput) ClusterIntranetDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringPtrOutput { return v.ClusterIntranetDomain }).(pulumi.StringPtrOutput)
}

// Subnet id who can access this independent cluster, this field must and can only set  when `clusterIntranet` is true. `clusterIntranetSubnetId` can not modify once be set.
func (o ClusterEndpointOutput) ClusterIntranetSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringPtrOutput { return v.ClusterIntranetSubnetId }).(pulumi.StringPtrOutput)
}

// Domain name for access.
func (o ClusterEndpointOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The LB parameter. Only used for public network access.
func (o ClusterEndpointOutput) ExtensiveParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringPtrOutput { return v.ExtensiveParameters }).(pulumi.StringPtrOutput)
}

// this argument was deprecated, use `clusterInternetSecurityGroup` instead. Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all. This field can only set when field `clusterDeployType` is 'MANAGED_CLUSTER' and `clusterInternet` is true. `managedClusterInternetSecurityPolicies` can not delete or empty once be set.
//
// Deprecated: this argument was deprecated, use `cluster_internet_security_group` instead.
func (o ClusterEndpointOutput) ManagedClusterInternetSecurityPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringArrayOutput { return v.ManagedClusterInternetSecurityPolicies }).(pulumi.StringArrayOutput)
}

// Password of account.
func (o ClusterEndpointOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The Intranet address used for access.
func (o ClusterEndpointOutput) PgwEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.PgwEndpoint }).(pulumi.StringOutput)
}

// User name of account.
func (o ClusterEndpointOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterEndpoint) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type ClusterEndpointArrayOutput struct{ *pulumi.OutputState }

func (ClusterEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterEndpoint)(nil)).Elem()
}

func (o ClusterEndpointArrayOutput) ToClusterEndpointArrayOutput() ClusterEndpointArrayOutput {
	return o
}

func (o ClusterEndpointArrayOutput) ToClusterEndpointArrayOutputWithContext(ctx context.Context) ClusterEndpointArrayOutput {
	return o
}

func (o ClusterEndpointArrayOutput) Index(i pulumi.IntInput) ClusterEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterEndpoint {
		return vs[0].([]*ClusterEndpoint)[vs[1].(int)]
	}).(ClusterEndpointOutput)
}

type ClusterEndpointMapOutput struct{ *pulumi.OutputState }

func (ClusterEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterEndpoint)(nil)).Elem()
}

func (o ClusterEndpointMapOutput) ToClusterEndpointMapOutput() ClusterEndpointMapOutput {
	return o
}

func (o ClusterEndpointMapOutput) ToClusterEndpointMapOutputWithContext(ctx context.Context) ClusterEndpointMapOutput {
	return o
}

func (o ClusterEndpointMapOutput) MapIndex(k pulumi.StringInput) ClusterEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterEndpoint {
		return vs[0].(map[string]*ClusterEndpoint)[vs[1].(string)]
	}).(ClusterEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEndpointInput)(nil)).Elem(), &ClusterEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEndpointArrayInput)(nil)).Elem(), ClusterEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterEndpointMapInput)(nil)).Elem(), ClusterEndpointMap{})
	pulumi.RegisterOutputType(ClusterEndpointOutput{})
	pulumi.RegisterOutputType(ClusterEndpointArrayOutput{})
	pulumi.RegisterOutputType(ClusterEndpointMapOutput{})
}
