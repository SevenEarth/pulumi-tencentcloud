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
    public sealed class IntegrationRealtimeTaskTaskInfoNodeNodeMappingSourceSchemaProperty
    {
        /// <summary>
        /// Attributes name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Attributes value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private IntegrationRealtimeTaskTaskInfoNodeNodeMappingSourceSchemaProperty(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
