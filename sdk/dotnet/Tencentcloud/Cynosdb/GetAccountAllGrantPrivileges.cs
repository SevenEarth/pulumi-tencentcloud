// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb
{
    public static class GetAccountAllGrantPrivileges
    {
        /// <summary>
        /// Use this data source to query detailed information of cynosdb account_all_grant_privileges
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
        ///         var accountAllGrantPrivileges = Output.Create(Tencentcloud.Cynosdb.GetAccountAllGrantPrivileges.InvokeAsync(new Tencentcloud.Cynosdb.GetAccountAllGrantPrivilegesArgs
        ///         {
        ///             Account = new Tencentcloud.Cynosdb.Inputs.GetAccountAllGrantPrivilegesAccountArgs
        ///             {
        ///                 AccountName = "keep_dts",
        ///                 Host = "%",
        ///             },
        ///             ClusterId = "cynosdbmysql-bws8h88b",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccountAllGrantPrivilegesResult> InvokeAsync(GetAccountAllGrantPrivilegesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountAllGrantPrivilegesResult>("tencentcloud:Cynosdb/getAccountAllGrantPrivileges:getAccountAllGrantPrivileges", args ?? new GetAccountAllGrantPrivilegesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cynosdb account_all_grant_privileges
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
        ///         var accountAllGrantPrivileges = Output.Create(Tencentcloud.Cynosdb.GetAccountAllGrantPrivileges.InvokeAsync(new Tencentcloud.Cynosdb.GetAccountAllGrantPrivilegesArgs
        ///         {
        ///             Account = new Tencentcloud.Cynosdb.Inputs.GetAccountAllGrantPrivilegesAccountArgs
        ///             {
        ///                 AccountName = "keep_dts",
        ///                 Host = "%",
        ///             },
        ///             ClusterId = "cynosdbmysql-bws8h88b",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAccountAllGrantPrivilegesResult> Invoke(GetAccountAllGrantPrivilegesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAccountAllGrantPrivilegesResult>("tencentcloud:Cynosdb/getAccountAllGrantPrivileges:getAccountAllGrantPrivileges", args ?? new GetAccountAllGrantPrivilegesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountAllGrantPrivilegesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// account information.
        /// </summary>
        [Input("account", required: true)]
        public Inputs.GetAccountAllGrantPrivilegesAccountArgs Account { get; set; } = null!;

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetAccountAllGrantPrivilegesArgs()
        {
        }
    }

    public sealed class GetAccountAllGrantPrivilegesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// account information.
        /// </summary>
        [Input("account", required: true)]
        public Input<Inputs.GetAccountAllGrantPrivilegesAccountInputArgs> Account { get; set; } = null!;

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetAccountAllGrantPrivilegesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountAllGrantPrivilegesResult
    {
        public readonly Outputs.GetAccountAllGrantPrivilegesAccountResult Account;
        public readonly string ClusterId;
        /// <summary>
        /// Database permissions note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccountAllGrantPrivilegesDatabasePrivilegeResult> DatabasePrivileges;
        /// <summary>
        /// Global permission note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<string> GlobalPrivileges;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Permission statement note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<string> PrivilegeStatements;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Database table permissions note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAccountAllGrantPrivilegesTablePrivilegeResult> TablePrivileges;

        [OutputConstructor]
        private GetAccountAllGrantPrivilegesResult(
            Outputs.GetAccountAllGrantPrivilegesAccountResult account,

            string clusterId,

            ImmutableArray<Outputs.GetAccountAllGrantPrivilegesDatabasePrivilegeResult> databasePrivileges,

            ImmutableArray<string> globalPrivileges,

            string id,

            ImmutableArray<string> privilegeStatements,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetAccountAllGrantPrivilegesTablePrivilegeResult> tablePrivileges)
        {
            Account = account;
            ClusterId = clusterId;
            DatabasePrivileges = databasePrivileges;
            GlobalPrivileges = globalPrivileges;
            Id = id;
            PrivilegeStatements = privilegeStatements;
            ResultOutputFile = resultOutputFile;
            TablePrivileges = tablePrivileges;
        }
    }
}
