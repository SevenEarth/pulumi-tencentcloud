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

    public sealed class WorkflowTriggerCosFileUploadTriggerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the COS Bucket bound to the workflow.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The input path directory of the workflow binding must be an absolute path, that is, start and end with `/`.
        /// </summary>
        [Input("dir")]
        public Input<string>? Dir { get; set; }

        [Input("formats")]
        private InputList<string>? _formats;

        /// <summary>
        /// A list of file formats that are allowed to be triggered by the workflow, if not filled in, it means that files of all formats can trigger the workflow.
        /// </summary>
        public InputList<string> Formats
        {
            get => _formats ?? (_formats = new InputList<string>());
            set => _formats = value;
        }

        /// <summary>
        /// The park to which the COS Bucket bound to the workflow belongs.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public WorkflowTriggerCosFileUploadTriggerGetArgs()
        {
        }
        public static new WorkflowTriggerCosFileUploadTriggerGetArgs Empty => new WorkflowTriggerCosFileUploadTriggerGetArgs();
    }
}
