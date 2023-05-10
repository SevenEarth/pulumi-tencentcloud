// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr
{
    /// <summary>
    /// Use this resource to create tcr vpc attachment to manage access of internal endpoint.
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
    ///         var foo = new Tencentcloud.Tcr.VpcAttachment("foo", new Tencentcloud.Tcr.VpcAttachmentArgs
    ///         {
    ///             InstanceId = "cls-satg5125",
    ///             SubnetId = "subnet-1uwh63so",
    ///             VpcId = "vpc-asg3sfa3",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tcr vpc attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tcr/vpcAttachment:VpcAttachment foo tcrId#vpcId#subnetId
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tcr/vpcAttachment:VpcAttachment")]
    public partial class VpcAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// IP address of the internal access.
        /// </summary>
        [Output("accessIp")]
        public Output<string> AccessIp { get; private set; } = null!;

        /// <summary>
        /// Whether to enable public domain dns. Default value is `false`.
        /// </summary>
        [Output("enablePublicDomainDns")]
        public Output<bool?> EnablePublicDomainDns { get; private set; } = null!;

        /// <summary>
        /// Whether to enable vpc domain dns. Default value is `false`.
        /// </summary>
        [Output("enableVpcDomainDns")]
        public Output<bool?> EnableVpcDomainDns { get; private set; } = null!;

        /// <summary>
        /// ID of the TCR instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// this argument was deprecated, use `region_name` instead. ID of region. Conflict with region_name, can not be set at the same time.
        /// </summary>
        [Output("regionId")]
        public Output<int?> RegionId { get; private set; } = null!;

        /// <summary>
        /// Name of region. Conflict with region_id, can not be set at the same time.
        /// </summary>
        [Output("regionName")]
        public Output<string?> RegionName { get; private set; } = null!;

        /// <summary>
        /// Status of the internal access.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// ID of subnet.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// ID of VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcAttachment(string name, VpcAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/vpcAttachment:VpcAttachment", name, args ?? new VpcAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcAttachment(string name, Input<string> id, VpcAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tcr/vpcAttachment:VpcAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcAttachment Get(string name, Input<string> id, VpcAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcAttachment(name, id, state, options);
        }
    }

    public sealed class VpcAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable public domain dns. Default value is `false`.
        /// </summary>
        [Input("enablePublicDomainDns")]
        public Input<bool>? EnablePublicDomainDns { get; set; }

        /// <summary>
        /// Whether to enable vpc domain dns. Default value is `false`.
        /// </summary>
        [Input("enableVpcDomainDns")]
        public Input<bool>? EnableVpcDomainDns { get; set; }

        /// <summary>
        /// ID of the TCR instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// this argument was deprecated, use `region_name` instead. ID of region. Conflict with region_name, can not be set at the same time.
        /// </summary>
        [Input("regionId")]
        public Input<int>? RegionId { get; set; }

        /// <summary>
        /// Name of region. Conflict with region_id, can not be set at the same time.
        /// </summary>
        [Input("regionName")]
        public Input<string>? RegionName { get; set; }

        /// <summary>
        /// ID of subnet.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        /// <summary>
        /// ID of VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public VpcAttachmentArgs()
        {
        }
    }

    public sealed class VpcAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address of the internal access.
        /// </summary>
        [Input("accessIp")]
        public Input<string>? AccessIp { get; set; }

        /// <summary>
        /// Whether to enable public domain dns. Default value is `false`.
        /// </summary>
        [Input("enablePublicDomainDns")]
        public Input<bool>? EnablePublicDomainDns { get; set; }

        /// <summary>
        /// Whether to enable vpc domain dns. Default value is `false`.
        /// </summary>
        [Input("enableVpcDomainDns")]
        public Input<bool>? EnableVpcDomainDns { get; set; }

        /// <summary>
        /// ID of the TCR instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// this argument was deprecated, use `region_name` instead. ID of region. Conflict with region_name, can not be set at the same time.
        /// </summary>
        [Input("regionId")]
        public Input<int>? RegionId { get; set; }

        /// <summary>
        /// Name of region. Conflict with region_id, can not be set at the same time.
        /// </summary>
        [Input("regionName")]
        public Input<string>? RegionName { get; set; }

        /// <summary>
        /// Status of the internal access.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// ID of subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// ID of VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public VpcAttachmentState()
        {
        }
    }
}
