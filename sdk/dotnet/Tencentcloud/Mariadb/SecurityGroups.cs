// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb
{
    /// <summary>
    /// Provides a resource to create a mariadb security_groups
    /// 
    /// &gt; **NOTE:** If you use this resource, please do not use security_group_ids in tencentcloud.Mariadb.Instance resource
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
    ///         var securityGroups = new Tencentcloud.Mariadb.SecurityGroups("securityGroups", new Tencentcloud.Mariadb.SecurityGroupsArgs
    ///         {
    ///             InstanceId = "tdsql-4pzs5b67",
    ///             Product = "mariadb",
    ///             SecurityGroupId = "sg-7kpsbxdb",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// mariadb security_groups can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Mariadb/securityGroups:SecurityGroups security_groups tdsql-4pzs5b67#sg-7kpsbxdb#mariadb
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mariadb/securityGroups:SecurityGroups")]
    public partial class SecurityGroups : Pulumi.CustomResource
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// product name, fixed to mariadb.
        /// </summary>
        [Output("product")]
        public Output<string> Product { get; private set; } = null!;

        /// <summary>
        /// security group id.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroups resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroups(string name, SecurityGroupsArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mariadb/securityGroups:SecurityGroups", name, args ?? new SecurityGroupsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroups(string name, Input<string> id, SecurityGroupsState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mariadb/securityGroups:SecurityGroups", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityGroups resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroups Get(string name, Input<string> id, SecurityGroupsState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroups(name, id, state, options);
        }
    }

    public sealed class SecurityGroupsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// product name, fixed to mariadb.
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        /// <summary>
        /// security group id.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        public SecurityGroupsArgs()
        {
        }
    }

    public sealed class SecurityGroupsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// product name, fixed to mariadb.
        /// </summary>
        [Input("product")]
        public Input<string>? Product { get; set; }

        /// <summary>
        /// security group id.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        public SecurityGroupsState()
        {
        }
    }
}
