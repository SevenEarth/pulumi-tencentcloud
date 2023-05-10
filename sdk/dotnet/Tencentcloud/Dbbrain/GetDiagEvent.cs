// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dbbrain
{
    public static class GetDiagEvent
    {
        /// <summary>
        /// Use this data source to query detailed information of dbbrain diag_event
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
        ///         var diagHistory = Output.Create(Tencentcloud.Dbbrain.GetDiagHistory.InvokeAsync(new Tencentcloud.Dbbrain.GetDiagHistoryArgs
        ///         {
        ///             InstanceId = "%s",
        ///             StartTime = "%s",
        ///             EndTime = "%s",
        ///             Product = "mysql",
        ///         }));
        ///         var diagEvent = diagHistory.Apply(diagHistory =&gt; Output.Create(Tencentcloud.Dbbrain.GetDiagEvent.InvokeAsync(new Tencentcloud.Dbbrain.GetDiagEventArgs
        ///         {
        ///             InstanceId = "%s",
        ///             EventId = diagHistory.Events?[0]?.EventId,
        ///             Product = "mysql",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDiagEventResult> InvokeAsync(GetDiagEventArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDiagEventResult>("tencentcloud:Dbbrain/getDiagEvent:getDiagEvent", args ?? new GetDiagEventArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dbbrain diag_event
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
        ///         var diagHistory = Output.Create(Tencentcloud.Dbbrain.GetDiagHistory.InvokeAsync(new Tencentcloud.Dbbrain.GetDiagHistoryArgs
        ///         {
        ///             InstanceId = "%s",
        ///             StartTime = "%s",
        ///             EndTime = "%s",
        ///             Product = "mysql",
        ///         }));
        ///         var diagEvent = diagHistory.Apply(diagHistory =&gt; Output.Create(Tencentcloud.Dbbrain.GetDiagEvent.InvokeAsync(new Tencentcloud.Dbbrain.GetDiagEventArgs
        ///         {
        ///             InstanceId = "%s",
        ///             EventId = diagHistory.Events?[0]?.EventId,
        ///             Product = "mysql",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDiagEventResult> Invoke(GetDiagEventInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDiagEventResult>("tencentcloud:Dbbrain/getDiagEvent:getDiagEvent", args ?? new GetDiagEventInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDiagEventArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Event ID. Obtain it through `Get Instance Diagnosis History DescribeDBDiagHistory`.
        /// </summary>
        [Input("eventId")]
        public int? EventId { get; set; }

        /// <summary>
        /// isntance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
        /// </summary>
        [Input("product")]
        public string? Product { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDiagEventArgs()
        {
        }
    }

    public sealed class GetDiagEventInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Event ID. Obtain it through `Get Instance Diagnosis History DescribeDBDiagHistory`.
        /// </summary>
        [Input("eventId")]
        public Input<int>? EventId { get; set; }

        /// <summary>
        /// isntance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Service product type, supported values include: `mysql` - cloud database MySQL, `cynosdb` - cloud database CynosDB for MySQL, the default is `mysql`.
        /// </summary>
        [Input("product")]
        public Input<string>? Product { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDiagEventInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDiagEventResult
    {
        /// <summary>
        /// diagnostic item.
        /// </summary>
        public readonly string DiagItem;
        /// <summary>
        /// Diagnostic type.
        /// </summary>
        public readonly string DiagType;
        /// <summary>
        /// End Time.
        /// </summary>
        public readonly string EndTime;
        public readonly int EventId;
        /// <summary>
        /// Diagnostic event details, output is empty if there is no additional explanatory information.
        /// </summary>
        public readonly string Explanation;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// reserved text. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string Metric;
        /// <summary>
        /// Diagnostic summary.
        /// </summary>
        public readonly string Outline;
        /// <summary>
        /// Diagnosed problem.
        /// </summary>
        public readonly string Problem;
        public readonly string? Product;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// severity. The severity is divided into 5 levels, according to the degree of impact from high to low: 1: Fatal, 2: Serious, 3: Warning, 4: Prompt, 5: Healthy.
        /// </summary>
        public readonly int Severity;
        /// <summary>
        /// Starting time.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// A diagnostic suggestion, or empty if there is no suggestion.
        /// </summary>
        public readonly string Suggestions;

        [OutputConstructor]
        private GetDiagEventResult(
            string diagItem,

            string diagType,

            string endTime,

            int eventId,

            string explanation,

            string id,

            string instanceId,

            string metric,

            string outline,

            string problem,

            string? product,

            string? resultOutputFile,

            int severity,

            string startTime,

            string suggestions)
        {
            DiagItem = diagItem;
            DiagType = diagType;
            EndTime = endTime;
            EventId = eventId;
            Explanation = explanation;
            Id = id;
            InstanceId = instanceId;
            Metric = metric;
            Outline = outline;
            Problem = problem;
            Product = product;
            ResultOutputFile = resultOutputFile;
            Severity = severity;
            StartTime = startTime;
            Suggestions = suggestions;
        }
    }
}
