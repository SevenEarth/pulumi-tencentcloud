// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Postgresql
{
    /// <summary>
    /// Provides a resource to create a postgresql security_group_config
    /// 
    /// ## Example Usage
    /// ### Set security group for the sepcified postgres instance
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var securityGroupConfig = new Tencentcloud.Postgresql.SecurityGroupConfig("securityGroupConfig", new Tencentcloud.Postgresql.SecurityGroupConfigArgs
    ///         {
    ///             SecurityGroupIdSets = 
    ///             {
    ///                 local.Sg_id,
    ///                 local.Sg_id2,
    ///             },
    ///             DbInstanceId = local.Pgsql_id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Set security group for the specified readonly group
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @group = new Tencentcloud.Postgresql.ReadonlyGroup("group", new Tencentcloud.Postgresql.ReadonlyGroupArgs
    ///         {
    ///             MasterDbInstanceId = local.Pgsql_id,
    ///             ProjectId = 0,
    ///             SubnetId = local.Subnet_id,
    ///             VpcId = local.Vpc_id,
    ///             ReplayLagEliminate = 1,
    ///             ReplayLatencyEliminate = 1,
    ///             MaxReplayLag = 100,
    ///             MaxReplayLatency = 512,
    ///             MinDelayEliminateReserve = 1,
    ///         });
    ///         var securityGroupConfig = new Tencentcloud.Postgresql.SecurityGroupConfig("securityGroupConfig", new Tencentcloud.Postgresql.SecurityGroupConfigArgs
    ///         {
    ///             SecurityGroupIdSets = 
    ///             {
    ///                 local.Sg_id,
    ///                 local.Sg_id2,
    ///             },
    ///             ReadOnlyGroupId = @group.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Postgresql/securityGroupConfig:SecurityGroupConfig")]
    public partial class SecurityGroupConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// Instance ID. Either this parameter or ReadOnlyGroupId must be passed in. If both parameters are passed in, ReadOnlyGroupId will be ignored.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string?> DbInstanceId { get; private set; } = null!;

        /// <summary>
        /// RO group ID. Either this parameter or DBInstanceId must be passed in. To query the security groups associated with the RO groups, only pass in ReadOnlyGroupId.
        /// </summary>
        [Output("readOnlyGroupId")]
        public Output<string?> ReadOnlyGroupId { get; private set; } = null!;

        /// <summary>
        /// Information of security groups in array.
        /// </summary>
        [Output("securityGroupIdSets")]
        public Output<ImmutableArray<string>> SecurityGroupIdSets { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroupConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroupConfig(string name, SecurityGroupConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/securityGroupConfig:SecurityGroupConfig", name, args ?? new SecurityGroupConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroupConfig(string name, Input<string> id, SecurityGroupConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Postgresql/securityGroupConfig:SecurityGroupConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityGroupConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroupConfig Get(string name, Input<string> id, SecurityGroupConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroupConfig(name, id, state, options);
        }
    }

    public sealed class SecurityGroupConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID. Either this parameter or ReadOnlyGroupId must be passed in. If both parameters are passed in, ReadOnlyGroupId will be ignored.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// RO group ID. Either this parameter or DBInstanceId must be passed in. To query the security groups associated with the RO groups, only pass in ReadOnlyGroupId.
        /// </summary>
        [Input("readOnlyGroupId")]
        public Input<string>? ReadOnlyGroupId { get; set; }

        [Input("securityGroupIdSets", required: true)]
        private InputList<string>? _securityGroupIdSets;

        /// <summary>
        /// Information of security groups in array.
        /// </summary>
        public InputList<string> SecurityGroupIdSets
        {
            get => _securityGroupIdSets ?? (_securityGroupIdSets = new InputList<string>());
            set => _securityGroupIdSets = value;
        }

        public SecurityGroupConfigArgs()
        {
        }
    }

    public sealed class SecurityGroupConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID. Either this parameter or ReadOnlyGroupId must be passed in. If both parameters are passed in, ReadOnlyGroupId will be ignored.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// RO group ID. Either this parameter or DBInstanceId must be passed in. To query the security groups associated with the RO groups, only pass in ReadOnlyGroupId.
        /// </summary>
        [Input("readOnlyGroupId")]
        public Input<string>? ReadOnlyGroupId { get; set; }

        [Input("securityGroupIdSets")]
        private InputList<string>? _securityGroupIdSets;

        /// <summary>
        /// Information of security groups in array.
        /// </summary>
        public InputList<string> SecurityGroupIdSets
        {
            get => _securityGroupIdSets ?? (_securityGroupIdSets = new InputList<string>());
            set => _securityGroupIdSets = value;
        }

        public SecurityGroupConfigState()
        {
        }
    }
}
