// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.ApiGateway.Inputs
{

    public sealed class ServiceApiListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the API.
        /// </summary>
        [Input("apiDesc")]
        public Input<string>? ApiDesc { get; set; }

        /// <summary>
        /// ID of the API.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// Name of the API.
        /// </summary>
        [Input("apiName")]
        public Input<string>? ApiName { get; set; }

        /// <summary>
        /// Method of the API.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// Path of the API.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public ServiceApiListArgs()
        {
        }
    }
}
