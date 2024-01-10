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

    public sealed class ScheduleTriggerCosFileUploadTriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the COS bucket bound to a workflow, such as `TopRankVideo-125xxx88`.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Input path directory bound to a workflow, such as `/movie/201907/`. If this parameter is left empty, the `/` root directory will be used.
        /// </summary>
        [Input("dir")]
        public Input<string>? Dir { get; set; }

        [Input("formats")]
        private InputList<string>? _formats;

        /// <summary>
        /// Format list of files that can trigger a workflow, such as [mp4, flv, mov]. If this parameter is left empty, files in all formats can trigger the workflow.
        /// </summary>
        public InputList<string> Formats
        {
            get => _formats ?? (_formats = new InputList<string>());
            set => _formats = value;
        }

        /// <summary>
        /// Region of the COS bucket bound to a workflow, such as `ap-chongiqng`.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public ScheduleTriggerCosFileUploadTriggerArgs()
        {
        }
    }
}
