// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ccn
{
    public static class GetTenantInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of vpc tenant_ccn
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
        ///         var tenantCcn = Output.Create(Tencentcloud.Ccn.GetTenantInstances.InvokeAsync(new Tencentcloud.Ccn.GetTenantInstancesArgs
        ///         {
        ///             CcnIds = 
        ///             {
        ///                 "ccn-39lqkygf",
        ///             },
        ///             IsSecurityLocks = 
        ///             {
        ///                 "true",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetTenantInstancesResult> InvokeAsync(GetTenantInstancesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTenantInstancesResult>("tencentcloud:Ccn/getTenantInstances:getTenantInstances", args ?? new GetTenantInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of vpc tenant_ccn
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
        ///         var tenantCcn = Output.Create(Tencentcloud.Ccn.GetTenantInstances.InvokeAsync(new Tencentcloud.Ccn.GetTenantInstancesArgs
        ///         {
        ///             CcnIds = 
        ///             {
        ///                 "ccn-39lqkygf",
        ///             },
        ///             IsSecurityLocks = 
        ///             {
        ///                 "true",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetTenantInstancesResult> Invoke(GetTenantInstancesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTenantInstancesResult>("tencentcloud:Ccn/getTenantInstances:getTenantInstances", args ?? new GetTenantInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTenantInstancesArgs : Pulumi.InvokeArgs
    {
        [Input("ccnIds")]
        private List<string>? _ccnIds;

        /// <summary>
        /// filter by ccn ids, like: ['ccn-12345678'].
        /// </summary>
        public List<string> CcnIds
        {
            get => _ccnIds ?? (_ccnIds = new List<string>());
            set => _ccnIds = value;
        }

        [Input("isSecurityLocks")]
        private List<string>? _isSecurityLocks;

        /// <summary>
        /// filter by locked, like ['true'].
        /// </summary>
        public List<string> IsSecurityLocks
        {
            get => _isSecurityLocks ?? (_isSecurityLocks = new List<string>());
            set => _isSecurityLocks = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("userAccountIds")]
        private List<string>? _userAccountIds;

        /// <summary>
        /// filter by ccn ids, like: ['12345678'].
        /// </summary>
        public List<string> UserAccountIds
        {
            get => _userAccountIds ?? (_userAccountIds = new List<string>());
            set => _userAccountIds = value;
        }

        public GetTenantInstancesArgs()
        {
        }
    }

    public sealed class GetTenantInstancesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("ccnIds")]
        private InputList<string>? _ccnIds;

        /// <summary>
        /// filter by ccn ids, like: ['ccn-12345678'].
        /// </summary>
        public InputList<string> CcnIds
        {
            get => _ccnIds ?? (_ccnIds = new InputList<string>());
            set => _ccnIds = value;
        }

        [Input("isSecurityLocks")]
        private InputList<string>? _isSecurityLocks;

        /// <summary>
        /// filter by locked, like ['true'].
        /// </summary>
        public InputList<string> IsSecurityLocks
        {
            get => _isSecurityLocks ?? (_isSecurityLocks = new InputList<string>());
            set => _isSecurityLocks = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("userAccountIds")]
        private InputList<string>? _userAccountIds;

        /// <summary>
        /// filter by ccn ids, like: ['12345678'].
        /// </summary>
        public InputList<string> UserAccountIds
        {
            get => _userAccountIds ?? (_userAccountIds = new InputList<string>());
            set => _userAccountIds = value;
        }

        public GetTenantInstancesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTenantInstancesResult
    {
        public readonly ImmutableArray<string> CcnIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> IsSecurityLocks;
        public readonly string? ResultOutputFile;
        public readonly ImmutableArray<string> UserAccountIds;

        [OutputConstructor]
        private GetTenantInstancesResult(
            ImmutableArray<string> ccnIds,

            string id,

            ImmutableArray<string> isSecurityLocks,

            string? resultOutputFile,

            ImmutableArray<string> userAccountIds)
        {
            CcnIds = ccnIds;
            Id = id;
            IsSecurityLocks = isSecurityLocks;
            ResultOutputFile = resultOutputFile;
            UserAccountIds = userAccountIds;
        }
    }
}
