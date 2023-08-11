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
    /// Provides a resource to create a vpc local_gateway
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
    ///         var vpc = new Tencentcloud.Vpc.Instance("vpc", new Tencentcloud.Vpc.InstanceArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///         });
    ///         var example = new Tencentcloud.Vpc.LocalGateway("example", new Tencentcloud.Vpc.LocalGatewayArgs
    ///         {
    ///             LocalGatewayName = "tf-example",
    ///             VpcId = vpc.Id,
    ///             CdcId = "cluster-j9gyu1iy",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// vpc local_gateway can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Vpc/localGateway:LocalGateway local_gateway local_gateway_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Vpc/localGateway:LocalGateway")]
    public partial class LocalGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// CDC instance ID.
        /// </summary>
        [Output("cdcId")]
        public Output<string> CdcId { get; private set; } = null!;

        /// <summary>
        /// Local gateway name.
        /// </summary>
        [Output("localGatewayName")]
        public Output<string> LocalGatewayName { get; private set; } = null!;

        /// <summary>
        /// VPC instance ID.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a LocalGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocalGateway(string name, LocalGatewayArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/localGateway:LocalGateway", name, args ?? new LocalGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocalGateway(string name, Input<string> id, LocalGatewayState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Vpc/localGateway:LocalGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LocalGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocalGateway Get(string name, Input<string> id, LocalGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new LocalGateway(name, id, state, options);
        }
    }

    public sealed class LocalGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CDC instance ID.
        /// </summary>
        [Input("cdcId", required: true)]
        public Input<string> CdcId { get; set; } = null!;

        /// <summary>
        /// Local gateway name.
        /// </summary>
        [Input("localGatewayName", required: true)]
        public Input<string> LocalGatewayName { get; set; } = null!;

        /// <summary>
        /// VPC instance ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public LocalGatewayArgs()
        {
        }
    }

    public sealed class LocalGatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CDC instance ID.
        /// </summary>
        [Input("cdcId")]
        public Input<string>? CdcId { get; set; }

        /// <summary>
        /// Local gateway name.
        /// </summary>
        [Input("localGatewayName")]
        public Input<string>? LocalGatewayName { get; set; }

        /// <summary>
        /// VPC instance ID.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public LocalGatewayState()
        {
        }
    }
}
