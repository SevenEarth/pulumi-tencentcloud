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
    public sealed class EditMediaOperationFileInfo
    {
        /// <summary>
        /// End time offset of video clipping in seconds.
        /// </summary>
        public readonly double? EndTimeOffset;
        /// <summary>
        /// Video input information.
        /// </summary>
        public readonly Outputs.EditMediaOperationFileInfoInputInfo InputInfo;
        /// <summary>
        /// Start time offset of video clipping in seconds.
        /// </summary>
        public readonly double? StartTimeOffset;

        [OutputConstructor]
        private EditMediaOperationFileInfo(
            double? endTimeOffset,

            Outputs.EditMediaOperationFileInfoInputInfo inputInfo,

            double? startTimeOffset)
        {
            EndTimeOffset = endTimeOffset;
            InputInfo = inputInfo;
            StartTimeOffset = startTimeOffset;
        }
    }
}