// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb.Outputs
{

    [OutputType]
    public sealed class GetAccountsAccountSetResult
    {
        /// <summary>
        /// Account name of database.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// Create time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The account description of database.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Host.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Maximum number of user connections.
        /// </summary>
        public readonly int MaxUserConnections;
        /// <summary>
        /// Update time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetAccountsAccountSetResult(
            string accountName,

            string createTime,

            string description,

            string host,

            int maxUserConnections,

            string updateTime)
        {
            AccountName = accountName;
            CreateTime = createTime;
            Description = description;
            Host = host;
            MaxUserConnections = maxUserConnections;
            UpdateTime = updateTime;
        }
    }
}
