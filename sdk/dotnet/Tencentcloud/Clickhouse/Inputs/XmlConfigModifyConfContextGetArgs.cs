// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clickhouse.Inputs
{

    public sealed class XmlConfigModifyConfContextGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration file name.
        /// </summary>
        [Input("fileName", required: true)]
        public Input<string> FileName { get; set; } = null!;

        /// <summary>
        /// Path to save configuration file.
        /// </summary>
        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        /// <summary>
        /// New content of configuration file, base64 encoded.
        /// </summary>
        [Input("newConfValue", required: true)]
        public Input<string> NewConfValue { get; set; } = null!;

        public XmlConfigModifyConfContextGetArgs()
        {
        }
        public static new XmlConfigModifyConfContextGetArgs Empty => new XmlConfigModifyConfContextGetArgs();
    }
}
