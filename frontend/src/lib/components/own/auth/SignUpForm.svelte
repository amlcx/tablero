<script lang="ts">
	// UI
	import { Button } from '$lib/components/ui/button/index'
	import * as Card from '$lib/components/ui/card/index'
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index'
	import * as Form from '$lib/components/ui/form/index'
	import * as ImageCropper from '$lib/components/ui/image-cropper/index'
	import { Input } from '$lib/components/ui/input/index'
	import * as Password from '$lib/components/ui/password/index'
	import Pencil from 'phosphor-svelte/lib/Pencil'

	// Functionality
	import { getFileFromUrl } from '$lib/components/ui/image-cropper/index'
	import { fileProxy, superForm, type Infer, type SuperValidated } from 'sveltekit-superforms'
	import { valibotClient } from 'sveltekit-superforms/adapters'
	import { signUpSchema, type SignUpSchema } from './schemas'
	import { toast } from 'svelte-sonner'
	import { goto } from '$app/navigation'

	let { data }: { data: { form: SuperValidated<Infer<SignUpSchema>> } } = $props()

	const form = superForm(data.form, {
		validators: valibotClient(signUpSchema),
		onUpdated({ form }) {
			if (form.message) {
				const { text } = form.message
				if (form.message.type === 'error') {
					toast.error(text)
				} else {
					toast.success(text)
                    goto('/app')
				}
			}
		}
	})

	const { form: formData, enhance } = form

	const file = fileProxy(form, 'avatar')

	let src = $state('')

	function onUnsupportedFile(f: File) {
		toast.error(`Unsupported file type: ${f.type}`)
	}

	async function onCropped(url: string) {
		src = url
		const f = await getFileFromUrl(url)
		file.set(f)
	}

    function clearSrc() {
        src = '';
    }
</script>

<Card.Root class="max-w-[350px] sm:max-w-[480px]">
	<Card.Header>
		<Card.Title class="title">Sign Up</Card.Title>
	</Card.Header>
	<Card.Content>
		<form
			action="?/signUp"
			method="post"
			enctype="multipart/form-data"
			class="flex flex-col gap-3"
			use:enhance
		>
			<Form.Field {form} name="email">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Email</Form.Label>
						<Input {...props} type="email" bind:value={$formData.email} />
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
				<Form.Description>Used for password recovery. Not visible to other users.</Form.Description>
			</Form.Field>

			<Form.Field {form} name="username">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Username</Form.Label>
						<Input {...props} type="text" bind:value={$formData.username} />
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
				<Form.Description
					>This is publicly visible, and is how other users will know you.</Form.Description
				>
			</Form.Field>

			<Form.Field {form} name="password">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Password</Form.Label>
						<Password.Root>
							<Password.Input {...props} bind:value={$formData.password}>
								<Password.ToggleVisibility />
							</Password.Input>
						</Password.Root>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
				<Form.Description>At least 8 characters in length.</Form.Description>
			</Form.Field>

			<Form.Field {form} name="confirmPassword">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Confirm Password</Form.Label>
						<Password.Root>
							<Password.Input {...props} bind:value={$formData.confirmPassword}>
								<Password.ToggleVisibility />
							</Password.Input>
						</Password.Root>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
				<Form.Description
					>Did we already mention at least 8 characters? Huh, please make sure they match.</Form.Description
				>
			</Form.Field>

			<Form.Field {form} name="avatar">
				<Form.Control>
					{#snippet children({ props })}
						<Form.Label>Avatar</Form.Label>
						<ImageCropper.Root bind:src {onUnsupportedFile} {onCropped}>
							<div class="relative">
								<ImageCropper.Preview />
								<DropdownMenu.Root>
									<DropdownMenu.Trigger>
										{#snippet child({ props })}
											<Button
												{...props}
												class="absolute -bottom-3 -left-3 rounded-full"
												variant="outline"
												size="icon"
											>
												<Pencil />
											</Button>
										{/snippet}
									</DropdownMenu.Trigger>
									<DropdownMenu.Content align="start">
										<ImageCropper.UploadTrigger>
											<DropdownMenu.Item>Upload a photo...</DropdownMenu.Item>
										</ImageCropper.UploadTrigger>
										<DropdownMenu.Item onclick={() => clearSrc()}>Remove photo</DropdownMenu.Item>
									</DropdownMenu.Content>
								</DropdownMenu.Root>
							</div>
							<ImageCropper.Dialog>
								<ImageCropper.Cropper />
								<ImageCropper.Controls>
									<ImageCropper.Cancel />
									<ImageCropper.Crop />
								</ImageCropper.Controls>
							</ImageCropper.Dialog>
						</ImageCropper.Root>
						<input type="file" {...props} hidden bind:files={$file}>
					{/snippet}
				</Form.Control>
				<Form.FieldErrors />
				<Form.Description>Keep it safe.</Form.Description>
			</Form.Field>

            <Form.Button>Let's go</Form.Button>
		</form>
	</Card.Content>
</Card.Root>
