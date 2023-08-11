// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor
{
    /// <summary>
    /// Provides a resource to create a tmp tke global notification
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
    ///         var defaultInstanceType = config.Get("defaultInstanceType") ?? "SA1.MEDIUM2";
    ///         var availabilityZoneFirst = config.Get("availabilityZoneFirst") ?? "ap-guangzhou-3";
    ///         var availabilityZoneSecond = config.Get("availabilityZoneSecond") ?? "ap-guangzhou-4";
    ///         var exampleClusterCidr = config.Get("exampleClusterCidr") ?? "10.31.0.0/16";
    ///         var vpcOne = Output.Create(Tencentcloud.Vpc.GetSubnets.InvokeAsync(new Tencentcloud.Vpc.GetSubnetsArgs
    ///         {
    ///             IsDefault = true,
    ///             AvailabilityZone = availabilityZoneFirst,
    ///         }));
    ///         var firstVpcId = vpcOne.Apply(vpcOne =&gt; vpcOne.InstanceLists?[0]?.VpcId);
    ///         var firstSubnetId = vpcOne.Apply(vpcOne =&gt; vpcOne.InstanceLists?[0]?.SubnetId);
    ///         var vpcTwo = Output.Create(Tencentcloud.Vpc.GetSubnets.InvokeAsync(new Tencentcloud.Vpc.GetSubnetsArgs
    ///         {
    ///             IsDefault = true,
    ///             AvailabilityZone = availabilityZoneSecond,
    ///         }));
    ///         var secondVpcId = vpcTwo.Apply(vpcTwo =&gt; vpcTwo.InstanceLists?[0]?.VpcId);
    ///         var secondSubnetId = vpcTwo.Apply(vpcTwo =&gt; vpcTwo.InstanceLists?[0]?.SubnetId);
    ///         var sg = new Tencentcloud.Security.Group("sg", new Tencentcloud.Security.GroupArgs
    ///         {
    ///         });
    ///         var sgId = sg.Id;
    ///         var @default = Output.Create(Tencentcloud.Images.GetInstance.InvokeAsync(new Tencentcloud.Images.GetInstanceArgs
    ///         {
    ///             ImageTypes = 
    ///             {
    ///                 "PUBLIC_IMAGE",
    ///             },
    ///             ImageNameRegex = "Final",
    ///         }));
    ///         var imageId = @default.Apply(@default =&gt; @default.ImageId);
    ///         var sgRule = new Tencentcloud.Security.GroupLiteRule("sgRule", new Tencentcloud.Security.GroupLiteRuleArgs
    ///         {
    ///             SecurityGroupId = sg.Id,
    ///             Ingresses = 
    ///             {
    ///                 "ACCEPT#10.0.0.0/16#ALL#ALL",
    ///                 "ACCEPT#172.16.0.0/22#ALL#ALL",
    ///                 "DROP#0.0.0.0/0#ALL#ALL",
    ///             },
    ///             Egresses = 
    ///             {
    ///                 "ACCEPT#172.16.0.0/22#ALL#ALL",
    ///             },
    ///         });
    ///         var example = new Tencentcloud.Kubernetes.Cluster("example", new Tencentcloud.Kubernetes.ClusterArgs
    ///         {
    ///             VpcId = firstVpcId,
    ///             ClusterCidr = exampleClusterCidr,
    ///             ClusterMaxPodNum = 32,
    ///             ClusterName = "tf_example_cluster",
    ///             ClusterDesc = "example for tke cluster",
    ///             ClusterMaxServiceNum = 32,
    ///             ClusterInternet = false,
    ///             ClusterInternetSecurityGroup = sgId,
    ///             ClusterVersion = "1.22.5",
    ///             ClusterDeployType = "MANAGED_CLUSTER",
    ///             WorkerConfigs = 
    ///             {
    ///                 new Tencentcloud.Kubernetes.Inputs.ClusterWorkerConfigArgs
    ///                 {
    ///                     Count = 1,
    ///                     AvailabilityZone = availabilityZoneFirst,
    ///                     InstanceType = defaultInstanceType,
    ///                     SystemDiskType = "CLOUD_SSD",
    ///                     SystemDiskSize = 60,
    ///                     InternetChargeType = "TRAFFIC_POSTPAID_BY_HOUR",
    ///                     InternetMaxBandwidthOut = 100,
    ///                     PublicIpAssigned = true,
    ///                     SubnetId = firstSubnetId,
    ///                     ImgId = imageId,
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
    ///                 new Tencentcloud.Kubernetes.Inputs.ClusterWorkerConfigArgs
    ///                 {
    ///                     Count = 1,
    ///                     AvailabilityZone = availabilityZoneSecond,
    ///                     InstanceType = defaultInstanceType,
    ///                     SystemDiskType = "CLOUD_SSD",
    ///                     SystemDiskSize = 60,
    ///                     InternetChargeType = "TRAFFIC_POSTPAID_BY_HOUR",
    ///                     InternetMaxBandwidthOut = 100,
    ///                     PublicIpAssigned = true,
    ///                     SubnetId = secondSubnetId,
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
    ///                     CamRoleName = "CVM_QcsRole",
    ///                     Password = "ZZXXccvv1212",
    ///                 },
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "test1", "test1" },
    ///                 { "test2", "test2" },
    ///             },
    ///         });
    ///         var zone = config.Get("zone") ?? "ap-guangzhou";
    ///         var clusterType = config.Get("clusterType") ?? "tke";
    ///         var fooTmpInstance = new Tencentcloud.Monitor.TmpInstance("fooTmpInstance", new Tencentcloud.Monitor.TmpInstanceArgs
    ///         {
    ///             InstanceName = "tf-tmp-instance",
    ///             VpcId = firstVpcId,
    ///             SubnetId = firstSubnetId,
    ///             DataRetentionTime = 30,
    ///             Zone = availabilityZoneSecond,
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "terraform" },
    ///             },
    ///         });
    ///         // tmp tke bind
    ///         var fooTmpTkeClusterAgent = new Tencentcloud.Monitor.TmpTkeClusterAgent("fooTmpTkeClusterAgent", new Tencentcloud.Monitor.TmpTkeClusterAgentArgs
    ///         {
    ///             InstanceId = fooTmpInstance.Id,
    ///             Agents = new Tencentcloud.Monitor.Inputs.TmpTkeClusterAgentAgentsArgs
    ///             {
    ///                 Region = zone,
    ///                 ClusterType = clusterType,
    ///                 ClusterId = example.Id,
    ///                 EnableExternal = false,
    ///             },
    ///         });
    ///         // create record rule
    ///         var basic = new Tencentcloud.Monitor.TmpTkeGlobalNotification("basic", new Tencentcloud.Monitor.TmpTkeGlobalNotificationArgs
    ///         {
    ///             InstanceId = fooTmpInstance.Id,
    ///             Notification = new Tencentcloud.Monitor.Inputs.TmpTkeGlobalNotificationNotificationArgs
    ///             {
    ///                 Enabled = true,
    ///                 Type = "webhook",
    ///                 AlertManagers = 
    ///                 {
    ///                     new Tencentcloud.Monitor.Inputs.TmpTkeGlobalNotificationNotificationAlertManagerArgs
    ///                     {
    ///                         ClusterId = "",
    ///                         ClusterType = "",
    ///                         Url = "",
    ///                     },
    ///                 },
    ///                 WebHook = "",
    ///                 RepeatInterval = "5m",
    ///                 TimeRangeStart = "00:00:00",
    ///                 TimeRangeEnd = "23:59:59",
    ///                 NotifyWays = 
    ///                 {
    ///                     "SMS",
    ///                     "EMAIL",
    ///                 },
    ///                 ReceiverGroups = {},
    ///                 PhoneNotifyOrders = {},
    ///                 PhoneCircleTimes = 0,
    ///                 PhoneInnerInterval = 0,
    ///                 PhoneCircleInterval = 0,
    ///                 PhoneArriveNotice = false,
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 fooTmpTkeClusterAgent,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/tmpTkeGlobalNotification:TmpTkeGlobalNotification")]
    public partial class TmpTkeGlobalNotification : Pulumi.CustomResource
    {
        /// <summary>
        /// Instance Id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Alarm notification channels.
        /// </summary>
        [Output("notification")]
        public Output<Outputs.TmpTkeGlobalNotificationNotification> Notification { get; private set; } = null!;


        /// <summary>
        /// Create a TmpTkeGlobalNotification resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TmpTkeGlobalNotification(string name, TmpTkeGlobalNotificationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeGlobalNotification:TmpTkeGlobalNotification", name, args ?? new TmpTkeGlobalNotificationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TmpTkeGlobalNotification(string name, Input<string> id, TmpTkeGlobalNotificationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeGlobalNotification:TmpTkeGlobalNotification", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TmpTkeGlobalNotification resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TmpTkeGlobalNotification Get(string name, Input<string> id, TmpTkeGlobalNotificationState? state = null, CustomResourceOptions? options = null)
        {
            return new TmpTkeGlobalNotification(name, id, state, options);
        }
    }

    public sealed class TmpTkeGlobalNotificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Alarm notification channels.
        /// </summary>
        [Input("notification", required: true)]
        public Input<Inputs.TmpTkeGlobalNotificationNotificationArgs> Notification { get; set; } = null!;

        public TmpTkeGlobalNotificationArgs()
        {
        }
    }

    public sealed class TmpTkeGlobalNotificationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Alarm notification channels.
        /// </summary>
        [Input("notification")]
        public Input<Inputs.TmpTkeGlobalNotificationNotificationGetArgs>? Notification { get; set; }

        public TmpTkeGlobalNotificationState()
        {
        }
    }
}
