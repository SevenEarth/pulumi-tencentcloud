// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Outputs
{

    [OutputType]
    public sealed class DomainCompressionCompressionRule
    {
        /// <summary>
        /// List of algorithms, available: `gzip` and `brotli`.
        /// </summary>
        public readonly ImmutableArray<string> Algorithms;
        /// <summary>
        /// Must be set as true, enables compression.
        /// </summary>
        public readonly bool Compress;
        /// <summary>
        /// List of file extensions like `jpg`, `txt`.
        /// </summary>
        public readonly ImmutableArray<string> FileExtensions;
        /// <summary>
        /// The maximum file size to trigger compression (in bytes).
        /// </summary>
        public readonly int MaxLength;
        /// <summary>
        /// The minimum file size to trigger compression (in bytes).
        /// </summary>
        public readonly int MinLength;
        /// <summary>
        /// List of rule paths for each `rule_type`: `*` for `all`, file ext like `jpg` for `file`, `/dir/like/` for `directory` and `/path/index.html` for `path`.
        /// </summary>
        public readonly ImmutableArray<string> RulePaths;
        /// <summary>
        /// Rule type, available: `all`, `file`, `directory`, `path`, `contentType`.
        /// </summary>
        public readonly string? RuleType;

        [OutputConstructor]
        private DomainCompressionCompressionRule(
            ImmutableArray<string> algorithms,

            bool compress,

            ImmutableArray<string> fileExtensions,

            int maxLength,

            int minLength,

            ImmutableArray<string> rulePaths,

            string? ruleType)
        {
            Algorithms = algorithms;
            Compress = compress;
            FileExtensions = fileExtensions;
            MaxLength = maxLength;
            MinLength = minLength;
            RulePaths = rulePaths;
            RuleType = ruleType;
        }
    }
}
