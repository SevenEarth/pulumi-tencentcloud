// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql
{
    /// <summary>
    /// Provides a resource to create a mysql switch_for_upgrade
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
    ///         var switchForUpgrade = new Tencentcloud.Mysql.SwitchForUpgrade("switchForUpgrade", new Tencentcloud.Mysql.SwitchForUpgradeArgs
    ///         {
    ///             InstanceId = "cdb-d9gbh7lt",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/switchForUpgrade:SwitchForUpgrade")]
    public partial class SwitchForUpgrade : Pulumi.CustomResource
    {
        /// <summary>
        /// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchForUpgrade resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchForUpgrade(string name, SwitchForUpgradeArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/switchForUpgrade:SwitchForUpgrade", name, args ?? new SwitchForUpgradeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchForUpgrade(string name, Input<string> id, SwitchForUpgradeState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/switchForUpgrade:SwitchForUpgrade", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SwitchForUpgrade resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchForUpgrade Get(string name, Input<string> id, SwitchForUpgradeState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchForUpgrade(name, id, state, options);
        }
    }

    public sealed class SwitchForUpgradeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public SwitchForUpgradeArgs()
        {
        }
    }

    public sealed class SwitchForUpgradeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance ID in the format of cdb-c1nl9rpv. It is the same as the instance ID displayed on the TencentDB Console page.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public SwitchForUpgradeState()
        {
        }
    }
}
