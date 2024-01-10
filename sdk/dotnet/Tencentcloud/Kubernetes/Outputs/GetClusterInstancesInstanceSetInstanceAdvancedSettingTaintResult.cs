// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes.Outputs
{

    [OutputType]
    public sealed class GetClusterInstancesInstanceSetInstanceAdvancedSettingTaintResult
    {
        /// <summary>
        /// Effect of taints mark.
        /// </summary>
        public readonly string Effect;
        /// <summary>
        /// Key of taints mark.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Value of taints mark.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetClusterInstancesInstanceSetInstanceAdvancedSettingTaintResult(
            string effect,

            string key,

            string value)
        {
            Effect = effect;
            Key = key;
            Value = value;
        }
    }
}
