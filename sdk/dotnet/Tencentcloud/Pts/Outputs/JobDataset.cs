// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts.Outputs
{

    [OutputType]
    public sealed class JobDataset
    {
        public readonly string? FileId;
        public readonly ImmutableArray<string> HeadLines;
        public readonly ImmutableArray<string> HeaderColumns;
        public readonly bool HeaderInFile;
        public readonly int? LineCount;
        public readonly string Name;
        public readonly int? Size;
        public readonly bool Split;
        public readonly ImmutableArray<string> TailLines;
        /// <summary>
        /// Scene Type.
        /// </summary>
        public readonly string? Type;
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private JobDataset(
            string? fileId,

            ImmutableArray<string> headLines,

            ImmutableArray<string> headerColumns,

            bool headerInFile,

            int? lineCount,

            string name,

            int? size,

            bool split,

            ImmutableArray<string> tailLines,

            string? type,

            string? updatedAt)
        {
            FileId = fileId;
            HeadLines = headLines;
            HeaderColumns = headerColumns;
            HeaderInFile = headerInFile;
            LineCount = lineCount;
            Name = name;
            Size = size;
            Split = split;
            TailLines = tailLines;
            Type = type;
            UpdatedAt = updatedAt;
        }
    }
}
