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

    public sealed class IntegrationRealtimeTaskTaskInfoMappingSourceSchemaGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Schema alias.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Schema comment.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Schema ID.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Schema name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("properties")]
        private InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingSourceSchemaPropertyGetArgs>? _properties;

        /// <summary>
        /// Schema extended attributes.
        /// </summary>
        public InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingSourceSchemaPropertyGetArgs> Properties
        {
            get => _properties ?? (_properties = new InputList<Inputs.IntegrationRealtimeTaskTaskInfoMappingSourceSchemaPropertyGetArgs>());
            set => _properties = value;
        }

        /// <summary>
        /// Schema type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Schema value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public IntegrationRealtimeTaskTaskInfoMappingSourceSchemaGetArgs()
        {
        }
        public static new IntegrationRealtimeTaskTaskInfoMappingSourceSchemaGetArgs Empty => new IntegrationRealtimeTaskTaskInfoMappingSourceSchemaGetArgs();
    }
}
