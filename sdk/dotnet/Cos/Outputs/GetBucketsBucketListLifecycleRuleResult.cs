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
    public sealed class GetBucketsBucketListLifecycleRuleResult
    {
        /// <summary>
        /// Specifies a period in the object's expire.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBucketsBucketListLifecycleRuleExpirationResult> Expirations;
        /// <summary>
        /// Object key prefix identifying one or more objects to which the rule applies.
        /// </summary>
        public readonly string FilterPrefix;
        /// <summary>
        /// Specifies when non current object versions shall expire.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBucketsBucketListLifecycleRuleNonCurrentExpirationResult> NonCurrentExpirations;
        /// <summary>
        /// Specifies when to transition objects of non current versions and the target storage class.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBucketsBucketListLifecycleRuleNonCurrentTransitionResult> NonCurrentTransitions;
        /// <summary>
        /// Specifies a period in the object's transitions.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetBucketsBucketListLifecycleRuleTransitionResult> Transitions;

        [OutputConstructor]
        private GetBucketsBucketListLifecycleRuleResult(
            ImmutableArray<Outputs.GetBucketsBucketListLifecycleRuleExpirationResult> expirations,

            string filterPrefix,

            ImmutableArray<Outputs.GetBucketsBucketListLifecycleRuleNonCurrentExpirationResult> nonCurrentExpirations,

            ImmutableArray<Outputs.GetBucketsBucketListLifecycleRuleNonCurrentTransitionResult> nonCurrentTransitions,

            ImmutableArray<Outputs.GetBucketsBucketListLifecycleRuleTransitionResult> transitions)
        {
            Expirations = expirations;
            FilterPrefix = filterPrefix;
            NonCurrentExpirations = nonCurrentExpirations;
            NonCurrentTransitions = nonCurrentTransitions;
            Transitions = transitions;
        }
    }
}
