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

    public sealed class TmpTkeAlertPolicyAlertRuleRuleAnnotationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of map.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Value of map.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public TmpTkeAlertPolicyAlertRuleRuleAnnotationGetArgs()
        {
        }
        public static new TmpTkeAlertPolicyAlertRuleRuleAnnotationGetArgs Empty => new TmpTkeAlertPolicyAlertRuleRuleAnnotationGetArgs();
    }
}
