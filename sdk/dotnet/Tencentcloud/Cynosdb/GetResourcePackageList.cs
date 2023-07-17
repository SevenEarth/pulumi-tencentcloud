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
    public static class GetResourcePackageList
    {
        /// <summary>
        /// Use this data source to query detailed information of cynosdb resource_package_list
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
        ///         var resourcePackageList = Output.Create(Tencentcloud.Cynosdb.GetResourcePackageList.InvokeAsync(new Tencentcloud.Cynosdb.GetResourcePackageListArgs
        ///         {
        ///             OrderBies = 
        ///             {
        ///                 "startTime",
        ///             },
        ///             OrderDirection = "DESC",
        ///             PackageIds = 
        ///             {
        ///                 "package-hy4d2ppl",
        ///             },
        ///             PackageNames = 
        ///             {
        ///                 "keep-package-disk",
        ///             },
        ///             PackageRegions = 
        ///             {
        ///                 "china",
        ///             },
        ///             PackageTypes = 
        ///             {
        ///                 "DISK",
        ///             },
        ///             Statuses = 
        ///             {
        ///                 "using",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetResourcePackageListResult> InvokeAsync(GetResourcePackageListArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetResourcePackageListResult>("tencentcloud:Cynosdb/getResourcePackageList:getResourcePackageList", args ?? new GetResourcePackageListArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cynosdb resource_package_list
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
        ///         var resourcePackageList = Output.Create(Tencentcloud.Cynosdb.GetResourcePackageList.InvokeAsync(new Tencentcloud.Cynosdb.GetResourcePackageListArgs
        ///         {
        ///             OrderBies = 
        ///             {
        ///                 "startTime",
        ///             },
        ///             OrderDirection = "DESC",
        ///             PackageIds = 
        ///             {
        ///                 "package-hy4d2ppl",
        ///             },
        ///             PackageNames = 
        ///             {
        ///                 "keep-package-disk",
        ///             },
        ///             PackageRegions = 
        ///             {
        ///                 "china",
        ///             },
        ///             PackageTypes = 
        ///             {
        ///                 "DISK",
        ///             },
        ///             Statuses = 
        ///             {
        ///                 "using",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetResourcePackageListResult> Invoke(GetResourcePackageListInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetResourcePackageListResult>("tencentcloud:Cynosdb/getResourcePackageList:getResourcePackageList", args ?? new GetResourcePackageListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetResourcePackageListArgs : Pulumi.InvokeArgs
    {
        [Input("orderBies")]
        private List<string>? _orderBies;

        /// <summary>
        /// Sorting conditions supported: startTime - effective time, expireTime - expiration time, packageUsedSpec - usage capacity, and packageTotalSpec - total storage capacity. Arrange in array order;.
        /// </summary>
        public List<string> OrderBies
        {
            get => _orderBies ?? (_orderBies = new List<string>());
            set => _orderBies = value;
        }

        /// <summary>
        /// Sort by, DESC Descending, ASC Ascending.
        /// </summary>
        [Input("orderDirection")]
        public string? OrderDirection { get; set; }

        [Input("packageIds")]
        private List<string>? _packageIds;

        /// <summary>
        /// Resource Package Unique ID.
        /// </summary>
        public List<string> PackageIds
        {
            get => _packageIds ?? (_packageIds = new List<string>());
            set => _packageIds = value;
        }

        [Input("packageNames")]
        private List<string>? _packageNames;

        /// <summary>
        /// Resource Package Name.
        /// </summary>
        public List<string> PackageNames
        {
            get => _packageNames ?? (_packageNames = new List<string>());
            set => _packageNames = value;
        }

        [Input("packageRegions")]
        private List<string>? _packageRegions;

        /// <summary>
        /// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and overseas.
        /// </summary>
        public List<string> PackageRegions
        {
            get => _packageRegions ?? (_packageRegions = new List<string>());
            set => _packageRegions = value;
        }

        [Input("packageTypes")]
        private List<string>? _packageTypes;

        /// <summary>
        /// Resource package type CCU - Compute resource package, DISK - Storage resource package.
        /// </summary>
        public List<string> PackageTypes
        {
            get => _packageTypes ?? (_packageTypes = new List<string>());
            set => _packageTypes = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        [Input("statuses")]
        private List<string>? _statuses;

        /// <summary>
        /// Resource package status creating - creating; Using - In use; Expired - has expired; Normal_ Finish - used up; Apply_ Refund - Applying for a refund; Refund - The fee has been refunded.
        /// </summary>
        public List<string> Statuses
        {
            get => _statuses ?? (_statuses = new List<string>());
            set => _statuses = value;
        }

        public GetResourcePackageListArgs()
        {
        }
    }

    public sealed class GetResourcePackageListInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("orderBies")]
        private InputList<string>? _orderBies;

        /// <summary>
        /// Sorting conditions supported: startTime - effective time, expireTime - expiration time, packageUsedSpec - usage capacity, and packageTotalSpec - total storage capacity. Arrange in array order;.
        /// </summary>
        public InputList<string> OrderBies
        {
            get => _orderBies ?? (_orderBies = new InputList<string>());
            set => _orderBies = value;
        }

        /// <summary>
        /// Sort by, DESC Descending, ASC Ascending.
        /// </summary>
        [Input("orderDirection")]
        public Input<string>? OrderDirection { get; set; }

        [Input("packageIds")]
        private InputList<string>? _packageIds;

        /// <summary>
        /// Resource Package Unique ID.
        /// </summary>
        public InputList<string> PackageIds
        {
            get => _packageIds ?? (_packageIds = new InputList<string>());
            set => _packageIds = value;
        }

        [Input("packageNames")]
        private InputList<string>? _packageNames;

        /// <summary>
        /// Resource Package Name.
        /// </summary>
        public InputList<string> PackageNames
        {
            get => _packageNames ?? (_packageNames = new InputList<string>());
            set => _packageNames = value;
        }

        [Input("packageRegions")]
        private InputList<string>? _packageRegions;

        /// <summary>
        /// Resource package usage region China - common in mainland China, overseas - common in Hong Kong, Macao, Taiwan, and overseas.
        /// </summary>
        public InputList<string> PackageRegions
        {
            get => _packageRegions ?? (_packageRegions = new InputList<string>());
            set => _packageRegions = value;
        }

        [Input("packageTypes")]
        private InputList<string>? _packageTypes;

        /// <summary>
        /// Resource package type CCU - Compute resource package, DISK - Storage resource package.
        /// </summary>
        public InputList<string> PackageTypes
        {
            get => _packageTypes ?? (_packageTypes = new InputList<string>());
            set => _packageTypes = value;
        }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        [Input("statuses")]
        private InputList<string>? _statuses;

        /// <summary>
        /// Resource package status creating - creating; Using - In use; Expired - has expired; Normal_ Finish - used up; Apply_ Refund - Applying for a refund; Refund - The fee has been refunded.
        /// </summary>
        public InputList<string> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<string>());
            set => _statuses = value;
        }

        public GetResourcePackageListInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetResourcePackageListResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> OrderBies;
        public readonly string? OrderDirection;
        /// <summary>
        /// Resource Package Unique ID Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<string> PackageIds;
        /// <summary>
        /// Resource package name note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<string> PackageNames;
        /// <summary>
        /// The resource package is used in China, which is commonly used in mainland China, and in overseas, which is commonly used in Hong Kong, Macao, Taiwan, and overseas. Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<string> PackageRegions;
        /// <summary>
        /// Resource package type CCU - Compute resource package, DISK - Store resource package Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<string> PackageTypes;
        /// <summary>
        /// Resource package details note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetResourcePackageListResourcePackageListResult> CynosdbResourcePackageList;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Resource package status creating - creating; Using - In use; Expired - has expired; Normal_ Finish - used up; Apply_ Refund - Applying for a refund; Refund - The fee has been refunded. Note: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly ImmutableArray<string> Statuses;

        [OutputConstructor]
        private GetResourcePackageListResult(
            string id,

            ImmutableArray<string> orderBies,

            string? orderDirection,

            ImmutableArray<string> packageIds,

            ImmutableArray<string> packageNames,

            ImmutableArray<string> packageRegions,

            ImmutableArray<string> packageTypes,

            ImmutableArray<Outputs.GetResourcePackageListResourcePackageListResult> resourcePackageLists,

            string? resultOutputFile,

            ImmutableArray<string> statuses)
        {
            Id = id;
            OrderBies = orderBies;
            OrderDirection = orderDirection;
            PackageIds = packageIds;
            PackageNames = packageNames;
            PackageRegions = packageRegions;
            PackageTypes = packageTypes;
            CynosdbResourcePackageList = resourcePackageLists;
            ResultOutputFile = resultOutputFile;
            Statuses = statuses;
        }
    }
}
