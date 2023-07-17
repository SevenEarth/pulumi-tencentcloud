// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls.Inputs
{

    public sealed class AlarmMonitorTimeGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// time period or point in time.
        /// </summary>
        [Input("time", required: true)]
        public Input<int> Time { get; set; } = null!;

        /// <summary>
        /// Period for periodic execution, Fixed for regular execution.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AlarmMonitorTimeGetArgs()
        {
        }
    }
}
