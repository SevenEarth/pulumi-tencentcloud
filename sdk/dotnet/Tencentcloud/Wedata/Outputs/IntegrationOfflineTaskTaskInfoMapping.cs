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
    public sealed class IntegrationOfflineTaskTaskInfoMapping
    {
        /// <summary>
        /// Node extension configuration information.
        /// </summary>
        public readonly ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoMappingExtConfig> ExtConfigs;
        /// <summary>
        /// Schema mapping information.
        /// </summary>
        public readonly ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoMappingSchemaMapping> SchemaMappings;
        /// <summary>
        /// Sink node ID.
        /// </summary>
        public readonly string? SinkId;
        /// <summary>
        /// Source node ID.
        /// </summary>
        public readonly string? SourceId;
        /// <summary>
        /// Source node schema information.
        /// </summary>
        public readonly ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoMappingSourceSchema> SourceSchemas;

        [OutputConstructor]
        private IntegrationOfflineTaskTaskInfoMapping(
            ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoMappingExtConfig> extConfigs,

            ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoMappingSchemaMapping> schemaMappings,

            string? sinkId,

            string? sourceId,

            ImmutableArray<Outputs.IntegrationOfflineTaskTaskInfoMappingSourceSchema> sourceSchemas)
        {
            ExtConfigs = extConfigs;
            SchemaMappings = schemaMappings;
            SinkId = sinkId;
            SourceId = sourceId;
            SourceSchemas = sourceSchemas;
        }
    }
}
