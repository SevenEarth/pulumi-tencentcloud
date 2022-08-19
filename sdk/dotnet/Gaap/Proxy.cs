// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Gaap
{
    /// <summary>
    /// Provides a resource to create a GAAP proxy.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Gaap.Proxy("foo", new Tencentcloud.Gaap.ProxyArgs
    ///         {
    ///             AccessRegion = "SouthChina",
    ///             Bandwidth = 10,
    ///             Concurrent = 2,
    ///             RealserverRegion = "NorthChina",
    ///             Tags = 
    ///             {
    ///                 { "test", "test" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// GAAP proxy can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Gaap/proxy:Proxy tencentcloud_gaap_proxy.foo link-11112222
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Gaap/proxy:Proxy")]
    public partial class Proxy : Pulumi.CustomResource
    {
        /// <summary>
        /// Access region of the GAAP proxy. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`, `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`, `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
        /// </summary>
        [Output("accessRegion")]
        public Output<string> AccessRegion { get; private set; } = null!;

        /// <summary>
        /// Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500` and `1000`.
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `90` and `100`.
        /// </summary>
        [Output("concurrent")]
        public Output<int> Concurrent { get; private set; } = null!;

        /// <summary>
        /// Creation time of the GAAP proxy.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Access domain of the GAAP proxy.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Indicates whether GAAP proxy is enabled, default value is `true`.
        /// </summary>
        [Output("enable")]
        public Output<bool?> Enable { get; private set; } = null!;

        /// <summary>
        /// Forwarding IP of the GAAP proxy.
        /// </summary>
        [Output("forwardIp")]
        public Output<string> ForwardIp { get; private set; } = null!;

        /// <summary>
        /// Access IP of the GAAP proxy.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// Name of the GAAP proxy, the maximum length is 30.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the project within the GAAP proxy, `0` means is default project.
        /// </summary>
        [Output("projectId")]
        public Output<int?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Region of the GAAP realserver. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`, `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`, `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
        /// </summary>
        [Output("realserverRegion")]
        public Output<string> RealserverRegion { get; private set; } = null!;

        /// <summary>
        /// Indicates whether GAAP proxy can scalable.
        /// </summary>
        [Output("scalable")]
        public Output<bool> Scalable { get; private set; } = null!;

        /// <summary>
        /// Status of the GAAP proxy.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Supported protocols of the GAAP proxy.
        /// </summary>
        [Output("supportProtocols")]
        public Output<ImmutableArray<string>> SupportProtocols { get; private set; } = null!;

        /// <summary>
        /// Tags of the GAAP proxy.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Proxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Proxy(string name, ProxyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Gaap/proxy:Proxy", name, args ?? new ProxyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Proxy(string name, Input<string> id, ProxyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Gaap/proxy:Proxy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Proxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Proxy Get(string name, Input<string> id, ProxyState? state = null, CustomResourceOptions? options = null)
        {
            return new Proxy(name, id, state, options);
        }
    }

    public sealed class ProxyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access region of the GAAP proxy. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`, `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`, `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
        /// </summary>
        [Input("accessRegion", required: true)]
        public Input<string> AccessRegion { get; set; } = null!;

        /// <summary>
        /// Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500` and `1000`.
        /// </summary>
        [Input("bandwidth", required: true)]
        public Input<int> Bandwidth { get; set; } = null!;

        /// <summary>
        /// Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `90` and `100`.
        /// </summary>
        [Input("concurrent", required: true)]
        public Input<int> Concurrent { get; set; } = null!;

        /// <summary>
        /// Indicates whether GAAP proxy is enabled, default value is `true`.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// Name of the GAAP proxy, the maximum length is 30.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the project within the GAAP proxy, `0` means is default project.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Region of the GAAP realserver. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`, `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`, `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
        /// </summary>
        [Input("realserverRegion", required: true)]
        public Input<string> RealserverRegion { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of the GAAP proxy.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ProxyArgs()
        {
        }
    }

    public sealed class ProxyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access region of the GAAP proxy. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`, `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`, `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
        /// </summary>
        [Input("accessRegion")]
        public Input<string>? AccessRegion { get; set; }

        /// <summary>
        /// Maximum bandwidth of the GAAP proxy, unit is Mbps. Valid value: `10`, `20`, `50`, `100`, `200`, `500` and `1000`.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// Maximum concurrency of the GAAP proxy, unit is 10k. Valid value: `2`, `5`, `10`, `20`, `30`, `40`, `50`, `60`, `70`, `80`, `90` and `100`.
        /// </summary>
        [Input("concurrent")]
        public Input<int>? Concurrent { get; set; }

        /// <summary>
        /// Creation time of the GAAP proxy.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Access domain of the GAAP proxy.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Indicates whether GAAP proxy is enabled, default value is `true`.
        /// </summary>
        [Input("enable")]
        public Input<bool>? Enable { get; set; }

        /// <summary>
        /// Forwarding IP of the GAAP proxy.
        /// </summary>
        [Input("forwardIp")]
        public Input<string>? ForwardIp { get; set; }

        /// <summary>
        /// Access IP of the GAAP proxy.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Name of the GAAP proxy, the maximum length is 30.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the project within the GAAP proxy, `0` means is default project.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Region of the GAAP realserver. Valid value: `NorthChina`, `EastChina`, `SouthChina`, `SouthwestChina`, `Hongkong`, `SL_TAIWAN`, `SoutheastAsia`, `Korea`, `SL_India`, `SL_Australia`, `Europe`, `SL_UK`, `SL_SouthAmerica`, `NorthAmerica`, `SL_MiddleUSA`, `Canada`, `SL_VIET`, `WestIndia`, `Thailand`, `Virginia`, `Russia`, `Japan` and `SL_Indonesia`.
        /// </summary>
        [Input("realserverRegion")]
        public Input<string>? RealserverRegion { get; set; }

        /// <summary>
        /// Indicates whether GAAP proxy can scalable.
        /// </summary>
        [Input("scalable")]
        public Input<bool>? Scalable { get; set; }

        /// <summary>
        /// Status of the GAAP proxy.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("supportProtocols")]
        private InputList<string>? _supportProtocols;

        /// <summary>
        /// Supported protocols of the GAAP proxy.
        /// </summary>
        public InputList<string> SupportProtocols
        {
            get => _supportProtocols ?? (_supportProtocols = new InputList<string>());
            set => _supportProtocols = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of the GAAP proxy.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ProxyState()
        {
        }
    }
}
