// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Inputs
{

    public sealed class ClusterOperationInfoAddInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reason for not showing.
        /// </summary>
        [Input("disabledReason")]
        public Input<string>? DisabledReason { get; set; }

        /// <summary>
        /// Is the button clickable.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// whether to show the button.
        /// </summary>
        [Input("supported")]
        public Input<bool>? Supported { get; set; }

        public ClusterOperationInfoAddInstanceArgs()
        {
        }
        public static new ClusterOperationInfoAddInstanceArgs Empty => new ClusterOperationInfoAddInstanceArgs();
    }
}
