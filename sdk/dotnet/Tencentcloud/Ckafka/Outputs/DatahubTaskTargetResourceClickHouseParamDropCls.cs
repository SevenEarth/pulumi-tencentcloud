// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class DatahubTaskTargetResourceClickHouseParamDropCls
    {
        /// <summary>
        /// cls log set.
        /// </summary>
        public readonly string? DropClsLogSet;
        /// <summary>
        /// Delivery account of cls.
        /// </summary>
        public readonly string? DropClsOwneruin;
        /// <summary>
        /// The region where the cls is delivered.
        /// </summary>
        public readonly string? DropClsRegion;
        /// <summary>
        /// topic of cls.
        /// </summary>
        public readonly string? DropClsTopicId;
        /// <summary>
        /// Whether to deliver to cls.
        /// </summary>
        public readonly bool? DropInvalidMessageToCls;

        [OutputConstructor]
        private DatahubTaskTargetResourceClickHouseParamDropCls(
            string? dropClsLogSet,

            string? dropClsOwneruin,

            string? dropClsRegion,

            string? dropClsTopicId,

            bool? dropInvalidMessageToCls)
        {
            DropClsLogSet = dropClsLogSet;
            DropClsOwneruin = dropClsOwneruin;
            DropClsRegion = dropClsRegion;
            DropClsTopicId = dropClsTopicId;
            DropInvalidMessageToCls = dropInvalidMessageToCls;
        }
    }
}
