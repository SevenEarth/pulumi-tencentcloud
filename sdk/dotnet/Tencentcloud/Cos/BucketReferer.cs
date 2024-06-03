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
    /// Provides a resource to create a cos bucket_referer
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
    ///     var bucketReferer = new Tencentcloud.Cos.BucketReferer("bucketReferer", new()
    ///     {
    ///         Bucket = "mycos-1258798060",
    ///         DomainLists = new[]
    ///         {
    ///             "127.0.0.1",
    ///             "terraform.com",
    ///         },
    ///         EmptyReferConfiguration = "Allow",
    ///         RefererType = "Black-List",
    ///         Status = "Enabled",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// cos bucket_referer can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Cos/bucketReferer:BucketReferer bucket_referer bucket_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cos/bucketReferer:BucketReferer")]
    public partial class BucketReferer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// A list of domain names in the blocklist/allowlist.
        /// </summary>
        [Output("domainLists")]
        public Output<ImmutableArray<string>> DomainLists { get; private set; } = null!;

        /// <summary>
        /// Whether to allow access with an empty referer. Enumerated values: `Allow`, `Deny` (default).
        /// </summary>
        [Output("emptyReferConfiguration")]
        public Output<string?> EmptyReferConfiguration { get; private set; } = null!;

        /// <summary>
        /// Hotlink protection type. Enumerated values: `Black-List`, `White-List`.
        /// </summary>
        [Output("refererType")]
        public Output<string> RefererType { get; private set; } = null!;

        /// <summary>
        /// Whether to enable hotlink protection. Enumerated values: `Enabled`, `Disabled`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a BucketReferer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketReferer(string name, BucketRefererArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cos/bucketReferer:BucketReferer", name, args ?? new BucketRefererArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketReferer(string name, Input<string> id, BucketRefererState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cos/bucketReferer:BucketReferer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BucketReferer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketReferer Get(string name, Input<string> id, BucketRefererState? state = null, CustomResourceOptions? options = null)
        {
            return new BucketReferer(name, id, state, options);
        }
    }

    public sealed class BucketRefererArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("domainLists", required: true)]
        private InputList<string>? _domainLists;

        /// <summary>
        /// A list of domain names in the blocklist/allowlist.
        /// </summary>
        public InputList<string> DomainLists
        {
            get => _domainLists ?? (_domainLists = new InputList<string>());
            set => _domainLists = value;
        }

        /// <summary>
        /// Whether to allow access with an empty referer. Enumerated values: `Allow`, `Deny` (default).
        /// </summary>
        [Input("emptyReferConfiguration")]
        public Input<string>? EmptyReferConfiguration { get; set; }

        /// <summary>
        /// Hotlink protection type. Enumerated values: `Black-List`, `White-List`.
        /// </summary>
        [Input("refererType", required: true)]
        public Input<string> RefererType { get; set; } = null!;

        /// <summary>
        /// Whether to enable hotlink protection. Enumerated values: `Enabled`, `Disabled`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        public BucketRefererArgs()
        {
        }
        public static new BucketRefererArgs Empty => new BucketRefererArgs();
    }

    public sealed class BucketRefererState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("domainLists")]
        private InputList<string>? _domainLists;

        /// <summary>
        /// A list of domain names in the blocklist/allowlist.
        /// </summary>
        public InputList<string> DomainLists
        {
            get => _domainLists ?? (_domainLists = new InputList<string>());
            set => _domainLists = value;
        }

        /// <summary>
        /// Whether to allow access with an empty referer. Enumerated values: `Allow`, `Deny` (default).
        /// </summary>
        [Input("emptyReferConfiguration")]
        public Input<string>? EmptyReferConfiguration { get; set; }

        /// <summary>
        /// Hotlink protection type. Enumerated values: `Black-List`, `White-List`.
        /// </summary>
        [Input("refererType")]
        public Input<string>? RefererType { get; set; }

        /// <summary>
        /// Whether to enable hotlink protection. Enumerated values: `Enabled`, `Disabled`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public BucketRefererState()
        {
        }
        public static new BucketRefererState Empty => new BucketRefererState();
    }
}
