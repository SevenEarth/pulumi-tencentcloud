// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ssl.Outputs
{

    [OutputType]
    public sealed class GetDescribeHostTeoInstanceListFilterResult
    {
        /// <summary>
        /// Filter parameter key.
        /// </summary>
        public readonly string FilterKey;
        /// <summary>
        /// Filter parameter value.
        /// </summary>
        public readonly string FilterValue;

        [OutputConstructor]
        private GetDescribeHostTeoInstanceListFilterResult(
            string filterKey,

            string filterValue)
        {
            FilterKey = filterKey;
            FilterValue = filterValue;
        }
    }
}
