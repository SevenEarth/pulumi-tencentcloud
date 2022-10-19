// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vod.Inputs
{

    public sealed class ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time point screen capturing template ID.
        /// </summary>
        [Input("definition", required: true)]
        public Input<string> Definition { get; set; } = null!;

        /// <summary>
        /// Screen capturing mode. Valid values: `Time`, `Percent`. `Time`: screen captures by time point, `Percent`: screen captures by percentage.
        /// </summary>
        [Input("positionType", required: true)]
        public Input<string> PositionType { get; set; } = null!;

        /// <summary>
        /// Screenshot position: For time point screen capturing, this means to take a screenshot at a specified time point (in seconds) and use it as the cover. For percentage screen capturing, this value means to take a screenshot at a specified percentage of the video duration and use it as the cover.
        /// </summary>
        [Input("positionValue", required: true)]
        public Input<double> PositionValue { get; set; } = null!;

        [Input("watermarkLists")]
        private InputList<Inputs.ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListWatermarkListGetArgs>? _watermarkLists;

        /// <summary>
        /// List of up to `10` image or text watermarks. Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public InputList<Inputs.ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListWatermarkListGetArgs> WatermarkLists
        {
            get => _watermarkLists ?? (_watermarkLists = new InputList<Inputs.ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListWatermarkListGetArgs>());
            set => _watermarkLists = value;
        }

        public ProcedureTemplateMediaProcessTaskCoverBySnapshotTaskListGetArgs()
        {
        }
    }
}