import { d as attr_style, c as pop, p as push } from "../../../chunks/index.js";
function _page($$payload, $$props) {
  push();
  let prog_bar_len = "";
  $$payload.out += `<div class="outer_container svelte-9dkli4"><h2>Batch Image Convert</h2> <div class="input_fields"><button>Select Input Folder</button> <h4>Current Input Folder:</h4> <button>Select Output Folder</button> <h4>Current Output Folder:</h4></div> <div class="status"><h4>Convertion Status:</h4> <div id="prog"><div id="prog_bar_lbl" class="svelte-9dkli4">Progress:</div> <div id="prog_border" class="svelte-9dkli4"><div id="prog_bar" class="svelte-9dkli4"${attr_style("", { width: prog_bar_len })}>0%</div></div></div></div> <div class="errors"></div></div>`;
  pop();
}
export {
  _page as default
};
