// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcm.Inputs
{

    public sealed class MeshConfigIstioSmartDnsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable auto allocate address.
        /// </summary>
        [Input("istioMetaDnsAutoAllocate")]
        public Input<bool>? IstioMetaDnsAutoAllocate { get; set; }

        /// <summary>
        /// Enable dns proxy.
        /// </summary>
        [Input("istioMetaDnsCapture")]
        public Input<bool>? IstioMetaDnsCapture { get; set; }

        public MeshConfigIstioSmartDnsArgs()
        {
        }
        public static new MeshConfigIstioSmartDnsArgs Empty => new MeshConfigIstioSmartDnsArgs();
    }
}
