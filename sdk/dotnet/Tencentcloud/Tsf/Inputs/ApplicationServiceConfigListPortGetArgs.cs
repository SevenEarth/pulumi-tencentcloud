// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Inputs
{

    public sealed class ApplicationServiceConfigListPortGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Port protocol.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Service port.
        /// </summary>
        [Input("targetPort", required: true)]
        public Input<int> TargetPort { get; set; } = null!;

        public ApplicationServiceConfigListPortGetArgs()
        {
        }
    }
}
