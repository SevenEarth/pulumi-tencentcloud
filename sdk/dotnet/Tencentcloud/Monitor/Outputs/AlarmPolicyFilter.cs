// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class AlarmPolicyFilter
    {
        /// <summary>
        /// JSON string generated by serializing the AlarmPolicyDimension two-dimensional array.
        /// </summary>
        public readonly string? Dimensions;
        /// <summary>
        /// Filter condition type. Valid values: DIMENSION (uses dimensions for filtering).
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AlarmPolicyFilter(
            string? dimensions,

            string type)
        {
            Dimensions = dimensions;
            Type = type;
        }
    }
}
