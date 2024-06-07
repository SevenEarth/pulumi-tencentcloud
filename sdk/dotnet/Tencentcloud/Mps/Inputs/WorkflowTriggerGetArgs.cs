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

    public sealed class WorkflowTriggerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mandatory and valid when Type is CosFileUpload, the rule is triggered for COS.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("cosFileUploadTrigger")]
        public Input<Inputs.WorkflowTriggerCosFileUploadTriggerGetArgs>? CosFileUploadTrigger { get; set; }

        /// <summary>
        /// The type of trigger, currently only supports CosFileUpload.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public WorkflowTriggerGetArgs()
        {
        }
        public static new WorkflowTriggerGetArgs Empty => new WorkflowTriggerGetArgs();
    }
}
