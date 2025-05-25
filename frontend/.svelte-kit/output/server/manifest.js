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
		client: {start:"_app/immutable/entry/start.DhlCHHe4.js",app:"_app/immutable/entry/app.C8siUw5h.js",imports:["_app/immutable/entry/start.DhlCHHe4.js","_app/immutable/chunks/CJ_UaEGy.js","_app/immutable/chunks/DQ0hcwHw.js","_app/immutable/chunks/CDfKoPyR.js","_app/immutable/entry/app.C8siUw5h.js","_app/immutable/chunks/DQ0hcwHw.js","_app/immutable/chunks/D-gfomyo.js","_app/immutable/chunks/D5ZE-k3H.js","_app/immutable/chunks/B0FDQ0ax.js","_app/immutable/chunks/CDfKoPyR.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
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
