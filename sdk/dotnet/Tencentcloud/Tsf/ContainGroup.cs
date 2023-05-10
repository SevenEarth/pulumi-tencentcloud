// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf
{
    /// <summary>
    /// Provides a resource to create a tsf contain_group
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var containGroup = new Tencentcloud.Tsf.ContainGroup("containGroup", new Tencentcloud.Tsf.ContainGroupArgs
    ///         {
    ///             AccessType = 0,
    ///             ApplicationId = "application-y5r4nejv",
    ///             ClusterId = "cls-2yu5kxr8",
    ///             CpuLimit = "0.5",
    ///             CpuRequest = "0.25",
    ///             GroupName = "terraform-test",
    ///             GroupResourceType = "DEF",
    ///             InstanceNum = 1,
    ///             MemLimit = "1280",
    ///             MemRequest = "640",
    ///             NamespaceId = "namespace-ydlezgxa",
    ///             ProtocolPorts = 
    ///             {
    ///                 new Tencentcloud.Tsf.Inputs.ContainGroupProtocolPortArgs
    ///                 {
    ///                     NodePort = 0,
    ///                     Port = 333,
    ///                     Protocol = "TCP",
    ///                     TargetPort = 333,
    ///                 },
    ///             },
    ///             UpdateIvl = 10,
    ///             UpdateType = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tsf contain_group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tsf/containGroup:ContainGroup contain_group contain_group_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tsf/containGroup:ContainGroup")]
    public partial class ContainGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// 0: public network 1: access within the cluster 2: NodePort.
        /// </summary>
        [Output("accessType")]
        public Output<int> AccessType { get; private set; } = null!;

        /// <summary>
        /// The maximum number of CPU cores for the agent container, corresponding to the limit of K8S.
        /// </summary>
        [Output("agentCpuLimit")]
        public Output<string?> AgentCpuLimit { get; private set; } = null!;

        /// <summary>
        /// The number of CPU cores allocated by the agent container, corresponding to the K8S request.
        /// </summary>
        [Output("agentCpuRequest")]
        public Output<string?> AgentCpuRequest { get; private set; } = null!;

        /// <summary>
        /// The maximum memory MiB of the agent container, corresponding to the limit of K8S.
        /// </summary>
        [Output("agentMemLimit")]
        public Output<string?> AgentMemLimit { get; private set; } = null!;

        /// <summary>
        /// The number of memory MiB allocated by the agent container, corresponding to the K8S request.
        /// </summary>
        [Output("agentMemRequest")]
        public Output<string?> AgentMemRequest { get; private set; } = null!;

        /// <summary>
        /// The application ID to which the group belongs.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// Application Name.
        /// </summary>
        [Output("applicationName")]
        public Output<string> ApplicationName { get; private set; } = null!;

        /// <summary>
        /// App types.
        /// </summary>
        [Output("applicationType")]
        public Output<string> ApplicationType { get; private set; } = null!;

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Service ip.
        /// </summary>
        [Output("clusterIp")]
        public Output<string> ClusterIp { get; private set; } = null!;

        /// <summary>
        /// cluster name.
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// The maximum number of allocated CPU cores, corresponding to the K8S limit.
        /// </summary>
        [Output("cpuLimit")]
        public Output<string?> CpuLimit { get; private set; } = null!;

        /// <summary>
        /// Initially allocated CPU cores, corresponding to K8S request.
        /// </summary>
        [Output("cpuRequest")]
        public Output<string?> CpuRequest { get; private set; } = null!;

        /// <summary>
        /// creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Total number of instances launched.
        /// </summary>
        [Output("currentNum")]
        public Output<int> CurrentNum { get; private set; } = null!;

        /// <summary>
        /// environment variable array object.
        /// </summary>
        [Output("envs")]
        public Output<ImmutableArray<Outputs.ContainGroupEnv>> Envs { get; private set; } = null!;

        /// <summary>
        /// Group remarks field, the length should not exceed 200 characters.
        /// </summary>
        [Output("groupComment")]
        public Output<string?> GroupComment { get; private set; } = null!;

        /// <summary>
        /// Deployment group ID.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// Group name field, length 1~60, beginning with a letter or underscore, can contain alphanumeric underscore.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// Deployment Group Resource Type.
        /// </summary>
        [Output("groupResourceType")]
        public Output<string> GroupResourceType { get; private set; } = null!;

        /// <summary>
        /// Deployment group health check settings.
        /// </summary>
        [Output("healthCheckSettings")]
        public Output<ImmutableArray<Outputs.ContainGroupHealthCheckSetting>> HealthCheckSettings { get; private set; } = null!;

        /// <summary>
        /// Number of deployment group instances.
        /// </summary>
        [Output("instanceCount")]
        public Output<int> InstanceCount { get; private set; } = null!;

        /// <summary>
        /// number of instances.
        /// </summary>
        [Output("instanceNum")]
        public Output<int> InstanceNum { get; private set; } = null!;

        /// <summary>
        /// The maximum number of CPU cores for the istioproxy container corresponds to the limit of K8S.
        /// </summary>
        [Output("istioCpuLimit")]
        public Output<string> IstioCpuLimit { get; private set; } = null!;

        /// <summary>
        /// The number of CPU cores allocated by the istioproxy container, corresponding to the K8S request.
        /// </summary>
        [Output("istioCpuRequest")]
        public Output<string> IstioCpuRequest { get; private set; } = null!;

        /// <summary>
        /// The maximum memory MiB of the istioproxy container corresponds to the limit of K8S.
        /// </summary>
        [Output("istioMemLimit")]
        public Output<string> IstioMemLimit { get; private set; } = null!;

        /// <summary>
        /// The number of memory MiB allocated by the istioproxy container, corresponding to the K8S request.
        /// </summary>
        [Output("istioMemRequest")]
        public Output<string> IstioMemRequest { get; private set; } = null!;

        /// <summary>
        /// load balancing ip.
        /// </summary>
        [Output("lbIp")]
        public Output<string> LbIp { get; private set; } = null!;

        /// <summary>
        /// The MaxSurge parameter of the kubernetes rolling update policy.
        /// </summary>
        [Output("maxSurge")]
        public Output<string> MaxSurge { get; private set; } = null!;

        /// <summary>
        /// The MaxUnavailable parameter of the kubernetes rolling update policy.
        /// </summary>
        [Output("maxUnavailable")]
        public Output<string> MaxUnavailable { get; private set; } = null!;

        /// <summary>
        /// Maximum allocated memory MiB, corresponding to K8S limit.
        /// </summary>
        [Output("memLimit")]
        public Output<string?> MemLimit { get; private set; } = null!;

        /// <summary>
        /// Initially allocated memory MiB, corresponding to K8S request.
        /// </summary>
        [Output("memRequest")]
        public Output<string?> MemRequest { get; private set; } = null!;

        /// <summary>
        /// pod error message description.
        /// </summary>
        [Output("message")]
        public Output<string> Message { get; private set; } = null!;

        /// <summary>
        /// Service type.
        /// </summary>
        [Output("microserviceType")]
        public Output<string> MicroserviceType { get; private set; } = null!;

        /// <summary>
        /// ID of the namespace to which the group belongs.
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// namespace name.
        /// </summary>
        [Output("namespaceName")]
        public Output<string> NamespaceName { get; private set; } = null!;

        /// <summary>
        /// Protocol Ports array.
        /// </summary>
        [Output("protocolPorts")]
        public Output<ImmutableArray<Outputs.ContainGroupProtocolPort>> ProtocolPorts { get; private set; } = null!;

        /// <summary>
        /// Mirror name, such as /tsf/nginx.
        /// </summary>
        [Output("reponame")]
        public Output<string> Reponame { get; private set; } = null!;

        /// <summary>
        /// mirror server.
        /// </summary>
        [Output("server")]
        public Output<string> Server { get; private set; } = null!;

        /// <summary>
        /// Deployment group status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// subnet ID.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Image version name.
        /// </summary>
        [Output("tagName")]
        public Output<string> TagName { get; private set; } = null!;

        /// <summary>
        /// Rolling update is required, update interval.
        /// </summary>
        [Output("updateIvl")]
        public Output<int> UpdateIvl { get; private set; } = null!;

        /// <summary>
        /// Update method: 0: fast update 1: rolling update.
        /// </summary>
        [Output("updateType")]
        public Output<int> UpdateType { get; private set; } = null!;

        /// <summary>
        /// Deployment group update timestamp.
        /// </summary>
        [Output("updatedTime")]
        public Output<int> UpdatedTime { get; private set; } = null!;


        /// <summary>
        /// Create a ContainGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainGroup(string name, ContainGroupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/containGroup:ContainGroup", name, args ?? new ContainGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainGroup(string name, Input<string> id, ContainGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tsf/containGroup:ContainGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContainGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainGroup Get(string name, Input<string> id, ContainGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ContainGroup(name, id, state, options);
        }
    }

    public sealed class ContainGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// 0: public network 1: access within the cluster 2: NodePort.
        /// </summary>
        [Input("accessType", required: true)]
        public Input<int> AccessType { get; set; } = null!;

        /// <summary>
        /// The maximum number of CPU cores for the agent container, corresponding to the limit of K8S.
        /// </summary>
        [Input("agentCpuLimit")]
        public Input<string>? AgentCpuLimit { get; set; }

        /// <summary>
        /// The number of CPU cores allocated by the agent container, corresponding to the K8S request.
        /// </summary>
        [Input("agentCpuRequest")]
        public Input<string>? AgentCpuRequest { get; set; }

        /// <summary>
        /// The maximum memory MiB of the agent container, corresponding to the limit of K8S.
        /// </summary>
        [Input("agentMemLimit")]
        public Input<string>? AgentMemLimit { get; set; }

        /// <summary>
        /// The number of memory MiB allocated by the agent container, corresponding to the K8S request.
        /// </summary>
        [Input("agentMemRequest")]
        public Input<string>? AgentMemRequest { get; set; }

        /// <summary>
        /// The application ID to which the group belongs.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The maximum number of allocated CPU cores, corresponding to the K8S limit.
        /// </summary>
        [Input("cpuLimit")]
        public Input<string>? CpuLimit { get; set; }

        /// <summary>
        /// Initially allocated CPU cores, corresponding to K8S request.
        /// </summary>
        [Input("cpuRequest")]
        public Input<string>? CpuRequest { get; set; }

        /// <summary>
        /// Group remarks field, the length should not exceed 200 characters.
        /// </summary>
        [Input("groupComment")]
        public Input<string>? GroupComment { get; set; }

        /// <summary>
        /// Group name field, length 1~60, beginning with a letter or underscore, can contain alphanumeric underscore.
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// Deployment Group Resource Type.
        /// </summary>
        [Input("groupResourceType")]
        public Input<string>? GroupResourceType { get; set; }

        /// <summary>
        /// number of instances.
        /// </summary>
        [Input("instanceNum", required: true)]
        public Input<int> InstanceNum { get; set; } = null!;

        /// <summary>
        /// The maximum number of CPU cores for the istioproxy container corresponds to the limit of K8S.
        /// </summary>
        [Input("istioCpuLimit")]
        public Input<string>? IstioCpuLimit { get; set; }

        /// <summary>
        /// The number of CPU cores allocated by the istioproxy container, corresponding to the K8S request.
        /// </summary>
        [Input("istioCpuRequest")]
        public Input<string>? IstioCpuRequest { get; set; }

        /// <summary>
        /// The maximum memory MiB of the istioproxy container corresponds to the limit of K8S.
        /// </summary>
        [Input("istioMemLimit")]
        public Input<string>? IstioMemLimit { get; set; }

        /// <summary>
        /// The number of memory MiB allocated by the istioproxy container, corresponding to the K8S request.
        /// </summary>
        [Input("istioMemRequest")]
        public Input<string>? IstioMemRequest { get; set; }

        /// <summary>
        /// Maximum allocated memory MiB, corresponding to K8S limit.
        /// </summary>
        [Input("memLimit")]
        public Input<string>? MemLimit { get; set; }

        /// <summary>
        /// Initially allocated memory MiB, corresponding to K8S request.
        /// </summary>
        [Input("memRequest")]
        public Input<string>? MemRequest { get; set; }

        /// <summary>
        /// ID of the namespace to which the group belongs.
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        [Input("protocolPorts", required: true)]
        private InputList<Inputs.ContainGroupProtocolPortArgs>? _protocolPorts;

        /// <summary>
        /// Protocol Ports array.
        /// </summary>
        public InputList<Inputs.ContainGroupProtocolPortArgs> ProtocolPorts
        {
            get => _protocolPorts ?? (_protocolPorts = new InputList<Inputs.ContainGroupProtocolPortArgs>());
            set => _protocolPorts = value;
        }

        /// <summary>
        /// subnet ID.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Rolling update is required, update interval.
        /// </summary>
        [Input("updateIvl")]
        public Input<int>? UpdateIvl { get; set; }

        /// <summary>
        /// Update method: 0: fast update 1: rolling update.
        /// </summary>
        [Input("updateType")]
        public Input<int>? UpdateType { get; set; }

        public ContainGroupArgs()
        {
        }
    }

    public sealed class ContainGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// 0: public network 1: access within the cluster 2: NodePort.
        /// </summary>
        [Input("accessType")]
        public Input<int>? AccessType { get; set; }

        /// <summary>
        /// The maximum number of CPU cores for the agent container, corresponding to the limit of K8S.
        /// </summary>
        [Input("agentCpuLimit")]
        public Input<string>? AgentCpuLimit { get; set; }

        /// <summary>
        /// The number of CPU cores allocated by the agent container, corresponding to the K8S request.
        /// </summary>
        [Input("agentCpuRequest")]
        public Input<string>? AgentCpuRequest { get; set; }

        /// <summary>
        /// The maximum memory MiB of the agent container, corresponding to the limit of K8S.
        /// </summary>
        [Input("agentMemLimit")]
        public Input<string>? AgentMemLimit { get; set; }

        /// <summary>
        /// The number of memory MiB allocated by the agent container, corresponding to the K8S request.
        /// </summary>
        [Input("agentMemRequest")]
        public Input<string>? AgentMemRequest { get; set; }

        /// <summary>
        /// The application ID to which the group belongs.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// Application Name.
        /// </summary>
        [Input("applicationName")]
        public Input<string>? ApplicationName { get; set; }

        /// <summary>
        /// App types.
        /// </summary>
        [Input("applicationType")]
        public Input<string>? ApplicationType { get; set; }

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Service ip.
        /// </summary>
        [Input("clusterIp")]
        public Input<string>? ClusterIp { get; set; }

        /// <summary>
        /// cluster name.
        /// </summary>
        [Input("clusterName")]
        public Input<string>? ClusterName { get; set; }

        /// <summary>
        /// The maximum number of allocated CPU cores, corresponding to the K8S limit.
        /// </summary>
        [Input("cpuLimit")]
        public Input<string>? CpuLimit { get; set; }

        /// <summary>
        /// Initially allocated CPU cores, corresponding to K8S request.
        /// </summary>
        [Input("cpuRequest")]
        public Input<string>? CpuRequest { get; set; }

        /// <summary>
        /// creation time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Total number of instances launched.
        /// </summary>
        [Input("currentNum")]
        public Input<int>? CurrentNum { get; set; }

        [Input("envs")]
        private InputList<Inputs.ContainGroupEnvGetArgs>? _envs;

        /// <summary>
        /// environment variable array object.
        /// </summary>
        public InputList<Inputs.ContainGroupEnvGetArgs> Envs
        {
            get => _envs ?? (_envs = new InputList<Inputs.ContainGroupEnvGetArgs>());
            set => _envs = value;
        }

        /// <summary>
        /// Group remarks field, the length should not exceed 200 characters.
        /// </summary>
        [Input("groupComment")]
        public Input<string>? GroupComment { get; set; }

        /// <summary>
        /// Deployment group ID.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// Group name field, length 1~60, beginning with a letter or underscore, can contain alphanumeric underscore.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// Deployment Group Resource Type.
        /// </summary>
        [Input("groupResourceType")]
        public Input<string>? GroupResourceType { get; set; }

        [Input("healthCheckSettings")]
        private InputList<Inputs.ContainGroupHealthCheckSettingGetArgs>? _healthCheckSettings;

        /// <summary>
        /// Deployment group health check settings.
        /// </summary>
        public InputList<Inputs.ContainGroupHealthCheckSettingGetArgs> HealthCheckSettings
        {
            get => _healthCheckSettings ?? (_healthCheckSettings = new InputList<Inputs.ContainGroupHealthCheckSettingGetArgs>());
            set => _healthCheckSettings = value;
        }

        /// <summary>
        /// Number of deployment group instances.
        /// </summary>
        [Input("instanceCount")]
        public Input<int>? InstanceCount { get; set; }

        /// <summary>
        /// number of instances.
        /// </summary>
        [Input("instanceNum")]
        public Input<int>? InstanceNum { get; set; }

        /// <summary>
        /// The maximum number of CPU cores for the istioproxy container corresponds to the limit of K8S.
        /// </summary>
        [Input("istioCpuLimit")]
        public Input<string>? IstioCpuLimit { get; set; }

        /// <summary>
        /// The number of CPU cores allocated by the istioproxy container, corresponding to the K8S request.
        /// </summary>
        [Input("istioCpuRequest")]
        public Input<string>? IstioCpuRequest { get; set; }

        /// <summary>
        /// The maximum memory MiB of the istioproxy container corresponds to the limit of K8S.
        /// </summary>
        [Input("istioMemLimit")]
        public Input<string>? IstioMemLimit { get; set; }

        /// <summary>
        /// The number of memory MiB allocated by the istioproxy container, corresponding to the K8S request.
        /// </summary>
        [Input("istioMemRequest")]
        public Input<string>? IstioMemRequest { get; set; }

        /// <summary>
        /// load balancing ip.
        /// </summary>
        [Input("lbIp")]
        public Input<string>? LbIp { get; set; }

        /// <summary>
        /// The MaxSurge parameter of the kubernetes rolling update policy.
        /// </summary>
        [Input("maxSurge")]
        public Input<string>? MaxSurge { get; set; }

        /// <summary>
        /// The MaxUnavailable parameter of the kubernetes rolling update policy.
        /// </summary>
        [Input("maxUnavailable")]
        public Input<string>? MaxUnavailable { get; set; }

        /// <summary>
        /// Maximum allocated memory MiB, corresponding to K8S limit.
        /// </summary>
        [Input("memLimit")]
        public Input<string>? MemLimit { get; set; }

        /// <summary>
        /// Initially allocated memory MiB, corresponding to K8S request.
        /// </summary>
        [Input("memRequest")]
        public Input<string>? MemRequest { get; set; }

        /// <summary>
        /// pod error message description.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// Service type.
        /// </summary>
        [Input("microserviceType")]
        public Input<string>? MicroserviceType { get; set; }

        /// <summary>
        /// ID of the namespace to which the group belongs.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// namespace name.
        /// </summary>
        [Input("namespaceName")]
        public Input<string>? NamespaceName { get; set; }

        [Input("protocolPorts")]
        private InputList<Inputs.ContainGroupProtocolPortGetArgs>? _protocolPorts;

        /// <summary>
        /// Protocol Ports array.
        /// </summary>
        public InputList<Inputs.ContainGroupProtocolPortGetArgs> ProtocolPorts
        {
            get => _protocolPorts ?? (_protocolPorts = new InputList<Inputs.ContainGroupProtocolPortGetArgs>());
            set => _protocolPorts = value;
        }

        /// <summary>
        /// Mirror name, such as /tsf/nginx.
        /// </summary>
        [Input("reponame")]
        public Input<string>? Reponame { get; set; }

        /// <summary>
        /// mirror server.
        /// </summary>
        [Input("server")]
        public Input<string>? Server { get; set; }

        /// <summary>
        /// Deployment group status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// subnet ID.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Image version name.
        /// </summary>
        [Input("tagName")]
        public Input<string>? TagName { get; set; }

        /// <summary>
        /// Rolling update is required, update interval.
        /// </summary>
        [Input("updateIvl")]
        public Input<int>? UpdateIvl { get; set; }

        /// <summary>
        /// Update method: 0: fast update 1: rolling update.
        /// </summary>
        [Input("updateType")]
        public Input<int>? UpdateType { get; set; }

        /// <summary>
        /// Deployment group update timestamp.
        /// </summary>
        [Input("updatedTime")]
        public Input<int>? UpdatedTime { get; set; }

        public ContainGroupState()
        {
        }
    }
}
