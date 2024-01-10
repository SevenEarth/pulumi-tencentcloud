// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dnspod.Outputs
{

    [OutputType]
    public sealed class GetDomainAnalyticsAliasDataResult
    {
        /// <summary>
        /// Subtotal of resolution volume for the current statistical dimension.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainAnalyticsAliasDataDataResult> Datas;
        /// <summary>
        /// Domain resolution volume statistics query information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainAnalyticsAliasDataInfoResult> Infos;

        [OutputConstructor]
        private GetDomainAnalyticsAliasDataResult(
            ImmutableArray<Outputs.GetDomainAnalyticsAliasDataDataResult> datas,

            ImmutableArray<Outputs.GetDomainAnalyticsAliasDataInfoResult> infos)
        {
            Datas = datas;
            Infos = infos;
        }
    }
}
