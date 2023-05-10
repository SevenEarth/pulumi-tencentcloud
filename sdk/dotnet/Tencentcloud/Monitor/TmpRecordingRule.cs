// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor
{
    /// <summary>
    /// Provides a resource to create a monitor tmp recordingRule
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
    ///         var recordingRule = new Tencentcloud.Monitor.TmpRecordingRule("recordingRule", new Tencentcloud.Monitor.TmpRecordingRuleArgs
    ///         {
    ///             Group = @"---
    /// name: example-test
    /// rules:
    ///   - record: job:http_inprogress_requests:sum
    ///     expr: sum by (job) (http_inprogress_requests)
    /// 
    /// ",
    ///             InstanceId = "prom-c89b3b3u",
    ///             RuleState = 2,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// monitor recordingRule can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Monitor/tmpRecordingRule:TmpRecordingRule recordingRule instanceId#recordingRule_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/tmpRecordingRule:TmpRecordingRule")]
    public partial class TmpRecordingRule : Pulumi.CustomResource
    {
        /// <summary>
        /// Recording rule group.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Recording rule name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Rule state.
        /// </summary>
        [Output("ruleState")]
        public Output<int?> RuleState { get; private set; } = null!;


        /// <summary>
        /// Create a TmpRecordingRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TmpRecordingRule(string name, TmpRecordingRuleArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpRecordingRule:TmpRecordingRule", name, args ?? new TmpRecordingRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TmpRecordingRule(string name, Input<string> id, TmpRecordingRuleState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpRecordingRule:TmpRecordingRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TmpRecordingRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TmpRecordingRule Get(string name, Input<string> id, TmpRecordingRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new TmpRecordingRule(name, id, state, options);
        }
    }

    public sealed class TmpRecordingRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Recording rule group.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Recording rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Rule state.
        /// </summary>
        [Input("ruleState")]
        public Input<int>? RuleState { get; set; }

        public TmpRecordingRuleArgs()
        {
        }
    }

    public sealed class TmpRecordingRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Recording rule group.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Recording rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Rule state.
        /// </summary>
        [Input("ruleState")]
        public Input<int>? RuleState { get; set; }

        public TmpRecordingRuleState()
        {
        }
    }
}
