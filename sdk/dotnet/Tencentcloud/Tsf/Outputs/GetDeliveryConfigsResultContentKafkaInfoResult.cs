// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Outputs
{

    [OutputType]
    public sealed class GetDeliveryConfigsResultContentKafkaInfoResult
    {
        /// <summary>
        /// Custom Line Rule.
        /// </summary>
        public readonly string CustomRule;
        /// <summary>
        /// Line Rule for log. Note: This field may return null, which means that no valid value was obtained.
        /// </summary>
        public readonly string LineRule;
        /// <summary>
        /// harvest log path. Note: This field may return null, which means that no valid value was obtained.
        /// </summary>
        public readonly ImmutableArray<string> Paths;
        /// <summary>
        /// Topic. Note: This field may return null, which means that no valid value was obtained.
        /// </summary>
        public readonly string Topic;

        [OutputConstructor]
        private GetDeliveryConfigsResultContentKafkaInfoResult(
            string customRule,

            string lineRule,

            ImmutableArray<string> paths,

            string topic)
        {
            CustomRule = customRule;
            LineRule = lineRule;
            Paths = paths;
            Topic = topic;
        }
    }
}
