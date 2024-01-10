// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cwp.Outputs
{

    [OutputType]
    public sealed class GetMachinesSimpleMachineTagResult
    {
        /// <summary>
        /// Only supported Keywords, Version and TagId.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Relevance tag id.
        /// </summary>
        public readonly int Rid;
        /// <summary>
        /// Tag ID.
        /// </summary>
        public readonly int TagId;

        [OutputConstructor]
        private GetMachinesSimpleMachineTagResult(
            string name,

            int rid,

            int tagId)
        {
            Name = name;
            Rid = rid;
            TagId = tagId;
        }
    }
}
