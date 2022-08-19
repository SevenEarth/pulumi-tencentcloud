// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cos.Outputs
{

    [OutputType]
    public sealed class BucketLifecycleRuleNonCurrentTransition
    {
        /// <summary>
        /// Number of days after non current object creation when the specific rule action takes effect.
        /// </summary>
        public readonly int? NonCurrentDays;
        /// <summary>
        /// Specifies the storage class to which you want the non current object to transition. Available values include `STANDARD`, `STANDARD_IA` and `ARCHIVE`.
        /// </summary>
        public readonly string StorageClass;

        [OutputConstructor]
        private BucketLifecycleRuleNonCurrentTransition(
            int? nonCurrentDays,

            string storageClass)
        {
            NonCurrentDays = nonCurrentDays;
            StorageClass = storageClass;
        }
    }
}
