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

    public sealed class CngwStrategyConfigBehaviorScaleUpPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// period of scale up
        /// Note: This field may return null, indicating that a valid value is not available.
        /// </summary>
        [Input("periodSeconds")]
        public Input<int>? PeriodSeconds { get; set; }

        /// <summary>
        /// type, default value: Pods
        /// Note: This field may return null, indicating that a valid value is not available.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// value
        /// Note: This field may return null, indicating that a valid value is not available.
        /// </summary>
        [Input("value")]
        public Input<int>? Value { get; set; }

        public CngwStrategyConfigBehaviorScaleUpPolicyArgs()
        {
        }
        public static new CngwStrategyConfigBehaviorScaleUpPolicyArgs Empty => new CngwStrategyConfigBehaviorScaleUpPolicyArgs();
    }
}
