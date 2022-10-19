// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Inputs
{

    public sealed class AlarmPolicyConditionsRuleFilterGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// JSON string generated by serializing the AlarmPolicyDimension two-dimensional array.
        /// </summary>
        [Input("dimensions")]
        public Input<string>? Dimensions { get; set; }

        /// <summary>
        /// Filter condition type. Valid values: DIMENSION (uses dimensions for filtering).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AlarmPolicyConditionsRuleFilterGetArgs()
        {
        }
    }
}