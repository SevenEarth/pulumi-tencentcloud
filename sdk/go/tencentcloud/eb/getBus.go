// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of eb bus
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Eb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Eb.NewEventBus(ctx, "foo", &Eb.EventBusArgs{
//				EventBusName: pulumi.String("tf-event_bus"),
//				Description:  pulumi.String("event bus desc"),
//				EnableStore:  pulumi.Bool(false),
//				SaveDays:     pulumi.Int(1),
//				Tags: pulumi.Map{
//					"createdBy": pulumi.Any("terraform"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Eb.GetBus(ctx, &eb.GetBusArgs{
//				OrderBy: pulumi.StringRef("AddTime"),
//				Order:   pulumi.StringRef("DESC"),
//				Filters: []eb.GetBusFilter{
//					{
//						Values: []string{
//							"Custom",
//						},
//						Name: "Type",
//					},
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
func GetBus(ctx *pulumi.Context, args *GetBusArgs, opts ...pulumi.InvokeOption) (*GetBusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBusResult
	err := ctx.Invoke("tencentcloud:Eb/getBus:getBus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBus.
type GetBusArgs struct {
	// Filter conditions. The upper limit of Filters per request is 10, and the upper limit of Filter.Values 5.
	Filters []GetBusFilter `pulumi:"filters"`
	// Return results in ascending or descending order, optional values ASC (ascending) and DESC (descending).
	Order *string `pulumi:"order"`
	// According to which field to sort the returned results, the following fields are supported: AddTime (creation time), ModTime (modification time).
	OrderBy *string `pulumi:"orderBy"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getBus.
type GetBusResult struct {
	// event set information.
	EventBuses []GetBusEventBus `pulumi:"eventBuses"`
	Filters    []GetBusFilter   `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Order            *string `pulumi:"order"`
	OrderBy          *string `pulumi:"orderBy"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetBusOutput(ctx *pulumi.Context, args GetBusOutputArgs, opts ...pulumi.InvokeOption) GetBusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBusResult, error) {
			args := v.(GetBusArgs)
			r, err := GetBus(ctx, &args, opts...)
			var s GetBusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBusResultOutput)
}

// A collection of arguments for invoking getBus.
type GetBusOutputArgs struct {
	// Filter conditions. The upper limit of Filters per request is 10, and the upper limit of Filter.Values 5.
	Filters GetBusFilterArrayInput `pulumi:"filters"`
	// Return results in ascending or descending order, optional values ASC (ascending) and DESC (descending).
	Order pulumi.StringPtrInput `pulumi:"order"`
	// According to which field to sort the returned results, the following fields are supported: AddTime (creation time), ModTime (modification time).
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetBusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBusArgs)(nil)).Elem()
}

// A collection of values returned by getBus.
type GetBusResultOutput struct{ *pulumi.OutputState }

func (GetBusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBusResult)(nil)).Elem()
}

func (o GetBusResultOutput) ToGetBusResultOutput() GetBusResultOutput {
	return o
}

func (o GetBusResultOutput) ToGetBusResultOutputWithContext(ctx context.Context) GetBusResultOutput {
	return o
}

// event set information.
func (o GetBusResultOutput) EventBuses() GetBusEventBusArrayOutput {
	return o.ApplyT(func(v GetBusResult) []GetBusEventBus { return v.EventBuses }).(GetBusEventBusArrayOutput)
}

func (o GetBusResultOutput) Filters() GetBusFilterArrayOutput {
	return o.ApplyT(func(v GetBusResult) []GetBusFilter { return v.Filters }).(GetBusFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBusResultOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBusResult) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o GetBusResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBusResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o GetBusResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBusResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBusResultOutput{})
}
