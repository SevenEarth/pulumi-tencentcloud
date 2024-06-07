// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts.Inputs
{

    public sealed class ScenarioSlaPolicyAlertChannelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AMP consumer ID.
        /// </summary>
        [Input("ampConsumerId")]
        public Input<string>? AmpConsumerId { get; set; }

        /// <summary>
        /// Notification template ID.
        /// </summary>
        [Input("noticeId")]
        public Input<string>? NoticeId { get; set; }

        public ScenarioSlaPolicyAlertChannelArgs()
        {
        }
        public static new ScenarioSlaPolicyAlertChannelArgs Empty => new ScenarioSlaPolicyAlertChannelArgs();
    }
}
