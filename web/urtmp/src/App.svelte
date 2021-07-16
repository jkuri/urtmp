<script>
  import mpegts from 'mpegts.js';

	export let name;

  fetch('/api/v1/streams')
    .then(res => res.json)
    .then(res => {
      console.log(res);
    })
</script>

<main>
	<h1>Hello!</h1>
  <div class="video">

  </div>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
