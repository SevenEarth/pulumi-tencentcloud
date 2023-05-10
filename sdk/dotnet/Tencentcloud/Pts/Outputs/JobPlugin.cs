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
    public sealed class JobPlugin
    {
        public readonly string? FileId;
        public readonly string? Name;
        public readonly int? Size;
        /// <summary>
        /// Scene Type.
        /// </summary>
        public readonly string? Type;
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private JobPlugin(
            string? fileId,

            string? name,

            int? size,

            string? type,

            string? updatedAt)
        {
            FileId = fileId;
            Name = name;
            Size = size;
            Type = type;
            UpdatedAt = updatedAt;
        }
    }
}
