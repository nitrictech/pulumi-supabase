// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("supabase");

/**
 * Supbase Personal Access Token for account
 */
export declare const token: string;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token") ?? (utilities.getEnv("SUPABASE_ACCESS_TOKEN") || "");
    },
    enumerable: true,
});

export declare const version: string | undefined;
Object.defineProperty(exports, "version", {
    get() {
        return __config.get("version");
    },
    enumerable: true,
});

