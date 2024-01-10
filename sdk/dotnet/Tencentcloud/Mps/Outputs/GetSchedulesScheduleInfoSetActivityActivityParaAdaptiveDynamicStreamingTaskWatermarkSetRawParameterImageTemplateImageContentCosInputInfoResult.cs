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
    public sealed class GetSchedulesScheduleInfoSetActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentCosInputInfoResult
    {
        /// <summary>
        /// Name of the COS bucket bound to a workflow, such as `TopRankVideo-125xxx88`.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// The path of the object to process, such as `/movie/201907/WildAnimal.mov`.
        /// </summary>
        public readonly string Object;
        /// <summary>
        /// Region of the COS bucket bound to a workflow, such as `ap-chongiqng`.
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private GetSchedulesScheduleInfoSetActivityActivityParaAdaptiveDynamicStreamingTaskWatermarkSetRawParameterImageTemplateImageContentCosInputInfoResult(
            string bucket,

            string @object,

            string region)
        {
            Bucket = bucket;
            Object = @object;
            Region = region;
        }
    }
}
