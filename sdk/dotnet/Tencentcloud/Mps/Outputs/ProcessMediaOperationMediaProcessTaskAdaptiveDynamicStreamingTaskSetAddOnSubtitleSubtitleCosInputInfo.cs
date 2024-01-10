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
    public sealed class ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetAddOnSubtitleSubtitleCosInputInfo
    {
        /// <summary>
        /// The COS bucket of the object to process, such as `TopRankVideo-125xxx88`.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// The path of the object to process, such as `/movie/201907/WildAnimal.mov`.
        /// </summary>
        public readonly string Object;
        /// <summary>
        /// The region of the COS bucket, such as `ap-chongqing`.
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private ProcessMediaOperationMediaProcessTaskAdaptiveDynamicStreamingTaskSetAddOnSubtitleSubtitleCosInputInfo(
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
