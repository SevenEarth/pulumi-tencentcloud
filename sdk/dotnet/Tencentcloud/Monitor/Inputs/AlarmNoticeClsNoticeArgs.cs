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

    public sealed class AlarmNoticeClsNoticeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Start-stop status, can not be transmitted, default enabled. 0= Disabled, 1= enabled.
        /// </summary>
        [Input("enable")]
        public Input<int>? Enable { get; set; }

        /// <summary>
        /// Log collection Id.
        /// </summary>
        [Input("logSetId", required: true)]
        public Input<string> LogSetId { get; set; } = null!;

        /// <summary>
        /// Regional.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// Theme Id.
        /// </summary>
        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        public AlarmNoticeClsNoticeArgs()
        {
        }
    }
}