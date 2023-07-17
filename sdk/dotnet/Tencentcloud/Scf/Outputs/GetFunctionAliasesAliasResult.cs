// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf.Outputs
{

    [OutputType]
    public sealed class GetFunctionAliasesAliasResult
    {
        /// <summary>
        /// Creation timeNote: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string AddTime;
        /// <summary>
        /// DescriptionNote: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// If this parameter is provided, only aliases associated with this function version will be returned.
        /// </summary>
        public readonly string FunctionVersion;
        /// <summary>
        /// Update timeNote: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string ModTime;
        /// <summary>
        /// Alias name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Routing information of aliasNote: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionAliasesAliasRoutingConfigResult> RoutingConfigs;

        [OutputConstructor]
        private GetFunctionAliasesAliasResult(
            string addTime,

            string description,

            string functionVersion,

            string modTime,

            string name,

            ImmutableArray<Outputs.GetFunctionAliasesAliasRoutingConfigResult> routingConfigs)
        {
            AddTime = addTime;
            Description = description;
            FunctionVersion = functionVersion;
            ModTime = modTime;
            Name = name;
            RoutingConfigs = routingConfigs;
        }
    }
}
