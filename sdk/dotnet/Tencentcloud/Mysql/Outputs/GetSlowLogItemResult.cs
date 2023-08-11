// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql.Outputs
{

    [OutputType]
    public sealed class GetSlowLogItemResult
    {
        /// <summary>
        /// Backup snapshot time, time format: 2016-03-17 02:10:37.
        /// </summary>
        public readonly string Date;
        /// <summary>
        /// External network download address.
        /// </summary>
        public readonly string InternetUrl;
        /// <summary>
        /// Intranet download address.
        /// </summary>
        public readonly string IntranetUrl;
        /// <summary>
        /// backup file name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Backup file size, unit: Byte.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// Log specific type, possible values: slowlog - slow log.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetSlowLogItemResult(
            string date,

            string internetUrl,

            string intranetUrl,

            string name,

            int size,

            string type)
        {
            Date = date;
            InternetUrl = internetUrl;
            IntranetUrl = intranetUrl;
            Name = name;
            Size = size;
            Type = type;
        }
    }
}