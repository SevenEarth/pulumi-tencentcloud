// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb
{
    /// <summary>
    /// Provides a resource to create a dcdb encrypt_attributes_config
    /// 
    /// &gt; **NOTE:**  This resource currently only supports the newly created MySQL 8.0.24 version.
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
    ///         var @internal = Output.Create(Tencentcloud.Security.GetGroups.InvokeAsync(new Tencentcloud.Security.GetGroupsArgs
    ///         {
    ///             Name = "default",
    ///         }));
    ///         var vpc = Output.Create(Tencentcloud.Vpc.GetInstances.InvokeAsync(new Tencentcloud.Vpc.GetInstancesArgs
    ///         {
    ///             Name = "Default-VPC",
    ///         }));
    ///         var subnet = vpc.Apply(vpc =&gt; Output.Create(Tencentcloud.Vpc.GetSubnets.InvokeAsync(new Tencentcloud.Vpc.GetSubnetsArgs
    ///         {
    ///             VpcId = vpc.InstanceLists?[0]?.VpcId,
    ///         })));
    ///         var vpcId = subnet.Apply(subnet =&gt; subnet.InstanceLists?[0]?.VpcId);
    ///         var subnetId = subnet.Apply(subnet =&gt; subnet.InstanceLists?[0]?.SubnetId);
    ///         var sgId = @internal.Apply(@internal =&gt; @internal.SecurityGroups?[0]?.SecurityGroupId);
    ///         var prepaidInstance = new Tencentcloud.Dcdb.DbInstance("prepaidInstance", new Tencentcloud.Dcdb.DbInstanceArgs
    ///         {
    ///             InstanceName = "test_dcdb_db_post_instance",
    ///             Zones = 
    ///             {
    ///                 @var.Default_az,
    ///             },
    ///             Period = 1,
    ///             ShardMemory = 2,
    ///             ShardStorage = 10,
    ///             ShardNodeCount = 2,
    ///             ShardCount = 2,
    ///             VpcId = vpcId,
    ///             SubnetId = subnetId,
    ///             DbVersionId = "8.0",
    ///             ResourceTags = 
    ///             {
    ///                 new Tencentcloud.Dcdb.Inputs.DbInstanceResourceTagArgs
    ///                 {
    ///                     TagKey = "aaa",
    ///                     TagValue = "bbb",
    ///                 },
    ///             },
    ///             SecurityGroupIds = 
    ///             {
    ///                 sgId,
    ///             },
    ///         });
    ///         var hourdbInstance = new Tencentcloud.Dcdb.HourdbInstance("hourdbInstance", new Tencentcloud.Dcdb.HourdbInstanceArgs
    ///         {
    ///             InstanceName = "test_dcdb_db_hourdb_instance",
    ///             Zones = 
    ///             {
    ///                 @var.Default_az,
    ///             },
    ///             ShardMemory = 2,
    ///             ShardStorage = 10,
    ///             ShardNodeCount = 2,
    ///             ShardCount = 2,
    ///             VpcId = vpcId,
    ///             SubnetId = subnetId,
    ///             SecurityGroupId = sgId,
    ///             DbVersionId = "8.0",
    ///             ResourceTags = 
    ///             {
    ///                 new Tencentcloud.Dcdb.Inputs.HourdbInstanceResourceTagArgs
    ///                 {
    ///                     TagKey = "aaa",
    ///                     TagValue = "bbb",
    ///                 },
    ///             },
    ///         });
    ///         var prepaidDcdbId = prepaidInstance.Id;
    ///         var hourdbDcdbId = hourdbInstance.Id;
    ///         // for postpaid instance
    ///         var configHourdb = new Tencentcloud.Dcdb.EncryptAttributesConfig("configHourdb", new Tencentcloud.Dcdb.EncryptAttributesConfigArgs
    ///         {
    ///             InstanceId = hourdbDcdbId,
    ///             EncryptEnabled = 1,
    ///         });
    ///         // for prepaid instance
    ///         var configPrepaid = new Tencentcloud.Dcdb.EncryptAttributesConfig("configPrepaid", new Tencentcloud.Dcdb.EncryptAttributesConfigArgs
    ///         {
    ///             InstanceId = prepaidDcdbId,
    ///             EncryptEnabled = 1,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// dcdb encrypt_attributes_config can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dcdb/encryptAttributesConfig:EncryptAttributesConfig encrypt_attributes_config encrypt_attributes_config_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dcdb/encryptAttributesConfig:EncryptAttributesConfig")]
    public partial class EncryptAttributesConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
        /// </summary>
        [Output("encryptEnabled")]
        public Output<int> EncryptEnabled { get; private set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a EncryptAttributesConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EncryptAttributesConfig(string name, EncryptAttributesConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/encryptAttributesConfig:EncryptAttributesConfig", name, args ?? new EncryptAttributesConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EncryptAttributesConfig(string name, Input<string> id, EncryptAttributesConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dcdb/encryptAttributesConfig:EncryptAttributesConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EncryptAttributesConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EncryptAttributesConfig Get(string name, Input<string> id, EncryptAttributesConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new EncryptAttributesConfig(name, id, state, options);
        }
    }

    public sealed class EncryptAttributesConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
        /// </summary>
        [Input("encryptEnabled", required: true)]
        public Input<int> EncryptEnabled { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public EncryptAttributesConfigArgs()
        {
        }
    }

    public sealed class EncryptAttributesConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
        /// </summary>
        [Input("encryptEnabled")]
        public Input<int>? EncryptEnabled { get; set; }

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public EncryptAttributesConfigState()
        {
        }
    }
}
