// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cos.Inputs
{

    public sealed class BucketLifecycleRuleNonCurrentExpirationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of days after non current object creation when the specific rule action takes effect. The maximum value is 3650.
        /// </summary>
        [Input("nonCurrentDays")]
        public Input<int>? NonCurrentDays { get; set; }

        public BucketLifecycleRuleNonCurrentExpirationArgs()
        {
        }
    }
}
