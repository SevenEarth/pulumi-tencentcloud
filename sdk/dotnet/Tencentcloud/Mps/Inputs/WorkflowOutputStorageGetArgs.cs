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

    public sealed class WorkflowOutputStorageGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Valid when Type is COS, this item is required, indicating the media processing COS output location.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("cosOutputStorage")]
        public Input<Inputs.WorkflowOutputStorageCosOutputStorageGetArgs>? CosOutputStorage { get; set; }

        /// <summary>
        /// The type of media processing output object storage location, now only supports COS.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public WorkflowOutputStorageGetArgs()
        {
        }
        public static new WorkflowOutputStorageGetArgs Empty => new WorkflowOutputStorageGetArgs();
    }
}
