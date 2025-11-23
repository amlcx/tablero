import { auth } from '$lib/auth'
import { toSvelteKitHandler } from 'better-auth/svelte-kit'
import type { RequestHandler } from './$types'

// Original Better Auth handler
const handler = toSvelteKitHandler(auth)

export const GET: RequestHandler = async (event) => {
	// --- Logging input ---
	console.log('[JWKS] Incoming request:', {
		method: event.request.method,
		url: event.url.toString(),
		headers: Object.fromEntries(event.request.headers)
	})

	let response: Response
	try {
		response = await handler(event)
	} catch (err) {
		console.error('[JWKS] Handler error:', err)
		throw err
	}

	// --- Logging output ---
	console.log('[JWKS] Response:', {
		status: response.status,
		headers: Object.fromEntries(response.headers)
	})

	return response
}
