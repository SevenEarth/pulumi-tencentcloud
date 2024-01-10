// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Oceanus.Outputs
{

    [OutputType]
    public sealed class GetTreeResourcesTreeInfoResult
    {
        /// <summary>
        /// Subdirectory Information.
        /// </summary>
        public readonly string Children;
        /// <summary>
        /// ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of items.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTreeResourcesTreeInfoItemResult> Items;
        /// <summary>
        /// Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Parent Id.
        /// </summary>
        public readonly string ParentId;

        [OutputConstructor]
        private GetTreeResourcesTreeInfoResult(
            string children,

            string id,

            ImmutableArray<Outputs.GetTreeResourcesTreeInfoItemResult> items,

            string name,

            string parentId)
        {
            Children = children;
            Id = id;
            Items = items;
            Name = name;
            ParentId = parentId;
        }
    }
}
