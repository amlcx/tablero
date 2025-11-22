import { BACKEND_URL } from '$env/static/private'
import { auth } from '$lib/auth'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ request }) => {
	try {
		const { token } = await auth.api.getToken({ headers: request.headers })
		console.log(token)

		const headers = new Headers()
		headers.set('Authorization', `Bearer: ${token}`)

		const resp = await fetch(`${BACKEND_URL}/ping`, { headers })
		const msg = await resp.json()

		console.log('MSG', msg)

		return {
			msg
		}
	} catch (e) {
		console.error(e)
	}
}
