<script lang="ts">
	import type { PageData } from "./$types"

    let { pageTitle }: PageData = $props();
</script>

<svelte:head>
    <title>{pageTitle}</title>
</svelte:head>

<h1 class="title">Welcome to tablero</h1>

<a href="/auth/sign-up">Sign up</a>
