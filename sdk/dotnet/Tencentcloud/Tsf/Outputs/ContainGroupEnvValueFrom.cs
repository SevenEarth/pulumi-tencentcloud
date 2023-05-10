// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Outputs
{

    [OutputType]
    public sealed class ContainGroupEnvValueFrom
    {
        /// <summary>
        /// FieldRef for k8s env.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainGroupEnvValueFromFieldRef> FieldReves;
        /// <summary>
        /// ResourceFieldRef of k8s env.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainGroupEnvValueFromResourceFieldRef> ResourceFieldReves;

        [OutputConstructor]
        private ContainGroupEnvValueFrom(
            ImmutableArray<Outputs.ContainGroupEnvValueFromFieldRef> fieldReves,

            ImmutableArray<Outputs.ContainGroupEnvValueFromResourceFieldRef> resourceFieldReves)
        {
            FieldReves = fieldReves;
            ResourceFieldReves = resourceFieldReves;
        }
    }
}