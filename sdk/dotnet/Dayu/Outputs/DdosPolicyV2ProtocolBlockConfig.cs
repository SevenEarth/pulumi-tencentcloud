// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu.Outputs
{

    [OutputType]
    public sealed class DdosPolicyV2ProtocolBlockConfig
    {
        /// <summary>
        /// ICMP block, value [0 (block off), 1 (block on)].
        /// </summary>
        public readonly int DropIcmp;
        /// <summary>
        /// Other block, value [0 (block off), 1 (block on)].
        /// </summary>
        public readonly int DropOther;
        /// <summary>
        /// TCP block, value [0 (block off), 1 (block on)].
        /// </summary>
        public readonly int DropTcp;
        /// <summary>
        /// UDP block, value [0 (block off), 1 (block on)].
        /// </summary>
        public readonly int DropUdp;

        [OutputConstructor]
        private DdosPolicyV2ProtocolBlockConfig(
            int dropIcmp,

            int dropOther,

            int dropTcp,

            int dropUdp)
        {
            DropIcmp = dropIcmp;
            DropOther = dropOther;
            DropTcp = dropTcp;
            DropUdp = dropUdp;
        }
    }
}
