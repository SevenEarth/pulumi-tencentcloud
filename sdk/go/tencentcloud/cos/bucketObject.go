// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cos

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a COS object resource to put an object(content or file) to the bucket.
//
// ## Example Usage
//
// ### Uploading a file to a bucket
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cos.NewBucketObject(ctx, "myobject", &Cos.BucketObjectArgs{
//				Bucket: pulumi.String("mycos-1258798060"),
//				Key:    pulumi.String("new_object_key"),
//				Source: pulumi.String("path/to/file"),
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
// ### Uploading a content to a bucket
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mycos, err := Cos.NewBucket(ctx, "mycos", &Cos.BucketArgs{
//				Bucket: pulumi.String("mycos-1258798060"),
//				Acl:    pulumi.String("public-read"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Cos.NewBucketObject(ctx, "myobject", &Cos.BucketObjectArgs{
//				Bucket:  mycos.Bucket,
//				Key:     pulumi.String("new_object_key"),
//				Content: pulumi.String("the content that you want to upload."),
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
type BucketObject struct {
	pulumi.CustomResourceState

	// The canned ACL to apply. Available values include `private`, `public-read`, and `public-read-write`. Defaults to `private`.
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// The name of a bucket to use. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain. For further details, RFC2616 can be referred.
	CacheControl pulumi.StringOutput `pulumi:"cacheControl"`
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// Specifies presentational information for the object.
	ContentDisposition pulumi.StringPtrOutput `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
	ContentEncoding pulumi.StringPtrOutput `pulumi:"contentEncoding"`
	// A standard MIME type describing the format of the object data.
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// The ETag generated for the object (an MD5 sum of the object content).
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the object once it is in the bucket.
	Key pulumi.StringOutput `pulumi:"key"`
	// The path to the source file being uploaded to the bucket.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// Object storage type, Available values include `STANDARD_IA`, `MAZ_STANDARD_IA`, `INTELLIGENT_TIERING`, `MAZ_INTELLIGENT_TIERING`, `ARCHIVE`, `DEEP_ARCHIVE`. For more information, please refer to: https://cloud.tencent.com/document/product/436/33417.
	StorageClass pulumi.StringOutput `pulumi:"storageClass"`
	// Tag of the object.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewBucketObject registers a new resource with the given unique name, arguments, and options.
func NewBucketObject(ctx *pulumi.Context,
	name string, args *BucketObjectArgs, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketObject
	err := ctx.RegisterResource("tencentcloud:Cos/bucketObject:BucketObject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketObject gets an existing BucketObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketObjectState, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	var resource BucketObject
	err := ctx.ReadResource("tencentcloud:Cos/bucketObject:BucketObject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketObject resources.
type bucketObjectState struct {
	// The canned ACL to apply. Available values include `private`, `public-read`, and `public-read-write`. Defaults to `private`.
	Acl *string `pulumi:"acl"`
	// The name of a bucket to use. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
	Bucket *string `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain. For further details, RFC2616 can be referred.
	CacheControl *string `pulumi:"cacheControl"`
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content *string `pulumi:"content"`
	// Specifies presentational information for the object.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// A standard MIME type describing the format of the object data.
	ContentType *string `pulumi:"contentType"`
	// The ETag generated for the object (an MD5 sum of the object content).
	Etag *string `pulumi:"etag"`
	// The name of the object once it is in the bucket.
	Key *string `pulumi:"key"`
	// The path to the source file being uploaded to the bucket.
	Source *string `pulumi:"source"`
	// Object storage type, Available values include `STANDARD_IA`, `MAZ_STANDARD_IA`, `INTELLIGENT_TIERING`, `MAZ_INTELLIGENT_TIERING`, `ARCHIVE`, `DEEP_ARCHIVE`. For more information, please refer to: https://cloud.tencent.com/document/product/436/33417.
	StorageClass *string `pulumi:"storageClass"`
	// Tag of the object.
	Tags map[string]interface{} `pulumi:"tags"`
}

type BucketObjectState struct {
	// The canned ACL to apply. Available values include `private`, `public-read`, and `public-read-write`. Defaults to `private`.
	Acl pulumi.StringPtrInput
	// The name of a bucket to use. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
	Bucket pulumi.StringPtrInput
	// Specifies caching behavior along the request/reply chain. For further details, RFC2616 can be referred.
	CacheControl pulumi.StringPtrInput
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content pulumi.StringPtrInput
	// Specifies presentational information for the object.
	ContentDisposition pulumi.StringPtrInput
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
	ContentEncoding pulumi.StringPtrInput
	// A standard MIME type describing the format of the object data.
	ContentType pulumi.StringPtrInput
	// The ETag generated for the object (an MD5 sum of the object content).
	Etag pulumi.StringPtrInput
	// The name of the object once it is in the bucket.
	Key pulumi.StringPtrInput
	// The path to the source file being uploaded to the bucket.
	Source pulumi.StringPtrInput
	// Object storage type, Available values include `STANDARD_IA`, `MAZ_STANDARD_IA`, `INTELLIGENT_TIERING`, `MAZ_INTELLIGENT_TIERING`, `ARCHIVE`, `DEEP_ARCHIVE`. For more information, please refer to: https://cloud.tencent.com/document/product/436/33417.
	StorageClass pulumi.StringPtrInput
	// Tag of the object.
	Tags pulumi.MapInput
}

func (BucketObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectState)(nil)).Elem()
}

type bucketObjectArgs struct {
	// The canned ACL to apply. Available values include `private`, `public-read`, and `public-read-write`. Defaults to `private`.
	Acl *string `pulumi:"acl"`
	// The name of a bucket to use. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
	Bucket string `pulumi:"bucket"`
	// Specifies caching behavior along the request/reply chain. For further details, RFC2616 can be referred.
	CacheControl *string `pulumi:"cacheControl"`
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content *string `pulumi:"content"`
	// Specifies presentational information for the object.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// A standard MIME type describing the format of the object data.
	ContentType *string `pulumi:"contentType"`
	// The ETag generated for the object (an MD5 sum of the object content).
	Etag *string `pulumi:"etag"`
	// The name of the object once it is in the bucket.
	Key string `pulumi:"key"`
	// The path to the source file being uploaded to the bucket.
	Source *string `pulumi:"source"`
	// Object storage type, Available values include `STANDARD_IA`, `MAZ_STANDARD_IA`, `INTELLIGENT_TIERING`, `MAZ_INTELLIGENT_TIERING`, `ARCHIVE`, `DEEP_ARCHIVE`. For more information, please refer to: https://cloud.tencent.com/document/product/436/33417.
	StorageClass *string `pulumi:"storageClass"`
	// Tag of the object.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a BucketObject resource.
type BucketObjectArgs struct {
	// The canned ACL to apply. Available values include `private`, `public-read`, and `public-read-write`. Defaults to `private`.
	Acl pulumi.StringPtrInput
	// The name of a bucket to use. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
	Bucket pulumi.StringInput
	// Specifies caching behavior along the request/reply chain. For further details, RFC2616 can be referred.
	CacheControl pulumi.StringPtrInput
	// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
	Content pulumi.StringPtrInput
	// Specifies presentational information for the object.
	ContentDisposition pulumi.StringPtrInput
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
	ContentEncoding pulumi.StringPtrInput
	// A standard MIME type describing the format of the object data.
	ContentType pulumi.StringPtrInput
	// The ETag generated for the object (an MD5 sum of the object content).
	Etag pulumi.StringPtrInput
	// The name of the object once it is in the bucket.
	Key pulumi.StringInput
	// The path to the source file being uploaded to the bucket.
	Source pulumi.StringPtrInput
	// Object storage type, Available values include `STANDARD_IA`, `MAZ_STANDARD_IA`, `INTELLIGENT_TIERING`, `MAZ_INTELLIGENT_TIERING`, `ARCHIVE`, `DEEP_ARCHIVE`. For more information, please refer to: https://cloud.tencent.com/document/product/436/33417.
	StorageClass pulumi.StringPtrInput
	// Tag of the object.
	Tags pulumi.MapInput
}

func (BucketObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectArgs)(nil)).Elem()
}

type BucketObjectInput interface {
	pulumi.Input

	ToBucketObjectOutput() BucketObjectOutput
	ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput
}

func (*BucketObject) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketObject)(nil)).Elem()
}

func (i *BucketObject) ToBucketObjectOutput() BucketObjectOutput {
	return i.ToBucketObjectOutputWithContext(context.Background())
}

func (i *BucketObject) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectOutput)
}

// BucketObjectArrayInput is an input type that accepts BucketObjectArray and BucketObjectArrayOutput values.
// You can construct a concrete instance of `BucketObjectArrayInput` via:
//
//	BucketObjectArray{ BucketObjectArgs{...} }
type BucketObjectArrayInput interface {
	pulumi.Input

	ToBucketObjectArrayOutput() BucketObjectArrayOutput
	ToBucketObjectArrayOutputWithContext(context.Context) BucketObjectArrayOutput
}

type BucketObjectArray []BucketObjectInput

func (BucketObjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketObject)(nil)).Elem()
}

func (i BucketObjectArray) ToBucketObjectArrayOutput() BucketObjectArrayOutput {
	return i.ToBucketObjectArrayOutputWithContext(context.Background())
}

func (i BucketObjectArray) ToBucketObjectArrayOutputWithContext(ctx context.Context) BucketObjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectArrayOutput)
}

// BucketObjectMapInput is an input type that accepts BucketObjectMap and BucketObjectMapOutput values.
// You can construct a concrete instance of `BucketObjectMapInput` via:
//
//	BucketObjectMap{ "key": BucketObjectArgs{...} }
type BucketObjectMapInput interface {
	pulumi.Input

	ToBucketObjectMapOutput() BucketObjectMapOutput
	ToBucketObjectMapOutputWithContext(context.Context) BucketObjectMapOutput
}

type BucketObjectMap map[string]BucketObjectInput

func (BucketObjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketObject)(nil)).Elem()
}

func (i BucketObjectMap) ToBucketObjectMapOutput() BucketObjectMapOutput {
	return i.ToBucketObjectMapOutputWithContext(context.Background())
}

func (i BucketObjectMap) ToBucketObjectMapOutputWithContext(ctx context.Context) BucketObjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectMapOutput)
}

type BucketObjectOutput struct{ *pulumi.OutputState }

func (BucketObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketObject)(nil)).Elem()
}

func (o BucketObjectOutput) ToBucketObjectOutput() BucketObjectOutput {
	return o
}

func (o BucketObjectOutput) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return o
}

// The canned ACL to apply. Available values include `private`, `public-read`, and `public-read-write`. Defaults to `private`.
func (o BucketObjectOutput) Acl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.Acl }).(pulumi.StringPtrOutput)
}

// The name of a bucket to use. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
func (o BucketObjectOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Specifies caching behavior along the request/reply chain. For further details, RFC2616 can be referred.
func (o BucketObjectOutput) CacheControl() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.CacheControl }).(pulumi.StringOutput)
}

// Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
func (o BucketObjectOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

// Specifies presentational information for the object.
func (o BucketObjectOutput) ContentDisposition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.ContentDisposition }).(pulumi.StringPtrOutput)
}

// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
func (o BucketObjectOutput) ContentEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.ContentEncoding }).(pulumi.StringPtrOutput)
}

// A standard MIME type describing the format of the object data.
func (o BucketObjectOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.ContentType }).(pulumi.StringOutput)
}

// The ETag generated for the object (an MD5 sum of the object content).
func (o BucketObjectOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the object once it is in the bucket.
func (o BucketObjectOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The path to the source file being uploaded to the bucket.
func (o BucketObjectOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

// Object storage type, Available values include `STANDARD_IA`, `MAZ_STANDARD_IA`, `INTELLIGENT_TIERING`, `MAZ_INTELLIGENT_TIERING`, `ARCHIVE`, `DEEP_ARCHIVE`. For more information, please refer to: https://cloud.tencent.com/document/product/436/33417.
func (o BucketObjectOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.StorageClass }).(pulumi.StringOutput)
}

// Tag of the object.
func (o BucketObjectOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type BucketObjectArrayOutput struct{ *pulumi.OutputState }

func (BucketObjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketObject)(nil)).Elem()
}

func (o BucketObjectArrayOutput) ToBucketObjectArrayOutput() BucketObjectArrayOutput {
	return o
}

func (o BucketObjectArrayOutput) ToBucketObjectArrayOutputWithContext(ctx context.Context) BucketObjectArrayOutput {
	return o
}

func (o BucketObjectArrayOutput) Index(i pulumi.IntInput) BucketObjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketObject {
		return vs[0].([]*BucketObject)[vs[1].(int)]
	}).(BucketObjectOutput)
}

type BucketObjectMapOutput struct{ *pulumi.OutputState }

func (BucketObjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketObject)(nil)).Elem()
}

func (o BucketObjectMapOutput) ToBucketObjectMapOutput() BucketObjectMapOutput {
	return o
}

func (o BucketObjectMapOutput) ToBucketObjectMapOutputWithContext(ctx context.Context) BucketObjectMapOutput {
	return o
}

func (o BucketObjectMapOutput) MapIndex(k pulumi.StringInput) BucketObjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketObject {
		return vs[0].(map[string]*BucketObject)[vs[1].(string)]
	}).(BucketObjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectInput)(nil)).Elem(), &BucketObject{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectArrayInput)(nil)).Elem(), BucketObjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectMapInput)(nil)).Elem(), BucketObjectMap{})
	pulumi.RegisterOutputType(BucketObjectOutput{})
	pulumi.RegisterOutputType(BucketObjectArrayOutput{})
	pulumi.RegisterOutputType(BucketObjectMapOutput{})
}
