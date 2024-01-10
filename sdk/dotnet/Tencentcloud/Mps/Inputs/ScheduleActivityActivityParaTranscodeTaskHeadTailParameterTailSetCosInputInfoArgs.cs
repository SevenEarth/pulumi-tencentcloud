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

    public sealed class ScheduleActivityActivityParaTranscodeTaskHeadTailParameterTailSetCosInputInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The COS bucket of the object to process, such as `TopRankVideo-125xxx88`.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The path of the object to process, such as `/movie/201907/WildAnimal.mov`.
        /// </summary>
        [Input("object", required: true)]
        public Input<string> Object { get; set; } = null!;

        /// <summary>
        /// The region of the COS bucket, such as `ap-chongqing`.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public ScheduleActivityActivityParaTranscodeTaskHeadTailParameterTailSetCosInputInfoArgs()
        {
        }
    }
}
