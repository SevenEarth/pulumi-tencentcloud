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
    public sealed class ContainGroupEnvValueFromFieldRef
    {
        /// <summary>
        /// FieldPath of k8s.
        /// </summary>
        public readonly string? FieldPath;

        [OutputConstructor]
        private ContainGroupEnvValueFromFieldRef(string? fieldPath)
        {
            FieldPath = fieldPath;
        }
    }
}
