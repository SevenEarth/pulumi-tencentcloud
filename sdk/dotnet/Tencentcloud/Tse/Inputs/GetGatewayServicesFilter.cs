// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Inputs
{

    public sealed class GetGatewayServicesFilterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// filter name.
        /// </summary>
        [Input("key")]
        public string? Key { get; set; }

        /// <summary>
        /// filter value.
        /// </summary>
        [Input("value")]
        public string? Value { get; set; }

        public GetGatewayServicesFilterArgs()
        {
        }
    }
}
