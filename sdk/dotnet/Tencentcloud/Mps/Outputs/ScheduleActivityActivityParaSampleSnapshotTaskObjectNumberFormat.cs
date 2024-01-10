// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Outputs
{

    [OutputType]
    public sealed class ScheduleActivityActivityParaSampleSnapshotTaskObjectNumberFormat
    {
        /// <summary>
        /// Increment of the `{number}` variable. Default value: 1.
        /// </summary>
        public readonly int? Increment;
        /// <summary>
        /// Start value of the `{number}` variable. Default value: 0.
        /// </summary>
        public readonly int? InitialValue;
        /// <summary>
        /// Minimum length of the `{number}` variable. A placeholder will be used if the variable length is below the minimum requirement. Default value: 1.
        /// </summary>
        public readonly int? MinLength;
        /// <summary>
        /// Placeholder used when the `{number}` variable length is below the minimum requirement. Default value: 0.
        /// </summary>
        public readonly string? PlaceHolder;

        [OutputConstructor]
        private ScheduleActivityActivityParaSampleSnapshotTaskObjectNumberFormat(
            int? increment,

            int? initialValue,

            int? minLength,

            string? placeHolder)
        {
            Increment = increment;
            InitialValue = initialValue;
            MinLength = minLength;
            PlaceHolder = placeHolder;
        }
    }
}
