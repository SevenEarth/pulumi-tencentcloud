// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eb.Inputs
{

    public sealed class EventTargetTargetDescriptionCkafkaTargetParamsRetryPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of retries.
        /// </summary>
        [Input("maxRetryAttempts", required: true)]
        public Input<int> MaxRetryAttempts { get; set; } = null!;

        /// <summary>
        /// Retry Interval Unit: Seconds.
        /// </summary>
        [Input("retryInterval", required: true)]
        public Input<int> RetryInterval { get; set; } = null!;

        public EventTargetTargetDescriptionCkafkaTargetParamsRetryPolicyArgs()
        {
        }
    }
}
