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
    public sealed class IntegrationTaskNodeNodeInfoExtConfig
    {
        /// <summary>
        /// Configuration name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Configuration value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private IntegrationTaskNodeNodeInfoExtConfig(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
