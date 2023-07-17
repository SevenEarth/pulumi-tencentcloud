// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb.Outputs
{

    [OutputType]
    public sealed class GetInstanceSpecsSpecResult
    {
        /// <summary>
        /// machine type.
        /// </summary>
        public readonly string Machine;
        /// <summary>
        /// list of machine specifications.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceSpecsSpecSpecInfoResult> SpecInfos;

        [OutputConstructor]
        private GetInstanceSpecsSpecResult(
            string machine,

            ImmutableArray<Outputs.GetInstanceSpecsSpecSpecInfoResult> specInfos)
        {
            Machine = machine;
            SpecInfos = specInfos;
        }
    }
}
