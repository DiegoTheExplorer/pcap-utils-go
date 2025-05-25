

export const index = 3;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/about/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/3.erxTgYix.js","_app/immutable/chunks/B0FDQ0ax.js","_app/immutable/chunks/DQ0hcwHw.js","_app/immutable/chunks/Cd-0_WxI.js"];
export const stylesheets = [];
export const fonts = [];
