// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc
{
    /// <summary>
    /// Provides a resource to create a vpc end_point_service_white_list
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
    ///         var endPointServiceWhiteList = new Tencentcloud.Vpc.EndPointServiceWhiteList("endPointServiceWhiteList", new Tencentcloud.Vpc.EndPointServiceWhiteListArgs
    ///         {
    ///             Description = "terraform for test",
    ///             EndPointServiceId = "vpcsvc-69y13tdb",
    ///             UserUin = "100020512675",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// vpc end_point_service_white_list can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Vpc/endPointServiceWhiteList:EndPointServiceWhiteList end_point_service_white_list end_point_service_white_list_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpc/endPointServiceWhiteList:EndPointServiceWhiteList")]
    public partial class EndPointServiceWhiteList : Pulumi.CustomResource
    {
        /// <summary>
        /// Create Time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of white list.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// ID of endpoint service.
        /// </summary>
        [Output("endPointServiceId")]
        public Output<string> EndPointServiceId { get; private set; } = null!;

        /// <summary>
        /// APPID.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// UIN.
        /// </summary>
        [Output("userUin")]
        public Output<string> UserUin { get; private set; } = null!;


        /// <summary>
        /// Create a EndPointServiceWhiteList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndPointServiceWhiteList(string name, EndPointServiceWhiteListArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/endPointServiceWhiteList:EndPointServiceWhiteList", name, args ?? new EndPointServiceWhiteListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndPointServiceWhiteList(string name, Input<string> id, EndPointServiceWhiteListState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/endPointServiceWhiteList:EndPointServiceWhiteList", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EndPointServiceWhiteList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndPointServiceWhiteList Get(string name, Input<string> id, EndPointServiceWhiteListState? state = null, CustomResourceOptions? options = null)
        {
            return new EndPointServiceWhiteList(name, id, state, options);
        }
    }

    public sealed class EndPointServiceWhiteListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of white list.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of endpoint service.
        /// </summary>
        [Input("endPointServiceId", required: true)]
        public Input<string> EndPointServiceId { get; set; } = null!;

        /// <summary>
        /// UIN.
        /// </summary>
        [Input("userUin", required: true)]
        public Input<string> UserUin { get; set; } = null!;

        public EndPointServiceWhiteListArgs()
        {
        }
    }

    public sealed class EndPointServiceWhiteListState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Create Time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description of white list.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// ID of endpoint service.
        /// </summary>
        [Input("endPointServiceId")]
        public Input<string>? EndPointServiceId { get; set; }

        /// <summary>
        /// APPID.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// UIN.
        /// </summary>
        [Input("userUin")]
        public Input<string>? UserUin { get; set; }

        public EndPointServiceWhiteListState()
        {
        }
    }
}
