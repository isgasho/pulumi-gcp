// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the rollout state.
 *
 * https://cloud.google.com/game-servers/docs/reference/rest/v1beta/GameServerDeploymentRollout
 */
export function getGameServerDeploymentRollout(args: GetGameServerDeploymentRolloutArgs, opts?: pulumi.InvokeOptions): Promise<GetGameServerDeploymentRolloutResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("gcp:gameservices/getGameServerDeploymentRollout:getGameServerDeploymentRollout", {
        "deploymentId": args.deploymentId,
    }, opts);
}

/**
 * A collection of arguments for invoking getGameServerDeploymentRollout.
 */
export interface GetGameServerDeploymentRolloutArgs {
    /**
     * The deployment to get the rollout state from. Only 1 rollout must be associated with each deployment.
     */
    readonly deploymentId: string;
}

/**
 * A collection of values returned by getGameServerDeploymentRollout.
 */
export interface GetGameServerDeploymentRolloutResult {
    readonly defaultGameServerConfig: string;
    readonly deploymentId: string;
    readonly gameServerConfigOverrides: outputs.gameservices.GetGameServerDeploymentRolloutGameServerConfigOverride[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project: string;
}
