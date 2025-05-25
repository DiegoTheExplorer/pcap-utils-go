

export const index = 2;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/2.ZHxSu5nW.js","_app/immutable/chunks/DMOJqsnG.js","_app/immutable/chunks/YKuBPo_k.js","_app/immutable/chunks/vAgtbmxX.js"];
export const stylesheets = ["_app/immutable/assets/2.Bq26kVWF.css"];
export const fonts = [];
