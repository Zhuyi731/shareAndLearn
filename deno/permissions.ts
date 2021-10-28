// 运行时询问权限
const desc1 = { name: "read", path: "/foo" } as const;
const status1 = await Deno.permissions.request(desc1);
console.log(status1);
// 查询
console.log(await Deno.permissions.query(desc1));

console.log(await Deno.permissions.revoke(desc1));

console.log(await Deno.permissions.query(desc1));

// deno run --allow-read=/foo permissions.ts
// const desc = { name: "read", path: "/foo/bar" } as const;
// console.log(await Deno.permissions.revoke(desc));
// PermissionStatus { state: "granted" }