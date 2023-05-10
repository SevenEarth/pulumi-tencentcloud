// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcm.Outputs
{

    [OutputType]
    public sealed class AccessLogConfigSelectedRange
    {
        /// <summary>
        /// Select all if true, default false.
        /// </summary>
        public readonly bool? All;
        /// <summary>
        /// Items.
        /// </summary>
        public readonly ImmutableArray<Outputs.AccessLogConfigSelectedRangeItem> Items;

        [OutputConstructor]
        private AccessLogConfigSelectedRange(
            bool? all,

            ImmutableArray<Outputs.AccessLogConfigSelectedRangeItem> items)
        {
            All = all;
            Items = items;
        }
    }
}
