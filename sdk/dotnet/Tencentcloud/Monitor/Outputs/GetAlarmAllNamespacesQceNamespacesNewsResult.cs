// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class GetAlarmAllNamespacesQceNamespacesNewsResult
    {
        /// <summary>
        /// List of supported regions.
        /// </summary>
        public readonly ImmutableArray<string> AvailableRegions;
        /// <summary>
        /// Configuration information.
        /// </summary>
        public readonly string Config;
        /// <summary>
        /// Unique representation in dashboard.
        /// </summary>
        public readonly string DashboardId;
        /// <summary>
        /// Namespace labeling.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Namespace name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Product Name.
        /// </summary>
        public readonly string ProductName;
        /// <summary>
        /// Sort Id.
        /// </summary>
        public readonly int SortId;
        /// <summary>
        /// Namespace value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetAlarmAllNamespacesQceNamespacesNewsResult(
            ImmutableArray<string> availableRegions,

            string config,

            string dashboardId,

            string id,

            string name,

            string productName,

            int sortId,

            string value)
        {
            AvailableRegions = availableRegions;
            Config = config;
            DashboardId = dashboardId;
            Id = id;
            Name = name;
            ProductName = productName;
            SortId = sortId;
            Value = value;
        }
    }
}
