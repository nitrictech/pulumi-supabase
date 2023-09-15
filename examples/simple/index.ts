import * as supabase from "@pulumi/supabase";

const random = new supabase.Random("my-random", { length: 24 });

export const output = random.result;