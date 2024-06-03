// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dnspod
{
    public static class GetRecordType
    {
        /// <summary>
        /// Use this data source to query detailed information of dnspod record_type
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
        ///     var recordType = Tencentcloud.Dnspod.GetRecordType.Invoke(new()
        ///     {
        ///         DomainGrade = "DP_FREE",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetRecordTypeResult> InvokeAsync(GetRecordTypeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRecordTypeResult>("tencentcloud:Dnspod/getRecordType:getRecordType", args ?? new GetRecordTypeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dnspod record_type
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
        ///     var recordType = Tencentcloud.Dnspod.GetRecordType.Invoke(new()
        ///     {
        ///         DomainGrade = "DP_FREE",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetRecordTypeResult> Invoke(GetRecordTypeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRecordTypeResult>("tencentcloud:Dnspod/getRecordType:getRecordType", args ?? new GetRecordTypeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRecordTypeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Domain level. + Old packages: D_FREE, D_PLUS, D_EXTRA, D_EXPERT, D_ULTRA correspond to free package, personal luxury, enterprise 1, enterprise 2, enterprise 3. + New packages: DP_FREE, DP_PLUS, DP_EXTRA, DP_EXPERT, DP_ULTRA correspond to new free, personal professional, enterprise basic, enterprise standard, enterprise flagship.
        /// </summary>
        [Input("domainGrade", required: true)]
        public string DomainGrade { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetRecordTypeArgs()
        {
        }
        public static new GetRecordTypeArgs Empty => new GetRecordTypeArgs();
    }

    public sealed class GetRecordTypeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Domain level. + Old packages: D_FREE, D_PLUS, D_EXTRA, D_EXPERT, D_ULTRA correspond to free package, personal luxury, enterprise 1, enterprise 2, enterprise 3. + New packages: DP_FREE, DP_PLUS, DP_EXTRA, DP_EXPERT, DP_ULTRA correspond to new free, personal professional, enterprise basic, enterprise standard, enterprise flagship.
        /// </summary>
        [Input("domainGrade", required: true)]
        public Input<string> DomainGrade { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetRecordTypeInvokeArgs()
        {
        }
        public static new GetRecordTypeInvokeArgs Empty => new GetRecordTypeInvokeArgs();
    }


    [OutputType]
    public sealed class GetRecordTypeResult
    {
        public readonly string DomainGrade;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Record type list.
        /// </summary>
        public readonly ImmutableArray<string> TypeLists;

        [OutputConstructor]
        private GetRecordTypeResult(
            string domainGrade,

            string id,

            string? resultOutputFile,

            ImmutableArray<string> typeLists)
        {
            DomainGrade = domainGrade;
            Id = id;
            ResultOutputFile = resultOutputFile;
            TypeLists = typeLists;
        }
    }
}
