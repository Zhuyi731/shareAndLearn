import { App } from "../../../deno-koa/mod.ts";

const app = new App();

app.use(async function (ctx: any) {
  ctx.status = 200;
  ctx.body = "Hello World";
});

app.listen({ port: 8000 });
console.log('App is listening @localhost:8000')