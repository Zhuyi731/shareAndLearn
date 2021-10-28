import { App } from "../../../../deno-koa/mod.ts";
import { router as post } from './routes/post.ts'
import { router as user } from './routes/user.ts'
const app = new App();

app.use(post)
app.use(user)

app.listen({ port: 8000 });
console.log('App is listening @localhost:8000')