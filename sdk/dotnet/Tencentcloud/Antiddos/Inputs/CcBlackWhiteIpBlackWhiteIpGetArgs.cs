// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Antiddos.Inputs
{

    public sealed class CcBlackWhiteIpBlackWhiteIpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ip address.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// ip mask.
        /// </summary>
        [Input("mask", required: true)]
        public Input<int> Mask { get; set; } = null!;

        public CcBlackWhiteIpBlackWhiteIpGetArgs()
        {
        }
        public static new CcBlackWhiteIpBlackWhiteIpGetArgs Empty => new CcBlackWhiteIpBlackWhiteIpGetArgs();
    }
}
