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

    public sealed class IntegrationRealtimeTaskTaskInfoMappingArgs : global::Pulumi.ResourceArgs
    {
        [Input("extConfigs")]
        private InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingExtConfigArgs>? _extConfigs;

        /// <summary>
        /// Node extension configuration information.
        /// </summary>
        public InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingExtConfigArgs> ExtConfigs
        {
            get => _extConfigs ?? (_extConfigs = new InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingExtConfigArgs>());
            set => _extConfigs = value;
        }

        [Input("schemaMappings")]
        private InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingSchemaMappingArgs>? _schemaMappings;

        /// <summary>
        /// Schema mapping information.
        /// </summary>
        public InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingSchemaMappingArgs> SchemaMappings
        {
            get => _schemaMappings ?? (_schemaMappings = new InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingSchemaMappingArgs>());
            set => _schemaMappings = value;
        }

        /// <summary>
        /// Sink node ID.
        /// </summary>
        [Input("sinkId")]
        public Input<string>? SinkId { get; set; }

        /// <summary>
        /// Source node ID.
        /// </summary>
        [Input("sourceId")]
        public Input<string>? SourceId { get; set; }

        [Input("sourceSchemas")]
        private InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingSourceSchemaArgs>? _sourceSchemas;

        /// <summary>
        /// Source node schema information.
        /// </summary>
        public InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingSourceSchemaArgs> SourceSchemas
        {
            get => _sourceSchemas ?? (_sourceSchemas = new InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingSourceSchemaArgs>());
            set => _sourceSchemas = value;
        }

        public IntegrationRealtimeTaskTaskInfoMappingArgs()
        {
        }
        public static new IntegrationRealtimeTaskTaskInfoMappingArgs Empty => new IntegrationRealtimeTaskTaskInfoMappingArgs();
    }
}
