import { Router } from '../router.ts'

const router = Router()

router.get('/user', async (ctx:any) => {
    ctx.body = "/user 请求"
})

export { router }