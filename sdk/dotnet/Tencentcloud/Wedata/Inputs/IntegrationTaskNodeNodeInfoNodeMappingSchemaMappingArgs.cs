// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Wedata.Inputs
{

    public sealed class IntegrationTaskNodeNodeInfoNodeMappingSchemaMappingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Schema ID from sink node.
        /// </summary>
        [Input("sinkSchemaId", required: true)]
        public Input<string> SinkSchemaId { get; set; } = null!;

        /// <summary>
        /// Schema ID from source node.
        /// </summary>
        [Input("sourceSchemaId", required: true)]
        public Input<string> SourceSchemaId { get; set; } = null!;

        public IntegrationTaskNodeNodeInfoNodeMappingSchemaMappingArgs()
        {
        }
    }
}
