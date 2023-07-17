// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos
{
    /// <summary>
    /// Provides a resource to create a cos bucket_inventory
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
    ///         var bucketInventory = new Tencentcloud.Cos.BucketInventory("bucketInventory", new Tencentcloud.Cos.BucketInventoryArgs
    ///         {
    ///             Bucket = "keep-test-xxxxxx",
    ///             Destination = new Tencentcloud.Cos.Inputs.BucketInventoryDestinationArgs
    ///             {
    ///                 AccountId = "",
    ///                 Bucket = "qcs::cos:ap-guangzhou::keep-test-xxxxxx",
    ///                 Format = "CSV",
    ///                 Prefix = "cos_bucket_inventory",
    ///             },
    ///             Filter = new Tencentcloud.Cos.Inputs.BucketInventoryFilterArgs
    ///             {
    ///                 Period = new Tencentcloud.Cos.Inputs.BucketInventoryFilterPeriodArgs
    ///                 {
    ///                     StartTime = "1687276800",
    ///                 },
    ///             },
    ///             IncludedObjectVersions = "Current",
    ///             IsEnabled = "true",
    ///             OptionalFields = new Tencentcloud.Cos.Inputs.BucketInventoryOptionalFieldsArgs
    ///             {
    ///                 Fields = 
    ///                 {
    ///                     "Size",
    ///                     "ETag",
    ///                 },
    ///             },
    ///             Schedule = new Tencentcloud.Cos.Inputs.BucketInventoryScheduleArgs
    ///             {
    ///                 Frequency = "Weekly",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// cos bucket_inventory can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cos/bucketInventory:BucketInventory bucket_inventory bucket_inventory_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cos/bucketInventory:BucketInventory")]
    public partial class BucketInventory : Pulumi.CustomResource
    {
        /// <summary>
        /// Bucket name.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Information about the inventory result destination.
        /// </summary>
        [Output("destination")]
        public Output<Outputs.BucketInventoryDestination> Destination { get; private set; } = null!;

        /// <summary>
        /// Filters objects prefixed with the specified value to analyze.
        /// </summary>
        [Output("filter")]
        public Output<Outputs.BucketInventoryFilter?> Filter { get; private set; } = null!;

        /// <summary>
        /// Whether to include object versions in the inventory. All or No.
        /// </summary>
        [Output("includedObjectVersions")]
        public Output<string> IncludedObjectVersions { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the inventory. true or false.
        /// </summary>
        [Output("isEnabled")]
        public Output<string> IsEnabled { get; private set; } = null!;

        /// <summary>
        /// Inventory Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Analysis items to include in the inventory result	.
        /// </summary>
        [Output("optionalFields")]
        public Output<Outputs.BucketInventoryOptionalFields?> OptionalFields { get; private set; } = null!;

        /// <summary>
        /// Inventory job cycle.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.BucketInventorySchedule> Schedule { get; private set; } = null!;


        /// <summary>
        /// Create a BucketInventory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketInventory(string name, BucketInventoryArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cos/bucketInventory:BucketInventory", name, args ?? new BucketInventoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketInventory(string name, Input<string> id, BucketInventoryState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cos/bucketInventory:BucketInventory", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketInventory resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketInventory Get(string name, Input<string> id, BucketInventoryState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketInventory(name, id, state, options);
        }
    }

    public sealed class BucketInventoryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Information about the inventory result destination.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.BucketInventoryDestinationArgs> Destination { get; set; } = null!;

        /// <summary>
        /// Filters objects prefixed with the specified value to analyze.
        /// </summary>
        [Input("filter")]
        public Input<Inputs.BucketInventoryFilterArgs>? Filter { get; set; }

        /// <summary>
        /// Whether to include object versions in the inventory. All or No.
        /// </summary>
        [Input("includedObjectVersions", required: true)]
        public Input<string> IncludedObjectVersions { get; set; } = null!;

        /// <summary>
        /// Whether to enable the inventory. true or false.
        /// </summary>
        [Input("isEnabled", required: true)]
        public Input<string> IsEnabled { get; set; } = null!;

        /// <summary>
        /// Inventory Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Analysis items to include in the inventory result	.
        /// </summary>
        [Input("optionalFields")]
        public Input<Inputs.BucketInventoryOptionalFieldsArgs>? OptionalFields { get; set; }

        /// <summary>
        /// Inventory job cycle.
        /// </summary>
        [Input("schedule", required: true)]
        public Input<Inputs.BucketInventoryScheduleArgs> Schedule { get; set; } = null!;

        public BucketInventoryArgs()
        {
        }
    }

    public sealed class BucketInventoryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bucket name.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Information about the inventory result destination.
        /// </summary>
        [Input("destination")]
        public Input<Inputs.BucketInventoryDestinationGetArgs>? Destination { get; set; }

        /// <summary>
        /// Filters objects prefixed with the specified value to analyze.
        /// </summary>
        [Input("filter")]
        public Input<Inputs.BucketInventoryFilterGetArgs>? Filter { get; set; }

        /// <summary>
        /// Whether to include object versions in the inventory. All or No.
        /// </summary>
        [Input("includedObjectVersions")]
        public Input<string>? IncludedObjectVersions { get; set; }

        /// <summary>
        /// Whether to enable the inventory. true or false.
        /// </summary>
        [Input("isEnabled")]
        public Input<string>? IsEnabled { get; set; }

        /// <summary>
        /// Inventory Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Analysis items to include in the inventory result	.
        /// </summary>
        [Input("optionalFields")]
        public Input<Inputs.BucketInventoryOptionalFieldsGetArgs>? OptionalFields { get; set; }

        /// <summary>
        /// Inventory job cycle.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.BucketInventoryScheduleGetArgs>? Schedule { get; set; }

        public BucketInventoryState()
        {
        }
    }
}
