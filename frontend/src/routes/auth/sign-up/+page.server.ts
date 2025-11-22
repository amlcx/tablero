import { fail, message, setError, superValidate } from 'sveltekit-superforms'
import type { Actions, PageServerLoad } from './$types'
import { valibot } from 'sveltekit-superforms/adapters'
import { signUpSchema } from '$lib/components/own/auth/schemas'
import { auth } from '$lib/auth'
import { saveAvatarToDisk } from '$lib/utils/avatar'
import { APIError } from 'better-auth'
import { normalizeErrorStatus } from '$lib/utils/statuses'

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(valibot(signUpSchema))
	}
}

export const actions = {
	signUp: async ({ request }) => {
		try {
			console.debug('server action triggered')
			const form = await superValidate(request, valibot(signUpSchema))

			if (!form.valid) {
				return fail(400, { form })
			}

			const { data } = form

			const usernameResponse = await auth.api.isUsernameAvailable({
				body: { username: data.username }
			})

			if (!usernameResponse.available) {
				return setError(form, 'username', 'This username is already taken')
			}

			let image: string | undefined = undefined

			if (data.avatar !== undefined) {
				image = await saveAvatarToDisk(data.avatar, data.username)
			}

			console.log('server action has generated avatar')

			const signUpResponse = await auth.api.signUpEmail({
				body: {
					email: data.email,
					name: data.username,
					password: data.password,
					username: data.username,
					image: image
				}
			})

			console.log('server action completed')

			return message(form, {
				type: 'success',
				text: `Hi, ${signUpResponse.user.name}. You've successfully registered.`
			})
		} catch (e) {
			console.error(e)
			const form = await superValidate(request, valibot(signUpSchema))

			if (e instanceof APIError) {
				return message(
					form,
					{ type: 'error', text: e.message },
					{ status: normalizeErrorStatus(e.statusCode) }
				)
			}

			return message(form, { type: 'error', text: 'Unexpected error' }, { status: 500 })
		}
	}
} satisfies Actions
