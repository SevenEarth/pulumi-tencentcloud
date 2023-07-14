// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Outputs
{

    [OutputType]
    public sealed class GetGroupsResultResult
    {
        /// <summary>
        /// Virtual machine deployment group list. Note: This field may return null, indicating that no valid value was found.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupsResultContentResult> Contents;
        /// <summary>
        /// Total count virtual machine deployment group. Note: This field may return null, indicating that no valid value was found.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private GetGroupsResultResult(
            ImmutableArray<Outputs.GetGroupsResultContentResult> contents,

            int totalCount)
        {
            Contents = contents;
            TotalCount = totalCount;
        }
    }
}
