// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tem
{
    /// <summary>
    /// Provides a resource to create a tem workload
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var workload = new Tencentcloud.Tem.Workload("workload", new()
    ///     {
    ///         ApplicationId = "app-j4d3x6kj",
    ///         CpuSpec = 1,
    ///         DeployMode = "IMAGE",
    ///         DeployVersion = "hello-world",
    ///         EnvironmentId = "en-85377m6j",
    ///         ImgRepo = "tem_demo/tem_demo",
    ///         InitPodNum = 1,
    ///         MemorySpec = 1,
    ///         RepoServer = "ccr.ccs.tencentyun.com",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// tem workload can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Tem/workload:Workload workload envirnomentId#applicationId
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tem/workload:Workload")]
    public partial class Workload : global::Pulumi.CustomResource
    {
        /// <summary>
        /// application ID.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// cpu.
        /// </summary>
        [Output("cpuSpec")]
        public Output<double> CpuSpec { get; private set; } = null!;

        /// <summary>
        /// deploy mode, support IMAGE.
        /// </summary>
        [Output("deployMode")]
        public Output<string> DeployMode { get; private set; } = null!;

        /// <summary>
        /// deploy strategy.
        /// </summary>
        [Output("deployStrategyConf")]
        public Output<Outputs.WorkloadDeployStrategyConf?> DeployStrategyConf { get; private set; } = null!;

        /// <summary>
        /// deploy version.
        /// </summary>
        [Output("deployVersion")]
        public Output<string> DeployVersion { get; private set; } = null!;

        /// <summary>
        /// .
        /// </summary>
        [Output("envConfs")]
        public Output<ImmutableArray<Outputs.WorkloadEnvConf>> EnvConfs { get; private set; } = null!;

        /// <summary>
        /// environment ID.
        /// </summary>
        [Output("environmentId")]
        public Output<string> EnvironmentId { get; private set; } = null!;

        /// <summary>
        /// repository name.
        /// </summary>
        [Output("imgRepo")]
        public Output<string> ImgRepo { get; private set; } = null!;

        /// <summary>
        /// initial pod number.
        /// </summary>
        [Output("initPodNum")]
        public Output<int> InitPodNum { get; private set; } = null!;

        /// <summary>
        /// liveness config.
        /// </summary>
        [Output("liveness")]
        public Output<Outputs.WorkloadLiveness?> Liveness { get; private set; } = null!;

        /// <summary>
        /// mem.
        /// </summary>
        [Output("memorySpec")]
        public Output<double> MemorySpec { get; private set; } = null!;

        /// <summary>
        /// mem.
        /// </summary>
        [Output("postStart")]
        public Output<string?> PostStart { get; private set; } = null!;

        /// <summary>
        /// mem.
        /// </summary>
        [Output("preStop")]
        public Output<string?> PreStop { get; private set; } = null!;

        /// <summary>
        /// .
        /// </summary>
        [Output("readiness")]
        public Output<Outputs.WorkloadReadiness?> Readiness { get; private set; } = null!;

        /// <summary>
        /// repo server addr when deploy by image.
        /// </summary>
        [Output("repoServer")]
        public Output<string?> RepoServer { get; private set; } = null!;

        /// <summary>
        /// repo type when deploy: 0: tcr personal; 1: tcr enterprise; 2: public repository; 3: tem host tcr; 4: demo repo.
        /// </summary>
        [Output("repoType")]
        public Output<int?> RepoType { get; private set; } = null!;

        /// <summary>
        /// security groups.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// .
        /// </summary>
        [Output("startupProbe")]
        public Output<Outputs.WorkloadStartupProbe?> StartupProbe { get; private set; } = null!;

        /// <summary>
        /// storage configuration.
        /// </summary>
        [Output("storageConfs")]
        public Output<ImmutableArray<Outputs.WorkloadStorageConf>> StorageConfs { get; private set; } = null!;

        /// <summary>
        /// storage mount configuration.
        /// </summary>
        [Output("storageMountConfs")]
        public Output<ImmutableArray<Outputs.WorkloadStorageMountConf>> StorageMountConfs { get; private set; } = null!;

        /// <summary>
        /// tcr instance id when deploy by image.
        /// </summary>
        [Output("tcrInstanceId")]
        public Output<string?> TcrInstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a Workload resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Workload(string name, WorkloadArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tem/workload:Workload", name, args ?? new WorkloadArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Workload(string name, Input<string> id, WorkloadState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tem/workload:Workload", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Workload resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Workload Get(string name, Input<string> id, WorkloadState? state = null, CustomResourceOptions? options = null)
        {
            return new Workload(name, id, state, options);
        }
    }

    public sealed class WorkloadArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// application ID.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// cpu.
        /// </summary>
        [Input("cpuSpec", required: true)]
        public Input<double> CpuSpec { get; set; } = null!;

        /// <summary>
        /// deploy mode, support IMAGE.
        /// </summary>
        [Input("deployMode", required: true)]
        public Input<string> DeployMode { get; set; } = null!;

        /// <summary>
        /// deploy strategy.
        /// </summary>
        [Input("deployStrategyConf")]
        public Input<Inputs.WorkloadDeployStrategyConfArgs>? DeployStrategyConf { get; set; }

        /// <summary>
        /// deploy version.
        /// </summary>
        [Input("deployVersion", required: true)]
        public Input<string> DeployVersion { get; set; } = null!;

        [Input("envConfs")]
        private InputList<Inputs.WorkloadEnvConfArgs>? _envConfs;

        /// <summary>
        /// .
        /// </summary>
        public InputList<Inputs.WorkloadEnvConfArgs> EnvConfs
        {
            get => _envConfs ?? (_envConfs = new InputList<Inputs.WorkloadEnvConfArgs>());
            set => _envConfs = value;
        }

        /// <summary>
        /// environment ID.
        /// </summary>
        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        /// <summary>
        /// repository name.
        /// </summary>
        [Input("imgRepo", required: true)]
        public Input<string> ImgRepo { get; set; } = null!;

        /// <summary>
        /// initial pod number.
        /// </summary>
        [Input("initPodNum", required: true)]
        public Input<int> InitPodNum { get; set; } = null!;

        /// <summary>
        /// liveness config.
        /// </summary>
        [Input("liveness")]
        public Input<Inputs.WorkloadLivenessArgs>? Liveness { get; set; }

        /// <summary>
        /// mem.
        /// </summary>
        [Input("memorySpec", required: true)]
        public Input<double> MemorySpec { get; set; } = null!;

        /// <summary>
        /// mem.
        /// </summary>
        [Input("postStart")]
        public Input<string>? PostStart { get; set; }

        /// <summary>
        /// mem.
        /// </summary>
        [Input("preStop")]
        public Input<string>? PreStop { get; set; }

        /// <summary>
        /// .
        /// </summary>
        [Input("readiness")]
        public Input<Inputs.WorkloadReadinessArgs>? Readiness { get; set; }

        /// <summary>
        /// repo server addr when deploy by image.
        /// </summary>
        [Input("repoServer")]
        public Input<string>? RepoServer { get; set; }

        /// <summary>
        /// repo type when deploy: 0: tcr personal; 1: tcr enterprise; 2: public repository; 3: tem host tcr; 4: demo repo.
        /// </summary>
        [Input("repoType")]
        public Input<int>? RepoType { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// security groups.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// .
        /// </summary>
        [Input("startupProbe")]
        public Input<Inputs.WorkloadStartupProbeArgs>? StartupProbe { get; set; }

        [Input("storageConfs")]
        private InputList<Inputs.WorkloadStorageConfArgs>? _storageConfs;

        /// <summary>
        /// storage configuration.
        /// </summary>
        public InputList<Inputs.WorkloadStorageConfArgs> StorageConfs
        {
            get => _storageConfs ?? (_storageConfs = new InputList<Inputs.WorkloadStorageConfArgs>());
            set => _storageConfs = value;
        }

        [Input("storageMountConfs")]
        private InputList<Inputs.WorkloadStorageMountConfArgs>? _storageMountConfs;

        /// <summary>
        /// storage mount configuration.
        /// </summary>
        public InputList<Inputs.WorkloadStorageMountConfArgs> StorageMountConfs
        {
            get => _storageMountConfs ?? (_storageMountConfs = new InputList<Inputs.WorkloadStorageMountConfArgs>());
            set => _storageMountConfs = value;
        }

        /// <summary>
        /// tcr instance id when deploy by image.
        /// </summary>
        [Input("tcrInstanceId")]
        public Input<string>? TcrInstanceId { get; set; }

        public WorkloadArgs()
        {
        }
        public static new WorkloadArgs Empty => new WorkloadArgs();
    }

    public sealed class WorkloadState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// application ID.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// cpu.
        /// </summary>
        [Input("cpuSpec")]
        public Input<double>? CpuSpec { get; set; }

        /// <summary>
        /// deploy mode, support IMAGE.
        /// </summary>
        [Input("deployMode")]
        public Input<string>? DeployMode { get; set; }

        /// <summary>
        /// deploy strategy.
        /// </summary>
        [Input("deployStrategyConf")]
        public Input<Inputs.WorkloadDeployStrategyConfGetArgs>? DeployStrategyConf { get; set; }

        /// <summary>
        /// deploy version.
        /// </summary>
        [Input("deployVersion")]
        public Input<string>? DeployVersion { get; set; }

        [Input("envConfs")]
        private InputList<Inputs.WorkloadEnvConfGetArgs>? _envConfs;

        /// <summary>
        /// .
        /// </summary>
        public InputList<Inputs.WorkloadEnvConfGetArgs> EnvConfs
        {
            get => _envConfs ?? (_envConfs = new InputList<Inputs.WorkloadEnvConfGetArgs>());
            set => _envConfs = value;
        }

        /// <summary>
        /// environment ID.
        /// </summary>
        [Input("environmentId")]
        public Input<string>? EnvironmentId { get; set; }

        /// <summary>
        /// repository name.
        /// </summary>
        [Input("imgRepo")]
        public Input<string>? ImgRepo { get; set; }

        /// <summary>
        /// initial pod number.
        /// </summary>
        [Input("initPodNum")]
        public Input<int>? InitPodNum { get; set; }

        /// <summary>
        /// liveness config.
        /// </summary>
        [Input("liveness")]
        public Input<Inputs.WorkloadLivenessGetArgs>? Liveness { get; set; }

        /// <summary>
        /// mem.
        /// </summary>
        [Input("memorySpec")]
        public Input<double>? MemorySpec { get; set; }

        /// <summary>
        /// mem.
        /// </summary>
        [Input("postStart")]
        public Input<string>? PostStart { get; set; }

        /// <summary>
        /// mem.
        /// </summary>
        [Input("preStop")]
        public Input<string>? PreStop { get; set; }

        /// <summary>
        /// .
        /// </summary>
        [Input("readiness")]
        public Input<Inputs.WorkloadReadinessGetArgs>? Readiness { get; set; }

        /// <summary>
        /// repo server addr when deploy by image.
        /// </summary>
        [Input("repoServer")]
        public Input<string>? RepoServer { get; set; }

        /// <summary>
        /// repo type when deploy: 0: tcr personal; 1: tcr enterprise; 2: public repository; 3: tem host tcr; 4: demo repo.
        /// </summary>
        [Input("repoType")]
        public Input<int>? RepoType { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// security groups.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// .
        /// </summary>
        [Input("startupProbe")]
        public Input<Inputs.WorkloadStartupProbeGetArgs>? StartupProbe { get; set; }

        [Input("storageConfs")]
        private InputList<Inputs.WorkloadStorageConfGetArgs>? _storageConfs;

        /// <summary>
        /// storage configuration.
        /// </summary>
        public InputList<Inputs.WorkloadStorageConfGetArgs> StorageConfs
        {
            get => _storageConfs ?? (_storageConfs = new InputList<Inputs.WorkloadStorageConfGetArgs>());
            set => _storageConfs = value;
        }

        [Input("storageMountConfs")]
        private InputList<Inputs.WorkloadStorageMountConfGetArgs>? _storageMountConfs;

        /// <summary>
        /// storage mount configuration.
        /// </summary>
        public InputList<Inputs.WorkloadStorageMountConfGetArgs> StorageMountConfs
        {
            get => _storageMountConfs ?? (_storageMountConfs = new InputList<Inputs.WorkloadStorageMountConfGetArgs>());
            set => _storageMountConfs = value;
        }

        /// <summary>
        /// tcr instance id when deploy by image.
        /// </summary>
        [Input("tcrInstanceId")]
        public Input<string>? TcrInstanceId { get; set; }

        public WorkloadState()
        {
        }
        public static new WorkloadState Empty => new WorkloadState();
    }
}
