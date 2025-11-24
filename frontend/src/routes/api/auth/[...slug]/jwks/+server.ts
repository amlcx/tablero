import { auth } from '$lib/auth'
import { toSvelteKitHandler } from 'better-auth/svelte-kit'

// Original Better Auth handler
const handler = toSvelteKitHandler(auth)

interface WithRequest {
	request: Request
}

export const fallback = (request: WithRequest) => {
	console.log(`[JWKS endpoint]`)
	return handler(request)
}
