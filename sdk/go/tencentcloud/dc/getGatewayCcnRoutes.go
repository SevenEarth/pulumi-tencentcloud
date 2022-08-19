// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of direct connect gateway route entries.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ccn"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		main, err := Ccn.NewInstance(ctx, "main", &Ccn.InstanceArgs{
// 			Description: pulumi.String("ci-temp-test-ccn-des"),
// 			Qos:         pulumi.String("AG"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ccnMain, err := Dc.NewGateway(ctx, "ccnMain", &Dc.GatewayArgs{
// 			NetworkInstanceId: main.ID(),
// 			NetworkType:       pulumi.String("CCN"),
// 			GatewayType:       pulumi.String("NORMAL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Dc.NewGatewayCcnRoute(ctx, "route1", &Dc.GatewayCcnRouteArgs{
// 			DcgId:     ccnMain.ID(),
// 			CidrBlock: pulumi.String("10.1.1.0/32"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Dc.NewGatewayCcnRoute(ctx, "route2", &Dc.GatewayCcnRouteArgs{
// 			DcgId:     ccnMain.ID(),
// 			CidrBlock: pulumi.String("192.1.1.0/32"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_ = Dc.GetGatewayCcnRoutesOutput(ctx, dc.GetGatewayCcnRoutesOutputArgs{
// 			DcgId: ccnMain.ID(),
// 		}, nil)
// 		return nil
// 	})
// }
// ```
func GetGatewayCcnRoutes(ctx *pulumi.Context, args *GetGatewayCcnRoutesArgs, opts ...pulumi.InvokeOption) (*GetGatewayCcnRoutesResult, error) {
	var rv GetGatewayCcnRoutesResult
	err := ctx.Invoke("tencentcloud:Dc/getGatewayCcnRoutes:getGatewayCcnRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGatewayCcnRoutes.
type GetGatewayCcnRoutesArgs struct {
	// ID of the DCG to be queried.
	DcgId string `pulumi:"dcgId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getGatewayCcnRoutes.
type GetGatewayCcnRoutesResult struct {
	// ID of the DCG.
	DcgId string `pulumi:"dcgId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Information list of the DCG route entries.
	InstanceLists    []GetGatewayCcnRoutesInstanceList `pulumi:"instanceLists"`
	ResultOutputFile *string                           `pulumi:"resultOutputFile"`
}

func GetGatewayCcnRoutesOutput(ctx *pulumi.Context, args GetGatewayCcnRoutesOutputArgs, opts ...pulumi.InvokeOption) GetGatewayCcnRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGatewayCcnRoutesResult, error) {
			args := v.(GetGatewayCcnRoutesArgs)
			r, err := GetGatewayCcnRoutes(ctx, &args, opts...)
			var s GetGatewayCcnRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGatewayCcnRoutesResultOutput)
}

// A collection of arguments for invoking getGatewayCcnRoutes.
type GetGatewayCcnRoutesOutputArgs struct {
	// ID of the DCG to be queried.
	DcgId pulumi.StringInput `pulumi:"dcgId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetGatewayCcnRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewayCcnRoutesArgs)(nil)).Elem()
}

// A collection of values returned by getGatewayCcnRoutes.
type GetGatewayCcnRoutesResultOutput struct{ *pulumi.OutputState }

func (GetGatewayCcnRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewayCcnRoutesResult)(nil)).Elem()
}

func (o GetGatewayCcnRoutesResultOutput) ToGetGatewayCcnRoutesResultOutput() GetGatewayCcnRoutesResultOutput {
	return o
}

func (o GetGatewayCcnRoutesResultOutput) ToGetGatewayCcnRoutesResultOutputWithContext(ctx context.Context) GetGatewayCcnRoutesResultOutput {
	return o
}

// ID of the DCG.
func (o GetGatewayCcnRoutesResultOutput) DcgId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewayCcnRoutesResult) string { return v.DcgId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGatewayCcnRoutesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewayCcnRoutesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Information list of the DCG route entries.
func (o GetGatewayCcnRoutesResultOutput) InstanceLists() GetGatewayCcnRoutesInstanceListArrayOutput {
	return o.ApplyT(func(v GetGatewayCcnRoutesResult) []GetGatewayCcnRoutesInstanceList { return v.InstanceLists }).(GetGatewayCcnRoutesInstanceListArrayOutput)
}

func (o GetGatewayCcnRoutesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewayCcnRoutesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGatewayCcnRoutesResultOutput{})
}
