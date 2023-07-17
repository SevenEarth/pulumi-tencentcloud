// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskTransformsParamRowParamGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// delimiter.
        /// </summary>
        [Input("entryDelimiter")]
        public Input<string>? EntryDelimiter { get; set; }

        /// <summary>
        /// key, value delimiter.
        /// </summary>
        [Input("keyValueDelimiter")]
        public Input<string>? KeyValueDelimiter { get; set; }

        /// <summary>
        /// row content, KEY_VALUE, VALUE.
        /// </summary>
        [Input("rowContent", required: true)]
        public Input<string> RowContent { get; set; } = null!;

        public DatahubTaskTransformsParamRowParamGetArgs()
        {
        }
    }
}
