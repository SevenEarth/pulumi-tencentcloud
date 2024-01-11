// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Antiddos.Outputs
{

    [OutputType]
    public sealed class DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValue
    {
        /// <summary>
        /// Speed limit value type, value [1 (packet rate pps) 2 (bandwidth bps)].
        /// </summary>
        public readonly int Type;
        /// <summary>
        /// value.
        /// </summary>
        public readonly int Value;

        [OutputConstructor]
        private DdosSpeedLimitConfigDdosSpeedLimitConfigSpeedValue(
            int type,

            int value)
        {
            Type = type;
            Value = value;
        }
    }
}