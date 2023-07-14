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

    public sealed class ImportImageTagSpecificationArgs : Pulumi.ResourceArgs
    {
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        [Input("tags", required: true)]
        private InputList<Inputs.ImportImageTagSpecificationTagArgs>? _tags;
        public InputList<Inputs.ImportImageTagSpecificationTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ImportImageTagSpecificationTagArgs>());
            set => _tags = value;
        }

        public ImportImageTagSpecificationArgs()
        {
        }
    }
}
