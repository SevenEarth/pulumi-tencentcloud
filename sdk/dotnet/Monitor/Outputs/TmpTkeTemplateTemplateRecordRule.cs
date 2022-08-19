// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class TmpTkeTemplateTemplateRecordRule
    {
        /// <summary>
        /// Config.
        /// </summary>
        public readonly string Config;
        /// <summary>
        /// Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Used for the argument, if the configuration comes to the template, the template id.
        /// </summary>
        public readonly string? TemplateId;

        [OutputConstructor]
        private TmpTkeTemplateTemplateRecordRule(
            string config,

            string name,

            string? templateId)
        {
            Config = config;
            Name = name;
            TemplateId = templateId;
        }
    }
}
