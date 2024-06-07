// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Inputs
{

    public sealed class ScheduleActivityActivityParaSampleSnapshotTaskObjectNumberFormatArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Increment of the `{number}` variable. Default value: 1.
        /// </summary>
        [Input("increment")]
        public Input<int>? Increment { get; set; }

        /// <summary>
        /// Start value of the `{number}` variable. Default value: 0.
        /// </summary>
        [Input("initialValue")]
        public Input<int>? InitialValue { get; set; }

        /// <summary>
        /// Minimum length of the `{number}` variable. A placeholder will be used if the variable length is below the minimum requirement. Default value: 1.
        /// </summary>
        [Input("minLength")]
        public Input<int>? MinLength { get; set; }

        /// <summary>
        /// Placeholder used when the `{number}` variable length is below the minimum requirement. Default value: 0.
        /// </summary>
        [Input("placeHolder")]
        public Input<string>? PlaceHolder { get; set; }

        public ScheduleActivityActivityParaSampleSnapshotTaskObjectNumberFormatArgs()
        {
        }
        public static new ScheduleActivityActivityParaSampleSnapshotTaskObjectNumberFormatArgs Empty => new ScheduleActivityActivityParaSampleSnapshotTaskObjectNumberFormatArgs();
    }
}
