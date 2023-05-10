// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Outputs
{

    [OutputType]
    public sealed class GetSyncJobsListTagResult
    {
        /// <summary>
        /// tag key.
        /// </summary>
        public readonly string TagKey;
        /// <summary>
        /// tag value.
        /// </summary>
        public readonly string TagValue;

        [OutputConstructor]
        private GetSyncJobsListTagResult(
            string tagKey,

            string tagValue)
        {
            TagKey = tagKey;
            TagValue = tagValue;
        }
    }
}
