// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cos.Inputs
{

    public sealed class BucketLifecycleRuleTransitionGetArgs : Pulumi.ResourceArgs
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
        /// Specifies the storage class to which you want the object to transition. Available values include `STANDARD`, `STANDARD_IA` and `ARCHIVE`.
        /// </summary>
        [Input("storageClass", required: true)]
        public Input<string> StorageClass { get; set; } = null!;

        public BucketLifecycleRuleTransitionGetArgs()
        {
        }
    }
}
