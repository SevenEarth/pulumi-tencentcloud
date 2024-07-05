
---
title: "getGrafanaPluginOverviews"
title_tag: "tencentcloud.Monitor.getGrafanaPluginOverviews"
meta_desc: "Documentation for the tencentcloud.Monitor.getGrafanaPluginOverviews function with examples, input properties, output properties, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by Pulumi Docs Generator. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->

Use this data source to query detailed information of monitor grafana_plugin_overviews

## Example Usage

<!--Start PulumiCodeChooser -->
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as tencentcloud from "@tencentcloud_iac/pulumi";

const grafanaPluginOverviews = tencentcloud.Monitor.getGrafanaPluginOverviews({});
```
```python
import pulumi
import pulumi_tencentcloud as tencentcloud

grafana_plugin_overviews = tencentcloud.Monitor.get_grafana_plugin_overviews()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Tencentcloud = Pulumi.Tencentcloud;

return await Deployment.RunAsync(() => 
{
    var grafanaPluginOverviews = Tencentcloud.Monitor.GetGrafanaPluginOverviews.Invoke();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := Monitor.GetGrafanaPluginOverviews(ctx, nil, nil)
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
import com.pulumi.tencentcloud.Monitor.MonitorFunctions;
import com.pulumi.tencentcloud.Monitor.inputs.GetGrafanaPluginOverviewsArgs;
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
        final var grafanaPluginOverviews = MonitorFunctions.getGrafanaPluginOverviews();

    }
}
```
```yaml
variables:
  grafanaPluginOverviews:
    fn::invoke:
      Function: tencentcloud:Monitor:getGrafanaPluginOverviews
      Arguments: {}
```
<!--End PulumiCodeChooser -->




## Using getGrafanaPluginOverviews {#using}

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
><span class="k">function </span>getGrafanaPluginOverviews<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetGrafanaPluginOverviewsArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Promise&lt;<span class="nx"><a href="#result">GetGrafanaPluginOverviewsResult</a></span>></span
><span class="k">
function </span>getGrafanaPluginOverviewsOutput<span class="p">(</span><span class="nx">args</span><span class="p">:</span> <span class="nx">GetGrafanaPluginOverviewsOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#InvokeOptions">InvokeOptions</a></span><span class="p">): Output&lt;<span class="nx"><a href="#result">GetGrafanaPluginOverviewsResult</a></span>></span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"
><span class="k">def </span>get_grafana_plugin_overviews<span class="p">(</span><span class="nx">result_output_file</span><span class="p">:</span> <span class="nx">Optional[str]</span> = None<span class="p">,</span>
                                 <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>GetGrafanaPluginOverviewsResult</span
><span class="k">
def </span>get_grafana_plugin_overviews_output<span class="p">(</span><span class="nx">result_output_file</span><span class="p">:</span> <span class="nx">Optional[pulumi.Input[str]]</span> = None<span class="p">,</span>
                                 <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.InvokeOptions">Optional[InvokeOptions]</a></span> = None<span class="p">) -&gt;</span> <span>Output[GetGrafanaPluginOverviewsResult]</span
></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"
><span class="k">func </span>GetGrafanaPluginOverviews<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetGrafanaPluginOverviewsArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) (*<span class="nx"><a href="#result">GetGrafanaPluginOverviewsResult</a></span>, error)</span
><span class="k">
func </span>GetGrafanaPluginOverviewsOutput<span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx">GetGrafanaPluginOverviewsOutputArgs</span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#InvokeOption">InvokeOption</a></span><span class="p">) GetGrafanaPluginOverviewsResultOutput</span
></code></pre></div>

&gt; Note: This function is named `GetGrafanaPluginOverviews` in the Go SDK.

</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public static class </span><span class="nx">GetGrafanaPluginOverviews </span><span class="p">
{</span><span class="k">
    public static </span>Task&lt;<span class="nx"><a href="#result">GetGrafanaPluginOverviewsResult</a></span>> <span class="p">InvokeAsync(</span><span class="nx">GetGrafanaPluginOverviewsArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="k">
    public static </span>Output&lt;<span class="nx"><a href="#result">GetGrafanaPluginOverviewsResult</a></span>> <span class="p">Invoke(</span><span class="nx">GetGrafanaPluginOverviewsInvokeArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.InvokeOptions.html">InvokeOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span><span class="p">
}</span></code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma"><code class="language-java" data-lang="java"><span class="k">public static CompletableFuture&lt;<span class="nx"><a href="#result">GetGrafanaPluginOverviewsResult</a></span>> </span>getGrafanaPluginOverviews<span class="p">(</span><span class="nx">GetGrafanaPluginOverviewsArgs</span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">InvokeOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
<span class="c">// Output-based functions aren't available in Java yet</span>
</code></pre></div>
</pulumi-choosable>
</div>


<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml"><span class="k">fn::invoke:</span>
<span class="k">&nbsp;&nbsp;function:</span> tencentcloud:Monitor/getGrafanaPluginOverviews:getGrafanaPluginOverviews
<span class="k">&nbsp;&nbsp;arguments:</span>
<span class="c">&nbsp;&nbsp;&nbsp;&nbsp;# arguments dictionary</span></code></pre></div>
</pulumi-choosable>
</div>



The following arguments are supported:


<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-optional"
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
<dl class="resources-properties"><dt class="property-optional"
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
<dl class="resources-properties"><dt class="property-optional"
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
<dl class="resources-properties"><dt class="property-optional"
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
<dl class="resources-properties"><dt class="property-optional"
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
<dl class="resources-properties"><dt class="property-optional"
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




## getGrafanaPluginOverviews Result {#result}

The following output properties are available:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
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
        <span id="pluginsets_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pluginsets_csharp" style="color: inherit; text-decoration: inherit;">Plugin<wbr>Sets</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#getgrafanapluginoverviewspluginset">List&lt;Tencent<wbr>Cloud<wbr>IAC.<wbr>Pulumi<wbr>Package.<wbr>Tencentcloud.<wbr>Monitor.<wbr>Outputs.<wbr>Get<wbr>Grafana<wbr>Plugin<wbr>Overviews<wbr>Plugin<wbr>Set&gt;</a></span>
    </dt>
    <dd><p>Plugin set.</p>
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
        <span id="id_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_go" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>The provider-assigned unique ID for this managed resource.</p>
</dd><dt class="property-"
            title="">
        <span id="pluginsets_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pluginsets_go" style="color: inherit; text-decoration: inherit;">Plugin<wbr>Sets</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#getgrafanapluginoverviewspluginset">[]Get<wbr>Grafana<wbr>Plugin<wbr>Overviews<wbr>Plugin<wbr>Set</a></span>
    </dt>
    <dd><p>Plugin set.</p>
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
        <span id="id_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_java" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>The provider-assigned unique ID for this managed resource.</p>
</dd><dt class="property-"
            title="">
        <span id="pluginsets_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pluginsets_java" style="color: inherit; text-decoration: inherit;">plugin<wbr>Sets</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#getgrafanapluginoverviewspluginset">List&lt;Get<wbr>Grafana<wbr>Plugin<wbr>Overviews<wbr>Plugin<wbr>Set&gt;</a></span>
    </dt>
    <dd><p>Plugin set.</p>
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
        <span id="id_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_nodejs" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>The provider-assigned unique ID for this managed resource.</p>
</dd><dt class="property-"
            title="">
        <span id="pluginsets_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pluginsets_nodejs" style="color: inherit; text-decoration: inherit;">plugin<wbr>Sets</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#getgrafanapluginoverviewspluginset">Get<wbr>Grafana<wbr>Plugin<wbr>Overviews<wbr>Plugin<wbr>Set[]</a></span>
    </dt>
    <dd><p>Plugin set.</p>
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
        <span id="id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_python" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd><p>The provider-assigned unique ID for this managed resource.</p>
</dd><dt class="property-"
            title="">
        <span id="plugin_sets_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#plugin_sets_python" style="color: inherit; text-decoration: inherit;">plugin_<wbr>sets</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#getgrafanapluginoverviewspluginset">Get<wbr>Grafana<wbr>Plugin<wbr>Overviews<wbr>Plugin<wbr>Set]</a></span>
    </dt>
    <dd><p>Plugin set.</p>
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
        <span id="id_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_yaml" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>The provider-assigned unique ID for this managed resource.</p>
</dd><dt class="property-"
            title="">
        <span id="pluginsets_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pluginsets_yaml" style="color: inherit; text-decoration: inherit;">plugin<wbr>Sets</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#getgrafanapluginoverviewspluginset">List&lt;Property Map&gt;</a></span>
    </dt>
    <dd><p>Plugin set.</p>
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


<h4 id="getgrafanapluginoverviewspluginset">Get<wbr>Grafana<wbr>Plugin<wbr>Overviews<wbr>Plugin<wbr>Set</h4>



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="pluginid_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pluginid_csharp" style="color: inherit; text-decoration: inherit;">Plugin<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Grafana plugin ID.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="version_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_csharp" style="color: inherit; text-decoration: inherit;">Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Grafana plugin version.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="pluginid_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pluginid_go" style="color: inherit; text-decoration: inherit;">Plugin<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Grafana plugin ID.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="version_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_go" style="color: inherit; text-decoration: inherit;">Version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Grafana plugin version.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="pluginid_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pluginid_java" style="color: inherit; text-decoration: inherit;">plugin<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>Grafana plugin ID.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="version_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_java" style="color: inherit; text-decoration: inherit;">version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>Grafana plugin version.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="pluginid_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pluginid_nodejs" style="color: inherit; text-decoration: inherit;">plugin<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Grafana plugin ID.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="version_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_nodejs" style="color: inherit; text-decoration: inherit;">version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd><p>Grafana plugin version.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="plugin_id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#plugin_id_python" style="color: inherit; text-decoration: inherit;">plugin_<wbr>id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd><p>Grafana plugin ID.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="version_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_python" style="color: inherit; text-decoration: inherit;">version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd><p>Grafana plugin version.</p>
</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-required"
            title="Required">
        <span id="pluginid_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#pluginid_yaml" style="color: inherit; text-decoration: inherit;">plugin<wbr>Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>Grafana plugin ID.</p>
</dd><dt class="property-required"
            title="Required">
        <span id="version_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#version_yaml" style="color: inherit; text-decoration: inherit;">version</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd><p>Grafana plugin version.</p>
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
