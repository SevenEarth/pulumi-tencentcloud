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
    public sealed class WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetObjectNumberFormat
    {
        /// <summary>
        /// The growth step of the `{number}` variable, the default is 1.
        /// </summary>
        public readonly int? Increment;
        /// <summary>
        /// The starting value of `{number}` variable, the default is 0.
        /// </summary>
        public readonly int? InitialValue;
        /// <summary>
        /// The minimum length of the `{number}` variable, if insufficient, placeholders will be filled. Default is 1.
        /// </summary>
        public readonly int? MinLength;
        /// <summary>
        /// When the length of the `{number}` variable is insufficient, a placeholder is added. Default is 0.
        /// </summary>
        public readonly string? PlaceHolder;

        [OutputConstructor]
        private WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetObjectNumberFormat(
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
