import * as path from "https://deno.land/std/path/mod.ts"
import __ from 'https://deno.land/x/dirname/mod.ts';

const { __dirname } = __(import.meta);
export function readFile(): string {
    return String(Deno.readFileSync(path.join(__dirname, './content.txt')))
}