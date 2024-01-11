// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cam
{
    public static class GetGroupUserAccount
    {
        /// <summary>
        /// Use this data source to query detailed information of cam group_user_account
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var groupUserAccount = Output.Create(Tencentcloud.Cam.GetGroupUserAccount.InvokeAsync(new Tencentcloud.Cam.GetGroupUserAccountArgs
        ///         {
        ///             SubUin = 100033690181,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupUserAccountResult> InvokeAsync(GetGroupUserAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupUserAccountResult>("tencentcloud:Cam/getGroupUserAccount:getGroupUserAccount", args ?? new GetGroupUserAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cam group_user_account
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var groupUserAccount = Output.Create(Tencentcloud.Cam.GetGroupUserAccount.InvokeAsync(new Tencentcloud.Cam.GetGroupUserAccountArgs
        ///         {
        ///             SubUin = 100033690181,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetGroupUserAccountResult> Invoke(GetGroupUserAccountInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGroupUserAccountResult>("tencentcloud:Cam/getGroupUserAccount:getGroupUserAccount", args ?? new GetGroupUserAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupUserAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Number per page. The default is 20.
        /// </summary>
        [Input("rp")]
        public int? Rp { get; set; }

        /// <summary>
        /// Sub-user uin.
        /// </summary>
        [Input("subUin")]
        public int? SubUin { get; set; }

        /// <summary>
        /// Sub-user uid.
        /// </summary>
        [Input("uid")]
        public int? Uid { get; set; }

        public GetGroupUserAccountArgs()
        {
        }
    }

    public sealed class GetGroupUserAccountInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Number per page. The default is 20.
        /// </summary>
        [Input("rp")]
        public Input<int>? Rp { get; set; }

        /// <summary>
        /// Sub-user uin.
        /// </summary>
        [Input("subUin")]
        public Input<int>? SubUin { get; set; }

        /// <summary>
        /// Sub-user uid.
        /// </summary>
        [Input("uid")]
        public Input<int>? Uid { get; set; }

        public GetGroupUserAccountInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupUserAccountResult
    {
        /// <summary>
        /// User group information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetGroupUserAccountGroupInfoResult> GroupInfos;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly int? Rp;
        public readonly int? SubUin;
        /// <summary>
        /// The total number of user groups the sub-user has joined.
        /// </summary>
        public readonly int TotalNum;
        public readonly int? Uid;

        [OutputConstructor]
        private GetGroupUserAccountResult(
            ImmutableArray<Outputs.GetGroupUserAccountGroupInfoResult> groupInfos,

            string id,

            string? resultOutputFile,

            int? rp,

            int? subUin,

            int totalNum,

            int? uid)
        {
            GroupInfos = groupInfos;
            Id = id;
            ResultOutputFile = resultOutputFile;
            Rp = rp;
            SubUin = subUin;
            TotalNum = totalNum;
            Uid = uid;
        }
    }
}