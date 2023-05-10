// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Outputs
{

    [OutputType]
    public sealed class LaunchTemplateVersionTagSpecificationTag
    {
        /// <summary>
        /// Tag key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Tag value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private LaunchTemplateVersionTagSpecificationTag(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
