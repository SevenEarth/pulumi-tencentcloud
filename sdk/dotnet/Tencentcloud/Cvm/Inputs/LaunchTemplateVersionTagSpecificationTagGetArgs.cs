// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Inputs
{

    public sealed class LaunchTemplateVersionTagSpecificationTagGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tag key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Tag value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public LaunchTemplateVersionTagSpecificationTagGetArgs()
        {
        }
        public static new LaunchTemplateVersionTagSpecificationTagGetArgs Empty => new LaunchTemplateVersionTagSpecificationTagGetArgs();
    }
}
