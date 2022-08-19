// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kubernetes.Outputs
{

    [OutputType]
    public sealed class NodePoolTaint
    {
        /// <summary>
        /// Effect of the taint. Valid values are: `NoSchedule`, `PreferNoSchedule`, `NoExecute`.
        /// </summary>
        public readonly string Effect;
        /// <summary>
        /// Key of the taint. The taint key name does not exceed 63 characters, only supports English, numbers,'/','-', and does not allow beginning with ('/').
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Value of the taint.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private NodePoolTaint(
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
