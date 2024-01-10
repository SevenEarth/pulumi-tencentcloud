// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ses
{
    public static class GetEmailIdentities
    {
        /// <summary>
        /// Use this data source to query detailed information of ses email_identities
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
        ///         var emailIdentities = Output.Create(Tencentcloud.Ses.GetEmailIdentities.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEmailIdentitiesResult> InvokeAsync(GetEmailIdentitiesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEmailIdentitiesResult>("tencentcloud:Ses/getEmailIdentities:getEmailIdentities", args ?? new GetEmailIdentitiesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of ses email_identities
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
        ///         var emailIdentities = Output.Create(Tencentcloud.Ses.GetEmailIdentities.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEmailIdentitiesResult> Invoke(GetEmailIdentitiesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetEmailIdentitiesResult>("tencentcloud:Ses/getEmailIdentities:getEmailIdentities", args ?? new GetEmailIdentitiesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEmailIdentitiesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetEmailIdentitiesArgs()
        {
        }
    }

    public sealed class GetEmailIdentitiesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetEmailIdentitiesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEmailIdentitiesResult
    {
        /// <summary>
        /// Sending domain name list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetEmailIdentitiesEmailIdentityResult> SesEmailIdentities;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Maximum daily sending volume for a single domain name.
        /// </summary>
        public readonly int MaxDailyQuota;
        /// <summary>
        /// Maximum credit rating.
        /// </summary>
        public readonly int MaxReputationLevel;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetEmailIdentitiesResult(
            ImmutableArray<Outputs.GetEmailIdentitiesEmailIdentityResult> emailIdentities,

            string id,

            int maxDailyQuota,

            int maxReputationLevel,

            string? resultOutputFile)
        {
            SesEmailIdentities = emailIdentities;
            Id = id;
            MaxDailyQuota = maxDailyQuota;
            MaxReputationLevel = maxReputationLevel;
            ResultOutputFile = resultOutputFile;
        }
    }
}
