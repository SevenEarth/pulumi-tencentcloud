// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clb.Inputs
{

    public sealed class GetListenersByTargetsBackendArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Private network IP to be queried, which can be of the CVM or ENI.
        /// </summary>
        [Input("privateIp", required: true)]
        public string PrivateIp { get; set; } = null!;

        /// <summary>
        /// VPC ID.
        /// </summary>
        [Input("vpcId", required: true)]
        public string VpcId { get; set; } = null!;

        public GetListenersByTargetsBackendArgs()
        {
        }
        public static new GetListenersByTargetsBackendArgs Empty => new GetListenersByTargetsBackendArgs();
    }
}
