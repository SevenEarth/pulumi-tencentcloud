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
    public sealed class WorkflowMediaProcessTaskTranscodeTaskSetOutputStorageCosOutputStorage
    {
        /// <summary>
        /// The target Bucket name of the file output generated by media processing, if not filled, it means the upper layer.
        /// </summary>
        public readonly string? Bucket;
        /// <summary>
        /// The park of the target Bucket for the output of the file generated by media processing. If not filled, it means inheriting from the upper layer.
        /// </summary>
        public readonly string? Region;

        [OutputConstructor]
        private WorkflowMediaProcessTaskTranscodeTaskSetOutputStorageCosOutputStorage(
            string? bucket,

            string? region)
        {
            Bucket = bucket;
            Region = region;
        }
    }
}
