// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eb.Outputs
{

    [OutputType]
    public sealed class EventTransformTransformationTransform
    {
        /// <summary>
        /// Describe how the data is transformed.
        /// </summary>
        public readonly ImmutableArray<Outputs.EventTransformTransformationTransformOutputStruct> OutputStructs;

        [OutputConstructor]
        private EventTransformTransformationTransform(ImmutableArray<Outputs.EventTransformTransformationTransformOutputStruct> outputStructs)
        {
            OutputStructs = outputStructs;
        }
    }
}
