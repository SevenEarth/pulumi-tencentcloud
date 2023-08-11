// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Sqlserver
{
    /// <summary>
    /// Provides a resource to create a sqlserver general_cloud_ro_instance
    /// 
    /// ## Example Usage
    /// ### If read_only_group_type value is 1 - Ship according to one instance and one read-only group:
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
    ///         var zones = Output.Create(Tencentcloud.Availability.GetZonesByProduct.InvokeAsync(new Tencentcloud.Availability.GetZonesByProductArgs
    ///         {
    ///             Product = "sqlserver",
    ///         }));
    ///         var vpc = new Tencentcloud.Vpc.Instance("vpc", new Tencentcloud.Vpc.InstanceArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///         });
    ///         var subnet = new Tencentcloud.Subnet.Instance("subnet", new Tencentcloud.Subnet.InstanceArgs
    ///         {
    ///             AvailabilityZone = zones.Apply(zones =&gt; zones.Zones?[4]?.Name),
    ///             VpcId = vpc.Id,
    ///             CidrBlock = "10.0.0.0/16",
    ///             IsMulticast = false,
    ///         });
    ///         var securityGroup = new Tencentcloud.Security.Group("securityGroup", new Tencentcloud.Security.GroupArgs
    ///         {
    ///             Description = "desc.",
    ///         });
    ///         var exampleGeneralCloudInstance = new Tencentcloud.Sqlserver.GeneralCloudInstance("exampleGeneralCloudInstance", new Tencentcloud.Sqlserver.GeneralCloudInstanceArgs
    ///         {
    ///             Zone = zones.Apply(zones =&gt; zones.Zones?[4]?.Name),
    ///             Memory = 4,
    ///             Storage = 100,
    ///             Cpu = 2,
    ///             MachineType = "CLOUD_HSSD",
    ///             InstanceChargeType = "POSTPAID",
    ///             ProjectId = 0,
    ///             SubnetId = subnet.Id,
    ///             VpcId = vpc.Id,
    ///             DbVersion = "2008R2",
    ///             SecurityGroupLists = 
    ///             {
    ///                 securityGroup.Id,
    ///             },
    ///             Weeklies = 
    ///             {
    ///                 1,
    ///                 2,
    ///                 3,
    ///                 5,
    ///                 6,
    ///                 7,
    ///             },
    ///             StartTime = "00:00",
    ///             Span = 6,
    ///             ResourceTags = 
    ///             {
    ///                 new Tencentcloud.Sqlserver.Inputs.GeneralCloudInstanceResourceTagArgs
    ///                 {
    ///                     TagKey = "test",
    ///                     TagValue = "test",
    ///                 },
    ///             },
    ///             Collation = "Chinese_PRC_CI_AS",
    ///             TimeZone = "China Standard Time",
    ///         });
    ///         var exampleGeneralCloudRoInstance = new Tencentcloud.Sqlserver.GeneralCloudRoInstance("exampleGeneralCloudRoInstance", new Tencentcloud.Sqlserver.GeneralCloudRoInstanceArgs
    ///         {
    ///             InstanceId = exampleGeneralCloudInstance.Id,
    ///             Zone = zones.Apply(zones =&gt; zones.Zones?[4]?.Name),
    ///             ReadOnlyGroupType = 1,
    ///             Memory = 4,
    ///             Storage = 100,
    ///             Cpu = 2,
    ///             MachineType = "CLOUD_BSSD",
    ///             InstanceChargeType = "POSTPAID",
    ///             SubnetId = subnet.Id,
    ///             VpcId = vpc.Id,
    ///             SecurityGroupLists = 
    ///             {
    ///                 securityGroup.Id,
    ///             },
    ///             Collation = "Chinese_PRC_CI_AS",
    ///             TimeZone = "China Standard Time",
    ///             ResourceTags = 
    ///             {
    ///                 { "test-key1", "test-value1" },
    ///                 { "test-key2", "test-value2" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### If read_only_group_type value is 2 - Ship after creating a read-only group, all instances are under this read-only group:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Tencentcloud.Sqlserver.GeneralCloudRoInstance("example", new Tencentcloud.Sqlserver.GeneralCloudRoInstanceArgs
    ///         {
    ///             InstanceId = tencentcloud_sqlserver_general_cloud_instance.Example.Id,
    ///             Zone = data.Tencentcloud_availability_zones_by_product.Zones.Zones[4].Name,
    ///             ReadOnlyGroupType = 2,
    ///             ReadOnlyGroupName = "test-ro-group",
    ///             ReadOnlyGroupIsOfflineDelay = 1,
    ///             ReadOnlyGroupMaxDelayTime = 10,
    ///             ReadOnlyGroupMinInGroup = 1,
    ///             Memory = 4,
    ///             Storage = 100,
    ///             Cpu = 2,
    ///             MachineType = "CLOUD_BSSD",
    ///             InstanceChargeType = "POSTPAID",
    ///             SubnetId = tencentcloud_subnet.Subnet.Id,
    ///             VpcId = tencentcloud_vpc.Vpc.Id,
    ///             SecurityGroupLists = 
    ///             {
    ///                 tencentcloud_security_group.Security_group.Id,
    ///             },
    ///             Collation = "Chinese_PRC_CI_AS",
    ///             TimeZone = "China Standard Time",
    ///             ResourceTags = 
    ///             {
    ///                 { "test-key1", "test-value1" },
    ///                 { "test-key2", "test-value2" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### If read_only_group_type value is 3 - All instances shipped are in the existing Some read-only groups below:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Tencentcloud.Sqlserver.GeneralCloudRoInstance("example", new Tencentcloud.Sqlserver.GeneralCloudRoInstanceArgs
    ///         {
    ///             InstanceId = tencentcloud_sqlserver_general_cloud_instance.Example.Id,
    ///             Zone = data.Tencentcloud_availability_zones_by_product.Zones.Zones[4].Name,
    ///             ReadOnlyGroupType = 3,
    ///             Memory = 4,
    ///             Storage = 100,
    ///             Cpu = 2,
    ///             MachineType = "CLOUD_BSSD",
    ///             ReadOnlyGroupId = "mssqlrg-clboghrj",
    ///             InstanceChargeType = "POSTPAID",
    ///             SubnetId = tencentcloud_subnet.Subnet.Id,
    ///             VpcId = tencentcloud_vpc.Vpc.Id,
    ///             SecurityGroupLists = 
    ///             {
    ///                 tencentcloud_security_group.Security_group.Id,
    ///             },
    ///             Collation = "Chinese_PRC_CI_AS",
    ///             TimeZone = "China Standard Time",
    ///             ResourceTags = 
    ///             {
    ///                 { "test-key1", "test-value1" },
    ///                 { "test-key2", "test-value2" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/generalCloudRoInstance:GeneralCloudRoInstance")]
    public partial class GeneralCloudRoInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// System character set collation, default: Chinese_PRC_CI_AS.
        /// </summary>
        [Output("collation")]
        public Output<string?> Collation { get; private set; } = null!;

        /// <summary>
        /// Number of instance cores.
        /// </summary>
        [Output("cpu")]
        public Output<int> Cpu { get; private set; } = null!;

        /// <summary>
        /// Payment mode, the value supports PREPAID (prepaid), POSTPAID (postpaid).
        /// </summary>
        [Output("instanceChargeType")]
        public Output<string?> InstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// Primary instance ID, in the format: mssql-3l3fgqn7.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The host disk type of the purchased instance, CLOUD_HSSD-enhanced SSD cloud disk for virtual machines, CLOUD_TSSD-extremely fast SSD cloud disk for virtual machines, CLOUD_BSSD-universal SSD cloud disk for virtual machines.
        /// </summary>
        [Output("machineType")]
        public Output<string> MachineType { get; private set; } = null!;

        /// <summary>
        /// Instance memory size, in GB.
        /// </summary>
        [Output("memory")]
        public Output<int> Memory { get; private set; } = null!;

        /// <summary>
        /// Purchase instance period, the default value is 1, which means one month. The value cannot exceed 48.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// Required when ReadOnlyGroupType=3, existing read-only group ID.
        /// </summary>
        [Output("readOnlyGroupId")]
        public Output<string> ReadOnlyGroupId { get; private set; } = null!;

        /// <summary>
        /// Required when ReadOnlyGroupType=2, whether to enable the delayed elimination function for the newly created read-only group, 1-on, 0-off. When the delay between the read-only replica and the primary instance is greater than the threshold, it will be automatically removed.
        /// </summary>
        [Output("readOnlyGroupIsOfflineDelay")]
        public Output<int> ReadOnlyGroupIsOfflineDelay { get; private set; } = null!;

        /// <summary>
        /// Mandatory when ReadOnlyGroupType=2 and ReadOnlyGroupIsOfflineDelay=1, the threshold for delay culling of newly created read-only groups.
        /// </summary>
        [Output("readOnlyGroupMaxDelayTime")]
        public Output<int> ReadOnlyGroupMaxDelayTime { get; private set; } = null!;

        /// <summary>
        /// Required when ReadOnlyGroupType=2 and ReadOnlyGroupIsOfflineDelay=1, the newly created read-only group retains at least the number of read-only replicas after delay elimination.
        /// </summary>
        [Output("readOnlyGroupMinInGroup")]
        public Output<int> ReadOnlyGroupMinInGroup { get; private set; } = null!;

        /// <summary>
        /// Required when ReadOnlyGroupType=2, the name of the newly created read-only group.
        /// </summary>
        [Output("readOnlyGroupName")]
        public Output<string> ReadOnlyGroupName { get; private set; } = null!;

        /// <summary>
        /// Read-only group type option, 1- Ship according to one instance and one read-only group, 2 - Ship after creating a read-only group, all instances are under this read-only group, 3 - All instances shipped are in the existing Some read-only groups below.
        /// </summary>
        [Output("readOnlyGroupType")]
        public Output<int> ReadOnlyGroupType { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("resourceTags")]
        public Output<ImmutableDictionary<string, object>?> ResourceTags { get; private set; } = null!;

        /// <summary>
        /// Primary read only instance ID, in the format: mssqlro-lbljc5qd.
        /// </summary>
        [Output("roInstanceId")]
        public Output<string> RoInstanceId { get; private set; } = null!;

        /// <summary>
        /// Security group list, fill in the security group ID in the form of sg-xxx.
        /// </summary>
        [Output("securityGroupLists")]
        public Output<ImmutableArray<string>> SecurityGroupLists { get; private set; } = null!;

        /// <summary>
        /// Instance disk size, in GB.
        /// </summary>
        [Output("storage")]
        public Output<int> Storage { get; private set; } = null!;

        /// <summary>
        /// VPC subnet ID, in the form of subnet-bdoe83fa; SubnetId and VpcId need to be set at the same time or not set at the same time.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// System time zone, default: China Standard Time.
        /// </summary>
        [Output("timeZone")]
        public Output<string?> TimeZone { get; private set; } = null!;

        /// <summary>
        /// VPC network ID, in the form of vpc-dsp338hz; SubnetId and VpcId need to be set at the same time or not set at the same time.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;

        /// <summary>
        /// Instance Availability Zone, similar to ap-guangzhou-1 (Guangzhou District 1); the instance sales area can be obtained through the interface DescribeZones.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a GeneralCloudRoInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GeneralCloudRoInstance(string name, GeneralCloudRoInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/generalCloudRoInstance:GeneralCloudRoInstance", name, args ?? new GeneralCloudRoInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GeneralCloudRoInstance(string name, Input<string> id, GeneralCloudRoInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/generalCloudRoInstance:GeneralCloudRoInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GeneralCloudRoInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GeneralCloudRoInstance Get(string name, Input<string> id, GeneralCloudRoInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new GeneralCloudRoInstance(name, id, state, options);
        }
    }

    public sealed class GeneralCloudRoInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// System character set collation, default: Chinese_PRC_CI_AS.
        /// </summary>
        [Input("collation")]
        public Input<string>? Collation { get; set; }

        /// <summary>
        /// Number of instance cores.
        /// </summary>
        [Input("cpu", required: true)]
        public Input<int> Cpu { get; set; } = null!;

        /// <summary>
        /// Payment mode, the value supports PREPAID (prepaid), POSTPAID (postpaid).
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// Primary instance ID, in the format: mssql-3l3fgqn7.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The host disk type of the purchased instance, CLOUD_HSSD-enhanced SSD cloud disk for virtual machines, CLOUD_TSSD-extremely fast SSD cloud disk for virtual machines, CLOUD_BSSD-universal SSD cloud disk for virtual machines.
        /// </summary>
        [Input("machineType", required: true)]
        public Input<string> MachineType { get; set; } = null!;

        /// <summary>
        /// Instance memory size, in GB.
        /// </summary>
        [Input("memory", required: true)]
        public Input<int> Memory { get; set; } = null!;

        /// <summary>
        /// Purchase instance period, the default value is 1, which means one month. The value cannot exceed 48.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Required when ReadOnlyGroupType=3, existing read-only group ID.
        /// </summary>
        [Input("readOnlyGroupId")]
        public Input<string>? ReadOnlyGroupId { get; set; }

        /// <summary>
        /// Required when ReadOnlyGroupType=2, whether to enable the delayed elimination function for the newly created read-only group, 1-on, 0-off. When the delay between the read-only replica and the primary instance is greater than the threshold, it will be automatically removed.
        /// </summary>
        [Input("readOnlyGroupIsOfflineDelay")]
        public Input<int>? ReadOnlyGroupIsOfflineDelay { get; set; }

        /// <summary>
        /// Mandatory when ReadOnlyGroupType=2 and ReadOnlyGroupIsOfflineDelay=1, the threshold for delay culling of newly created read-only groups.
        /// </summary>
        [Input("readOnlyGroupMaxDelayTime")]
        public Input<int>? ReadOnlyGroupMaxDelayTime { get; set; }

        /// <summary>
        /// Required when ReadOnlyGroupType=2 and ReadOnlyGroupIsOfflineDelay=1, the newly created read-only group retains at least the number of read-only replicas after delay elimination.
        /// </summary>
        [Input("readOnlyGroupMinInGroup")]
        public Input<int>? ReadOnlyGroupMinInGroup { get; set; }

        /// <summary>
        /// Required when ReadOnlyGroupType=2, the name of the newly created read-only group.
        /// </summary>
        [Input("readOnlyGroupName")]
        public Input<string>? ReadOnlyGroupName { get; set; }

        /// <summary>
        /// Read-only group type option, 1- Ship according to one instance and one read-only group, 2 - Ship after creating a read-only group, all instances are under this read-only group, 3 - All instances shipped are in the existing Some read-only groups below.
        /// </summary>
        [Input("readOnlyGroupType", required: true)]
        public Input<int> ReadOnlyGroupType { get; set; } = null!;

        [Input("resourceTags")]
        private InputMap<object>? _resourceTags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> ResourceTags
        {
            get => _resourceTags ?? (_resourceTags = new InputMap<object>());
            set => _resourceTags = value;
        }

        [Input("securityGroupLists")]
        private InputList<string>? _securityGroupLists;

        /// <summary>
        /// Security group list, fill in the security group ID in the form of sg-xxx.
        /// </summary>
        public InputList<string> SecurityGroupLists
        {
            get => _securityGroupLists ?? (_securityGroupLists = new InputList<string>());
            set => _securityGroupLists = value;
        }

        /// <summary>
        /// Instance disk size, in GB.
        /// </summary>
        [Input("storage", required: true)]
        public Input<int> Storage { get; set; } = null!;

        /// <summary>
        /// VPC subnet ID, in the form of subnet-bdoe83fa; SubnetId and VpcId need to be set at the same time or not set at the same time.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// System time zone, default: China Standard Time.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        /// <summary>
        /// VPC network ID, in the form of vpc-dsp338hz; SubnetId and VpcId need to be set at the same time or not set at the same time.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Instance Availability Zone, similar to ap-guangzhou-1 (Guangzhou District 1); the instance sales area can be obtained through the interface DescribeZones.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GeneralCloudRoInstanceArgs()
        {
        }
    }

    public sealed class GeneralCloudRoInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// System character set collation, default: Chinese_PRC_CI_AS.
        /// </summary>
        [Input("collation")]
        public Input<string>? Collation { get; set; }

        /// <summary>
        /// Number of instance cores.
        /// </summary>
        [Input("cpu")]
        public Input<int>? Cpu { get; set; }

        /// <summary>
        /// Payment mode, the value supports PREPAID (prepaid), POSTPAID (postpaid).
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// Primary instance ID, in the format: mssql-3l3fgqn7.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The host disk type of the purchased instance, CLOUD_HSSD-enhanced SSD cloud disk for virtual machines, CLOUD_TSSD-extremely fast SSD cloud disk for virtual machines, CLOUD_BSSD-universal SSD cloud disk for virtual machines.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        /// <summary>
        /// Instance memory size, in GB.
        /// </summary>
        [Input("memory")]
        public Input<int>? Memory { get; set; }

        /// <summary>
        /// Purchase instance period, the default value is 1, which means one month. The value cannot exceed 48.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// Required when ReadOnlyGroupType=3, existing read-only group ID.
        /// </summary>
        [Input("readOnlyGroupId")]
        public Input<string>? ReadOnlyGroupId { get; set; }

        /// <summary>
        /// Required when ReadOnlyGroupType=2, whether to enable the delayed elimination function for the newly created read-only group, 1-on, 0-off. When the delay between the read-only replica and the primary instance is greater than the threshold, it will be automatically removed.
        /// </summary>
        [Input("readOnlyGroupIsOfflineDelay")]
        public Input<int>? ReadOnlyGroupIsOfflineDelay { get; set; }

        /// <summary>
        /// Mandatory when ReadOnlyGroupType=2 and ReadOnlyGroupIsOfflineDelay=1, the threshold for delay culling of newly created read-only groups.
        /// </summary>
        [Input("readOnlyGroupMaxDelayTime")]
        public Input<int>? ReadOnlyGroupMaxDelayTime { get; set; }

        /// <summary>
        /// Required when ReadOnlyGroupType=2 and ReadOnlyGroupIsOfflineDelay=1, the newly created read-only group retains at least the number of read-only replicas after delay elimination.
        /// </summary>
        [Input("readOnlyGroupMinInGroup")]
        public Input<int>? ReadOnlyGroupMinInGroup { get; set; }

        /// <summary>
        /// Required when ReadOnlyGroupType=2, the name of the newly created read-only group.
        /// </summary>
        [Input("readOnlyGroupName")]
        public Input<string>? ReadOnlyGroupName { get; set; }

        /// <summary>
        /// Read-only group type option, 1- Ship according to one instance and one read-only group, 2 - Ship after creating a read-only group, all instances are under this read-only group, 3 - All instances shipped are in the existing Some read-only groups below.
        /// </summary>
        [Input("readOnlyGroupType")]
        public Input<int>? ReadOnlyGroupType { get; set; }

        [Input("resourceTags")]
        private InputMap<object>? _resourceTags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> ResourceTags
        {
            get => _resourceTags ?? (_resourceTags = new InputMap<object>());
            set => _resourceTags = value;
        }

        /// <summary>
        /// Primary read only instance ID, in the format: mssqlro-lbljc5qd.
        /// </summary>
        [Input("roInstanceId")]
        public Input<string>? RoInstanceId { get; set; }

        [Input("securityGroupLists")]
        private InputList<string>? _securityGroupLists;

        /// <summary>
        /// Security group list, fill in the security group ID in the form of sg-xxx.
        /// </summary>
        public InputList<string> SecurityGroupLists
        {
            get => _securityGroupLists ?? (_securityGroupLists = new InputList<string>());
            set => _securityGroupLists = value;
        }

        /// <summary>
        /// Instance disk size, in GB.
        /// </summary>
        [Input("storage")]
        public Input<int>? Storage { get; set; }

        /// <summary>
        /// VPC subnet ID, in the form of subnet-bdoe83fa; SubnetId and VpcId need to be set at the same time or not set at the same time.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// System time zone, default: China Standard Time.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        /// <summary>
        /// VPC network ID, in the form of vpc-dsp338hz; SubnetId and VpcId need to be set at the same time or not set at the same time.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Instance Availability Zone, similar to ap-guangzhou-1 (Guangzhou District 1); the instance sales area can be obtained through the interface DescribeZones.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GeneralCloudRoInstanceState()
        {
        }
    }
}
