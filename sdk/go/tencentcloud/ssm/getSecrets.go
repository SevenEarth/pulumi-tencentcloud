// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of SSM secret
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ssm"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ssm.GetSecrets(ctx, &ssm.GetSecretsArgs{
//				OrderType:  pulumi.IntRef(1),
//				SecretName: pulumi.StringRef("test"),
//				State:      pulumi.IntRef(1),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSecrets(ctx *pulumi.Context, args *GetSecretsArgs, opts ...pulumi.InvokeOption) (*GetSecretsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSecretsResult
	err := ctx.Invoke("tencentcloud:Ssm/getSecrets:getSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecrets.
type GetSecretsArgs struct {
	// The order to sort the create time of secret. `0` - desc, `1` - asc. Default value is `0`.
	OrderType *int `pulumi:"orderType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Secret name used to filter result.
	SecretName *string `pulumi:"secretName"`
	// Filter by state of secret. `0` - all secrets are queried, `1` - only Enabled secrets are queried, `2` - only Disabled secrets are queried, `3` - only PendingDelete secrets are queried.
	State *int `pulumi:"state"`
	// Tags to filter secret.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getSecrets.
type GetSecretsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	OrderType        *int    `pulumi:"orderType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// A list of SSM secrets.
	SecretLists []GetSecretsSecretList `pulumi:"secretLists"`
	// Name of secret.
	SecretName *string                `pulumi:"secretName"`
	State      *int                   `pulumi:"state"`
	Tags       map[string]interface{} `pulumi:"tags"`
}

func GetSecretsOutput(ctx *pulumi.Context, args GetSecretsOutputArgs, opts ...pulumi.InvokeOption) GetSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecretsResult, error) {
			args := v.(GetSecretsArgs)
			r, err := GetSecrets(ctx, &args, opts...)
			var s GetSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecretsResultOutput)
}

// A collection of arguments for invoking getSecrets.
type GetSecretsOutputArgs struct {
	// The order to sort the create time of secret. `0` - desc, `1` - asc. Default value is `0`.
	OrderType pulumi.IntPtrInput `pulumi:"orderType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Secret name used to filter result.
	SecretName pulumi.StringPtrInput `pulumi:"secretName"`
	// Filter by state of secret. `0` - all secrets are queried, `1` - only Enabled secrets are queried, `2` - only Disabled secrets are queried, `3` - only PendingDelete secrets are queried.
	State pulumi.IntPtrInput `pulumi:"state"`
	// Tags to filter secret.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsArgs)(nil)).Elem()
}

// A collection of values returned by getSecrets.
type GetSecretsResultOutput struct{ *pulumi.OutputState }

func (GetSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsResult)(nil)).Elem()
}

func (o GetSecretsResultOutput) ToGetSecretsResultOutput() GetSecretsResultOutput {
	return o
}

func (o GetSecretsResultOutput) ToGetSecretsResultOutputWithContext(ctx context.Context) GetSecretsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSecretsResultOutput) OrderType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSecretsResult) *int { return v.OrderType }).(pulumi.IntPtrOutput)
}

func (o GetSecretsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecretsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// A list of SSM secrets.
func (o GetSecretsResultOutput) SecretLists() GetSecretsSecretListArrayOutput {
	return o.ApplyT(func(v GetSecretsResult) []GetSecretsSecretList { return v.SecretLists }).(GetSecretsSecretListArrayOutput)
}

// Name of secret.
func (o GetSecretsResultOutput) SecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecretsResult) *string { return v.SecretName }).(pulumi.StringPtrOutput)
}

func (o GetSecretsResultOutput) State() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSecretsResult) *int { return v.State }).(pulumi.IntPtrOutput)
}

func (o GetSecretsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetSecretsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecretsResultOutput{})
}
