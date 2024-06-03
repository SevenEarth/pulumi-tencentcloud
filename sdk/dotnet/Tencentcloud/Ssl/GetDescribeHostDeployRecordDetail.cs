// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ssl
{
    public static class GetDescribeHostDeployRecordDetail
    {
        /// <summary>
        /// Use this data source to query detailed information of ssl describe_host_deploy_record_detail
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var describeHostDeployRecordDetail = Tencentcloud.Ssl.GetDescribeHostDeployRecordDetail.Invoke(new()
        ///     {
        ///         DeployRecordId = "",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetDescribeHostDeployRecordDetailResult> InvokeAsync(GetDescribeHostDeployRecordDetailArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDescribeHostDeployRecordDetailResult>("tencentcloud:Ssl/getDescribeHostDeployRecordDetail:getDescribeHostDeployRecordDetail", args ?? new GetDescribeHostDeployRecordDetailArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ssl describe_host_deploy_record_detail
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var describeHostDeployRecordDetail = Tencentcloud.Ssl.GetDescribeHostDeployRecordDetail.Invoke(new()
        ///     {
        ///         DeployRecordId = "",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetDescribeHostDeployRecordDetailResult> Invoke(GetDescribeHostDeployRecordDetailInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDescribeHostDeployRecordDetailResult>("tencentcloud:Ssl/getDescribeHostDeployRecordDetail:getDescribeHostDeployRecordDetail", args ?? new GetDescribeHostDeployRecordDetailInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeHostDeployRecordDetailArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Deployment record ID.
        /// </summary>
        [Input("deployRecordId", required: true)]
        public string DeployRecordId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDescribeHostDeployRecordDetailArgs()
        {
        }
        public static new GetDescribeHostDeployRecordDetailArgs Empty => new GetDescribeHostDeployRecordDetailArgs();
    }

    public sealed class GetDescribeHostDeployRecordDetailInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Deployment record ID.
        /// </summary>
        [Input("deployRecordId", required: true)]
        public Input<string> DeployRecordId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDescribeHostDeployRecordDetailInvokeArgs()
        {
        }
        public static new GetDescribeHostDeployRecordDetailInvokeArgs Empty => new GetDescribeHostDeployRecordDetailInvokeArgs();
    }


    [OutputType]
    public sealed class GetDescribeHostDeployRecordDetailResult
    {
        /// <summary>
        /// Certificate deployment record listNote: This field may return NULL, indicating that the valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDescribeHostDeployRecordDetailDeployRecordDetailListResult> DeployRecordDetailLists;
        public readonly string DeployRecordId;
        /// <summary>
        /// Total number of failuresNote: This field may return NULL, indicating that the valid value cannot be obtained.
        /// </summary>
        public readonly int FailedTotalCount;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Total number of deploymentNote: This field may return NULL, indicating that the valid value cannot be obtained.
        /// </summary>
        public readonly int RunningTotalCount;
        /// <summary>
        /// Total successNote: This field may return NULL, indicating that the valid value cannot be obtained.
        /// </summary>
        public readonly int SuccessTotalCount;

        [OutputConstructor]
        private GetDescribeHostDeployRecordDetailResult(
            ImmutableArray<Outputs.GetDescribeHostDeployRecordDetailDeployRecordDetailListResult> deployRecordDetailLists,

            string deployRecordId,

            int failedTotalCount,

            string id,

            string? resultOutputFile,

            int runningTotalCount,

            int successTotalCount)
        {
            DeployRecordDetailLists = deployRecordDetailLists;
            DeployRecordId = deployRecordId;
            FailedTotalCount = failedTotalCount;
            Id = id;
            ResultOutputFile = resultOutputFile;
            RunningTotalCount = runningTotalCount;
            SuccessTotalCount = successTotalCount;
        }
    }
}
