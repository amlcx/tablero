<script>
	import SiteHeader from '$lib/components/own/site-header.svelte'

    let { children } = $props();
</script>

<SiteHeader />

{@render children()}