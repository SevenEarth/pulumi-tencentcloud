// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cdn
{
    /// <summary>
    /// Provide a resource to invoke a Url Purge Request.
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
    ///         var foo = new Tencentcloud.Cdn.UrlPurge("foo", new Tencentcloud.Cdn.UrlPurgeArgs
    ///         {
    ///             Urls = 
    ///             {
    ///                 "https://www.example.com/a",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// argument to request new purge task with same urls
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Cdn.UrlPurge("foo", new Tencentcloud.Cdn.UrlPurgeArgs
    ///         {
    ///             Redo = 1,
    ///             Urls = 
    ///             {
    ///                 "https://www.example.com/a",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cdn/urlPurge:UrlPurge")]
    public partial class UrlPurge : Pulumi.CustomResource
    {
        /// <summary>
        /// Specify purge area. NOTE: only purge same area cache contents.
        /// </summary>
        [Output("area")]
        public Output<string?> Area { get; private set; } = null!;

        /// <summary>
        /// logs of latest purge task.
        /// </summary>
        [Output("purgeHistories")]
        public Output<ImmutableArray<Outputs.UrlPurgePurgeHistory>> PurgeHistories { get; private set; } = null!;

        /// <summary>
        /// Change to purge again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
        /// </summary>
        [Output("redo")]
        public Output<int?> Redo { get; private set; } = null!;

        /// <summary>
        /// Task id of last operation.
        /// </summary>
        [Output("taskId")]
        public Output<string> TaskId { get; private set; } = null!;

        /// <summary>
        /// Whether to encode urls, if set to `true` will auto encode instead of manual process.
        /// </summary>
        [Output("urlEncode")]
        public Output<bool?> UrlEncode { get; private set; } = null!;

        /// <summary>
        /// List of url to purge. NOTE: urls need include protocol prefix `http://` or `https://`.
        /// </summary>
        [Output("urls")]
        public Output<ImmutableArray<string>> Urls { get; private set; } = null!;


        /// <summary>
        /// Create a UrlPurge resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UrlPurge(string name, UrlPurgeArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cdn/urlPurge:UrlPurge", name, args ?? new UrlPurgeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UrlPurge(string name, Input<string> id, UrlPurgeState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cdn/urlPurge:UrlPurge", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UrlPurge resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UrlPurge Get(string name, Input<string> id, UrlPurgeState? state = null, CustomResourceOptions? options = null)
        {
            return new UrlPurge(name, id, state, options);
        }
    }

    public sealed class UrlPurgeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify purge area. NOTE: only purge same area cache contents.
        /// </summary>
        [Input("area")]
        public Input<string>? Area { get; set; }

        /// <summary>
        /// Change to purge again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
        /// </summary>
        [Input("redo")]
        public Input<int>? Redo { get; set; }

        /// <summary>
        /// Whether to encode urls, if set to `true` will auto encode instead of manual process.
        /// </summary>
        [Input("urlEncode")]
        public Input<bool>? UrlEncode { get; set; }

        [Input("urls", required: true)]
        private InputList<string>? _urls;

        /// <summary>
        /// List of url to purge. NOTE: urls need include protocol prefix `http://` or `https://`.
        /// </summary>
        public InputList<string> Urls
        {
            get => _urls ?? (_urls = new InputList<string>());
            set => _urls = value;
        }

        public UrlPurgeArgs()
        {
        }
    }

    public sealed class UrlPurgeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify purge area. NOTE: only purge same area cache contents.
        /// </summary>
        [Input("area")]
        public Input<string>? Area { get; set; }

        [Input("purgeHistories")]
        private InputList<Inputs.UrlPurgePurgeHistoryGetArgs>? _purgeHistories;

        /// <summary>
        /// logs of latest purge task.
        /// </summary>
        public InputList<Inputs.UrlPurgePurgeHistoryGetArgs> PurgeHistories
        {
            get => _purgeHistories ?? (_purgeHistories = new InputList<Inputs.UrlPurgePurgeHistoryGetArgs>());
            set => _purgeHistories = value;
        }

        /// <summary>
        /// Change to purge again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
        /// </summary>
        [Input("redo")]
        public Input<int>? Redo { get; set; }

        /// <summary>
        /// Task id of last operation.
        /// </summary>
        [Input("taskId")]
        public Input<string>? TaskId { get; set; }

        /// <summary>
        /// Whether to encode urls, if set to `true` will auto encode instead of manual process.
        /// </summary>
        [Input("urlEncode")]
        public Input<bool>? UrlEncode { get; set; }

        [Input("urls")]
        private InputList<string>? _urls;

        /// <summary>
        /// List of url to purge. NOTE: urls need include protocol prefix `http://` or `https://`.
        /// </summary>
        public InputList<string> Urls
        {
            get => _urls ?? (_urls = new InputList<string>());
            set => _urls = value;
        }

        public UrlPurgeState()
        {
        }
    }
}
