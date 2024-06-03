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
    /// Provides a resource to create a sqlserver config_instance_security_groups
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
    ///     var configInstanceSecurityGroups = new Tencentcloud.Sqlserver.ConfigInstanceSecurityGroups("configInstanceSecurityGroups", new()
    ///     {
    ///         InstanceId = "mssql-qelbzgwf",
    ///         SecurityGroupIdSets = new[]
    ///         {
    ///             "sg-mayqdlt1",
    ///             "sg-5aubsf8n",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// sqlserver config_instance_security_groups can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Sqlserver/configInstanceSecurityGroups:ConfigInstanceSecurityGroups config_instance_security_groups config_instance_security_groups_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Sqlserver/configInstanceSecurityGroups:ConfigInstanceSecurityGroups")]
    public partial class ConfigInstanceSecurityGroups : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// A list of security group IDs to modify, an array of one or more security group IDs.
        /// </summary>
        [Output("securityGroupIdSets")]
        public Output<ImmutableArray<string>> SecurityGroupIdSets { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigInstanceSecurityGroups resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigInstanceSecurityGroups(string name, ConfigInstanceSecurityGroupsArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/configInstanceSecurityGroups:ConfigInstanceSecurityGroups", name, args ?? new ConfigInstanceSecurityGroupsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigInstanceSecurityGroups(string name, Input<string> id, ConfigInstanceSecurityGroupsState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Sqlserver/configInstanceSecurityGroups:ConfigInstanceSecurityGroups", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigInstanceSecurityGroups resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigInstanceSecurityGroups Get(string name, Input<string> id, ConfigInstanceSecurityGroupsState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigInstanceSecurityGroups(name, id, state, options);
        }
    }

    public sealed class ConfigInstanceSecurityGroupsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("securityGroupIdSets", required: true)]
        private InputList<string>? _securityGroupIdSets;

        /// <summary>
        /// A list of security group IDs to modify, an array of one or more security group IDs.
        /// </summary>
        public InputList<string> SecurityGroupIdSets
        {
            get => _securityGroupIdSets ?? (_securityGroupIdSets = new InputList<string>());
            set => _securityGroupIdSets = value;
        }

        public ConfigInstanceSecurityGroupsArgs()
        {
        }
        public static new ConfigInstanceSecurityGroupsArgs Empty => new ConfigInstanceSecurityGroupsArgs();
    }

    public sealed class ConfigInstanceSecurityGroupsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        [Input("securityGroupIdSets")]
        private InputList<string>? _securityGroupIdSets;

        /// <summary>
        /// A list of security group IDs to modify, an array of one or more security group IDs.
        /// </summary>
        public InputList<string> SecurityGroupIdSets
        {
            get => _securityGroupIdSets ?? (_securityGroupIdSets = new InputList<string>());
            set => _securityGroupIdSets = value;
        }

        public ConfigInstanceSecurityGroupsState()
        {
        }
        public static new ConfigInstanceSecurityGroupsState Empty => new ConfigInstanceSecurityGroupsState();
    }
}
