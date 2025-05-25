import "clsx";
function _layout($$payload, $$props) {
  let { children } = $$props;
  $$payload.out += `<nav class="svelte-1lyq3l9"><ul class="svelte-1lyq3l9"><li class="svelte-1lyq3l9"><a href="/" class="svelte-1lyq3l9">Home</a></li> <li class="svelte-1lyq3l9"><a href="/name_verification" class="svelte-1lyq3l9">Name Verification</a></li> <li class="svelte-1lyq3l9"><a href="/batch_img_convert" class="svelte-1lyq3l9">Batch Image Convert</a></li></ul></nav> `;
  children($$payload);
  $$payload.out += `<!---->`;
}
export {
  _layout as default
};
