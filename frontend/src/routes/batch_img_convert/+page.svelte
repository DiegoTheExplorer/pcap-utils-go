<script lang="ts">
	import { beforeNavigate } from '$app/navigation';
	import { Get_dir_path } from '$lib/wailsjs/go/main/App.js';
	import { Batch_img_conv } from '$lib/wailsjs/go/main/App.js';
	import { EventsOn } from '$lib/wailsjs/runtime/runtime';

	// Reactive page element values
	let conversion_percentage = $state(0);
	let prog_bar_len = $derived.by(() => {
		let num_px = conversion_percentage * 200;
		let width_str = num_px.toString() + 'px';
		return width_str;
	});
	let inp_dir = $state('');
	let out_dir = $state('');
	let disable_conv_btn = $state(false);
	let missing_input = $derived.by(() => {
		if (out_dir == '' || inp_dir == '') {
			return true;
		}
		return false;
	});
	let disable_routing = $state(false);
	let unconverted_files = $state(['']);

	beforeNavigate(({ cancel }) => {
		// calling cancel will prevent routing
		if (disable_routing) {
			cancel();
		}
	});

	// Event Handlers
	EventsOn('conversion-progress', (prog) => {
		conversion_percentage = prog >= 1 ? 1 : prog * 0.1;
	});

	async function update_dir_path(is_inp_dir: boolean) {
		if (is_inp_dir) {
			inp_dir = await Get_dir_path('Select Input Folder');
		} else {
			out_dir = await Get_dir_path('Select Output Folder');
		}
	}

	async function batch_img_conv() {
		disable_conv_btn = true;
		disable_routing = true;
		Batch_img_conv(inp_dir, out_dir)
			.then((res) => {
				unconverted_files = res;
				disable_conv_btn = false;
				disable_routing = false;
			})
			.catch((e) => {
				// TODO: add error display to gui and logging
				console.log(e);
				disable_conv_btn = false;
				disable_routing = false;
			});
	}
</script>

<div class="outer_container">
	<h2>Batch Image Convert</h2>
	<div class="input_fields">
		<button onclick={() => update_dir_path(true)}> Select Input Folder </button>
		<h4>Current Input Folder: {inp_dir}</h4>
		<button onclick={() => update_dir_path(false)}> Select Output Folder</button>
		<h4>Current Output Folder: {out_dir}</h4>
		<button disabled={missing_input || disable_conv_btn} onclick={() => batch_img_conv()}>
			Start Conversion
		</button>
	</div>

	<div class="status">
		<h4>Conversion Status:</h4>
		<div id="prog">
			<div id="prog_bar_lbl">Progress:</div>
			<div id="prog_border">
				<div id="prog_bar" style:width={prog_bar_len}>{conversion_percentage * 100}%</div>
			</div>
		</div>
	</div>

	<div class="errors">
		<h3>Unconverted files or sub directories</h3>
		<ul>
			{#each unconverted_files as fname}
				<li>{fname}</li>
			{/each}
		</ul>
	</div>
</div>

<style>
	.outer_container {
		margin-left: 10px;
	}
	#prog_bar_lbl {
		font-size: 1.5em;
		color: black;
		padding-right: 10px;
	}

	#prog_border {
		border-style: solid;
		border-radius: 5px;
		width: 200px;
		height: 25px;
	}

	#prog_bar {
		width: 0px;
		height: 25px;
		background-color: #6f894e;
		border-radius: 5x;
		display: block;
		text-align: center;
		color: lightcoral;
	}
</style>
