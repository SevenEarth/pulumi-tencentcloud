// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Outputs
{

    [OutputType]
    public sealed class GetCompareTasksListConfigObjectItemResult
    {
        public readonly string? DbMode;
        public readonly string? DbName;
        public readonly string? SchemaName;
        public readonly string? TableMode;
        public readonly ImmutableArray<Outputs.GetCompareTasksListConfigObjectItemTableResult> Tables;
        public readonly string? ViewMode;
        public readonly ImmutableArray<Outputs.GetCompareTasksListConfigObjectItemViewResult> Views;

        [OutputConstructor]
        private GetCompareTasksListConfigObjectItemResult(
            string? dbMode,

            string? dbName,

            string? schemaName,

            string? tableMode,

            ImmutableArray<Outputs.GetCompareTasksListConfigObjectItemTableResult> tables,

            string? viewMode,

            ImmutableArray<Outputs.GetCompareTasksListConfigObjectItemViewResult> views)
        {
            DbMode = dbMode;
            DbName = dbName;
            SchemaName = schemaName;
            TableMode = tableMode;
            Tables = tables;
            ViewMode = viewMode;
            Views = views;
        }
    }
}
