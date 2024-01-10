// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps
{
    public static class GetParseNotification
    {
        /// <summary>
        /// Use this data source to query detailed information of mps parse_notification
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
        ///         var parseNotification = Output.Create(Tencentcloud.Mps.GetParseNotification.InvokeAsync(new Tencentcloud.Mps.GetParseNotificationArgs
        ///         {
        ///             Content = "your_content",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetParseNotificationResult> InvokeAsync(GetParseNotificationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetParseNotificationResult>("tencentcloud:Mps/getParseNotification:getParseNotification", args ?? new GetParseNotificationArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mps parse_notification
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
        ///         var parseNotification = Output.Create(Tencentcloud.Mps.GetParseNotification.InvokeAsync(new Tencentcloud.Mps.GetParseNotificationArgs
        ///         {
        ///             Content = "your_content",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetParseNotificationResult> Invoke(GetParseNotificationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetParseNotificationResult>("tencentcloud:Mps/getParseNotification:getParseNotification", args ?? new GetParseNotificationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetParseNotificationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Event notification obtained from CMQ.
        /// </summary>
        [Input("content", required: true)]
        public string Content { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetParseNotificationArgs()
        {
        }
    }

    public sealed class GetParseNotificationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Event notification obtained from CMQ.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetParseNotificationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetParseNotificationResult
    {
        public readonly string Content;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetParseNotificationResult(
            string content,

            string id,

            string? resultOutputFile)
        {
            Content = content;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}
