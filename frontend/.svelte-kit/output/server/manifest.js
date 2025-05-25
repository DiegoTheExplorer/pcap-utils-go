export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png"]),
	mimeTypes: {".png":"image/png"},
	_: {
		client: {start:"_app/immutable/entry/start.CEIJDqrQ.js",app:"_app/immutable/entry/app.DOd91kOv.js",imports:["_app/immutable/entry/start.CEIJDqrQ.js","_app/immutable/chunks/DFnvpG1K.js","_app/immutable/chunks/BbAI_zX7.js","_app/immutable/chunks/DIeogL5L.js","_app/immutable/chunks/CLbpfrfV.js","_app/immutable/entry/app.DOd91kOv.js","_app/immutable/chunks/BbAI_zX7.js","_app/immutable/chunks/DIeogL5L.js","_app/immutable/chunks/BJFeLwTw.js","_app/immutable/chunks/zKezSjUH.js","_app/immutable/chunks/CWj6FrbW.js","_app/immutable/chunks/CLbpfrfV.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js'))
		],
		routes: [
			
		],
		prerendered_routes: new Set(["/"]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();
