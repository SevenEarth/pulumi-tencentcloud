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

    public sealed class WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetObjectNumberFormatGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The growth step of the `{number}` variable, the default is 1.
        /// </summary>
        [Input("increment")]
        public Input<int>? Increment { get; set; }

        /// <summary>
        /// The starting value of `{number}` variable, the default is 0.
        /// </summary>
        [Input("initialValue")]
        public Input<int>? InitialValue { get; set; }

        /// <summary>
        /// The minimum length of the `{number}` variable, if insufficient, placeholders will be filled. Default is 1.
        /// </summary>
        [Input("minLength")]
        public Input<int>? MinLength { get; set; }

        /// <summary>
        /// When the length of the `{number}` variable is insufficient, a placeholder is added. Default is 0.
        /// </summary>
        [Input("placeHolder")]
        public Input<string>? PlaceHolder { get; set; }

        public WorkflowMediaProcessTaskSnapshotByTimeOffsetTaskSetObjectNumberFormatGetArgs()
        {
        }
    }
}
