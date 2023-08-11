// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tmp tke template attachment
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Images"
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Kubernetes"
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Images"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Kubernetes"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		defaultInstanceType := "SA1.MEDIUM2"
// 		if param := cfg.Get("defaultInstanceType"); param != "" {
// 			defaultInstanceType = param
// 		}
// 		availabilityZoneFirst := "ap-guangzhou-3"
// 		if param := cfg.Get("availabilityZoneFirst"); param != "" {
// 			availabilityZoneFirst = param
// 		}
// 		availabilityZoneSecond := "ap-guangzhou-4"
// 		if param := cfg.Get("availabilityZoneSecond"); param != "" {
// 			availabilityZoneSecond = param
// 		}
// 		exampleClusterCidr := "10.31.0.0/16"
// 		if param := cfg.Get("exampleClusterCidr"); param != "" {
// 			exampleClusterCidr = param
// 		}
// 		vpcOne, err := Vpc.GetSubnets(ctx, &vpc.GetSubnetsArgs{
// 			IsDefault:        pulumi.BoolRef(true),
// 			AvailabilityZone: pulumi.StringRef(availabilityZoneFirst),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		firstVpcId := vpcOne.InstanceLists[0].VpcId
// 		firstSubnetId := vpcOne.InstanceLists[0].SubnetId
// 		vpcTwo, err := Vpc.GetSubnets(ctx, &vpc.GetSubnetsArgs{
// 			IsDefault:        pulumi.BoolRef(true),
// 			AvailabilityZone: pulumi.StringRef(availabilityZoneSecond),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_ := vpcTwo.InstanceLists[0].VpcId
// 		secondSubnetId := vpcTwo.InstanceLists[0].SubnetId
// 		sg, err := Security.NewGroup(ctx, "sg", nil)
// 		if err != nil {
// 			return err
// 		}
// 		sgId := sg.ID()
// 		_default, err := Images.GetInstance(ctx, &images.GetInstanceArgs{
// 			ImageTypes: []string{
// 				"PUBLIC_IMAGE",
// 			},
// 			ImageNameRegex: pulumi.StringRef("Final"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		imageId := _default.ImageId
// 		_, err = Security.NewGroupLiteRule(ctx, "sgRule", &Security.GroupLiteRuleArgs{
// 			SecurityGroupId: sg.ID(),
// 			Ingresses: pulumi.StringArray{
// 				pulumi.String("ACCEPT#10.0.0.0/16#ALL#ALL"),
// 				pulumi.String("ACCEPT#172.16.0.0/22#ALL#ALL"),
// 				pulumi.String("DROP#0.0.0.0/0#ALL#ALL"),
// 			},
// 			Egresses: pulumi.StringArray{
// 				pulumi.String("ACCEPT#172.16.0.0/22#ALL#ALL"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example, err := Kubernetes.NewCluster(ctx, "example", &Kubernetes.ClusterArgs{
// 			VpcId:                        pulumi.String(firstVpcId),
// 			ClusterCidr:                  pulumi.String(exampleClusterCidr),
// 			ClusterMaxPodNum:             pulumi.Int(32),
// 			ClusterName:                  pulumi.String("tf_example_cluster"),
// 			ClusterDesc:                  pulumi.String("example for tke cluster"),
// 			ClusterMaxServiceNum:         pulumi.Int(32),
// 			ClusterInternet:              pulumi.Bool(false),
// 			ClusterInternetSecurityGroup: pulumi.String(sgId),
// 			ClusterVersion:               pulumi.String("1.22.5"),
// 			ClusterDeployType:            pulumi.String("MANAGED_CLUSTER"),
// 			WorkerConfigs: kubernetes.ClusterWorkerConfigArray{
// 				&kubernetes.ClusterWorkerConfigArgs{
// 					Count:                   pulumi.Int(1),
// 					AvailabilityZone:        pulumi.String(availabilityZoneFirst),
// 					InstanceType:            pulumi.String(defaultInstanceType),
// 					SystemDiskType:          pulumi.String("CLOUD_SSD"),
// 					SystemDiskSize:          pulumi.Int(60),
// 					InternetChargeType:      pulumi.String("TRAFFIC_POSTPAID_BY_HOUR"),
// 					InternetMaxBandwidthOut: pulumi.Int(100),
// 					PublicIpAssigned:        pulumi.Bool(true),
// 					SubnetId:                pulumi.String(firstSubnetId),
// 					ImgId:                   pulumi.String(imageId),
// 					DataDisks: kubernetes.ClusterWorkerConfigDataDiskArray{
// 						&kubernetes.ClusterWorkerConfigDataDiskArgs{
// 							DiskType: pulumi.String("CLOUD_PREMIUM"),
// 							DiskSize: pulumi.Int(50),
// 						},
// 					},
// 					EnhancedSecurityService: pulumi.Bool(false),
// 					EnhancedMonitorService:  pulumi.Bool(false),
// 					UserData:                pulumi.String("dGVzdA=="),
// 					Password:                pulumi.String("ZZXXccvv1212"),
// 				},
// 				&kubernetes.ClusterWorkerConfigArgs{
// 					Count:                   pulumi.Int(1),
// 					AvailabilityZone:        pulumi.String(availabilityZoneSecond),
// 					InstanceType:            pulumi.String(defaultInstanceType),
// 					SystemDiskType:          pulumi.String("CLOUD_SSD"),
// 					SystemDiskSize:          pulumi.Int(60),
// 					InternetChargeType:      pulumi.String("TRAFFIC_POSTPAID_BY_HOUR"),
// 					InternetMaxBandwidthOut: pulumi.Int(100),
// 					PublicIpAssigned:        pulumi.Bool(true),
// 					SubnetId:                pulumi.String(secondSubnetId),
// 					DataDisks: kubernetes.ClusterWorkerConfigDataDiskArray{
// 						&kubernetes.ClusterWorkerConfigDataDiskArgs{
// 							DiskType: pulumi.String("CLOUD_PREMIUM"),
// 							DiskSize: pulumi.Int(50),
// 						},
// 					},
// 					EnhancedSecurityService: pulumi.Bool(false),
// 					EnhancedMonitorService:  pulumi.Bool(false),
// 					UserData:                pulumi.String("dGVzdA=="),
// 					CamRoleName:             pulumi.String("CVM_QcsRole"),
// 					Password:                pulumi.String("ZZXXccvv1212"),
// 				},
// 			},
// 			Labels: pulumi.AnyMap{
// 				"test1": pulumi.Any("test1"),
// 				"test2": pulumi.Any("test2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		zone := "ap-guangzhou"
// 		if param := cfg.Get("zone"); param != "" {
// 			zone = param
// 		}
// 		clusterType := "tke"
// 		if param := cfg.Get("clusterType"); param != "" {
// 			clusterType = param
// 		}
// 		fooTmpInstance, err := Monitor.NewTmpInstance(ctx, "fooTmpInstance", &Monitor.TmpInstanceArgs{
// 			InstanceName:      pulumi.String("tf-tmp-instance"),
// 			VpcId:             pulumi.String(firstVpcId),
// 			SubnetId:          pulumi.String(firstSubnetId),
// 			DataRetentionTime: pulumi.Int(30),
// 			Zone:              pulumi.String(availabilityZoneSecond),
// 			Tags: pulumi.AnyMap{
// 				"createdBy": pulumi.Any("terraform"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooTmpTkeClusterAgent, err := Monitor.NewTmpTkeClusterAgent(ctx, "fooTmpTkeClusterAgent", &Monitor.TmpTkeClusterAgentArgs{
// 			InstanceId: fooTmpInstance.ID(),
// 			Agents: &monitor.TmpTkeClusterAgentAgentsArgs{
// 				Region:         pulumi.String(zone),
// 				ClusterType:    pulumi.String(clusterType),
// 				ClusterId:      example.ID(),
// 				EnableExternal: pulumi.Bool(false),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooTmpTkeTemplate, err := Monitor.NewTmpTkeTemplate(ctx, "fooTmpTkeTemplate", &Monitor.TmpTkeTemplateArgs{
// 			Template: &monitor.TmpTkeTemplateTemplateArgs{
// 				Name:     pulumi.String("tf-template"),
// 				Level:    pulumi.String("cluster"),
// 				Describe: pulumi.String("template"),
// 				ServiceMonitors: monitor.TmpTkeTemplateTemplateServiceMonitorArray{
// 					&monitor.TmpTkeTemplateTemplateServiceMonitorArgs{
// 						Name:   pulumi.String("tf-ServiceMonitor"),
// 						Config: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "apiVersion: monitoring.coreos.com/v1\n", "kind: ServiceMonitor\n", "metadata:\n", "  name: example-service-monitor\n", "  namespace: monitoring\n", "  labels:\n", "    k8s-app: example-service\n", "spec:\n", "  selector:\n", "    matchLabels:\n", "      k8s-app: example-service\n", "  namespaceSelector:\n", "    matchNames:\n", "      - default\n", "  endpoints:\n", "  - port: http-metrics\n", "    interval: 30s\n", "    path: /metrics\n", "    scheme: http\n", "    bearerTokenFile: /var/run/secrets/kubernetes.io/serviceaccount/token\n", "    tlsConfig:\n", "      insecureSkipVerify: true\n")),
// 					},
// 				},
// 				PodMonitors: monitor.TmpTkeTemplateTemplatePodMonitorArray{
// 					&monitor.TmpTkeTemplateTemplatePodMonitorArgs{
// 						Name:   pulumi.String("tf-PodMonitors"),
// 						Config: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "apiVersion: monitoring.coreos.com/v1\n", "kind: PodMonitor\n", "metadata:\n", "  name: example-pod-monitor\n", "  namespace: monitoring\n", "  labels:\n", "    k8s-app: example-pod\n", "spec:\n", "  selector:\n", "    matchLabels:\n", "      k8s-app: example-pod\n", "  namespaceSelector:\n", "    matchNames:\n", "      - default\n", "  podMetricsEndpoints:\n", "  - port: http-metrics\n", "    interval: 30s\n", "    path: /metrics\n", "    scheme: http\n", "    bearerTokenFile: /var/run/secrets/kubernetes.io/serviceaccount/token\n", "    tlsConfig:\n", "      insecureSkipVerify: true\n")),
// 					},
// 					&monitor.TmpTkeTemplateTemplatePodMonitorArgs{
// 						Name:   pulumi.String("tf-RawJobs"),
// 						Config: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v", "scrape_configs:\n", "  - job_name: 'example-job'\n", "    scrape_interval: 30s\n", "    static_configs:\n", "      - targets: ['example-service.default.svc.cluster.local:8080']\n", "    metrics_path: /metrics\n", "    scheme: http\n", "    bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token\n", "    tls_config:\n", "      insecure_skip_verify: true\n")),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Monitor.NewTmpTkeTemplateAttachment(ctx, "tempAttachment", &Monitor.TmpTkeTemplateAttachmentArgs{
// 			TemplateId: fooTmpTkeTemplate.ID(),
// 			Targets: &monitor.TmpTkeTemplateAttachmentTargetsArgs{
// 				ClusterType: pulumi.String(clusterType),
// 				ClusterId:   example.ID(),
// 				Region:      pulumi.String(zone),
// 				InstanceId:  fooTmpInstance.ID(),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			fooTmpTkeClusterAgent,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type TmpTkeTemplateAttachment struct {
	pulumi.CustomResourceState

	// Sync target details.
	Targets TmpTkeTemplateAttachmentTargetsOutput `pulumi:"targets"`
	// The ID of the template, which is used for the outgoing reference.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
}

// NewTmpTkeTemplateAttachment registers a new resource with the given unique name, arguments, and options.
func NewTmpTkeTemplateAttachment(ctx *pulumi.Context,
	name string, args *TmpTkeTemplateAttachmentArgs, opts ...pulumi.ResourceOption) (*TmpTkeTemplateAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Targets == nil {
		return nil, errors.New("invalid value for required argument 'Targets'")
	}
	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TmpTkeTemplateAttachment
	err := ctx.RegisterResource("tencentcloud:Monitor/tmpTkeTemplateAttachment:TmpTkeTemplateAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTmpTkeTemplateAttachment gets an existing TmpTkeTemplateAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTmpTkeTemplateAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TmpTkeTemplateAttachmentState, opts ...pulumi.ResourceOption) (*TmpTkeTemplateAttachment, error) {
	var resource TmpTkeTemplateAttachment
	err := ctx.ReadResource("tencentcloud:Monitor/tmpTkeTemplateAttachment:TmpTkeTemplateAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TmpTkeTemplateAttachment resources.
type tmpTkeTemplateAttachmentState struct {
	// Sync target details.
	Targets *TmpTkeTemplateAttachmentTargets `pulumi:"targets"`
	// The ID of the template, which is used for the outgoing reference.
	TemplateId *string `pulumi:"templateId"`
}

type TmpTkeTemplateAttachmentState struct {
	// Sync target details.
	Targets TmpTkeTemplateAttachmentTargetsPtrInput
	// The ID of the template, which is used for the outgoing reference.
	TemplateId pulumi.StringPtrInput
}

func (TmpTkeTemplateAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpTkeTemplateAttachmentState)(nil)).Elem()
}

type tmpTkeTemplateAttachmentArgs struct {
	// Sync target details.
	Targets TmpTkeTemplateAttachmentTargets `pulumi:"targets"`
	// The ID of the template, which is used for the outgoing reference.
	TemplateId string `pulumi:"templateId"`
}

// The set of arguments for constructing a TmpTkeTemplateAttachment resource.
type TmpTkeTemplateAttachmentArgs struct {
	// Sync target details.
	Targets TmpTkeTemplateAttachmentTargetsInput
	// The ID of the template, which is used for the outgoing reference.
	TemplateId pulumi.StringInput
}

func (TmpTkeTemplateAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpTkeTemplateAttachmentArgs)(nil)).Elem()
}

type TmpTkeTemplateAttachmentInput interface {
	pulumi.Input

	ToTmpTkeTemplateAttachmentOutput() TmpTkeTemplateAttachmentOutput
	ToTmpTkeTemplateAttachmentOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentOutput
}

func (*TmpTkeTemplateAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpTkeTemplateAttachment)(nil)).Elem()
}

func (i *TmpTkeTemplateAttachment) ToTmpTkeTemplateAttachmentOutput() TmpTkeTemplateAttachmentOutput {
	return i.ToTmpTkeTemplateAttachmentOutputWithContext(context.Background())
}

func (i *TmpTkeTemplateAttachment) ToTmpTkeTemplateAttachmentOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeTemplateAttachmentOutput)
}

// TmpTkeTemplateAttachmentArrayInput is an input type that accepts TmpTkeTemplateAttachmentArray and TmpTkeTemplateAttachmentArrayOutput values.
// You can construct a concrete instance of `TmpTkeTemplateAttachmentArrayInput` via:
//
//          TmpTkeTemplateAttachmentArray{ TmpTkeTemplateAttachmentArgs{...} }
type TmpTkeTemplateAttachmentArrayInput interface {
	pulumi.Input

	ToTmpTkeTemplateAttachmentArrayOutput() TmpTkeTemplateAttachmentArrayOutput
	ToTmpTkeTemplateAttachmentArrayOutputWithContext(context.Context) TmpTkeTemplateAttachmentArrayOutput
}

type TmpTkeTemplateAttachmentArray []TmpTkeTemplateAttachmentInput

func (TmpTkeTemplateAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpTkeTemplateAttachment)(nil)).Elem()
}

func (i TmpTkeTemplateAttachmentArray) ToTmpTkeTemplateAttachmentArrayOutput() TmpTkeTemplateAttachmentArrayOutput {
	return i.ToTmpTkeTemplateAttachmentArrayOutputWithContext(context.Background())
}

func (i TmpTkeTemplateAttachmentArray) ToTmpTkeTemplateAttachmentArrayOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeTemplateAttachmentArrayOutput)
}

// TmpTkeTemplateAttachmentMapInput is an input type that accepts TmpTkeTemplateAttachmentMap and TmpTkeTemplateAttachmentMapOutput values.
// You can construct a concrete instance of `TmpTkeTemplateAttachmentMapInput` via:
//
//          TmpTkeTemplateAttachmentMap{ "key": TmpTkeTemplateAttachmentArgs{...} }
type TmpTkeTemplateAttachmentMapInput interface {
	pulumi.Input

	ToTmpTkeTemplateAttachmentMapOutput() TmpTkeTemplateAttachmentMapOutput
	ToTmpTkeTemplateAttachmentMapOutputWithContext(context.Context) TmpTkeTemplateAttachmentMapOutput
}

type TmpTkeTemplateAttachmentMap map[string]TmpTkeTemplateAttachmentInput

func (TmpTkeTemplateAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpTkeTemplateAttachment)(nil)).Elem()
}

func (i TmpTkeTemplateAttachmentMap) ToTmpTkeTemplateAttachmentMapOutput() TmpTkeTemplateAttachmentMapOutput {
	return i.ToTmpTkeTemplateAttachmentMapOutputWithContext(context.Background())
}

func (i TmpTkeTemplateAttachmentMap) ToTmpTkeTemplateAttachmentMapOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpTkeTemplateAttachmentMapOutput)
}

type TmpTkeTemplateAttachmentOutput struct{ *pulumi.OutputState }

func (TmpTkeTemplateAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpTkeTemplateAttachment)(nil)).Elem()
}

func (o TmpTkeTemplateAttachmentOutput) ToTmpTkeTemplateAttachmentOutput() TmpTkeTemplateAttachmentOutput {
	return o
}

func (o TmpTkeTemplateAttachmentOutput) ToTmpTkeTemplateAttachmentOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentOutput {
	return o
}

// Sync target details.
func (o TmpTkeTemplateAttachmentOutput) Targets() TmpTkeTemplateAttachmentTargetsOutput {
	return o.ApplyT(func(v *TmpTkeTemplateAttachment) TmpTkeTemplateAttachmentTargetsOutput { return v.Targets }).(TmpTkeTemplateAttachmentTargetsOutput)
}

// The ID of the template, which is used for the outgoing reference.
func (o TmpTkeTemplateAttachmentOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpTkeTemplateAttachment) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

type TmpTkeTemplateAttachmentArrayOutput struct{ *pulumi.OutputState }

func (TmpTkeTemplateAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpTkeTemplateAttachment)(nil)).Elem()
}

func (o TmpTkeTemplateAttachmentArrayOutput) ToTmpTkeTemplateAttachmentArrayOutput() TmpTkeTemplateAttachmentArrayOutput {
	return o
}

func (o TmpTkeTemplateAttachmentArrayOutput) ToTmpTkeTemplateAttachmentArrayOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentArrayOutput {
	return o
}

func (o TmpTkeTemplateAttachmentArrayOutput) Index(i pulumi.IntInput) TmpTkeTemplateAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TmpTkeTemplateAttachment {
		return vs[0].([]*TmpTkeTemplateAttachment)[vs[1].(int)]
	}).(TmpTkeTemplateAttachmentOutput)
}

type TmpTkeTemplateAttachmentMapOutput struct{ *pulumi.OutputState }

func (TmpTkeTemplateAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpTkeTemplateAttachment)(nil)).Elem()
}

func (o TmpTkeTemplateAttachmentMapOutput) ToTmpTkeTemplateAttachmentMapOutput() TmpTkeTemplateAttachmentMapOutput {
	return o
}

func (o TmpTkeTemplateAttachmentMapOutput) ToTmpTkeTemplateAttachmentMapOutputWithContext(ctx context.Context) TmpTkeTemplateAttachmentMapOutput {
	return o
}

func (o TmpTkeTemplateAttachmentMapOutput) MapIndex(k pulumi.StringInput) TmpTkeTemplateAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TmpTkeTemplateAttachment {
		return vs[0].(map[string]*TmpTkeTemplateAttachment)[vs[1].(string)]
	}).(TmpTkeTemplateAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeTemplateAttachmentInput)(nil)).Elem(), &TmpTkeTemplateAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeTemplateAttachmentArrayInput)(nil)).Elem(), TmpTkeTemplateAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpTkeTemplateAttachmentMapInput)(nil)).Elem(), TmpTkeTemplateAttachmentMap{})
	pulumi.RegisterOutputType(TmpTkeTemplateAttachmentOutput{})
	pulumi.RegisterOutputType(TmpTkeTemplateAttachmentArrayOutput{})
	pulumi.RegisterOutputType(TmpTkeTemplateAttachmentMapOutput{})
}
