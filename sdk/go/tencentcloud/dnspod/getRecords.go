// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnspod

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query dnspod record list.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dnspod"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dnspod"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		record, err := Dnspod.GetRecords(ctx, &dnspod.GetRecordsArgs{
// 			Domain:    pulumi.StringRef("example.com"),
// 			Subdomain: pulumi.StringRef("www"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("result", record.Results)
// 		return nil
// 	})
// }
// ```
//
// Use verbose filter
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dnspod"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dnspod"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		record, err := Dnspod.GetRecords(ctx, &dnspod.GetRecordsArgs{
// 			Domain:     pulumi.StringRef("example.com"),
// 			Subdomain:  pulumi.StringRef("www"),
// 			Limit:      pulumi.IntRef(100),
// 			RecordType: pulumi.StringRef("TXT"),
// 			SortField:  pulumi.StringRef("updated_on"),
// 			SortType:   pulumi.StringRef("DESC"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("result", record.Results)
// 		return nil
// 	})
// }
// ```
func LookupRecords(ctx *pulumi.Context, args *LookupRecordsArgs, opts ...pulumi.InvokeOption) (*LookupRecordsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRecordsResult
	err := ctx.Invoke("tencentcloud:Dnspod/getRecords:getRecords", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRecords.
type LookupRecordsArgs struct {
	// The domain for which DNS records are to be obtained.
	Domain *string `pulumi:"domain"`
	// The ID of the domain for which DNS records are to be obtained. If DomainId is passed in, the system will omit the parameter domain.
	DomainId *string `pulumi:"domainId"`
	// The group ID.
	GroupId *string `pulumi:"groupId"`
	// The keyword for searching for DNS records. Host headers and record values are supported.
	Keyword *string `pulumi:"keyword"`
	// The limit. It defaults to 100 and can be up to 3,000.
	Limit *int `pulumi:"limit"`
	// The offset. Default value: 0.
	Offset *int `pulumi:"offset"`
	// The split zone name.
	RecordLine *string `pulumi:"recordLine"`
	// The split zone ID. If `recordLineId` is passed in, the system will omit the parameter `recordLine`.
	RecordLineId *string `pulumi:"recordLineId"`
	// The type of DNS record, such as A, CNAME, NS, AAAA, explicit URL, implicit URL, CAA, or SPF record.
	RecordType *string `pulumi:"recordType"`
	// Used for store query result as JSON.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// The sorting field. Available values: name, line, type, value, weight, mx, and ttl,updated_on.
	SortField *string `pulumi:"sortField"`
	// The sorting type. Valid values: ASC (ascending, default), DESC (descending).
	SortType *string `pulumi:"sortType"`
	// The host header of a DNS record. If this parameter is passed in, only the DNS record corresponding to this host header will be returned.
	Subdomain *string `pulumi:"subdomain"`
}

// A collection of values returned by getRecords.
type LookupRecordsResult struct {
	Domain   *string `pulumi:"domain"`
	DomainId *string `pulumi:"domainId"`
	GroupId  *string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Keyword *string `pulumi:"keyword"`
	Limit   *int    `pulumi:"limit"`
	Offset  *int    `pulumi:"offset"`
	// Count info of the queried record list.
	RecordCountInfos []GetRecordsRecordCountInfo `pulumi:"recordCountInfos"`
	RecordLine       *string                     `pulumi:"recordLine"`
	RecordLineId     *string                     `pulumi:"recordLineId"`
	RecordType       *string                     `pulumi:"recordType"`
	ResultOutputFile *string                     `pulumi:"resultOutputFile"`
	// The record list result.
	Results   []GetRecordsResult `pulumi:"results"`
	SortField *string            `pulumi:"sortField"`
	SortType  *string            `pulumi:"sortType"`
	Subdomain *string            `pulumi:"subdomain"`
}

func LookupRecordsOutput(ctx *pulumi.Context, args LookupRecordsOutputArgs, opts ...pulumi.InvokeOption) LookupRecordsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRecordsResult, error) {
			args := v.(LookupRecordsArgs)
			r, err := LookupRecords(ctx, &args, opts...)
			var s LookupRecordsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRecordsResultOutput)
}

// A collection of arguments for invoking getRecords.
type LookupRecordsOutputArgs struct {
	// The domain for which DNS records are to be obtained.
	Domain pulumi.StringPtrInput `pulumi:"domain"`
	// The ID of the domain for which DNS records are to be obtained. If DomainId is passed in, the system will omit the parameter domain.
	DomainId pulumi.StringPtrInput `pulumi:"domainId"`
	// The group ID.
	GroupId pulumi.StringPtrInput `pulumi:"groupId"`
	// The keyword for searching for DNS records. Host headers and record values are supported.
	Keyword pulumi.StringPtrInput `pulumi:"keyword"`
	// The limit. It defaults to 100 and can be up to 3,000.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// The offset. Default value: 0.
	Offset pulumi.IntPtrInput `pulumi:"offset"`
	// The split zone name.
	RecordLine pulumi.StringPtrInput `pulumi:"recordLine"`
	// The split zone ID. If `recordLineId` is passed in, the system will omit the parameter `recordLine`.
	RecordLineId pulumi.StringPtrInput `pulumi:"recordLineId"`
	// The type of DNS record, such as A, CNAME, NS, AAAA, explicit URL, implicit URL, CAA, or SPF record.
	RecordType pulumi.StringPtrInput `pulumi:"recordType"`
	// Used for store query result as JSON.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// The sorting field. Available values: name, line, type, value, weight, mx, and ttl,updated_on.
	SortField pulumi.StringPtrInput `pulumi:"sortField"`
	// The sorting type. Valid values: ASC (ascending, default), DESC (descending).
	SortType pulumi.StringPtrInput `pulumi:"sortType"`
	// The host header of a DNS record. If this parameter is passed in, only the DNS record corresponding to this host header will be returned.
	Subdomain pulumi.StringPtrInput `pulumi:"subdomain"`
}

func (LookupRecordsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordsArgs)(nil)).Elem()
}

// A collection of values returned by getRecords.
type LookupRecordsResultOutput struct{ *pulumi.OutputState }

func (LookupRecordsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordsResult)(nil)).Elem()
}

func (o LookupRecordsResultOutput) ToLookupRecordsResultOutput() LookupRecordsResultOutput {
	return o
}

func (o LookupRecordsResultOutput) ToLookupRecordsResultOutputWithContext(ctx context.Context) LookupRecordsResultOutput {
	return o
}

func (o LookupRecordsResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRecordsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRecordsResultOutput) Keyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.Keyword }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o LookupRecordsResultOutput) Offset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *int { return v.Offset }).(pulumi.IntPtrOutput)
}

// Count info of the queried record list.
func (o LookupRecordsResultOutput) RecordCountInfos() GetRecordsRecordCountInfoArrayOutput {
	return o.ApplyT(func(v LookupRecordsResult) []GetRecordsRecordCountInfo { return v.RecordCountInfos }).(GetRecordsRecordCountInfoArrayOutput)
}

func (o LookupRecordsResultOutput) RecordLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.RecordLine }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) RecordLineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.RecordLineId }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) RecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.RecordType }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// The record list result.
func (o LookupRecordsResultOutput) Results() GetRecordsResultArrayOutput {
	return o.ApplyT(func(v LookupRecordsResult) []GetRecordsResult { return v.Results }).(GetRecordsResultArrayOutput)
}

func (o LookupRecordsResultOutput) SortField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.SortField }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) SortType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.SortType }).(pulumi.StringPtrOutput)
}

func (o LookupRecordsResultOutput) Subdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordsResult) *string { return v.Subdomain }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRecordsResultOutput{})
}
