<script lang="ts">
	// Reactive page element values
	let convertion_percentage = $state(0);
	let prog_bar_len = $state('');
</script>

<div class="outer_container">
	<h2>Batch Image Convert</h2>
	<div class="input_fields">
		<h4>Current Input Folder:</h4>
		<button> Select Input Folder </button>
		<h4>Current Output Folder:</h4>
		<button> Select Output Folder </button>
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
