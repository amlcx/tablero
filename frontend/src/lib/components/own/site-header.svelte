<script lang="ts">
	import * as Popover from '$lib/components/ui/popover/index'
	import * as Tooltip from '$lib/components/ui/tooltip/index'
	import House from 'phosphor-svelte/lib/House'
	import Bell from 'phosphor-svelte/lib/Bell'
	import TagSimple from 'phosphor-svelte/lib/TagSimple'
</script>

<header
	class="sticky top-0 z-50 flex h-(--header-height) shrink-0 items-center gap-2 border-b bg-background py-2 transition-[width,height] ease-linear"
>
	<div class="relative flex w-full flex-row items-center justify-between">
		<div class="ml-1 flex flex-row items-center justify-center space-x-2.5 px-2">
			<!-- Left icons -->
			<Tooltip.Provider>
				<Tooltip.Root>
					<Tooltip.Trigger>
						<a href="/app">
							<House />
						</a>
					</Tooltip.Trigger>
					<Tooltip.Content class="rounded-sm">
						<p>Home</p>
					</Tooltip.Content>
				</Tooltip.Root>
			</Tooltip.Provider>

			<Tooltip.Provider>
				<Tooltip.Root>
					<Tooltip.Trigger>
						<a href="/app/categories">
							<TagSimple />
						</a>
					</Tooltip.Trigger>
					<Tooltip.Content class="rounded-sm">
						<p>Categories</p>
					</Tooltip.Content>
				</Tooltip.Root>
			</Tooltip.Provider>
		</div>

		<!-- Center content with absolute positioning -->
		<div class="absolute left-1/2 flex h-full -translate-x-1/2 items-center">
			<p class="font-mono leading-none">tablero</p>
		</div>

		<div class="mr-3 flex flex-col justify-center">
			<!-- Right icon -->
			<Popover.Root>
				<Popover.Trigger>
					<Bell />
				</Popover.Trigger>
				<Popover.Content class="rounded-sm bg-background">
					<span class="font-mono">Nothing to see here...</span>
				</Popover.Content>
			</Popover.Root>
		</div>
	</div>
</header>
