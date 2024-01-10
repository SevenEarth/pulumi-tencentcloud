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

    public sealed class IntegrationTaskNodeNodeInfoNodeMappingArgs : Pulumi.ResourceArgs
    {
        [Input("extConfigs")]
        private InputList<Inputs.IntegrationTaskNodeNodeInfoNodeMappingExtConfigArgs>? _extConfigs;

        /// <summary>
        /// Node extension configuration information.
        /// </summary>
        public InputList<Inputs.IntegrationTaskNodeNodeInfoNodeMappingExtConfigArgs> ExtConfigs
        {
            get => _extConfigs ?? (_extConfigs = new InputList<Inputs.IntegrationTaskNodeNodeInfoNodeMappingExtConfigArgs>());
            set => _extConfigs = value;
        }

        [Input("schemaMappings")]
        private InputList<Inputs.IntegrationTaskNodeNodeInfoNodeMappingSchemaMappingArgs>? _schemaMappings;

        /// <summary>
        /// Schema mapping information.
        /// </summary>
        public InputList<Inputs.IntegrationTaskNodeNodeInfoNodeMappingSchemaMappingArgs> SchemaMappings
        {
            get => _schemaMappings ?? (_schemaMappings = new InputList<Inputs.IntegrationTaskNodeNodeInfoNodeMappingSchemaMappingArgs>());
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
        private InputList<Inputs.IntegrationTaskNodeNodeInfoNodeMappingSourceSchemaArgs>? _sourceSchemas;

        /// <summary>
        /// Source node schema information.
        /// </summary>
        public InputList<Inputs.IntegrationTaskNodeNodeInfoNodeMappingSourceSchemaArgs> SourceSchemas
        {
            get => _sourceSchemas ?? (_sourceSchemas = new InputList<Inputs.IntegrationTaskNodeNodeInfoNodeMappingSourceSchemaArgs>());
            set => _sourceSchemas = value;
        }

        public IntegrationTaskNodeNodeInfoNodeMappingArgs()
        {
        }
    }
}
