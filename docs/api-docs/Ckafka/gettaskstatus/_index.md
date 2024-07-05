
---
title: "getTaskStatus"
title_tag: "tencentcloud.Ckafka.getTaskStatus"
meta_desc: "Documentation for the tencentcloud.Ckafka.getTaskStatus function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Use this data source to query detailed information of ckafka task_status

## Example Usage

<!--Start PulumiCodeChooser -->
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as tencentcloud from "@tencentcloud_iac/pulumi";

const taskStatus = tencentcloud.Ckafka.getTaskStatus({
    flowId: 123456,
});
```
```python
import pulumi
import pulumi_tencentcloud as tencentcloud

task_status = tencentcloud.Ckafka.get_task_status(flow_id=123456)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Tencentcloud = Pulumi.Tencentcloud;

return await Deployment.RunAsync(() => 
{
    var taskStatus = Tencentcloud.Ckafka.GetTaskStatus.Invoke(new()
    {
        FlowId = 123456,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ckafka"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := Ckafka.GetTaskStatus(ctx, &ckafka.GetTaskStatusArgs{
			FlowId: 123456,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.tencentcloud.Ckafka.CkafkaFunctions;
import com.pulumi.tencentcloud.Ckafka.inputs.GetTaskStatusArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        final var taskStatus = CkafkaFunctions.getTaskStatus(GetTaskStatusArgs.builder()
            .flowId(123456)
            .build());

    }
}
```
```yaml
variables:
  taskStatus:
    fn::invoke:
      Function: tencentcloud:Ckafka:getTaskStatus
      Arguments:
        flowId: 123456
```
<!--End PulumiCodeChooser -->




## Using getTaskStatus {#using}

Two invocation forms are available. The direct form accepts plain
arguments and either blocks until the result value is available, or
returns a Promise-wrapped result. The output form accepts
Input-wrapped arguments and returns an Output-wrapped result.

<div>
<pulumi-chooser type="language" options="typescript,python,go,csharp,java,yaml"></pulumi-chooser>
</div>


<div>
<pulumi-choosable type="language" values="javascript,typescript">
<div class="highlight"
><pre class="chroma"><code class="language-typescript" data-lang="typescript"
><span class="k">function </span>getTaskStatus<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetTaskStatusArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetTaskStatusResult</a></span>></span
><span class="k">
function </span>getTaskStatusOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetTaskStatusOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetTaskStatusResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_task_status<span class="p">(</span><span class="nx">flow_id</span><span class="p">:</span> <span class="nx">Optional[int]</span> = None<span class="p">,</span>
                    <span class="nx">result_output_file</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                    <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetTaskStatusResult</span
><span class="k">
def </span>get_task_status_output<span class="p">(</span><span class="nx">flow_id</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[int]]</span> = None<span class="p">,</span>
                    <span class="nx">result_output_file</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                    <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetTaskStatusResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>GetTaskStatus<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetTaskStatusArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">GetTaskStatusResult</a></span>, error)</span
><span class="k">
func </span>GetTaskStatusOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetTaskStatusOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) GetTaskStatusResultOutput</span
></code></pre></div>

&gt; Note: This function is named `GetTaskStatus` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetTaskStatus </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetTaskStatusResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">GetTaskStatusArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetTaskStatusResult</a></span>> <span class="p">Invoke(</span><span class="nx">GetTaskStatusInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetTaskStatusResult</a></span>> </span>getTaskStatus<span class="p">(</span><span class="nx">GetTaskStatusArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="c">// Output-based functions aren't available in Java yet</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> tencentcloud:Ckafka/getTaskStatus:getTaskStatus
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="flowid_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowid_csharp" style="color: inherit; text-decoration: inherit;">Flow<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd><p>FlowId.</p>
</dd><dt class="property-optional"
            title="Optional">
        <span id="resultoutputfile_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resultoutputfile_csharp" style="color: inherit; text-decoration: inherit;">Result<wbr>Output<wbr>File</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Used to save results.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="flowid_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowid_go" style="color: inherit; text-decoration: inherit;">Flow<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd><p>FlowId.</p>
</dd><dt class="property-optional"
            title="Optional">
        <span id="resultoutputfile_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resultoutputfile_go" style="color: inherit; text-decoration: inherit;">Result<wbr>Output<wbr>File</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Used to save results.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="flowid_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowid_java" style="color: inherit; text-decoration: inherit;">flow<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Integer</span>
    </dt>
    <dd><p>FlowId.</p>
</dd><dt class="property-optional"
            title="Optional">
        <span id="resultoutputfile_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resultoutputfile_java" style="color: inherit; text-decoration: inherit;">result<wbr>Output<wbr>File</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>Used to save results.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="flowid_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowid_nodejs" style="color: inherit; text-decoration: inherit;">flow<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">number</span>
    </dt>
    <dd><p>FlowId.</p>
</dd><dt class="property-optional"
            title="Optional">
        <span id="resultoutputfile_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resultoutputfile_nodejs" style="color: inherit; text-decoration: inherit;">result<wbr>Output<wbr>File</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Used to save results.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="flow_id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flow_id_python" style="color: inherit; text-decoration: inherit;">flow_<wbr>id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd><p>FlowId.</p>
</dd><dt class="property-optional"
            title="Optional">
        <span id="result_output_file_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#result_output_file_python" style="color: inherit; text-decoration: inherit;">result_<wbr>output_<wbr>file</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd><p>Used to save results.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="flowid_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowid_yaml" style="color: inherit; text-decoration: inherit;">flow<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Number</span>
    </dt>
    <dd><p>FlowId.</p>
</dd><dt class="property-optional"
            title="Optional">
        <span id="resultoutputfile_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resultoutputfile_yaml" style="color: inherit; text-decoration: inherit;">result<wbr>Output<wbr>File</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>Used to save results.</p>
</dd></dl>
</pulumi-choosable>
</div>




## getTaskStatus Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="flowid_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowid_csharp" style="color: inherit; text-decoration: inherit;">Flow<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="id_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_csharp" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>The provider-assigned unique ID for this managed resource.</p>
</dd><dt class="property-"
            title="">
        <span id="results_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#results_csharp" style="color: inherit; text-decoration: inherit;">Results</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#gettaskstatusresult">List&lt;Tencent<wbr>Cloud<wbr>IAC.<wbr>Pulumi<wbr>Package.<wbr>Tencentcloud.<wbr>Ckafka.<wbr>Outputs.<wbr>Get<wbr>Task<wbr>Status<wbr>Result&gt;</a></span>
    </dt>
    <dd><p>Result.</p>
</dd><dt class="property-"
            title="">
        <span id="resultoutputfile_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resultoutputfile_csharp" style="color: inherit; text-decoration: inherit;">Result<wbr>Output<wbr>File</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="flowid_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowid_go" style="color: inherit; text-decoration: inherit;">Flow<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="id_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_go" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>The provider-assigned unique ID for this managed resource.</p>
</dd><dt class="property-"
            title="">
        <span id="results_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#results_go" style="color: inherit; text-decoration: inherit;">Results</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#gettaskstatusresult">[]Get<wbr>Task<wbr>Status<wbr>Result</a></span>
    </dt>
    <dd><p>Result.</p>
</dd><dt class="property-"
            title="">
        <span id="resultoutputfile_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resultoutputfile_go" style="color: inherit; text-decoration: inherit;">Result<wbr>Output<wbr>File</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="flowid_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowid_java" style="color: inherit; text-decoration: inherit;">flow<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Integer</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="id_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_java" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>The provider-assigned unique ID for this managed resource.</p>
</dd><dt class="property-"
            title="">
        <span id="results_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#results_java" style="color: inherit; text-decoration: inherit;">results</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#gettaskstatusresult">List&lt;Get<wbr>Task<wbr>Status<wbr>Result&gt;</a></span>
    </dt>
    <dd><p>Result.</p>
</dd><dt class="property-"
            title="">
        <span id="resultoutputfile_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resultoutputfile_java" style="color: inherit; text-decoration: inherit;">result<wbr>Output<wbr>File</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="flowid_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowid_nodejs" style="color: inherit; text-decoration: inherit;">flow<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">number</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="id_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_nodejs" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>The provider-assigned unique ID for this managed resource.</p>
</dd><dt class="property-"
            title="">
        <span id="results_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#results_nodejs" style="color: inherit; text-decoration: inherit;">results</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#gettaskstatusresult">Get<wbr>Task<wbr>Status<wbr>Result[]</a></span>
    </dt>
    <dd><p>Result.</p>
</dd><dt class="property-"
            title="">
        <span id="resultoutputfile_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resultoutputfile_nodejs" style="color: inherit; text-decoration: inherit;">result<wbr>Output<wbr>File</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="flow_id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flow_id_python" style="color: inherit; text-decoration: inherit;">flow_<wbr>id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_python" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd><p>The provider-assigned unique ID for this managed resource.</p>
</dd><dt class="property-"
            title="">
        <span id="results_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#results_python" style="color: inherit; text-decoration: inherit;">results</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#gettaskstatusresult">Get<wbr>Task<wbr>Status<wbr>Result]</a></span>
    </dt>
    <dd><p>Result.</p>
</dd><dt class="property-"
            title="">
        <span id="result_output_file_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#result_output_file_python" style="color: inherit; text-decoration: inherit;">result_<wbr>output_<wbr>file</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="flowid_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#flowid_yaml" style="color: inherit; text-decoration: inherit;">flow<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Number</span>
    </dt>
    <dd></dd><dt class="property-"
            title="">
        <span id="id_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_yaml" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>The provider-assigned unique ID for this managed resource.</p>
</dd><dt class="property-"
            title="">
        <span id="results_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#results_yaml" style="color: inherit; text-decoration: inherit;">results</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#gettaskstatusresult">List&lt;Property Map&gt;</a></span>
    </dt>
    <dd><p>Result.</p>
</dd><dt class="property-"
            title="">
        <span id="resultoutputfile_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#resultoutputfile_yaml" style="color: inherit; text-decoration: inherit;">result<wbr>Output<wbr>File</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>




## Supporting Types


<h4 id="gettaskstatusresult">Get<wbr>Task<wbr>Status<wbr>Result</h4>



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="output_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#output_csharp" style="color: inherit; text-decoration: inherit;">Output</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>OutPut Info.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="status_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#status_csharp" style="color: inherit; text-decoration: inherit;">Status</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd><p>Status.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="output_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#output_go" style="color: inherit; text-decoration: inherit;">Output</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>OutPut Info.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="status_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#status_go" style="color: inherit; text-decoration: inherit;">Status</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd><p>Status.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="output_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#output_java" style="color: inherit; text-decoration: inherit;">output</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>OutPut Info.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="status_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#status_java" style="color: inherit; text-decoration: inherit;">status</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Integer</span>
    </dt>
    <dd><p>Status.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="output_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#output_nodejs" style="color: inherit; text-decoration: inherit;">output</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>OutPut Info.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="status_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#status_nodejs" style="color: inherit; text-decoration: inherit;">status</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">number</span>
    </dt>
    <dd><p>Status.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="output_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#output_python" style="color: inherit; text-decoration: inherit;">output</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd><p>OutPut Info.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="status_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#status_python" style="color: inherit; text-decoration: inherit;">status</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">int</span>
    </dt>
    <dd><p>Status.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="output_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#output_yaml" style="color: inherit; text-decoration: inherit;">output</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>OutPut Info.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="status_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#status_yaml" style="color: inherit; text-decoration: inherit;">status</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Number</span>
    </dt>
    <dd><p>Status.</p>
</dd></dl>
</pulumi-choosable>
</div>





<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="https://github.com/tencentcloudstack/pulumi-tencentcloud">tencentcloud tencentcloudstack/pulumi-tencentcloud</a></dd>
	<dt>License</dt>
	<dd>Apache-2.0</dd>
	<dt>Notes</dt>
	<dd><p>This Pulumi package is based on the <a href="https://github.com/tencentcloudstack/terraform-provider-tencentcloud"><code>tencentcloud</code> Terraform Provider</a>.</p>
</dd>
</dl>
