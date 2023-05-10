// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tem

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tem workload
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tem"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tem.NewWorkload(ctx, "workload", &Tem.WorkloadArgs{
// 			ApplicationId: pulumi.String("app-j4d3x6kj"),
// 			CpuSpec:       pulumi.Float64(1),
// 			DeployMode:    pulumi.String("IMAGE"),
// 			DeployVersion: pulumi.String("hello-world"),
// 			EnvironmentId: pulumi.String("en-85377m6j"),
// 			ImgRepo:       pulumi.String("tem_demo/tem_demo"),
// 			InitPodNum:    pulumi.Int(1),
// 			MemorySpec:    pulumi.Float64(1),
// 			RepoServer:    pulumi.String("ccr.ccs.tencentyun.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// tem workload can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Tem/workload:Workload workload envirnomentId#applicationId
// ```
type Workload struct {
	pulumi.CustomResourceState

	// application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// cpu.
	CpuSpec pulumi.Float64Output `pulumi:"cpuSpec"`
	// deploy mode, support IMAGE.
	DeployMode pulumi.StringOutput `pulumi:"deployMode"`
	// deploy strategy.
	DeployStrategyConf WorkloadDeployStrategyConfPtrOutput `pulumi:"deployStrategyConf"`
	// deploy version.
	DeployVersion pulumi.StringOutput `pulumi:"deployVersion"`
	// .
	EnvConfs WorkloadEnvConfArrayOutput `pulumi:"envConfs"`
	// environment ID.
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// repository name.
	ImgRepo pulumi.StringOutput `pulumi:"imgRepo"`
	// initial pod number.
	InitPodNum pulumi.IntOutput `pulumi:"initPodNum"`
	// liveness config.
	Liveness WorkloadLivenessPtrOutput `pulumi:"liveness"`
	// mem.
	MemorySpec pulumi.Float64Output `pulumi:"memorySpec"`
	// mem.
	PostStart pulumi.StringPtrOutput `pulumi:"postStart"`
	// mem.
	PreStop pulumi.StringPtrOutput `pulumi:"preStop"`
	// .
	Readiness WorkloadReadinessPtrOutput `pulumi:"readiness"`
	// repo server addr when deploy by image.
	RepoServer pulumi.StringPtrOutput `pulumi:"repoServer"`
	// repo type when deploy: 0: tcr personal; 1: tcr enterprise; 2: public repository; 3: tem host tcr; 4: demo repo.
	RepoType pulumi.IntPtrOutput `pulumi:"repoType"`
	// security groups.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// .
	StartupProbe WorkloadStartupProbePtrOutput `pulumi:"startupProbe"`
	// storage configuration.
	StorageConfs WorkloadStorageConfArrayOutput `pulumi:"storageConfs"`
	// storage mount configuration.
	StorageMountConfs WorkloadStorageMountConfArrayOutput `pulumi:"storageMountConfs"`
	// tcr instance id when deploy by image.
	TcrInstanceId pulumi.StringPtrOutput `pulumi:"tcrInstanceId"`
}

// NewWorkload registers a new resource with the given unique name, arguments, and options.
func NewWorkload(ctx *pulumi.Context,
	name string, args *WorkloadArgs, opts ...pulumi.ResourceOption) (*Workload, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.CpuSpec == nil {
		return nil, errors.New("invalid value for required argument 'CpuSpec'")
	}
	if args.DeployMode == nil {
		return nil, errors.New("invalid value for required argument 'DeployMode'")
	}
	if args.DeployVersion == nil {
		return nil, errors.New("invalid value for required argument 'DeployVersion'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.ImgRepo == nil {
		return nil, errors.New("invalid value for required argument 'ImgRepo'")
	}
	if args.InitPodNum == nil {
		return nil, errors.New("invalid value for required argument 'InitPodNum'")
	}
	if args.MemorySpec == nil {
		return nil, errors.New("invalid value for required argument 'MemorySpec'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Workload
	err := ctx.RegisterResource("tencentcloud:Tem/workload:Workload", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkload gets an existing Workload resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkload(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadState, opts ...pulumi.ResourceOption) (*Workload, error) {
	var resource Workload
	err := ctx.ReadResource("tencentcloud:Tem/workload:Workload", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workload resources.
type workloadState struct {
	// application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// cpu.
	CpuSpec *float64 `pulumi:"cpuSpec"`
	// deploy mode, support IMAGE.
	DeployMode *string `pulumi:"deployMode"`
	// deploy strategy.
	DeployStrategyConf *WorkloadDeployStrategyConf `pulumi:"deployStrategyConf"`
	// deploy version.
	DeployVersion *string `pulumi:"deployVersion"`
	// .
	EnvConfs []WorkloadEnvConf `pulumi:"envConfs"`
	// environment ID.
	EnvironmentId *string `pulumi:"environmentId"`
	// repository name.
	ImgRepo *string `pulumi:"imgRepo"`
	// initial pod number.
	InitPodNum *int `pulumi:"initPodNum"`
	// liveness config.
	Liveness *WorkloadLiveness `pulumi:"liveness"`
	// mem.
	MemorySpec *float64 `pulumi:"memorySpec"`
	// mem.
	PostStart *string `pulumi:"postStart"`
	// mem.
	PreStop *string `pulumi:"preStop"`
	// .
	Readiness *WorkloadReadiness `pulumi:"readiness"`
	// repo server addr when deploy by image.
	RepoServer *string `pulumi:"repoServer"`
	// repo type when deploy: 0: tcr personal; 1: tcr enterprise; 2: public repository; 3: tem host tcr; 4: demo repo.
	RepoType *int `pulumi:"repoType"`
	// security groups.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// .
	StartupProbe *WorkloadStartupProbe `pulumi:"startupProbe"`
	// storage configuration.
	StorageConfs []WorkloadStorageConf `pulumi:"storageConfs"`
	// storage mount configuration.
	StorageMountConfs []WorkloadStorageMountConf `pulumi:"storageMountConfs"`
	// tcr instance id when deploy by image.
	TcrInstanceId *string `pulumi:"tcrInstanceId"`
}

type WorkloadState struct {
	// application ID.
	ApplicationId pulumi.StringPtrInput
	// cpu.
	CpuSpec pulumi.Float64PtrInput
	// deploy mode, support IMAGE.
	DeployMode pulumi.StringPtrInput
	// deploy strategy.
	DeployStrategyConf WorkloadDeployStrategyConfPtrInput
	// deploy version.
	DeployVersion pulumi.StringPtrInput
	// .
	EnvConfs WorkloadEnvConfArrayInput
	// environment ID.
	EnvironmentId pulumi.StringPtrInput
	// repository name.
	ImgRepo pulumi.StringPtrInput
	// initial pod number.
	InitPodNum pulumi.IntPtrInput
	// liveness config.
	Liveness WorkloadLivenessPtrInput
	// mem.
	MemorySpec pulumi.Float64PtrInput
	// mem.
	PostStart pulumi.StringPtrInput
	// mem.
	PreStop pulumi.StringPtrInput
	// .
	Readiness WorkloadReadinessPtrInput
	// repo server addr when deploy by image.
	RepoServer pulumi.StringPtrInput
	// repo type when deploy: 0: tcr personal; 1: tcr enterprise; 2: public repository; 3: tem host tcr; 4: demo repo.
	RepoType pulumi.IntPtrInput
	// security groups.
	SecurityGroupIds pulumi.StringArrayInput
	// .
	StartupProbe WorkloadStartupProbePtrInput
	// storage configuration.
	StorageConfs WorkloadStorageConfArrayInput
	// storage mount configuration.
	StorageMountConfs WorkloadStorageMountConfArrayInput
	// tcr instance id when deploy by image.
	TcrInstanceId pulumi.StringPtrInput
}

func (WorkloadState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadState)(nil)).Elem()
}

type workloadArgs struct {
	// application ID.
	ApplicationId string `pulumi:"applicationId"`
	// cpu.
	CpuSpec float64 `pulumi:"cpuSpec"`
	// deploy mode, support IMAGE.
	DeployMode string `pulumi:"deployMode"`
	// deploy strategy.
	DeployStrategyConf *WorkloadDeployStrategyConf `pulumi:"deployStrategyConf"`
	// deploy version.
	DeployVersion string `pulumi:"deployVersion"`
	// .
	EnvConfs []WorkloadEnvConf `pulumi:"envConfs"`
	// environment ID.
	EnvironmentId string `pulumi:"environmentId"`
	// repository name.
	ImgRepo string `pulumi:"imgRepo"`
	// initial pod number.
	InitPodNum int `pulumi:"initPodNum"`
	// liveness config.
	Liveness *WorkloadLiveness `pulumi:"liveness"`
	// mem.
	MemorySpec float64 `pulumi:"memorySpec"`
	// mem.
	PostStart *string `pulumi:"postStart"`
	// mem.
	PreStop *string `pulumi:"preStop"`
	// .
	Readiness *WorkloadReadiness `pulumi:"readiness"`
	// repo server addr when deploy by image.
	RepoServer *string `pulumi:"repoServer"`
	// repo type when deploy: 0: tcr personal; 1: tcr enterprise; 2: public repository; 3: tem host tcr; 4: demo repo.
	RepoType *int `pulumi:"repoType"`
	// security groups.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// .
	StartupProbe *WorkloadStartupProbe `pulumi:"startupProbe"`
	// storage configuration.
	StorageConfs []WorkloadStorageConf `pulumi:"storageConfs"`
	// storage mount configuration.
	StorageMountConfs []WorkloadStorageMountConf `pulumi:"storageMountConfs"`
	// tcr instance id when deploy by image.
	TcrInstanceId *string `pulumi:"tcrInstanceId"`
}

// The set of arguments for constructing a Workload resource.
type WorkloadArgs struct {
	// application ID.
	ApplicationId pulumi.StringInput
	// cpu.
	CpuSpec pulumi.Float64Input
	// deploy mode, support IMAGE.
	DeployMode pulumi.StringInput
	// deploy strategy.
	DeployStrategyConf WorkloadDeployStrategyConfPtrInput
	// deploy version.
	DeployVersion pulumi.StringInput
	// .
	EnvConfs WorkloadEnvConfArrayInput
	// environment ID.
	EnvironmentId pulumi.StringInput
	// repository name.
	ImgRepo pulumi.StringInput
	// initial pod number.
	InitPodNum pulumi.IntInput
	// liveness config.
	Liveness WorkloadLivenessPtrInput
	// mem.
	MemorySpec pulumi.Float64Input
	// mem.
	PostStart pulumi.StringPtrInput
	// mem.
	PreStop pulumi.StringPtrInput
	// .
	Readiness WorkloadReadinessPtrInput
	// repo server addr when deploy by image.
	RepoServer pulumi.StringPtrInput
	// repo type when deploy: 0: tcr personal; 1: tcr enterprise; 2: public repository; 3: tem host tcr; 4: demo repo.
	RepoType pulumi.IntPtrInput
	// security groups.
	SecurityGroupIds pulumi.StringArrayInput
	// .
	StartupProbe WorkloadStartupProbePtrInput
	// storage configuration.
	StorageConfs WorkloadStorageConfArrayInput
	// storage mount configuration.
	StorageMountConfs WorkloadStorageMountConfArrayInput
	// tcr instance id when deploy by image.
	TcrInstanceId pulumi.StringPtrInput
}

func (WorkloadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadArgs)(nil)).Elem()
}

type WorkloadInput interface {
	pulumi.Input

	ToWorkloadOutput() WorkloadOutput
	ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput
}

func (*Workload) ElementType() reflect.Type {
	return reflect.TypeOf((**Workload)(nil)).Elem()
}

func (i *Workload) ToWorkloadOutput() WorkloadOutput {
	return i.ToWorkloadOutputWithContext(context.Background())
}

func (i *Workload) ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadOutput)
}

// WorkloadArrayInput is an input type that accepts WorkloadArray and WorkloadArrayOutput values.
// You can construct a concrete instance of `WorkloadArrayInput` via:
//
//          WorkloadArray{ WorkloadArgs{...} }
type WorkloadArrayInput interface {
	pulumi.Input

	ToWorkloadArrayOutput() WorkloadArrayOutput
	ToWorkloadArrayOutputWithContext(context.Context) WorkloadArrayOutput
}

type WorkloadArray []WorkloadInput

func (WorkloadArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workload)(nil)).Elem()
}

func (i WorkloadArray) ToWorkloadArrayOutput() WorkloadArrayOutput {
	return i.ToWorkloadArrayOutputWithContext(context.Background())
}

func (i WorkloadArray) ToWorkloadArrayOutputWithContext(ctx context.Context) WorkloadArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadArrayOutput)
}

// WorkloadMapInput is an input type that accepts WorkloadMap and WorkloadMapOutput values.
// You can construct a concrete instance of `WorkloadMapInput` via:
//
//          WorkloadMap{ "key": WorkloadArgs{...} }
type WorkloadMapInput interface {
	pulumi.Input

	ToWorkloadMapOutput() WorkloadMapOutput
	ToWorkloadMapOutputWithContext(context.Context) WorkloadMapOutput
}

type WorkloadMap map[string]WorkloadInput

func (WorkloadMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workload)(nil)).Elem()
}

func (i WorkloadMap) ToWorkloadMapOutput() WorkloadMapOutput {
	return i.ToWorkloadMapOutputWithContext(context.Background())
}

func (i WorkloadMap) ToWorkloadMapOutputWithContext(ctx context.Context) WorkloadMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadMapOutput)
}

type WorkloadOutput struct{ *pulumi.OutputState }

func (WorkloadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workload)(nil)).Elem()
}

func (o WorkloadOutput) ToWorkloadOutput() WorkloadOutput {
	return o
}

func (o WorkloadOutput) ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput {
	return o
}

// application ID.
func (o WorkloadOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// cpu.
func (o WorkloadOutput) CpuSpec() pulumi.Float64Output {
	return o.ApplyT(func(v *Workload) pulumi.Float64Output { return v.CpuSpec }).(pulumi.Float64Output)
}

// deploy mode, support IMAGE.
func (o WorkloadOutput) DeployMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringOutput { return v.DeployMode }).(pulumi.StringOutput)
}

// deploy strategy.
func (o WorkloadOutput) DeployStrategyConf() WorkloadDeployStrategyConfPtrOutput {
	return o.ApplyT(func(v *Workload) WorkloadDeployStrategyConfPtrOutput { return v.DeployStrategyConf }).(WorkloadDeployStrategyConfPtrOutput)
}

// deploy version.
func (o WorkloadOutput) DeployVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringOutput { return v.DeployVersion }).(pulumi.StringOutput)
}

// .
func (o WorkloadOutput) EnvConfs() WorkloadEnvConfArrayOutput {
	return o.ApplyT(func(v *Workload) WorkloadEnvConfArrayOutput { return v.EnvConfs }).(WorkloadEnvConfArrayOutput)
}

// environment ID.
func (o WorkloadOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// repository name.
func (o WorkloadOutput) ImgRepo() pulumi.StringOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringOutput { return v.ImgRepo }).(pulumi.StringOutput)
}

// initial pod number.
func (o WorkloadOutput) InitPodNum() pulumi.IntOutput {
	return o.ApplyT(func(v *Workload) pulumi.IntOutput { return v.InitPodNum }).(pulumi.IntOutput)
}

// liveness config.
func (o WorkloadOutput) Liveness() WorkloadLivenessPtrOutput {
	return o.ApplyT(func(v *Workload) WorkloadLivenessPtrOutput { return v.Liveness }).(WorkloadLivenessPtrOutput)
}

// mem.
func (o WorkloadOutput) MemorySpec() pulumi.Float64Output {
	return o.ApplyT(func(v *Workload) pulumi.Float64Output { return v.MemorySpec }).(pulumi.Float64Output)
}

// mem.
func (o WorkloadOutput) PostStart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringPtrOutput { return v.PostStart }).(pulumi.StringPtrOutput)
}

// mem.
func (o WorkloadOutput) PreStop() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringPtrOutput { return v.PreStop }).(pulumi.StringPtrOutput)
}

// .
func (o WorkloadOutput) Readiness() WorkloadReadinessPtrOutput {
	return o.ApplyT(func(v *Workload) WorkloadReadinessPtrOutput { return v.Readiness }).(WorkloadReadinessPtrOutput)
}

// repo server addr when deploy by image.
func (o WorkloadOutput) RepoServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringPtrOutput { return v.RepoServer }).(pulumi.StringPtrOutput)
}

// repo type when deploy: 0: tcr personal; 1: tcr enterprise; 2: public repository; 3: tem host tcr; 4: demo repo.
func (o WorkloadOutput) RepoType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Workload) pulumi.IntPtrOutput { return v.RepoType }).(pulumi.IntPtrOutput)
}

// security groups.
func (o WorkloadOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// .
func (o WorkloadOutput) StartupProbe() WorkloadStartupProbePtrOutput {
	return o.ApplyT(func(v *Workload) WorkloadStartupProbePtrOutput { return v.StartupProbe }).(WorkloadStartupProbePtrOutput)
}

// storage configuration.
func (o WorkloadOutput) StorageConfs() WorkloadStorageConfArrayOutput {
	return o.ApplyT(func(v *Workload) WorkloadStorageConfArrayOutput { return v.StorageConfs }).(WorkloadStorageConfArrayOutput)
}

// storage mount configuration.
func (o WorkloadOutput) StorageMountConfs() WorkloadStorageMountConfArrayOutput {
	return o.ApplyT(func(v *Workload) WorkloadStorageMountConfArrayOutput { return v.StorageMountConfs }).(WorkloadStorageMountConfArrayOutput)
}

// tcr instance id when deploy by image.
func (o WorkloadOutput) TcrInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workload) pulumi.StringPtrOutput { return v.TcrInstanceId }).(pulumi.StringPtrOutput)
}

type WorkloadArrayOutput struct{ *pulumi.OutputState }

func (WorkloadArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Workload)(nil)).Elem()
}

func (o WorkloadArrayOutput) ToWorkloadArrayOutput() WorkloadArrayOutput {
	return o
}

func (o WorkloadArrayOutput) ToWorkloadArrayOutputWithContext(ctx context.Context) WorkloadArrayOutput {
	return o
}

func (o WorkloadArrayOutput) Index(i pulumi.IntInput) WorkloadOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Workload {
		return vs[0].([]*Workload)[vs[1].(int)]
	}).(WorkloadOutput)
}

type WorkloadMapOutput struct{ *pulumi.OutputState }

func (WorkloadMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Workload)(nil)).Elem()
}

func (o WorkloadMapOutput) ToWorkloadMapOutput() WorkloadMapOutput {
	return o
}

func (o WorkloadMapOutput) ToWorkloadMapOutputWithContext(ctx context.Context) WorkloadMapOutput {
	return o
}

func (o WorkloadMapOutput) MapIndex(k pulumi.StringInput) WorkloadOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Workload {
		return vs[0].(map[string]*Workload)[vs[1].(string)]
	}).(WorkloadOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadInput)(nil)).Elem(), &Workload{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadArrayInput)(nil)).Elem(), WorkloadArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadMapInput)(nil)).Elem(), WorkloadMap{})
	pulumi.RegisterOutputType(WorkloadOutput{})
	pulumi.RegisterOutputType(WorkloadArrayOutput{})
	pulumi.RegisterOutputType(WorkloadMapOutput{})
}
