<script lang="ts">
	let { children } = $props();
</script>

<nav>
	<ul>
		<li><a href="/">Home</a></li>
		<li><a href="/name_verification">Name Verification</a></li>
		<li><a href="/batch_img_convert">Batch Image Convert</a></li>
	</ul>
</nav>

{@render children()}

<style>
	nav {
		background-color: #1f1f28;
		padding: 0;
		margin: 0;
		height: 3em;
	}
	ul {
		list-style-type: none;
		margin: 0;
		padding: 0.5em 0.5em 0.5em 0em;
		overflow: hidden;
		li {
			float: left;
			a {
				display: block;
				text-align: center;
				padding: 0.5em;
				text-decoration: none;
				color: white;
			}
			a:hover {
				color: #b35b79;
			}
		}
	}
</style>
