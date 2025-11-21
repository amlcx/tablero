<script>
	import SignUpForm from '$lib/components/own/auth/SignUpForm.svelte'

    let { data } = $props();
</script>

<SignUpForm {data} />