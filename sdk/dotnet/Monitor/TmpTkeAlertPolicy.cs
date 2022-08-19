// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor
{
    /// <summary>
    /// Provides a resource to create a tke tmpAlertPolicy
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/tmpTkeAlertPolicy:TmpTkeAlertPolicy")]
    public partial class TmpTkeAlertPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Alarm notification channels.
        /// </summary>
        [Output("alertRule")]
        public Output<Outputs.TmpTkeAlertPolicyAlertRule> AlertRule { get; private set; } = null!;

        /// <summary>
        /// Instance Id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a TmpTkeAlertPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TmpTkeAlertPolicy(string name, TmpTkeAlertPolicyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeAlertPolicy:TmpTkeAlertPolicy", name, args ?? new TmpTkeAlertPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TmpTkeAlertPolicy(string name, Input<string> id, TmpTkeAlertPolicyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeAlertPolicy:TmpTkeAlertPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TmpTkeAlertPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TmpTkeAlertPolicy Get(string name, Input<string> id, TmpTkeAlertPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new TmpTkeAlertPolicy(name, id, state, options);
        }
    }

    public sealed class TmpTkeAlertPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alarm notification channels.
        /// </summary>
        [Input("alertRule", required: true)]
        public Input<Inputs.TmpTkeAlertPolicyAlertRuleArgs> AlertRule { get; set; } = null!;

        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public TmpTkeAlertPolicyArgs()
        {
        }
    }

    public sealed class TmpTkeAlertPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alarm notification channels.
        /// </summary>
        [Input("alertRule")]
        public Input<Inputs.TmpTkeAlertPolicyAlertRuleGetArgs>? AlertRule { get; set; }

        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public TmpTkeAlertPolicyState()
        {
        }
    }
}
