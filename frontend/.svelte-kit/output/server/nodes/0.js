

export const index = 0;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/fallbacks/layout.svelte.js')).default;
export const universal = {
  "prerender": true,
  "ssr": false
};
export const universal_id = "src/routes/+layout.ts";
export const imports = ["_app/immutable/nodes/0.m5QRi5Pk.js","_app/immutable/chunks/CWj6FrbW.js","_app/immutable/chunks/BbAI_zX7.js","_app/immutable/chunks/DIeogL5L.js","_app/immutable/chunks/zKezSjUH.js"];
export const stylesheets = [];
export const fonts = [];
