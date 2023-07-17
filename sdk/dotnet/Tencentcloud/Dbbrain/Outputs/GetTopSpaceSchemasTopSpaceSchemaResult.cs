// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dbbrain.Outputs
{

    [OutputType]
    public sealed class GetTopSpaceSchemasTopSpaceSchemaResult
    {
        /// <summary>
        /// Fragmentation space (MB).
        /// </summary>
        public readonly double DataFree;
        /// <summary>
        /// data space (MB).
        /// </summary>
        public readonly double DataLength;
        /// <summary>
        /// Fragmentation rate (%).
        /// </summary>
        public readonly double FragRatio;
        /// <summary>
        /// Index space (MB).
        /// </summary>
        public readonly double IndexLength;
        /// <summary>
        /// The sum (MB) of the independent physical file sizes corresponding to all tables in the library. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly double PhysicalFileSize;
        /// <summary>
        /// Number of lines.
        /// </summary>
        public readonly int TableRows;
        /// <summary>
        /// library name.
        /// </summary>
        public readonly string TableSchema;
        /// <summary>
        /// Total space used (MB).
        /// </summary>
        public readonly double TotalLength;

        [OutputConstructor]
        private GetTopSpaceSchemasTopSpaceSchemaResult(
            double dataFree,

            double dataLength,

            double fragRatio,

            double indexLength,

            double physicalFileSize,

            int tableRows,

            string tableSchema,

            double totalLength)
        {
            DataFree = dataFree;
            DataLength = dataLength;
            FragRatio = fragRatio;
            IndexLength = indexLength;
            PhysicalFileSize = physicalFileSize;
            TableRows = tableRows;
            TableSchema = tableSchema;
            TotalLength = totalLength;
        }
    }
}
