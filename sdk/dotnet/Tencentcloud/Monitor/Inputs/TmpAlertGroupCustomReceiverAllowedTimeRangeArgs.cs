// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Inputs
{

    public sealed class TmpAlertGroupCustomReceiverAllowedTimeRangeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time range end, seconds since 0 o'clock.
        /// </summary>
        [Input("end")]
        public Input<string>? End { get; set; }

        /// <summary>
        /// Time range start, seconds since 0 o'clock.
        /// </summary>
        [Input("start")]
        public Input<string>? Start { get; set; }

        public TmpAlertGroupCustomReceiverAllowedTimeRangeArgs()
        {
        }
    }
}