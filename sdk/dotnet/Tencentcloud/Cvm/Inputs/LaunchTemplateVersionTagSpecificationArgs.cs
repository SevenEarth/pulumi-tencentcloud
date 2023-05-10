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

    public sealed class LaunchTemplateVersionTagSpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of resource that the tag is bound to.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        [Input("tags", required: true)]
        private InputList<Inputs.LaunchTemplateVersionTagSpecificationTagArgs>? _tags;

        /// <summary>
        /// List of tags.
        /// </summary>
        public InputList<Inputs.LaunchTemplateVersionTagSpecificationTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.LaunchTemplateVersionTagSpecificationTagArgs>());
            set => _tags = value;
        }

        public LaunchTemplateVersionTagSpecificationArgs()
        {
        }
    }
}
