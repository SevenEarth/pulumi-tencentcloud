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

    public sealed class IntegrationOfflineTaskTaskInfoMappingExtConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public IntegrationOfflineTaskTaskInfoMappingExtConfigArgs()
        {
        }
        public static new IntegrationOfflineTaskTaskInfoMappingExtConfigArgs Empty => new IntegrationOfflineTaskTaskInfoMappingExtConfigArgs();
    }
}
