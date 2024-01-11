// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dlc.Outputs
{

    [OutputType]
    public sealed class GetDescribeDataEnginePythonSparkImagesPythonSparkImageResult
    {
        /// <summary>
        /// Engine Image version id.
        /// </summary>
        public readonly string ChildImageVersionId;
        /// <summary>
        /// Spark image create time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Spark image description information.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Spark image unique id.
        /// </summary>
        public readonly string SparkImageId;
        /// <summary>
        /// Spark image name.
        /// </summary>
        public readonly string SparkImageVersion;
        /// <summary>
        /// Spark image update time.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetDescribeDataEnginePythonSparkImagesPythonSparkImageResult(
            string childImageVersionId,

            string createTime,

            string description,

            string sparkImageId,

            string sparkImageVersion,

            string updateTime)
        {
            ChildImageVersionId = childImageVersionId;
            CreateTime = createTime;
            Description = description;
            SparkImageId = sparkImageId;
            SparkImageVersion = sparkImageVersion;
            UpdateTime = updateTime;
        }
    }
}