// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ssl.Inputs
{

    public sealed class GetDescribeHostClbInstanceListFilterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter parameter key.
        /// </summary>
        [Input("filterKey", required: true)]
        public string FilterKey { get; set; } = null!;

        /// <summary>
        /// Filter parameter value.
        /// </summary>
        [Input("filterValue", required: true)]
        public string FilterValue { get; set; } = null!;

        public GetDescribeHostClbInstanceListFilterArgs()
        {
        }
        public static new GetDescribeHostClbInstanceListFilterArgs Empty => new GetDescribeHostClbInstanceListFilterArgs();
    }
}
