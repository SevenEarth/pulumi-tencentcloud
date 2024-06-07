// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Inputs
{

    public sealed class RenewInstanceInstanceChargePrepaidGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Subscription period; unit: month; valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60. Note: This field may return null, indicating that no valid value is found.
        /// </summary>
        [Input("period", required: true)]
        public Input<int> Period { get; set; } = null!;

        /// <summary>
        /// Auto renewal flag. Valid values:
        /// </summary>
        [Input("renewFlag")]
        public Input<string>? RenewFlag { get; set; }

        public RenewInstanceInstanceChargePrepaidGetArgs()
        {
        }
        public static new RenewInstanceInstanceChargePrepaidGetArgs Empty => new RenewInstanceInstanceChargePrepaidGetArgs();
    }
}
