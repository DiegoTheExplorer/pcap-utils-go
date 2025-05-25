

export const index = 5;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/name_verification/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/5.hQBsNdqO.js","_app/immutable/chunks/DMOJqsnG.js","_app/immutable/chunks/YKuBPo_k.js","_app/immutable/chunks/vAgtbmxX.js"];
export const stylesheets = [];
export const fonts = [];
