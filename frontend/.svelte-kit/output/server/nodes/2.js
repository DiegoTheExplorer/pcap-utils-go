

export const index = 2;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/2.C1uqBupw.js","_app/immutable/chunks/B0FDQ0ax.js","_app/immutable/chunks/DQ0hcwHw.js","_app/immutable/chunks/Cd-0_WxI.js"];
export const stylesheets = ["_app/immutable/assets/2.Bq26kVWF.css"];
export const fonts = [];
