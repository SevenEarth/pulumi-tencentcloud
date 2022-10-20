// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway
{
    /// <summary>
    /// Use this resource to create IP strategy of API gateway.
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
    ///         var service = new Tencentcloud.ApiGateway.Service("service", new Tencentcloud.ApiGateway.ServiceArgs
    ///         {
    ///             ServiceName = "niceservice",
    ///             Protocol = "http&amp;https",
    ///             ServiceDesc = "your nice service",
    ///             NetTypes = 
    ///             {
    ///                 "INNER",
    ///                 "OUTER",
    ///             },
    ///             IpVersion = "IPv4",
    ///         });
    ///         var test = new Tencentcloud.ApiGateway.IpStrategy("test", new Tencentcloud.ApiGateway.IpStrategyArgs
    ///         {
    ///             ServiceId = service.Id,
    ///             StrategyName = "tf_test",
    ///             StrategyType = "BLACK",
    ///             StrategyData = "9.9.9.9",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// IP strategy of API gateway can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:ApiGateway/ipStrategy:IpStrategy test service-ohxqslqe#IPStrategy-q1lk8ud2
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:ApiGateway/ipStrategy:IpStrategy")]
    public partial class IpStrategy : Pulumi.CustomResource
    {
        /// <summary>
        /// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the API gateway service.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;

        /// <summary>
        /// IP address data.
        /// </summary>
        [Output("strategyData")]
        public Output<string> StrategyData { get; private set; } = null!;

        /// <summary>
        /// IP policy ID.
        /// </summary>
        [Output("strategyId")]
        public Output<string> StrategyId { get; private set; } = null!;

        /// <summary>
        /// User defined strategy name.
        /// </summary>
        [Output("strategyName")]
        public Output<string> StrategyName { get; private set; } = null!;

        /// <summary>
        /// Blacklist or whitelist.
        /// </summary>
        [Output("strategyType")]
        public Output<string> StrategyType { get; private set; } = null!;


        /// <summary>
        /// Create a IpStrategy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpStrategy(string name, IpStrategyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/ipStrategy:IpStrategy", name, args ?? new IpStrategyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpStrategy(string name, Input<string> id, IpStrategyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/ipStrategy:IpStrategy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpStrategy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpStrategy Get(string name, Input<string> id, IpStrategyState? state = null, CustomResourceOptions? options = null)
        {
            return new IpStrategy(name, id, state, options);
        }
    }

    public sealed class IpStrategyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the API gateway service.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        /// <summary>
        /// IP address data.
        /// </summary>
        [Input("strategyData", required: true)]
        public Input<string> StrategyData { get; set; } = null!;

        /// <summary>
        /// User defined strategy name.
        /// </summary>
        [Input("strategyName", required: true)]
        public Input<string> StrategyName { get; set; } = null!;

        /// <summary>
        /// Blacklist or whitelist.
        /// </summary>
        [Input("strategyType", required: true)]
        public Input<string> StrategyType { get; set; } = null!;

        public IpStrategyArgs()
        {
        }
    }

    public sealed class IpStrategyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The ID of the API gateway service.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// IP address data.
        /// </summary>
        [Input("strategyData")]
        public Input<string>? StrategyData { get; set; }

        /// <summary>
        /// IP policy ID.
        /// </summary>
        [Input("strategyId")]
        public Input<string>? StrategyId { get; set; }

        /// <summary>
        /// User defined strategy name.
        /// </summary>
        [Input("strategyName")]
        public Input<string>? StrategyName { get; set; }

        /// <summary>
        /// Blacklist or whitelist.
        /// </summary>
        [Input("strategyType")]
        public Input<string>? StrategyType { get; set; }

        public IpStrategyState()
        {
        }
    }
}
