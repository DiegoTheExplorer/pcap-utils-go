<script lang="ts">
	// Reactive page element values
	import { Get_dir_path } from '../../lib/wailsjs/go/main/App.js';
	let convertion_percentage = $state(0);
	let prog_bar_len = $state('');
	let inp_dir = $state('');
	let out_dir = $state('');

	async function update_dir_path(is_inp_dir: boolean) {
		if (is_inp_dir) {
			inp_dir = await Get_dir_path('Select Input Folder');
			console.log(inp_dir);
		} else {
			out_dir = await Get_dir_path('Select Output Folder');
		}
	}
</script>

<div class="outer_container">
	<h2>Batch Image Convert</h2>
	<div class="input_fields">
		<button onclick={() => update_dir_path(true)}> Select Input Folder </button>
		<h4>Current Input Folder: {inp_dir}</h4>
		<button onclick={() => update_dir_path(false)}> Select Output Folder</button>
		<h4>Current Output Folder: {out_dir}</h4>
	</div>

	<div class="status">
		<h4>Convertion Status:</h4>
		<div id="prog">
			<div id="prog_bar_lbl">Progress:</div>
			<div id="prog_border">
				<div id="prog_bar" style:width={prog_bar_len}>{convertion_percentage}%</div>
			</div>
		</div>
	</div>

	<div class="errors"></div>
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
