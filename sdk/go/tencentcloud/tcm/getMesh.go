// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of tcm mesh
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcm"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tcm.GetMesh(ctx, &tcm.GetMeshArgs{
//				MeshClusters: []string{
//					"cls-xxxx",
//				},
//				MeshIds: []string{
//					"mesh-xxxxxx",
//				},
//				MeshNames: []string{
//					"KEEP_MASH",
//				},
//				Tags: []string{
//					"key",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupMesh(ctx *pulumi.Context, args *LookupMeshArgs, opts ...pulumi.InvokeOption) (*LookupMeshResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMeshResult
	err := ctx.Invoke("tencentcloud:Tcm/getMesh:getMesh", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMesh.
type LookupMeshArgs struct {
	// Mesh name.
	MeshClusters []string `pulumi:"meshClusters"`
	// Mesh instance Id.
	MeshIds []string `pulumi:"meshIds"`
	// Display name.
	MeshNames []string `pulumi:"meshNames"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// tag key.
	Tags []string `pulumi:"tags"`
}

// A collection of values returned by getMesh.
type LookupMeshResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id           string   `pulumi:"id"`
	MeshClusters []string `pulumi:"meshClusters"`
	// Mesh instance Id.
	MeshIds []string `pulumi:"meshIds"`
	// The mesh information is queriedNote: This field may return null, indicating that a valid value is not available.
	MeshLists        []GetMeshMeshList `pulumi:"meshLists"`
	MeshNames        []string          `pulumi:"meshNames"`
	ResultOutputFile *string           `pulumi:"resultOutputFile"`
	Tags             []string          `pulumi:"tags"`
}

func LookupMeshOutput(ctx *pulumi.Context, args LookupMeshOutputArgs, opts ...pulumi.InvokeOption) LookupMeshResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMeshResult, error) {
			args := v.(LookupMeshArgs)
			r, err := LookupMesh(ctx, &args, opts...)
			var s LookupMeshResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMeshResultOutput)
}

// A collection of arguments for invoking getMesh.
type LookupMeshOutputArgs struct {
	// Mesh name.
	MeshClusters pulumi.StringArrayInput `pulumi:"meshClusters"`
	// Mesh instance Id.
	MeshIds pulumi.StringArrayInput `pulumi:"meshIds"`
	// Display name.
	MeshNames pulumi.StringArrayInput `pulumi:"meshNames"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// tag key.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
}

func (LookupMeshOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMeshArgs)(nil)).Elem()
}

// A collection of values returned by getMesh.
type LookupMeshResultOutput struct{ *pulumi.OutputState }

func (LookupMeshResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMeshResult)(nil)).Elem()
}

func (o LookupMeshResultOutput) ToLookupMeshResultOutput() LookupMeshResultOutput {
	return o
}

func (o LookupMeshResultOutput) ToLookupMeshResultOutputWithContext(ctx context.Context) LookupMeshResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMeshResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMeshResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMeshResultOutput) MeshClusters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMeshResult) []string { return v.MeshClusters }).(pulumi.StringArrayOutput)
}

// Mesh instance Id.
func (o LookupMeshResultOutput) MeshIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMeshResult) []string { return v.MeshIds }).(pulumi.StringArrayOutput)
}

// The mesh information is queriedNote: This field may return null, indicating that a valid value is not available.
func (o LookupMeshResultOutput) MeshLists() GetMeshMeshListArrayOutput {
	return o.ApplyT(func(v LookupMeshResult) []GetMeshMeshList { return v.MeshLists }).(GetMeshMeshListArrayOutput)
}

func (o LookupMeshResultOutput) MeshNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMeshResult) []string { return v.MeshNames }).(pulumi.StringArrayOutput)
}

func (o LookupMeshResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMeshResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupMeshResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMeshResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMeshResultOutput{})
}
