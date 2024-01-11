// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Wedata.Outputs
{

    [OutputType]
    public sealed class IntegrationRealtimeTaskTaskInfoNodeNodeMappingSchemaMapping
    {
        /// <summary>
        /// Schema ID from sink node.
        /// </summary>
        public readonly string SinkSchemaId;
        /// <summary>
        /// Schema ID from source node.
        /// </summary>
        public readonly string SourceSchemaId;

        [OutputConstructor]
        private IntegrationRealtimeTaskTaskInfoNodeNodeMappingSchemaMapping(
            string sinkSchemaId,

            string sourceSchemaId)
        {
            SinkSchemaId = sinkSchemaId;
            SourceSchemaId = sourceSchemaId;
        }
    }
}