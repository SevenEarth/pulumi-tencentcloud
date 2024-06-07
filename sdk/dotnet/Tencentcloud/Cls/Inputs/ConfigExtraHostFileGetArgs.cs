// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls.Inputs
{

    public sealed class ConfigExtraHostFileGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("customLabels")]
        private InputList<string>? _customLabels;

        /// <summary>
        /// Metadata info.
        /// </summary>
        public InputList<string> CustomLabels
        {
            get => _customLabels ?? (_customLabels = new InputList<string>());
            set => _customLabels = value;
        }

        /// <summary>
        /// Log file name.
        /// </summary>
        [Input("filePattern", required: true)]
        public Input<string> FilePattern { get; set; } = null!;

        /// <summary>
        /// Log file dir.
        /// </summary>
        [Input("logPath", required: true)]
        public Input<string> LogPath { get; set; } = null!;

        public ConfigExtraHostFileGetArgs()
        {
        }
        public static new ConfigExtraHostFileGetArgs Empty => new ConfigExtraHostFileGetArgs();
    }
}
