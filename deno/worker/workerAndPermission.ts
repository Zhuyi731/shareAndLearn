import * as path from "https://deno.land/std/path/mod.ts"
import __ from 'https://deno.land/x/dirname/mod.ts';

const { __dirname } = __(import.meta);

const worker = new Worker(new URL("./worker.ts", import.meta.url).href, {
    type: "module",
    deno: {
        namespace: true,
        permissions: {
            read: true,
            // read: [
            //     new URL("./content.txt", import.meta.url),
            // ]
        },
    },
});

worker.postMessage({ filename: path.join(__dirname, "./content.txt") });