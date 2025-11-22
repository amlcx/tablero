import { auth } from '$lib/auth'
import { createApiClients } from '$lib/server/api/clients'
import { error } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import { ConnectError } from '@connectrpc/connect'
import { normalizeConnectErrorCode } from '$lib/utils/statuses'

export const load: PageServerLoad = async ({ request }) => {
	try {
		const { token } = await auth.api.getToken({ headers: request.headers })

		const headers = new Headers()

		headers.set('Authorization', `Bearer: ${token}`)

		const { greetApi } = createApiClients()

		const resp = await greetApi.greet(
			{
				greet: {
					id: Bun.randomUUIDv7(),
					name: 'Frontend'
				}
			},
			{ headers: headers }
		)

		const { msg } = resp

		return { msg }
	} catch (e) {
		console.log(e)

		if (e instanceof ConnectError) {
			const code = normalizeConnectErrorCode(e.code)
			error(code, e.message)
		}

		error(500, 'Unknown error')
	}
}
