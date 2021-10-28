import {
    importw,
    release,
    worker,
} from "https://deno.land/x/importw@1.1.0/mod.ts";

// Import module from within a worker.
const { log, add, [release]: terminate, [worker]: workerRef } = await importw(
    "https://deno.land/x/importw@1.1.0/examples/basic/exampleMod.ts",
    {
        name: "exampleWorker",
        deno: false,
    },
);

// Have access to the underlying Worker
console.log(workerRef.constructor.name); // Worker

// Run code within the Worker
await log(`add(40, 2) in a worker:`, await add(40, 2));

// Gracefully release Worker resources
await terminate();