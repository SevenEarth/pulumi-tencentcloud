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

    public sealed class ImportOpenApiMicroServiceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Micro service cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Microservice name.
        /// </summary>
        [Input("microServiceName")]
        public Input<string>? MicroServiceName { get; set; }

        /// <summary>
        /// Microservice namespace.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        public ImportOpenApiMicroServiceArgs()
        {
        }
    }
}
