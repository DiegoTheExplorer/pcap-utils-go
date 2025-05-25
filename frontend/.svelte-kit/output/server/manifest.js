export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png","images/chowchow.webp","images/chowchow2.jpg","images/chowchow_nobg.png"]),
	mimeTypes: {".png":"image/png",".webp":"image/webp",".jpg":"image/jpeg"},
	_: {
		client: {start:"_app/immutable/entry/start.DzSRqSTB.js",app:"_app/immutable/entry/app.BI3Aq_2B.js",imports:["_app/immutable/entry/start.DzSRqSTB.js","_app/immutable/chunks/DTj9pe6C.js","_app/immutable/chunks/YKuBPo_k.js","_app/immutable/chunks/FUM7f750.js","_app/immutable/entry/app.BI3Aq_2B.js","_app/immutable/chunks/YKuBPo_k.js","_app/immutable/chunks/B4fhUwB3.js","_app/immutable/chunks/DMOJqsnG.js","_app/immutable/chunks/FUM7f750.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js'))
		],
		routes: [
			
		],
		prerendered_routes: new Set(["/","/about","/batch_img_convert","/name_verification"]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();
