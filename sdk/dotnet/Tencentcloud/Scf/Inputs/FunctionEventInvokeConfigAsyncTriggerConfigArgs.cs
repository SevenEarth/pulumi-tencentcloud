// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf.Inputs
{

    public sealed class FunctionEventInvokeConfigAsyncTriggerConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Message retention period.
        /// </summary>
        [Input("msgTtl", required: true)]
        public Input<int> MsgTtl { get; set; } = null!;

        [Input("retryConfigs", required: true)]
        private InputList<Inputs.FunctionEventInvokeConfigAsyncTriggerConfigRetryConfigArgs>? _retryConfigs;

        /// <summary>
        /// Async retry configuration of function upon user error.
        /// </summary>
        public InputList<Inputs.FunctionEventInvokeConfigAsyncTriggerConfigRetryConfigArgs> RetryConfigs
        {
            get => _retryConfigs ?? (_retryConfigs = new InputList<Inputs.FunctionEventInvokeConfigAsyncTriggerConfigRetryConfigArgs>());
            set => _retryConfigs = value;
        }

        public FunctionEventInvokeConfigAsyncTriggerConfigArgs()
        {
        }
    }
}
