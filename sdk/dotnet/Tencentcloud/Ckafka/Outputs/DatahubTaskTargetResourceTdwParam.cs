// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class DatahubTaskTargetResourceTdwParam
    {
        /// <summary>
        /// Tdw bid.
        /// </summary>
        public readonly string Bid;
        /// <summary>
        /// default true.
        /// </summary>
        public readonly bool? IsDomestic;
        /// <summary>
        /// TDW address, defalt tl-tdbank-tdmanager.tencent-distribute.com.
        /// </summary>
        public readonly string? TdwHost;
        /// <summary>
        /// TDW port, default 8099.
        /// </summary>
        public readonly int? TdwPort;
        /// <summary>
        /// Tdw tid.
        /// </summary>
        public readonly string Tid;

        [OutputConstructor]
        private DatahubTaskTargetResourceTdwParam(
            string bid,

            bool? isDomestic,

            string? tdwHost,

            int? tdwPort,

            string tid)
        {
            Bid = bid;
            IsDomestic = isDomestic;
            TdwHost = tdwHost;
            TdwPort = tdwPort;
            Tid = tid;
        }
    }
}
