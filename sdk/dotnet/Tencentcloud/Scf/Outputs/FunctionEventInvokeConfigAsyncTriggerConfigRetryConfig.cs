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
    public sealed class FunctionEventInvokeConfigAsyncTriggerConfigRetryConfig
    {
        /// <summary>
        /// Number of retry attempts.
        /// </summary>
        public readonly int RetryNum;

        [OutputConstructor]
        private FunctionEventInvokeConfigAsyncTriggerConfigRetryConfig(int retryNum)
        {
            RetryNum = retryNum;
        }
    }
}