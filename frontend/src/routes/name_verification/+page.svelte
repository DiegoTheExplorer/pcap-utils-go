<script lang="ts">
	import {
		Get_file_path,
		Read_verified_names,
		Read_gforms_names,
		Verify_names
	} from '$lib/wailsjs/go/main/App.js';

	import { name_verification } from '$lib/wailsjs/go/models';

	let verified_names_fp = $state('');
	let verified_names_list = $state(['']);
	let gforms_names_fp = $state('');
	let gforms_names_list: name_verification.CrimName[] = $state([]);
	let for_man_review_list: name_verification.ForManualReview[] = $state([]);

	async function getFilePath(out_target: number) {
		const verified_names = 0;
		const gforms_names = 1;
		let file_path = await Get_file_path('Select a file');
		if (file_path == '') return;

		switch (out_target) {
			case verified_names:
				verified_names_fp = file_path;
				Read_verified_names(file_path, 100)
					.then((res) => {
						verified_names_list = res;
					})
					.catch((err) => {
						// TODO: Add gui display for this error
						console.log(err);
					});
				break;
			case gforms_names:
				gforms_names_fp = file_path;
				Read_gforms_names(file_path, 100)
					.then((res) => {
						gforms_names_list = res;
						// console.log(res);
					})
					.catch((err) => {
						// TODO: Add gui display for this error
						console.log(err);
					});
				break;
			default:
				return;
		}
	}

	async function verifyNames() {
		Verify_names(verified_names_fp, gforms_names_fp, '')
			.then((res) => {
				for_man_review_list = res;
				console.log(res);
			})
			.catch((err) => {
				console.log(err);
			});
	}
</script>

<h2>Name Verification</h2>
<div class="sub_app_con">
	<div class="verified_names_con input_con">
		<h3>Verified names selection</h3>
		<button onclick={() => getFilePath(0)}> Choose input text file </button>
		<p>Chosen file name: {verified_names_fp}</p>
	</div>
	<div class="gforms_names_con input_con">
		<h3>Gfroms names selection</h3>
		<button onclick={() => getFilePath(1)}> Choose input text file </button>
		<p>Chosen file name: {gforms_names_fp}</p>
	</div>
	<div class="for_manual_rev_con input_con">
		<h3>For manual review file</h3>
		<label for="output_filename">Output file name: </label>
		<input type="text" id="output_filename" /><br />
		<button onclick={() => verifyNames()}>Compare names</button>
	</div>
	<div class="verified_names_preview preview_con">
		<h3>Verified names preview</h3>
		<div class="preview">
			<ul>
				{#each verified_names_list as name}
					<li>{name}</li>
				{/each}
			</ul>
		</div>
	</div>
	<div class="gforms_names_preview preview_con">
		<h3>Gforms names preview</h3>
		<div class="preview">
			<!-- TODO: separate the table into its own component -->
			<table>
				<thead>
					<tr>
						<td>Last Name</td>
						<td>First Name</td>
						<td>Middle Name</td>
					</tr>
				</thead>
				<tbody>
					{#each gforms_names_list as crim_name}
						<tr>
							<td>{crim_name.lname}</td>
							<td>{crim_name.fname}</td>
							<td>{crim_name.mname}</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	</div>
	<div class="for_manual_rev preview_con">
		<h3>For manual review preview</h3>
		<div class="preview">
			<ul>
				{#each for_man_review_list as res}
					<li>
						Submitted name: {res.submitted_name}<br />

						{#if res.matches == null}
							^ No Matches Found ^
						{:else}
							Matches:
							<ul>
								{#each res.matches as match}
									<li>{match.Target}</li>
								{/each}
							</ul>
						{/if}
					</li>
				{/each}
			</ul>
		</div>
	</div>
</div>

<style>
	.sub_app_con {
		height: 100%;
		grid-template-columns: auto auto auto;
		display: grid;
		border-style: solid;
		border-color: blue;
	}

	.input_con {
		border-style: solid;
		border-color: green;
	}

	.preview_con {
		border-style: solid;
		border-color: red;
	}

	.preview {
		max-height: 300px;
		overflow-y: scroll;
	}

	thead {
		font-weight: bold;
	}

	table,
	td,
	tr {
		border: 2px solid black;
	}
</style>
