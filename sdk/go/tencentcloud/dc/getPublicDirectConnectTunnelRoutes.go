// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dc publicDirectConnectTunnelRoutes
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dc.GetPublicDirectConnectTunnelRoutes(ctx, &dc.GetPublicDirectConnectTunnelRoutesArgs{
//				DirectConnectTunnelId: "dcx-4z49tnws",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPublicDirectConnectTunnelRoutes(ctx *pulumi.Context, args *GetPublicDirectConnectTunnelRoutesArgs, opts ...pulumi.InvokeOption) (*GetPublicDirectConnectTunnelRoutesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetPublicDirectConnectTunnelRoutesResult
	err := ctx.Invoke("tencentcloud:Dc/getPublicDirectConnectTunnelRoutes:getPublicDirectConnectTunnelRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicDirectConnectTunnelRoutes.
type GetPublicDirectConnectTunnelRoutesArgs struct {
	// direct connect tunnel id.
	DirectConnectTunnelId string `pulumi:"directConnectTunnelId"`
	// filter condition: route-type: route type, value: BGP/STATIC route-subnet: route cidr, value such as: 192.68.1.0/24.
	Filters []GetPublicDirectConnectTunnelRoutesFilter `pulumi:"filters"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getPublicDirectConnectTunnelRoutes.
type GetPublicDirectConnectTunnelRoutesResult struct {
	DirectConnectTunnelId string                                     `pulumi:"directConnectTunnelId"`
	Filters               []GetPublicDirectConnectTunnelRoutesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Internet tunnel route list.
	Routes []GetPublicDirectConnectTunnelRoutesRoute `pulumi:"routes"`
}

func GetPublicDirectConnectTunnelRoutesOutput(ctx *pulumi.Context, args GetPublicDirectConnectTunnelRoutesOutputArgs, opts ...pulumi.InvokeOption) GetPublicDirectConnectTunnelRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPublicDirectConnectTunnelRoutesResult, error) {
			args := v.(GetPublicDirectConnectTunnelRoutesArgs)
			r, err := GetPublicDirectConnectTunnelRoutes(ctx, &args, opts...)
			var s GetPublicDirectConnectTunnelRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPublicDirectConnectTunnelRoutesResultOutput)
}

// A collection of arguments for invoking getPublicDirectConnectTunnelRoutes.
type GetPublicDirectConnectTunnelRoutesOutputArgs struct {
	// direct connect tunnel id.
	DirectConnectTunnelId pulumi.StringInput `pulumi:"directConnectTunnelId"`
	// filter condition: route-type: route type, value: BGP/STATIC route-subnet: route cidr, value such as: 192.68.1.0/24.
	Filters GetPublicDirectConnectTunnelRoutesFilterArrayInput `pulumi:"filters"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetPublicDirectConnectTunnelRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicDirectConnectTunnelRoutesArgs)(nil)).Elem()
}

// A collection of values returned by getPublicDirectConnectTunnelRoutes.
type GetPublicDirectConnectTunnelRoutesResultOutput struct{ *pulumi.OutputState }

func (GetPublicDirectConnectTunnelRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicDirectConnectTunnelRoutesResult)(nil)).Elem()
}

func (o GetPublicDirectConnectTunnelRoutesResultOutput) ToGetPublicDirectConnectTunnelRoutesResultOutput() GetPublicDirectConnectTunnelRoutesResultOutput {
	return o
}

func (o GetPublicDirectConnectTunnelRoutesResultOutput) ToGetPublicDirectConnectTunnelRoutesResultOutputWithContext(ctx context.Context) GetPublicDirectConnectTunnelRoutesResultOutput {
	return o
}

func (o GetPublicDirectConnectTunnelRoutesResultOutput) DirectConnectTunnelId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicDirectConnectTunnelRoutesResult) string { return v.DirectConnectTunnelId }).(pulumi.StringOutput)
}

func (o GetPublicDirectConnectTunnelRoutesResultOutput) Filters() GetPublicDirectConnectTunnelRoutesFilterArrayOutput {
	return o.ApplyT(func(v GetPublicDirectConnectTunnelRoutesResult) []GetPublicDirectConnectTunnelRoutesFilter {
		return v.Filters
	}).(GetPublicDirectConnectTunnelRoutesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPublicDirectConnectTunnelRoutesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicDirectConnectTunnelRoutesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPublicDirectConnectTunnelRoutesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPublicDirectConnectTunnelRoutesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Internet tunnel route list.
func (o GetPublicDirectConnectTunnelRoutesResultOutput) Routes() GetPublicDirectConnectTunnelRoutesRouteArrayOutput {
	return o.ApplyT(func(v GetPublicDirectConnectTunnelRoutesResult) []GetPublicDirectConnectTunnelRoutesRoute {
		return v.Routes
	}).(GetPublicDirectConnectTunnelRoutesRouteArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPublicDirectConnectTunnelRoutesResultOutput{})
}
