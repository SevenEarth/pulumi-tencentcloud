// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class RealtimeLogDeliveryClsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tencent Cloud CLS log set ID.
        /// </summary>
        [Input("logSetId", required: true)]
        public Input<string> LogSetId { get; set; } = null!;

        /// <summary>
        /// The region where the Tencent Cloud CLS log set is located.
        /// </summary>
        [Input("logSetRegion", required: true)]
        public Input<string> LogSetRegion { get; set; } = null!;

        /// <summary>
        /// Tencent Cloud CLS log topic ID.
        /// </summary>
        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        public RealtimeLogDeliveryClsGetArgs()
        {
        }
        public static new RealtimeLogDeliveryClsGetArgs Empty => new RealtimeLogDeliveryClsGetArgs();
    }
}
