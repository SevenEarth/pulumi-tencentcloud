// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Chdfs.Inputs
{

    public sealed class LifeCycleRuleLifeCycleRuleTransitionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// trigger days(n day).
        /// </summary>
        [Input("days", required: true)]
        public Input<int> Days { get; set; } = null!;

        /// <summary>
        /// transition type, 1: archive, 2: delete, 3: low rate.
        /// </summary>
        [Input("type", required: true)]
        public Input<int> Type { get; set; } = null!;

        public LifeCycleRuleLifeCycleRuleTransitionGetArgs()
        {
        }
    }
}
