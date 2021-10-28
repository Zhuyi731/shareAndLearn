import { assertEquals } from "https://deno.land/std@0.113.0/testing/asserts.ts";
import { delay } from "https://deno.land/std@0.113.0/async/delay.ts";

function add(a: number, b: number): number {
    return a + b
}
// 通过命令行 deno test xxx.ts 
Deno.test({
    name: '测试用例',
    fn() {
        const sum = add(1, 2)
        assertEquals(sum, 3)
    }
})

// 异步测试
Deno.test("async test", async () => {
    const x = add(3 ,12);

    // await some async task
    await delay(100);

    if (x !== 15) {
        throw Error("x should be equal to 15");
    }
});