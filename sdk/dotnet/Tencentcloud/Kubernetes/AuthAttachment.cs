// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes
{
    /// <summary>
    /// Provide a resource to configure kubernetes cluster authentication info.
    /// 
    /// &gt; **NOTE:** Only available for cluster version &gt;= 1.20
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-3";
    ///         var clusterCidr = config.Get("clusterCidr") ?? "172.16.0.0/16";
    ///         var defaultInstanceType = config.Get("defaultInstanceType") ?? "S1.SMALL1";
    ///         var @default = Output.Create(Tencentcloud.Images.GetInstance.InvokeAsync(new Tencentcloud.Images.GetInstanceArgs
    ///         {
    ///             ImageTypes = 
    ///             {
    ///                 "PUBLIC_IMAGE",
    ///             },
    ///             OsName = "centos",
    ///         }));
    ///         var vpc = Output.Create(Tencentcloud.Vpc.GetSubnets.InvokeAsync(new Tencentcloud.Vpc.GetSubnetsArgs
    ///         {
    ///             IsDefault = true,
    ///             AvailabilityZone = availabilityZone,
    ///         }));
    ///         var managedCluster = new Tencentcloud.Kubernetes.Cluster("managedCluster", new Tencentcloud.Kubernetes.ClusterArgs
    ///         {
    ///             VpcId = vpc.Apply(vpc =&gt; vpc.InstanceLists?[0]?.VpcId),
    ///             ClusterCidr = "10.31.0.0/16",
    ///             ClusterMaxPodNum = 32,
    ///             ClusterName = "keep",
    ///             ClusterDesc = "test cluster desc",
    ///             ClusterVersion = "1.20.6",
    ///             ClusterMaxServiceNum = 32,
    ///             WorkerConfigs = 
    ///             {
    ///                 new Tencentcloud.Kubernetes.Inputs.ClusterWorkerConfigArgs
    ///                 {
    ///                     Count = 1,
    ///                     AvailabilityZone = availabilityZone,
    ///                     InstanceType = defaultInstanceType,
    ///                     SystemDiskType = "CLOUD_SSD",
    ///                     SystemDiskSize = 60,
    ///                     InternetChargeType = "TRAFFIC_POSTPAID_BY_HOUR",
    ///                     InternetMaxBandwidthOut = 100,
    ///                     PublicIpAssigned = true,
    ///                     SubnetId = vpc.Apply(vpc =&gt; vpc.InstanceLists?[0]?.SubnetId),
    ///                     DataDisks = 
    ///                     {
    ///                         new Tencentcloud.Kubernetes.Inputs.ClusterWorkerConfigDataDiskArgs
    ///                         {
    ///                             DiskType = "CLOUD_PREMIUM",
    ///                             DiskSize = 50,
    ///                         },
    ///                     },
    ///                     EnhancedSecurityService = false,
    ///                     EnhancedMonitorService = false,
    ///                     UserData = "dGVzdA==",
    ///                     Password = "ZZXXccvv1212",
    ///                 },
    ///             },
    ///             ClusterDeployType = "MANAGED_CLUSTER",
    ///         });
    ///         var testAuthAttach = new Tencentcloud.Kubernetes.AuthAttachment("testAuthAttach", new Tencentcloud.Kubernetes.AuthAttachmentArgs
    ///         {
    ///             ClusterId = managedCluster.Id,
    ///             JwksUri = managedCluster.Id.Apply(id =&gt; $"https://{id}.ccs.tencent-cloud.com/openid/v1/jwks"),
    ///             Issuer = managedCluster.Id.Apply(id =&gt; $"https://{id}.ccs.tencent-cloud.com"),
    ///             AutoCreateDiscoveryAnonymousAuth = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Use the TKE default issuer and jwks_uri
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var availabilityZone = config.Get("availabilityZone") ?? "ap-guangzhou-3";
    ///         var clusterCidr = config.Get("clusterCidr") ?? "172.16.0.0/16";
    ///         var defaultInstanceType = config.Get("defaultInstanceType") ?? "S1.SMALL1";
    ///         var @default = Output.Create(Tencentcloud.Images.GetInstance.InvokeAsync(new Tencentcloud.Images.GetInstanceArgs
    ///         {
    ///             ImageTypes = 
    ///             {
    ///                 "PUBLIC_IMAGE",
    ///             },
    ///             OsName = "centos",
    ///         }));
    ///         var vpc = Output.Create(Tencentcloud.Vpc.GetSubnets.InvokeAsync(new Tencentcloud.Vpc.GetSubnetsArgs
    ///         {
    ///             IsDefault = true,
    ///             AvailabilityZone = availabilityZone,
    ///         }));
    ///         var managedCluster = new Tencentcloud.Kubernetes.Cluster("managedCluster", new Tencentcloud.Kubernetes.ClusterArgs
    ///         {
    ///             VpcId = vpc.Apply(vpc =&gt; vpc.InstanceLists?[0]?.VpcId),
    ///             ClusterCidr = "10.31.0.0/16",
    ///             ClusterMaxPodNum = 32,
    ///             ClusterName = "keep",
    ///             ClusterDesc = "test cluster desc",
    ///             ClusterVersion = "1.20.6",
    ///             ClusterMaxServiceNum = 32,
    ///             WorkerConfigs = 
    ///             {
    ///                 new Tencentcloud.Kubernetes.Inputs.ClusterWorkerConfigArgs
    ///                 {
    ///                     Count = 1,
    ///                     AvailabilityZone = availabilityZone,
    ///                     InstanceType = defaultInstanceType,
    ///                     SystemDiskType = "CLOUD_SSD",
    ///                     SystemDiskSize = 60,
    ///                     InternetChargeType = "TRAFFIC_POSTPAID_BY_HOUR",
    ///                     InternetMaxBandwidthOut = 100,
    ///                     PublicIpAssigned = true,
    ///                     SubnetId = vpc.Apply(vpc =&gt; vpc.InstanceLists?[0]?.SubnetId),
    ///                     DataDisks = 
    ///                     {
    ///                         new Tencentcloud.Kubernetes.Inputs.ClusterWorkerConfigDataDiskArgs
    ///                         {
    ///                             DiskType = "CLOUD_PREMIUM",
    ///                             DiskSize = 50,
    ///                         },
    ///                     },
    ///                     EnhancedSecurityService = false,
    ///                     EnhancedMonitorService = false,
    ///                     UserData = "dGVzdA==",
    ///                     Password = "ZZXXccvv1212",
    ///                 },
    ///             },
    ///             ClusterDeployType = "MANAGED_CLUSTER",
    ///         });
    ///         // if you want to use tke default issuer and jwks_uri, please set use_tke_default to true and set issuer to empty string.
    ///         var testUseTkeDefaultAuthAttach = new Tencentcloud.Kubernetes.AuthAttachment("testUseTkeDefaultAuthAttach", new Tencentcloud.Kubernetes.AuthAttachmentArgs
    ///         {
    ///             ClusterId = managedCluster.Id,
    ///             AutoCreateDiscoveryAnonymousAuth = true,
    ///             UseTkeDefault = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Kubernetes/authAttachment:AuthAttachment")]
    public partial class AuthAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// If set to `true`, the rbac rule will be created automatically which allow anonymous user to access '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        /// </summary>
        [Output("autoCreateDiscoveryAnonymousAuth")]
        public Output<bool?> AutoCreateDiscoveryAnonymousAuth { get; private set; } = null!;

        /// <summary>
        /// ID of clusters.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Specify service-account-issuer. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Output("issuer")]
        public Output<string?> Issuer { get; private set; } = null!;

        /// <summary>
        /// Specify service-account-jwks-uri. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Output("jwksUri")]
        public Output<string?> JwksUri { get; private set; } = null!;

        /// <summary>
        /// The default issuer of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.
        /// </summary>
        [Output("tkeDefaultIssuer")]
        public Output<string> TkeDefaultIssuer { get; private set; } = null!;

        /// <summary>
        /// The default jwks_uri of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.
        /// </summary>
        [Output("tkeDefaultJwksUri")]
        public Output<string> TkeDefaultJwksUri { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, the issuer and jwks_uri will be generated automatically by tke, please do not set issuer and jwks_uri.
        /// </summary>
        [Output("useTkeDefault")]
        public Output<bool?> UseTkeDefault { get; private set; } = null!;


        /// <summary>
        /// Create a AuthAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthAttachment(string name, AuthAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/authAttachment:AuthAttachment", name, args ?? new AuthAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthAttachment(string name, Input<string> id, AuthAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/authAttachment:AuthAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthAttachment Get(string name, Input<string> id, AuthAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthAttachment(name, id, state, options);
        }
    }

    public sealed class AuthAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to `true`, the rbac rule will be created automatically which allow anonymous user to access '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        /// </summary>
        [Input("autoCreateDiscoveryAnonymousAuth")]
        public Input<bool>? AutoCreateDiscoveryAnonymousAuth { get; set; }

        /// <summary>
        /// ID of clusters.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Specify service-account-issuer. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// Specify service-account-jwks-uri. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Input("jwksUri")]
        public Input<string>? JwksUri { get; set; }

        /// <summary>
        /// If set to `true`, the issuer and jwks_uri will be generated automatically by tke, please do not set issuer and jwks_uri.
        /// </summary>
        [Input("useTkeDefault")]
        public Input<bool>? UseTkeDefault { get; set; }

        public AuthAttachmentArgs()
        {
        }
    }

    public sealed class AuthAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set to `true`, the rbac rule will be created automatically which allow anonymous user to access '/.well-known/openid-configuration' and '/openid/v1/jwks'.
        /// </summary>
        [Input("autoCreateDiscoveryAnonymousAuth")]
        public Input<bool>? AutoCreateDiscoveryAnonymousAuth { get; set; }

        /// <summary>
        /// ID of clusters.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Specify service-account-issuer. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// Specify service-account-jwks-uri. If use_tke_default is set to `true`, please do not set this field.
        /// </summary>
        [Input("jwksUri")]
        public Input<string>? JwksUri { get; set; }

        /// <summary>
        /// The default issuer of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.
        /// </summary>
        [Input("tkeDefaultIssuer")]
        public Input<string>? TkeDefaultIssuer { get; set; }

        /// <summary>
        /// The default jwks_uri of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.
        /// </summary>
        [Input("tkeDefaultJwksUri")]
        public Input<string>? TkeDefaultJwksUri { get; set; }

        /// <summary>
        /// If set to `true`, the issuer and jwks_uri will be generated automatically by tke, please do not set issuer and jwks_uri.
        /// </summary>
        [Input("useTkeDefault")]
        public Input<bool>? UseTkeDefault { get; set; }

        public AuthAttachmentState()
        {
        }
    }
}
