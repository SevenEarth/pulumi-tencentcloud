// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dlc
{
    public static class GetDescribeUserRoles
    {
        /// <summary>
        /// Use this data source to query detailed information of dlc describe_user_roles
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
        ///     var describeUserRoles = Tencentcloud.Dlc.GetDescribeUserRoles.Invoke(new()
        ///     {
        ///         Fuzzy = "1",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetDescribeUserRolesResult> InvokeAsync(GetDescribeUserRolesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDescribeUserRolesResult>("tencentcloud:Dlc/getDescribeUserRoles:getDescribeUserRoles", args ?? new GetDescribeUserRolesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dlc describe_user_roles
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
        ///     var describeUserRoles = Tencentcloud.Dlc.GetDescribeUserRoles.Invoke(new()
        ///     {
        ///         Fuzzy = "1",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetDescribeUserRolesResult> Invoke(GetDescribeUserRolesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDescribeUserRolesResult>("tencentcloud:Dlc/getDescribeUserRoles:getDescribeUserRoles", args ?? new GetDescribeUserRolesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeUserRolesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// List according to ARN blur.
        /// </summary>
        [Input("fuzzy")]
        public string? Fuzzy { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// The return results are sorted according to this field.
        /// </summary>
        [Input("sortBy")]
        public string? SortBy { get; set; }

        /// <summary>
        /// Positive or inverted, such as DESC.
        /// </summary>
        [Input("sorting")]
        public string? Sorting { get; set; }

        public GetDescribeUserRolesArgs()
        {
        }
        public static new GetDescribeUserRolesArgs Empty => new GetDescribeUserRolesArgs();
    }

    public sealed class GetDescribeUserRolesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// List according to ARN blur.
        /// </summary>
        [Input("fuzzy")]
        public Input<string>? Fuzzy { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// The return results are sorted according to this field.
        /// </summary>
        [Input("sortBy")]
        public Input<string>? SortBy { get; set; }

        /// <summary>
        /// Positive or inverted, such as DESC.
        /// </summary>
        [Input("sorting")]
        public Input<string>? Sorting { get; set; }

        public GetDescribeUserRolesInvokeArgs()
        {
        }
        public static new GetDescribeUserRolesInvokeArgs Empty => new GetDescribeUserRolesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDescribeUserRolesResult
    {
        public readonly string? Fuzzy;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        public readonly string? SortBy;
        public readonly string? Sorting;
        /// <summary>
        /// User role information.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDescribeUserRolesUserRoleResult> UserRoles;

        [OutputConstructor]
        private GetDescribeUserRolesResult(
            string? fuzzy,

            string id,

            string? resultOutputFile,

            string? sortBy,

            string? sorting,

            ImmutableArray<Outputs.GetDescribeUserRolesUserRoleResult> userRoles)
        {
            Fuzzy = fuzzy;
            Id = id;
            ResultOutputFile = resultOutputFile;
            SortBy = sortBy;
            Sorting = sorting;
            UserRoles = userRoles;
        }
    }
}
