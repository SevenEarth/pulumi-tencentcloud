// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway.Inputs
{

    public sealed class ImportOpenApiServiceTsfLoadBalanceConfArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is load balancing enabled.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("isLoadBalance")]
        public Input<bool>? IsLoadBalance { get; set; }

        /// <summary>
        /// Load balancing method.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// Whether to enable session persistence.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("sessionStickRequired")]
        public Input<bool>? SessionStickRequired { get; set; }

        /// <summary>
        /// Session hold timeout.Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("sessionStickTimeout")]
        public Input<int>? SessionStickTimeout { get; set; }

        public ImportOpenApiServiceTsfLoadBalanceConfArgs()
        {
        }
    }
}
