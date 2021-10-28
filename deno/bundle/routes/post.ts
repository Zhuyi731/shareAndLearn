import { Router } from '../router.ts'

const router = Router()

router.get('/post', async (ctx: any) => {
    ctx.status = 200
    ctx.body = "/post 请求"
})

export { router }