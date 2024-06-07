// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Inputs
{

    public sealed class CngwRouteRateLimitLimitDetailQpsThresholdArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the max threshold.
        /// </summary>
        [Input("max", required: true)]
        public Input<int> Max { get; set; } = null!;

        /// <summary>
        /// qps threshold unit.Reference value:`second`,`minute`,`hour`,`day`,`month`,`year`.
        /// </summary>
        [Input("unit", required: true)]
        public Input<string> Unit { get; set; } = null!;

        public CngwRouteRateLimitLimitDetailQpsThresholdArgs()
        {
        }
        public static new CngwRouteRateLimitLimitDetailQpsThresholdArgs Empty => new CngwRouteRateLimitLimitDetailQpsThresholdArgs();
    }
}
