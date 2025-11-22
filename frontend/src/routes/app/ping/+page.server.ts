import { BACKEND_URL } from '$env/static/private'
import { auth } from '$lib/auth'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ request }) => {
	const { token } = await auth.api.getToken({ headers: request.headers })

	const headers = new Headers()
	headers.set('Authorization', `Bearer: ${token}`)

	const resp = await fetch(`${BACKEND_URL}/ping`, { headers })
	const msg = await resp.json()

	return {
		msg
	}
}
