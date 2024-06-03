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
    public static class GetDescribeHostUpdateRecord
    {
        /// <summary>
        /// Use this data source to query detailed information of ssl describe_host_update_record
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
        ///     var describeHostUpdateRecord = Tencentcloud.Ssl.GetDescribeHostUpdateRecord.Invoke(new()
        ///     {
        ///         OldCertificateId = "8u8DII0l",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetDescribeHostUpdateRecordResult> InvokeAsync(GetDescribeHostUpdateRecordArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDescribeHostUpdateRecordResult>("tencentcloud:Ssl/getDescribeHostUpdateRecord:getDescribeHostUpdateRecord", args ?? new GetDescribeHostUpdateRecordArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ssl describe_host_update_record
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
        ///     var describeHostUpdateRecord = Tencentcloud.Ssl.GetDescribeHostUpdateRecord.Invoke(new()
        ///     {
        ///         OldCertificateId = "8u8DII0l",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetDescribeHostUpdateRecordResult> Invoke(GetDescribeHostUpdateRecordInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDescribeHostUpdateRecordResult>("tencentcloud:Ssl/getDescribeHostUpdateRecord:getDescribeHostUpdateRecord", args ?? new GetDescribeHostUpdateRecordInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeHostUpdateRecordArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// New certificate ID.
        /// </summary>
        [Input("certificateId")]
        public string? CertificateId { get; set; }

        /// <summary>
        /// Original certificate ID.
        /// </summary>
        [Input("oldCertificateId")]
        public string? OldCertificateId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDescribeHostUpdateRecordArgs()
        {
        }
        public static new GetDescribeHostUpdateRecordArgs Empty => new GetDescribeHostUpdateRecordArgs();
    }

    public sealed class GetDescribeHostUpdateRecordInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// New certificate ID.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// Original certificate ID.
        /// </summary>
        [Input("oldCertificateId")]
        public Input<string>? OldCertificateId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDescribeHostUpdateRecordInvokeArgs()
        {
        }
        public static new GetDescribeHostUpdateRecordInvokeArgs Empty => new GetDescribeHostUpdateRecordInvokeArgs();
    }


    [OutputType]
    public sealed class GetDescribeHostUpdateRecordResult
    {
        public readonly string? CertificateId;
        /// <summary>
        /// Certificate deployment record listNote: This field may return NULL, indicating that the valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDescribeHostUpdateRecordDeployRecordListResult> DeployRecordLists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OldCertificateId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDescribeHostUpdateRecordResult(
            string? certificateId,

            ImmutableArray<Outputs.GetDescribeHostUpdateRecordDeployRecordListResult> deployRecordLists,

            string id,

            string? oldCertificateId,

            string? resultOutputFile)
        {
            CertificateId = certificateId;
            DeployRecordLists = deployRecordLists;
            Id = id;
            OldCertificateId = oldCertificateId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
