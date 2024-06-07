// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos.Inputs
{

    public sealed class BucketLifecycleRuleExpirationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the date after which you want the corresponding action to take effect.
        /// </summary>
        [Input("date")]
        public Input<string>? Date { get; set; }

        /// <summary>
        /// Specifies the number of days after object creation when the specific rule action takes effect.
        /// </summary>
        [Input("days")]
        public Input<int>? Days { get; set; }

        /// <summary>
        /// Indicates whether the delete marker of an expired object will be removed.
        /// </summary>
        [Input("deleteMarker")]
        public Input<bool>? DeleteMarker { get; set; }

        public BucketLifecycleRuleExpirationArgs()
        {
        }
        public static new BucketLifecycleRuleExpirationArgs Empty => new BucketLifecycleRuleExpirationArgs();
    }
}
