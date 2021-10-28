

export function Router() {
    const layers: any[] = []

    const route = async function (ctx: any, next: any) {
        layers.forEach(async (layer:any) => {
            console.log(ctx.req.method,layer)
            console.log(ctx.request.originalUrl,layer.path)

            if (ctx.req.method === layer.method && ctx.request.originalUrl === layer.path) {
                await layer.fn(ctx, next)
            }
        })
        await next()
    }
    route.get = function (p: string, l: any) {
        layers.push({ path: p, method: 'GET', fn: l })
    }
    return route
}