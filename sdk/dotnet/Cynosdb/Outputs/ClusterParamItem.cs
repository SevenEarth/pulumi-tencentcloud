// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cynosdb.Outputs
{

    [OutputType]
    public sealed class ClusterParamItem
    {
        /// <summary>
        /// Param expected value to set.
        /// </summary>
        public readonly string CurrentValue;
        /// <summary>
        /// Name of param, e.g. `character_set_server`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Param old value, indicates the value which already set, this value is required when modifying current_value.
        /// </summary>
        public readonly string? OldValue;

        [OutputConstructor]
        private ClusterParamItem(
            string currentValue,

            string name,

            string? oldValue)
        {
            CurrentValue = currentValue;
            Name = name;
            OldValue = oldValue;
        }
    }
}
