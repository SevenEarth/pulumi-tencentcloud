// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Outputs
{

    [OutputType]
    public sealed class CngwStrategyConfigBehavior
    {
        /// <summary>
        /// configuration of down scale
        /// Note: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly Outputs.CngwStrategyConfigBehaviorScaleDown? ScaleDown;
        /// <summary>
        /// configuration of up scale
        /// Note: This field may return null, indicating that a valid value is not available.
        /// </summary>
        public readonly Outputs.CngwStrategyConfigBehaviorScaleUp? ScaleUp;

        [OutputConstructor]
        private CngwStrategyConfigBehavior(
            Outputs.CngwStrategyConfigBehaviorScaleDown? scaleDown,

            Outputs.CngwStrategyConfigBehaviorScaleUp? scaleUp)
        {
            ScaleDown = scaleDown;
            ScaleUp = scaleUp;
        }
    }
}
