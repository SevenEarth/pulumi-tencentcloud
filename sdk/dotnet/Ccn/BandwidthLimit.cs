// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ccn
{
    /// <summary>
    /// Provides a resource to limit CCN bandwidth.
    /// 
    /// ## Example Usage
    /// 
    /// Set the upper limit of regional outbound bandwidth
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var otherRegion1 = config.Get("otherRegion1") ?? "ap-shanghai";
    ///         var main = new Tencentcloud.Ccn.Instance("main", new Tencentcloud.Ccn.InstanceArgs
    ///         {
    ///             Description = "ci-temp-test-ccn-des",
    ///             Qos = "AG",
    ///         });
    ///         var limit1 = new Tencentcloud.Ccn.BandwidthLimit("limit1", new Tencentcloud.Ccn.BandwidthLimitArgs
    ///         {
    ///             CcnId = main.Id,
    ///             Region = otherRegion1,
    ///             BandwidthLimit = 500,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Set the upper limit between regions
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var config = new Config();
    ///         var otherRegion1 = config.Get("otherRegion1") ?? "ap-shanghai";
    ///         var otherRegion2 = config.Get("otherRegion2") ?? "ap-nanjing";
    ///         var main = new Tencentcloud.Ccn.Instance("main", new Tencentcloud.Ccn.InstanceArgs
    ///         {
    ///             Description = "ci-temp-test-ccn-des",
    ///             Qos = "AG",
    ///             BandwidthLimitType = "INTER_REGION_LIMIT",
    ///         });
    ///         var limit1 = new Tencentcloud.Ccn.BandwidthLimit("limit1", new Tencentcloud.Ccn.BandwidthLimitArgs
    ///         {
    ///             CcnId = main.Id,
    ///             Region = otherRegion1,
    ///             DstRegion = otherRegion2,
    ///             BandwidthLimit = 100,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ccn/bandwidthLimit:BandwidthLimit")]
    public partial class BandwidthLimit : Pulumi.CustomResource
    {
        /// <summary>
        /// Limitation of bandwidth.
        /// </summary>
        [Output("bandwidthLimit")]
        public Output<int> CcnBandwidthLimit { get; private set; } = null!;

        /// <summary>
        /// ID of the CCN.
        /// </summary>
        [Output("ccnId")]
        public Output<string> CcnId { get; private set; } = null!;

        /// <summary>
        /// Destination area restriction. If the `CCN` rate limit type is `OUTER_REGION_LIMIT`, this value does not need to be set.
        /// </summary>
        [Output("dstRegion")]
        public Output<string?> DstRegion { get; private set; } = null!;

        /// <summary>
        /// Limitation of region.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a BandwidthLimit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BandwidthLimit(string name, BandwidthLimitArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ccn/bandwidthLimit:BandwidthLimit", name, args ?? new BandwidthLimitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BandwidthLimit(string name, Input<string> id, BandwidthLimitState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ccn/bandwidthLimit:BandwidthLimit", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BandwidthLimit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BandwidthLimit Get(string name, Input<string> id, BandwidthLimitState? state = null, CustomResourceOptions? options = null)
        {
            return new BandwidthLimit(name, id, state, options);
        }
    }

    public sealed class BandwidthLimitArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Limitation of bandwidth.
        /// </summary>
        [Input("bandwidthLimit")]
        public Input<int>? CcnBandwidthLimit { get; set; }

        /// <summary>
        /// ID of the CCN.
        /// </summary>
        [Input("ccnId", required: true)]
        public Input<string> CcnId { get; set; } = null!;

        /// <summary>
        /// Destination area restriction. If the `CCN` rate limit type is `OUTER_REGION_LIMIT`, this value does not need to be set.
        /// </summary>
        [Input("dstRegion")]
        public Input<string>? DstRegion { get; set; }

        /// <summary>
        /// Limitation of region.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public BandwidthLimitArgs()
        {
        }
    }

    public sealed class BandwidthLimitState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Limitation of bandwidth.
        /// </summary>
        [Input("bandwidthLimit")]
        public Input<int>? CcnBandwidthLimit { get; set; }

        /// <summary>
        /// ID of the CCN.
        /// </summary>
        [Input("ccnId")]
        public Input<string>? CcnId { get; set; }

        /// <summary>
        /// Destination area restriction. If the `CCN` rate limit type is `OUTER_REGION_LIMIT`, this value does not need to be set.
        /// </summary>
        [Input("dstRegion")]
        public Input<string>? DstRegion { get; set; }

        /// <summary>
        /// Limitation of region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public BandwidthLimitState()
        {
        }
    }
}
