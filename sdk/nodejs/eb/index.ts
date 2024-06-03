// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { EventBusArgs, EventBusState } from "./eventBus";
export type EventBus = import("./eventBus").EventBus;
export const EventBus: typeof import("./eventBus").EventBus = null as any;
utilities.lazyLoad(exports, ["EventBus"], () => require("./eventBus"));

export { EventConnectorArgs, EventConnectorState } from "./eventConnector";
export type EventConnector = import("./eventConnector").EventConnector;
export const EventConnector: typeof import("./eventConnector").EventConnector = null as any;
utilities.lazyLoad(exports, ["EventConnector"], () => require("./eventConnector"));

export { EventRuleArgs, EventRuleState } from "./eventRule";
export type EventRule = import("./eventRule").EventRule;
export const EventRule: typeof import("./eventRule").EventRule = null as any;
utilities.lazyLoad(exports, ["EventRule"], () => require("./eventRule"));

export { EventTargetArgs, EventTargetState } from "./eventTarget";
export type EventTarget = import("./eventTarget").EventTarget;
export const EventTarget: typeof import("./eventTarget").EventTarget = null as any;
utilities.lazyLoad(exports, ["EventTarget"], () => require("./eventTarget"));

export { EventTransformArgs, EventTransformState } from "./eventTransform";
export type EventTransform = import("./eventTransform").EventTransform;
export const EventTransform: typeof import("./eventTransform").EventTransform = null as any;
utilities.lazyLoad(exports, ["EventTransform"], () => require("./eventTransform"));

export { GetBusArgs, GetBusResult, GetBusOutputArgs } from "./getBus";
export const getBus: typeof import("./getBus").getBus = null as any;
export const getBusOutput: typeof import("./getBus").getBusOutput = null as any;
utilities.lazyLoad(exports, ["getBus","getBusOutput"], () => require("./getBus"));

export { GetEventRulesArgs, GetEventRulesResult, GetEventRulesOutputArgs } from "./getEventRules";
export const getEventRules: typeof import("./getEventRules").getEventRules = null as any;
export const getEventRulesOutput: typeof import("./getEventRules").getEventRulesOutput = null as any;
utilities.lazyLoad(exports, ["getEventRules","getEventRulesOutput"], () => require("./getEventRules"));

export { GetPlateformEventTemplateArgs, GetPlateformEventTemplateResult, GetPlateformEventTemplateOutputArgs } from "./getPlateformEventTemplate";
export const getPlateformEventTemplate: typeof import("./getPlateformEventTemplate").getPlateformEventTemplate = null as any;
export const getPlateformEventTemplateOutput: typeof import("./getPlateformEventTemplate").getPlateformEventTemplateOutput = null as any;
utilities.lazyLoad(exports, ["getPlateformEventTemplate","getPlateformEventTemplateOutput"], () => require("./getPlateformEventTemplate"));

export { GetPlatformEventNamesArgs, GetPlatformEventNamesResult, GetPlatformEventNamesOutputArgs } from "./getPlatformEventNames";
export const getPlatformEventNames: typeof import("./getPlatformEventNames").getPlatformEventNames = null as any;
export const getPlatformEventNamesOutput: typeof import("./getPlatformEventNames").getPlatformEventNamesOutput = null as any;
utilities.lazyLoad(exports, ["getPlatformEventNames","getPlatformEventNamesOutput"], () => require("./getPlatformEventNames"));

export { GetPlatformEventPatternsArgs, GetPlatformEventPatternsResult, GetPlatformEventPatternsOutputArgs } from "./getPlatformEventPatterns";
export const getPlatformEventPatterns: typeof import("./getPlatformEventPatterns").getPlatformEventPatterns = null as any;
export const getPlatformEventPatternsOutput: typeof import("./getPlatformEventPatterns").getPlatformEventPatternsOutput = null as any;
utilities.lazyLoad(exports, ["getPlatformEventPatterns","getPlatformEventPatternsOutput"], () => require("./getPlatformEventPatterns"));

export { GetPlatformProductsArgs, GetPlatformProductsResult, GetPlatformProductsOutputArgs } from "./getPlatformProducts";
export const getPlatformProducts: typeof import("./getPlatformProducts").getPlatformProducts = null as any;
export const getPlatformProductsOutput: typeof import("./getPlatformProducts").getPlatformProductsOutput = null as any;
utilities.lazyLoad(exports, ["getPlatformProducts","getPlatformProductsOutput"], () => require("./getPlatformProducts"));

export { GetSearchArgs, GetSearchResult, GetSearchOutputArgs } from "./getSearch";
export const getSearch: typeof import("./getSearch").getSearch = null as any;
export const getSearchOutput: typeof import("./getSearch").getSearchOutput = null as any;
utilities.lazyLoad(exports, ["getSearch","getSearchOutput"], () => require("./getSearch"));

export { PutEventsArgs, PutEventsState } from "./putEvents";
export type PutEvents = import("./putEvents").PutEvents;
export const PutEvents: typeof import("./putEvents").PutEvents = null as any;
utilities.lazyLoad(exports, ["PutEvents"], () => require("./putEvents"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Eb/eventBus:EventBus":
                return new EventBus(name, <any>undefined, { urn })
            case "tencentcloud:Eb/eventConnector:EventConnector":
                return new EventConnector(name, <any>undefined, { urn })
            case "tencentcloud:Eb/eventRule:EventRule":
                return new EventRule(name, <any>undefined, { urn })
            case "tencentcloud:Eb/eventTarget:EventTarget":
                return new EventTarget(name, <any>undefined, { urn })
            case "tencentcloud:Eb/eventTransform:EventTransform":
                return new EventTransform(name, <any>undefined, { urn })
            case "tencentcloud:Eb/putEvents:PutEvents":
                return new PutEvents(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Eb/eventBus", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Eb/eventConnector", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Eb/eventRule", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Eb/eventTarget", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Eb/eventTransform", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Eb/putEvents", _module)
