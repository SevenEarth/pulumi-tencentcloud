// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Inputs
{

    public sealed class ProcedureTemplateMediaProcessTaskAdaptiveDynamicStreamingTaskListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Adaptive bitrate streaming template ID.
        /// </summary>
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        [Input("watermarkLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskAdaptiveDynamicStreamingTaskListWatermarkListArgs>? _watermarkLists;

        /// <summary>
        /// List of up to `10` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskAdaptiveDynamicStreamingTaskListWatermarkListArgs> WatermarkLists
        {
            get => _watermarkLists ?? (_watermarkLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskAdaptiveDynamicStreamingTaskListWatermarkListArgs>());
            set => _watermarkLists = value;
        }

        public ProcedureTemplateMediaProcessTaskAdaptiveDynamicStreamingTaskListArgs()
        {
        }
    }
}
