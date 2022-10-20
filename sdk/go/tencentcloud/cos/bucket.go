// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a COS resource to create a COS bucket and set its attributes.
//
// ## Example Usage
//
// # Private Bucket
//
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
//			_, err := Cos.NewBucket(ctx, "mycos", &Cos.BucketArgs{
//				Acl:    pulumi.String("private"),
//				Bucket: pulumi.String("mycos-1258798060"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Creation of multiple available zone bucket
//
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
//			_, err := Cos.NewBucket(ctx, "mycos", &Cos.BucketArgs{
//				Acl:              pulumi.String("private"),
//				Bucket:           pulumi.String("mycos-1258798060"),
//				MultiAz:          pulumi.Bool(true),
//				VersioningEnable: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Using verbose acl
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cos.NewBucket(ctx, "withAclBody", &Cos.BucketArgs{
//				AclBody: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "<AccessControlPolicy>\n", "    <Owner>\n", "        <ID>qcs::cam::uin/100000000001:uin/100000000001</ID>\n", "    </Owner>\n", "    <AccessControlList>\n", "        <Grant>\n", "            <Grantee xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"Group\">\n", "                <URI>http://cam.qcloud.com/groups/global/AllUsers</URI>\n", "            </Grantee>\n", "            <Permission>READ</Permission>\n", "        </Grant>\n", "        <Grant>\n", "            <Grantee xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"CanonicalUser\">\n", "                <ID>qcs::cam::uin/100000000001:uin/100000000001</ID>\n", "                <DisplayName>qcs::cam::uin/100000000001:uin/100000000001</DisplayName>\n", "            </Grantee>\n", "            <Permission>WRITE</Permission>\n", "        </Grant>\n", "        <Grant>\n", "            <Grantee xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"CanonicalUser\">\n", "                <ID>qcs::cam::uin/100000000001:uin/100000000001</ID>\n", "                <DisplayName>qcs::cam::uin/100000000001:uin/100000000001</DisplayName>\n", "            </Grantee>\n", "            <Permission>READ_ACP</Permission>\n", "        </Grant>\n", "    </AccessControlList>\n", "</AccessControlPolicy>\n", "\n")),
//				Bucket:  pulumi.String("mycos-1258798060"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Static Website
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cos.NewBucket(ctx, "mycos", &Cos.BucketArgs{
//				Bucket: pulumi.String("mycos-1258798060"),
//				Website: &cos.BucketWebsiteArgs{
//					ErrorDocument: pulumi.String("error.html"),
//					IndexDocument: pulumi.String("index.html"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Using CORS
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cos.NewBucket(ctx, "mycos", &Cos.BucketArgs{
//				Acl:    pulumi.String("public-read-write"),
//				Bucket: pulumi.String("mycos-1258798060"),
//				CorsRules: cos.BucketCorsRuleArray{
//					&cos.BucketCorsRuleArgs{
//						AllowedHeaders: pulumi.StringArray{
//							pulumi.String("*"),
//						},
//						AllowedMethods: pulumi.StringArray{
//							pulumi.String("PUT"),
//							pulumi.String("POST"),
//						},
//						AllowedOrigins: pulumi.StringArray{
//							pulumi.String("http://*.abc.com"),
//						},
//						ExposeHeaders: pulumi.StringArray{
//							pulumi.String("Etag"),
//						},
//						MaxAgeSeconds: pulumi.Int(300),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Using object lifecycle
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cos.NewBucket(ctx, "mycos", &Cos.BucketArgs{
//				Acl:    pulumi.String("public-read-write"),
//				Bucket: pulumi.String("mycos-1258798060"),
//				LifecycleRules: cos.BucketLifecycleRuleArray{
//					&cos.BucketLifecycleRuleArgs{
//						Expiration: &cos.BucketLifecycleRuleExpirationArgs{
//							Days: pulumi.Int(90),
//						},
//						FilterPrefix: pulumi.String("path1/"),
//						Transitions: cos.BucketLifecycleRuleTransitionArray{
//							&cos.BucketLifecycleRuleTransitionArgs{
//								Date:         pulumi.String("2019-06-01"),
//								StorageClass: pulumi.String("STANDARD_IA"),
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Using custom origin domain settings
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cos.NewBucket(ctx, "withOrigin", &Cos.BucketArgs{
//				Acl:    pulumi.String("private"),
//				Bucket: pulumi.String("mycos-1258798060"),
//				OriginDomainRules: cos.BucketOriginDomainRuleArray{
//					&cos.BucketOriginDomainRuleArgs{
//						Domain: pulumi.String("abc.example.com"),
//						Status: pulumi.String("ENABLE"),
//						Type:   pulumi.String("REST"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Using origin-pull settings
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cos.NewBucket(ctx, "withOrigin", &Cos.BucketArgs{
//				Acl:    pulumi.String("private"),
//				Bucket: pulumi.String("mycos-1258798060"),
//				OriginPullRules: cos.BucketOriginPullRuleArray{
//					&cos.BucketOriginPullRuleArgs{
//						CustomHttpHeaders: pulumi.AnyMap{
//							"x-custom-header": pulumi.Any("custom_value"),
//						},
//						FollowHttpHeaders: pulumi.StringArray{
//							pulumi.String("origin"),
//							pulumi.String("host"),
//						},
//						FollowQueryString: pulumi.Bool(true),
//						FollowRedirection: pulumi.Bool(true),
//						Host:              pulumi.String("abc.example.com"),
//						Prefix:            pulumi.String("/"),
//						Priority:          pulumi.Int(1),
//						Protocol:          pulumi.String("FOLLOW"),
//						SyncBackToSource:  pulumi.Bool(false),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Using replication
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			replica1, err := Cos.NewBucket(ctx, "replica1", &Cos.BucketArgs{
//				Acl:              pulumi.String("private"),
//				Bucket:           pulumi.String("tf-replica-foo-1234567890"),
//				VersioningEnable: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Cos.NewBucket(ctx, "withReplication", &Cos.BucketArgs{
//				Acl:         pulumi.String("private"),
//				Bucket:      pulumi.String("tf-bucket-replica-1234567890"),
//				ReplicaRole: pulumi.String("qcs::cam::uin/100000000001:uin/100000000001"),
//				ReplicaRules: cos.BucketReplicaRuleArray{
//					&cos.BucketReplicaRuleArgs{
//						DestinationBucket: replica1.Bucket.ApplyT(func(bucket string) (string, error) {
//							return fmt.Sprintf("%v%v%v%v", "qcs::cos:", "%", "s::", bucket), nil
//						}).(pulumi.StringOutput),
//						Id:     pulumi.String("test-rep1"),
//						Prefix: pulumi.String("dist"),
//						Status: pulumi.String("Enabled"),
//					},
//				},
//				VersioningEnable: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # Setting log status
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cosLogGrantRole, err := Cam.NewRole(ctx, "cosLogGrantRole", &Cam.RoleArgs{
//				Document:    pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"version\": \"2.0\",\n", "  \"statement\": [\n", "    {\n", "      \"action\": [\n", "        \"name/sts:AssumeRole\"\n", "      ],\n", "      \"effect\": \"allow\",\n", "      \"principal\": {\n", "        \"service\": [\n", "          \"cls.cloud.tencent.com\"\n", "        ]\n", "      }\n", "    }\n", "  ]\n", "}\n")),
//				Description: pulumi.String("cos log enable grant"),
//			})
//			if err != nil {
//				return err
//			}
//			cosAccess, err := Cam.GetPolicies(ctx, &cam.GetPoliciesArgs{
//				Name: pulumi.StringRef("QcloudCOSAccessForCLSRole"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Cam.NewRolePolicyAttachment(ctx, "cosLogGrantRolePolicyAttachment", &Cam.RolePolicyAttachmentArgs{
//				RoleId:   cosLogGrantRole.ID(),
//				PolicyId: pulumi.String(cosAccess.PolicyLists[0].PolicyId),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Cos.NewBucket(ctx, "mylog", &Cos.BucketArgs{
//				Bucket: pulumi.String("mylog-1258798060"),
//				Acl:    pulumi.String("private"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Cos.NewBucket(ctx, "mycos", &Cos.BucketArgs{
//				Bucket:          pulumi.String("mycos-1258798060"),
//				Acl:             pulumi.String("private"),
//				LogEnable:       pulumi.Bool(true),
//				LogTargetBucket: pulumi.String("mylog-1258798060"),
//				LogPrefix:       pulumi.String("MyLogPrefix"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// COS bucket can be imported, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Cos/bucket:Bucket bucket bucket-name
//
// ```
type Bucket struct {
	pulumi.CustomResourceState

	// The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// ACL XML body for multiple grant info. NOTE: this argument will overwrite `acl`. Check https://intl.cloud.tencent.com/document/product/436/7737 for more detail.
	AclBody pulumi.StringPtrOutput `pulumi:"aclBody"`
	// The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRules BucketCorsRuleArrayOutput `pulumi:"corsRules"`
	// The URL of this cos bucket.
	CosBucketUrl pulumi.StringOutput `pulumi:"cosBucketUrl"`
	// The server-side encryption algorithm to use. Valid value is `AES256`.
	EncryptionAlgorithm pulumi.StringPtrOutput `pulumi:"encryptionAlgorithm"`
	// A configuration of object lifecycle management (documented below).
	LifecycleRules BucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// Indicate the access log of this bucket to be saved or not. Default is `false`. If set `true`, the access log will be saved with `logTargetBucket`. To enable log, the full access of log service must be granted. [Full Access Role Policy](https://intl.cloud.tencent.com/document/product/436/16920).
	LogEnable pulumi.BoolPtrOutput `pulumi:"logEnable"`
	// The prefix log name which saves the access log of this bucket per 5 minutes. Eg. `MyLogPrefix/`. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`.
	LogPrefix pulumi.StringOutput `pulumi:"logPrefix"`
	// The target bucket name which saves the access log of this bucket per 5 minutes. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`. User must have full access on this bucket.
	LogTargetBucket pulumi.StringOutput `pulumi:"logTargetBucket"`
	// Indicates whether to create a bucket of multi available zone. NOTE: If set to true, the versioning must enable.
	MultiAz pulumi.BoolPtrOutput `pulumi:"multiAz"`
	// Bucket Origin Domain settings.
	OriginDomainRules BucketOriginDomainRuleArrayOutput `pulumi:"originDomainRules"`
	// Bucket Origin-Pull settings.
	OriginPullRules BucketOriginPullRuleArrayOutput `pulumi:"originPullRules"`
	// Request initiator identifier, format: `qcs::cam::uin/<owneruin>:uin/<subuin>`. NOTE: only `versioningEnable` is true can configure this argument.
	ReplicaRole pulumi.StringPtrOutput `pulumi:"replicaRole"`
	// List of replica rule. NOTE: only `versioningEnable` is true and `replicaRole` set can configure this argument.
	ReplicaRules BucketReplicaRuleArrayOutput `pulumi:"replicaRules"`
	// The tags of a bucket.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Enable bucket versioning.
	VersioningEnable pulumi.BoolPtrOutput `pulumi:"versioningEnable"`
	// A website object(documented below).
	Website BucketWebsitePtrOutput `pulumi:"website"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Bucket
	err := ctx.RegisterResource("tencentcloud:Cos/bucket:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("tencentcloud:Cos/bucket:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
	// The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.
	Acl *string `pulumi:"acl"`
	// ACL XML body for multiple grant info. NOTE: this argument will overwrite `acl`. Check https://intl.cloud.tencent.com/document/product/436/7737 for more detail.
	AclBody *string `pulumi:"aclBody"`
	// The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
	Bucket *string `pulumi:"bucket"`
	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// The URL of this cos bucket.
	CosBucketUrl *string `pulumi:"cosBucketUrl"`
	// The server-side encryption algorithm to use. Valid value is `AES256`.
	EncryptionAlgorithm *string `pulumi:"encryptionAlgorithm"`
	// A configuration of object lifecycle management (documented below).
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// Indicate the access log of this bucket to be saved or not. Default is `false`. If set `true`, the access log will be saved with `logTargetBucket`. To enable log, the full access of log service must be granted. [Full Access Role Policy](https://intl.cloud.tencent.com/document/product/436/16920).
	LogEnable *bool `pulumi:"logEnable"`
	// The prefix log name which saves the access log of this bucket per 5 minutes. Eg. `MyLogPrefix/`. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`.
	LogPrefix *string `pulumi:"logPrefix"`
	// The target bucket name which saves the access log of this bucket per 5 minutes. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`. User must have full access on this bucket.
	LogTargetBucket *string `pulumi:"logTargetBucket"`
	// Indicates whether to create a bucket of multi available zone. NOTE: If set to true, the versioning must enable.
	MultiAz *bool `pulumi:"multiAz"`
	// Bucket Origin Domain settings.
	OriginDomainRules []BucketOriginDomainRule `pulumi:"originDomainRules"`
	// Bucket Origin-Pull settings.
	OriginPullRules []BucketOriginPullRule `pulumi:"originPullRules"`
	// Request initiator identifier, format: `qcs::cam::uin/<owneruin>:uin/<subuin>`. NOTE: only `versioningEnable` is true can configure this argument.
	ReplicaRole *string `pulumi:"replicaRole"`
	// List of replica rule. NOTE: only `versioningEnable` is true and `replicaRole` set can configure this argument.
	ReplicaRules []BucketReplicaRule `pulumi:"replicaRules"`
	// The tags of a bucket.
	Tags map[string]interface{} `pulumi:"tags"`
	// Enable bucket versioning.
	VersioningEnable *bool `pulumi:"versioningEnable"`
	// A website object(documented below).
	Website *BucketWebsite `pulumi:"website"`
}

type BucketState struct {
	// The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.
	Acl pulumi.StringPtrInput
	// ACL XML body for multiple grant info. NOTE: this argument will overwrite `acl`. Check https://intl.cloud.tencent.com/document/product/436/7737 for more detail.
	AclBody pulumi.StringPtrInput
	// The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
	Bucket pulumi.StringPtrInput
	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRules BucketCorsRuleArrayInput
	// The URL of this cos bucket.
	CosBucketUrl pulumi.StringPtrInput
	// The server-side encryption algorithm to use. Valid value is `AES256`.
	EncryptionAlgorithm pulumi.StringPtrInput
	// A configuration of object lifecycle management (documented below).
	LifecycleRules BucketLifecycleRuleArrayInput
	// Indicate the access log of this bucket to be saved or not. Default is `false`. If set `true`, the access log will be saved with `logTargetBucket`. To enable log, the full access of log service must be granted. [Full Access Role Policy](https://intl.cloud.tencent.com/document/product/436/16920).
	LogEnable pulumi.BoolPtrInput
	// The prefix log name which saves the access log of this bucket per 5 minutes. Eg. `MyLogPrefix/`. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`.
	LogPrefix pulumi.StringPtrInput
	// The target bucket name which saves the access log of this bucket per 5 minutes. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`. User must have full access on this bucket.
	LogTargetBucket pulumi.StringPtrInput
	// Indicates whether to create a bucket of multi available zone. NOTE: If set to true, the versioning must enable.
	MultiAz pulumi.BoolPtrInput
	// Bucket Origin Domain settings.
	OriginDomainRules BucketOriginDomainRuleArrayInput
	// Bucket Origin-Pull settings.
	OriginPullRules BucketOriginPullRuleArrayInput
	// Request initiator identifier, format: `qcs::cam::uin/<owneruin>:uin/<subuin>`. NOTE: only `versioningEnable` is true can configure this argument.
	ReplicaRole pulumi.StringPtrInput
	// List of replica rule. NOTE: only `versioningEnable` is true and `replicaRole` set can configure this argument.
	ReplicaRules BucketReplicaRuleArrayInput
	// The tags of a bucket.
	Tags pulumi.MapInput
	// Enable bucket versioning.
	VersioningEnable pulumi.BoolPtrInput
	// A website object(documented below).
	Website BucketWebsitePtrInput
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.
	Acl *string `pulumi:"acl"`
	// ACL XML body for multiple grant info. NOTE: this argument will overwrite `acl`. Check https://intl.cloud.tencent.com/document/product/436/7737 for more detail.
	AclBody *string `pulumi:"aclBody"`
	// The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
	Bucket string `pulumi:"bucket"`
	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRules []BucketCorsRule `pulumi:"corsRules"`
	// The server-side encryption algorithm to use. Valid value is `AES256`.
	EncryptionAlgorithm *string `pulumi:"encryptionAlgorithm"`
	// A configuration of object lifecycle management (documented below).
	LifecycleRules []BucketLifecycleRule `pulumi:"lifecycleRules"`
	// Indicate the access log of this bucket to be saved or not. Default is `false`. If set `true`, the access log will be saved with `logTargetBucket`. To enable log, the full access of log service must be granted. [Full Access Role Policy](https://intl.cloud.tencent.com/document/product/436/16920).
	LogEnable *bool `pulumi:"logEnable"`
	// The prefix log name which saves the access log of this bucket per 5 minutes. Eg. `MyLogPrefix/`. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`.
	LogPrefix *string `pulumi:"logPrefix"`
	// The target bucket name which saves the access log of this bucket per 5 minutes. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`. User must have full access on this bucket.
	LogTargetBucket *string `pulumi:"logTargetBucket"`
	// Indicates whether to create a bucket of multi available zone. NOTE: If set to true, the versioning must enable.
	MultiAz *bool `pulumi:"multiAz"`
	// Bucket Origin Domain settings.
	OriginDomainRules []BucketOriginDomainRule `pulumi:"originDomainRules"`
	// Bucket Origin-Pull settings.
	OriginPullRules []BucketOriginPullRule `pulumi:"originPullRules"`
	// Request initiator identifier, format: `qcs::cam::uin/<owneruin>:uin/<subuin>`. NOTE: only `versioningEnable` is true can configure this argument.
	ReplicaRole *string `pulumi:"replicaRole"`
	// List of replica rule. NOTE: only `versioningEnable` is true and `replicaRole` set can configure this argument.
	ReplicaRules []BucketReplicaRule `pulumi:"replicaRules"`
	// The tags of a bucket.
	Tags map[string]interface{} `pulumi:"tags"`
	// Enable bucket versioning.
	VersioningEnable *bool `pulumi:"versioningEnable"`
	// A website object(documented below).
	Website *BucketWebsite `pulumi:"website"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.
	Acl pulumi.StringPtrInput
	// ACL XML body for multiple grant info. NOTE: this argument will overwrite `acl`. Check https://intl.cloud.tencent.com/document/product/436/7737 for more detail.
	AclBody pulumi.StringPtrInput
	// The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
	Bucket pulumi.StringInput
	// A rule of Cross-Origin Resource Sharing (documented below).
	CorsRules BucketCorsRuleArrayInput
	// The server-side encryption algorithm to use. Valid value is `AES256`.
	EncryptionAlgorithm pulumi.StringPtrInput
	// A configuration of object lifecycle management (documented below).
	LifecycleRules BucketLifecycleRuleArrayInput
	// Indicate the access log of this bucket to be saved or not. Default is `false`. If set `true`, the access log will be saved with `logTargetBucket`. To enable log, the full access of log service must be granted. [Full Access Role Policy](https://intl.cloud.tencent.com/document/product/436/16920).
	LogEnable pulumi.BoolPtrInput
	// The prefix log name which saves the access log of this bucket per 5 minutes. Eg. `MyLogPrefix/`. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`.
	LogPrefix pulumi.StringPtrInput
	// The target bucket name which saves the access log of this bucket per 5 minutes. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`. User must have full access on this bucket.
	LogTargetBucket pulumi.StringPtrInput
	// Indicates whether to create a bucket of multi available zone. NOTE: If set to true, the versioning must enable.
	MultiAz pulumi.BoolPtrInput
	// Bucket Origin Domain settings.
	OriginDomainRules BucketOriginDomainRuleArrayInput
	// Bucket Origin-Pull settings.
	OriginPullRules BucketOriginPullRuleArrayInput
	// Request initiator identifier, format: `qcs::cam::uin/<owneruin>:uin/<subuin>`. NOTE: only `versioningEnable` is true can configure this argument.
	ReplicaRole pulumi.StringPtrInput
	// List of replica rule. NOTE: only `versioningEnable` is true and `replicaRole` set can configure this argument.
	ReplicaRules BucketReplicaRuleArrayInput
	// The tags of a bucket.
	Tags pulumi.MapInput
	// Enable bucket versioning.
	VersioningEnable pulumi.BoolPtrInput
	// A website object(documented below).
	Website BucketWebsitePtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

// BucketArrayInput is an input type that accepts BucketArray and BucketArrayOutput values.
// You can construct a concrete instance of `BucketArrayInput` via:
//
//	BucketArray{ BucketArgs{...} }
type BucketArrayInput interface {
	pulumi.Input

	ToBucketArrayOutput() BucketArrayOutput
	ToBucketArrayOutputWithContext(context.Context) BucketArrayOutput
}

type BucketArray []BucketInput

func (BucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (i BucketArray) ToBucketArrayOutput() BucketArrayOutput {
	return i.ToBucketArrayOutputWithContext(context.Background())
}

func (i BucketArray) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketArrayOutput)
}

// BucketMapInput is an input type that accepts BucketMap and BucketMapOutput values.
// You can construct a concrete instance of `BucketMapInput` via:
//
//	BucketMap{ "key": BucketArgs{...} }
type BucketMapInput interface {
	pulumi.Input

	ToBucketMapOutput() BucketMapOutput
	ToBucketMapOutputWithContext(context.Context) BucketMapOutput
}

type BucketMap map[string]BucketInput

func (BucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (i BucketMap) ToBucketMapOutput() BucketMapOutput {
	return i.ToBucketMapOutputWithContext(context.Background())
}

func (i BucketMap) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketMapOutput)
}

type BucketOutput struct{ *pulumi.OutputState }

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bucket)(nil)).Elem()
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

// The canned ACL to apply. Valid values: private, public-read, and public-read-write. Defaults to private.
func (o BucketOutput) Acl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.Acl }).(pulumi.StringPtrOutput)
}

// ACL XML body for multiple grant info. NOTE: this argument will overwrite `acl`. Check https://intl.cloud.tencent.com/document/product/436/7737 for more detail.
func (o BucketOutput) AclBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.AclBody }).(pulumi.StringPtrOutput)
}

// The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
func (o BucketOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// A rule of Cross-Origin Resource Sharing (documented below).
func (o BucketOutput) CorsRules() BucketCorsRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketCorsRuleArrayOutput { return v.CorsRules }).(BucketCorsRuleArrayOutput)
}

// The URL of this cos bucket.
func (o BucketOutput) CosBucketUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.CosBucketUrl }).(pulumi.StringOutput)
}

// The server-side encryption algorithm to use. Valid value is `AES256`.
func (o BucketOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.EncryptionAlgorithm }).(pulumi.StringPtrOutput)
}

// A configuration of object lifecycle management (documented below).
func (o BucketOutput) LifecycleRules() BucketLifecycleRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketLifecycleRuleArrayOutput { return v.LifecycleRules }).(BucketLifecycleRuleArrayOutput)
}

// Indicate the access log of this bucket to be saved or not. Default is `false`. If set `true`, the access log will be saved with `logTargetBucket`. To enable log, the full access of log service must be granted. [Full Access Role Policy](https://intl.cloud.tencent.com/document/product/436/16920).
func (o BucketOutput) LogEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.LogEnable }).(pulumi.BoolPtrOutput)
}

// The prefix log name which saves the access log of this bucket per 5 minutes. Eg. `MyLogPrefix/`. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`.
func (o BucketOutput) LogPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.LogPrefix }).(pulumi.StringOutput)
}

// The target bucket name which saves the access log of this bucket per 5 minutes. The log access file format is `logTargetBucket`/`logPrefix`{YYYY}/{MM}/{DD}/{time}_{random}_{index}.gz. Only valid when `logEnable` is `true`. User must have full access on this bucket.
func (o BucketOutput) LogTargetBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringOutput { return v.LogTargetBucket }).(pulumi.StringOutput)
}

// Indicates whether to create a bucket of multi available zone. NOTE: If set to true, the versioning must enable.
func (o BucketOutput) MultiAz() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.MultiAz }).(pulumi.BoolPtrOutput)
}

// Bucket Origin Domain settings.
func (o BucketOutput) OriginDomainRules() BucketOriginDomainRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketOriginDomainRuleArrayOutput { return v.OriginDomainRules }).(BucketOriginDomainRuleArrayOutput)
}

// Bucket Origin-Pull settings.
func (o BucketOutput) OriginPullRules() BucketOriginPullRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketOriginPullRuleArrayOutput { return v.OriginPullRules }).(BucketOriginPullRuleArrayOutput)
}

// Request initiator identifier, format: `qcs::cam::uin/<owneruin>:uin/<subuin>`. NOTE: only `versioningEnable` is true can configure this argument.
func (o BucketOutput) ReplicaRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.StringPtrOutput { return v.ReplicaRole }).(pulumi.StringPtrOutput)
}

// List of replica rule. NOTE: only `versioningEnable` is true and `replicaRole` set can configure this argument.
func (o BucketOutput) ReplicaRules() BucketReplicaRuleArrayOutput {
	return o.ApplyT(func(v *Bucket) BucketReplicaRuleArrayOutput { return v.ReplicaRules }).(BucketReplicaRuleArrayOutput)
}

// The tags of a bucket.
func (o BucketOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Bucket) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Enable bucket versioning.
func (o BucketOutput) VersioningEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bucket) pulumi.BoolPtrOutput { return v.VersioningEnable }).(pulumi.BoolPtrOutput)
}

// A website object(documented below).
func (o BucketOutput) Website() BucketWebsitePtrOutput {
	return o.ApplyT(func(v *Bucket) BucketWebsitePtrOutput { return v.Website }).(BucketWebsitePtrOutput)
}

type BucketArrayOutput struct{ *pulumi.OutputState }

func (BucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bucket)(nil)).Elem()
}

func (o BucketArrayOutput) ToBucketArrayOutput() BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) ToBucketArrayOutputWithContext(ctx context.Context) BucketArrayOutput {
	return o
}

func (o BucketArrayOutput) Index(i pulumi.IntInput) BucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].([]*Bucket)[vs[1].(int)]
	}).(BucketOutput)
}

type BucketMapOutput struct{ *pulumi.OutputState }

func (BucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bucket)(nil)).Elem()
}

func (o BucketMapOutput) ToBucketMapOutput() BucketMapOutput {
	return o
}

func (o BucketMapOutput) ToBucketMapOutputWithContext(ctx context.Context) BucketMapOutput {
	return o
}

func (o BucketMapOutput) MapIndex(k pulumi.StringInput) BucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bucket {
		return vs[0].(map[string]*Bucket)[vs[1].(string)]
	}).(BucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketInput)(nil)).Elem(), &Bucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketArrayInput)(nil)).Elem(), BucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketMapInput)(nil)).Elem(), BucketMap{})
	pulumi.RegisterOutputType(BucketOutput{})
	pulumi.RegisterOutputType(BucketArrayOutput{})
	pulumi.RegisterOutputType(BucketMapOutput{})
}
