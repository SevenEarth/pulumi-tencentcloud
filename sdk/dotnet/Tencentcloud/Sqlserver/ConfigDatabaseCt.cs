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
    /// Provides a resource to create a sqlserver config_database_ct
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
    ///         var configDatabaseCt = new Tencentcloud.Sqlserver.ConfigDatabaseCt("configDatabaseCt", new Tencentcloud.Sqlserver.ConfigDatabaseCtArgs
    ///         {
    ///             DbName = "keep_pubsub_db2",
    ///             InstanceId = "mssql-qelbzgwf",
    ///             ModifyType = "disable",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// sqlserver tencentcloud_sqlserver_config_database_ct can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Sqlserver/configDatabaseCt:ConfigDatabaseCt config_database_ct config_database_ct_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/configDatabaseCt:ConfigDatabaseCt")]
    public partial class ConfigDatabaseCt : Pulumi.CustomResource
    {
        /// <summary>
        /// Retention period (in days) of change tracking information when CT is enabled. Value range: 3-30. Default value: 3.
        /// </summary>
        [Output("changeRetentionDay")]
        public Output<int> ChangeRetentionDay { get; private set; } = null!;

        /// <summary>
        /// database name.
        /// </summary>
        [Output("dbName")]
        public Output<string> DbName { get; private set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Enable or disable CT. Valid values: enable, disable.
        /// </summary>
        [Output("modifyType")]
        public Output<string> ModifyType { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigDatabaseCt resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigDatabaseCt(string name, ConfigDatabaseCtArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/configDatabaseCt:ConfigDatabaseCt", name, args ?? new ConfigDatabaseCtArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigDatabaseCt(string name, Input<string> id, ConfigDatabaseCtState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/configDatabaseCt:ConfigDatabaseCt", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigDatabaseCt resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigDatabaseCt Get(string name, Input<string> id, ConfigDatabaseCtState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigDatabaseCt(name, id, state, options);
        }
    }

    public sealed class ConfigDatabaseCtArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Retention period (in days) of change tracking information when CT is enabled. Value range: 3-30. Default value: 3.
        /// </summary>
        [Input("changeRetentionDay")]
        public Input<int>? ChangeRetentionDay { get; set; }

        /// <summary>
        /// database name.
        /// </summary>
        [Input("dbName", required: true)]
        public Input<string> DbName { get; set; } = null!;

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Enable or disable CT. Valid values: enable, disable.
        /// </summary>
        [Input("modifyType", required: true)]
        public Input<string> ModifyType { get; set; } = null!;

        public ConfigDatabaseCtArgs()
        {
        }
    }

    public sealed class ConfigDatabaseCtState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Retention period (in days) of change tracking information when CT is enabled. Value range: 3-30. Default value: 3.
        /// </summary>
        [Input("changeRetentionDay")]
        public Input<int>? ChangeRetentionDay { get; set; }

        /// <summary>
        /// database name.
        /// </summary>
        [Input("dbName")]
        public Input<string>? DbName { get; set; }

        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Enable or disable CT. Valid values: enable, disable.
        /// </summary>
        [Input("modifyType")]
        public Input<string>? ModifyType { get; set; }

        public ConfigDatabaseCtState()
        {
        }
    }
}
