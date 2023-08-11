// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo
{
    [TencentcloudResourceType("tencentcloud:Teo/ddosPolicy:DdosPolicy")]
    public partial class DdosPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// DDoS Configuration of the zone.
        /// </summary>
        [Output("ddosRule")]
        public Output<Outputs.DdosPolicyDdosRule> DdosRule { get; private set; } = null!;

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Output("policyId")]
        public Output<int> PolicyId { get; private set; } = null!;

        /// <summary>
        /// Site ID.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a DdosPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DdosPolicy(string name, DdosPolicyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/ddosPolicy:DdosPolicy", name, args ?? new DdosPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DdosPolicy(string name, Input<string> id, DdosPolicyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Teo/ddosPolicy:DdosPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DdosPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DdosPolicy Get(string name, Input<string> id, DdosPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new DdosPolicy(name, id, state, options);
        }
    }

    public sealed class DdosPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// DDoS Configuration of the zone.
        /// </summary>
        [Input("ddosRule")]
        public Input<Inputs.DdosPolicyDdosRuleArgs>? DdosRule { get; set; }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<int> PolicyId { get; set; } = null!;

        /// <summary>
        /// Site ID.
        /// </summary>
        [Input("zoneId", required: true)]
        public Input<string> ZoneId { get; set; } = null!;

        public DdosPolicyArgs()
        {
        }
    }

    public sealed class DdosPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// DDoS Configuration of the zone.
        /// </summary>
        [Input("ddosRule")]
        public Input<Inputs.DdosPolicyDdosRuleGetArgs>? DdosRule { get; set; }

        /// <summary>
        /// Policy ID.
        /// </summary>
        [Input("policyId")]
        public Input<int>? PolicyId { get; set; }

        /// <summary>
        /// Site ID.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public DdosPolicyState()
        {
        }
    }
}
