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

    public sealed class TracingConfigZipkinGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Zipkin address.
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        public TracingConfigZipkinGetArgs()
        {
        }
        public static new TracingConfigZipkinGetArgs Empty => new TracingConfigZipkinGetArgs();
    }
}
