<script>
	let { data } = $props()
</script>

<p>Hello {data.user.username} from <span class="font-mono">/app</span></p>

<img src={data.user.image} alt={`${data.user.username}'s avatar`}>