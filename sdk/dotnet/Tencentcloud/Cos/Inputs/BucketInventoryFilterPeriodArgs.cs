// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos.Inputs
{

    public sealed class BucketInventoryFilterPeriodArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation end time of the objects to analyze. The parameter is a timestamp in seconds, for example, 1568688762.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Creation start time of the objects to analyze. The parameter is a timestamp in seconds, for example, 1568688761.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public BucketInventoryFilterPeriodArgs()
        {
        }
    }
}
