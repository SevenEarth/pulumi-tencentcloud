// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tsf apiGroup
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tsf.GetApiGroup(ctx, &tsf.GetApiGroupArgs{
//				AuthType:          pulumi.StringRef("none"),
//				GatewayInstanceId: pulumi.StringRef("gw-ins-lvdypq5k"),
//				GroupType:         pulumi.StringRef("ms"),
//				OrderBy:           pulumi.StringRef("created_time"),
//				OrderType:         pulumi.IntRef(0),
//				SearchWord:        pulumi.StringRef("xxx01"),
//				Status:            pulumi.StringRef("released"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupApiGroup(ctx *pulumi.Context, args *LookupApiGroupArgs, opts ...pulumi.InvokeOption) (*LookupApiGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupApiGroupResult
	err := ctx.Invoke("tencentcloud:Tsf/getApiGroup:getApiGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApiGroup.
type LookupApiGroupArgs struct {
	// Authentication type. secret: Secret key authentication; none: No authentication.
	AuthType *string `pulumi:"authType"`
	// Gateway Instance Id.
	GatewayInstanceId *string `pulumi:"gatewayInstanceId"`
	// Group type. ms: Microservice group; external: External API group.
	GroupType *string `pulumi:"groupType"`
	// Sorting field: createdTime or group_context.
	OrderBy *string `pulumi:"orderBy"`
	// Sorting type: 0 (ASC) or 1 (DESC).
	OrderType *int `pulumi:"orderType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// search word.
	SearchWord *string `pulumi:"searchWord"`
	// Publishing status. drafted: Not published. released: Published.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getApiGroup.
type LookupApiGroupResult struct {
	// Authentication type. secret: key authentication; none: no authentication.Note: This field may return null, indicating that no valid values can be obtained.
	AuthType *string `pulumi:"authType"`
	// Gateway Instance Id.Note: This field may return null, indicating that no valid values can be obtained.
	GatewayInstanceId *string `pulumi:"gatewayInstanceId"`
	// Group type.Note: This field may return null, indicating that no valid values can be obtained.
	GroupType *string `pulumi:"groupType"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	OrderBy          *string `pulumi:"orderBy"`
	OrderType        *int    `pulumi:"orderType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Pagination structure.Note: This field may return null, indicating that no valid values can be obtained.
	Results    []GetApiGroupResult `pulumi:"results"`
	SearchWord *string             `pulumi:"searchWord"`
	// Release status. drafted: not released. released: released.Note: This field may return null, indicating that no valid values can be obtained.
	Status *string `pulumi:"status"`
}

func LookupApiGroupOutput(ctx *pulumi.Context, args LookupApiGroupOutputArgs, opts ...pulumi.InvokeOption) LookupApiGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiGroupResult, error) {
			args := v.(LookupApiGroupArgs)
			r, err := LookupApiGroup(ctx, &args, opts...)
			var s LookupApiGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiGroupResultOutput)
}

// A collection of arguments for invoking getApiGroup.
type LookupApiGroupOutputArgs struct {
	// Authentication type. secret: Secret key authentication; none: No authentication.
	AuthType pulumi.StringPtrInput `pulumi:"authType"`
	// Gateway Instance Id.
	GatewayInstanceId pulumi.StringPtrInput `pulumi:"gatewayInstanceId"`
	// Group type. ms: Microservice group; external: External API group.
	GroupType pulumi.StringPtrInput `pulumi:"groupType"`
	// Sorting field: createdTime or group_context.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Sorting type: 0 (ASC) or 1 (DESC).
	OrderType pulumi.IntPtrInput `pulumi:"orderType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// search word.
	SearchWord pulumi.StringPtrInput `pulumi:"searchWord"`
	// Publishing status. drafted: Not published. released: Published.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (LookupApiGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiGroupArgs)(nil)).Elem()
}

// A collection of values returned by getApiGroup.
type LookupApiGroupResultOutput struct{ *pulumi.OutputState }

func (LookupApiGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiGroupResult)(nil)).Elem()
}

func (o LookupApiGroupResultOutput) ToLookupApiGroupResultOutput() LookupApiGroupResultOutput {
	return o
}

func (o LookupApiGroupResultOutput) ToLookupApiGroupResultOutputWithContext(ctx context.Context) LookupApiGroupResultOutput {
	return o
}

// Authentication type. secret: key authentication; none: no authentication.Note: This field may return null, indicating that no valid values can be obtained.
func (o LookupApiGroupResultOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGroupResult) *string { return v.AuthType }).(pulumi.StringPtrOutput)
}

// Gateway Instance Id.Note: This field may return null, indicating that no valid values can be obtained.
func (o LookupApiGroupResultOutput) GatewayInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGroupResult) *string { return v.GatewayInstanceId }).(pulumi.StringPtrOutput)
}

// Group type.Note: This field may return null, indicating that no valid values can be obtained.
func (o LookupApiGroupResultOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGroupResult) *string { return v.GroupType }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupApiGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiGroupResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGroupResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o LookupApiGroupResultOutput) OrderType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupApiGroupResult) *int { return v.OrderType }).(pulumi.IntPtrOutput)
}

func (o LookupApiGroupResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGroupResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Pagination structure.Note: This field may return null, indicating that no valid values can be obtained.
func (o LookupApiGroupResultOutput) Results() GetApiGroupResultArrayOutput {
	return o.ApplyT(func(v LookupApiGroupResult) []GetApiGroupResult { return v.Results }).(GetApiGroupResultArrayOutput)
}

func (o LookupApiGroupResultOutput) SearchWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGroupResult) *string { return v.SearchWord }).(pulumi.StringPtrOutput)
}

// Release status. drafted: not released. released: released.Note: This field may return null, indicating that no valid values can be obtained.
func (o LookupApiGroupResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiGroupResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiGroupResultOutput{})
}
