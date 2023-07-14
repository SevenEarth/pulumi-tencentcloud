// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf.Outputs
{

    [OutputType]
    public sealed class GetAccountInfoAccountLimitNamespaceResult
    {
        /// <summary>
        /// Concurrency.
        /// </summary>
        public readonly int ConcurrentExecutions;
        /// <summary>
        /// Number of functions in namespace.
        /// </summary>
        public readonly int FunctionsCount;
        /// <summary>
        /// Initialization timeout limit.
        /// </summary>
        public readonly int InitTimeoutLimit;
        /// <summary>
        /// Upper limit of message retention time for async retry.
        /// </summary>
        public readonly int MaxMsgTtl;
        /// <summary>
        /// Lower limit of message retention time for async retry.
        /// </summary>
        public readonly int MinMsgTtl;
        /// <summary>
        /// Namespace name.
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// Limit of async retry attempt quantity.
        /// </summary>
        public readonly int RetryNumLimit;
        /// <summary>
        /// Test event limit Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly int TestModelLimit;
        /// <summary>
        /// Timeout limit.
        /// </summary>
        public readonly int TimeoutLimit;
        /// <summary>
        /// Trigger information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccountInfoAccountLimitNamespaceTriggerResult> Triggers;

        [OutputConstructor]
        private GetAccountInfoAccountLimitNamespaceResult(
            int concurrentExecutions,

            int functionsCount,

            int initTimeoutLimit,

            int maxMsgTtl,

            int minMsgTtl,

            string @namespace,

            int retryNumLimit,

            int testModelLimit,

            int timeoutLimit,

            ImmutableArray<Outputs.GetAccountInfoAccountLimitNamespaceTriggerResult> triggers)
        {
            ConcurrentExecutions = concurrentExecutions;
            FunctionsCount = functionsCount;
            InitTimeoutLimit = initTimeoutLimit;
            MaxMsgTtl = maxMsgTtl;
            MinMsgTtl = minMsgTtl;
            Namespace = @namespace;
            RetryNumLimit = retryNumLimit;
            TestModelLimit = testModelLimit;
            TimeoutLimit = timeoutLimit;
            Triggers = triggers;
        }
    }
}
