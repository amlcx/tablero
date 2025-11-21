import { redirect, type Handle } from '@sveltejs/kit'
import { auth } from '$lib/auth'
import { svelteKitHandler } from 'better-auth/svelte-kit'
import { building } from '$app/environment'
import { sequence } from '@sveltejs/kit/hooks'

const betterAuth: Handle = async ({ event, resolve }) => {
	const session = await auth.api.getSession({
		headers: event.request.headers
	})

	if (session) {
		event.locals.user = session.user
		event.locals.session = session.session
	}

	return svelteKitHandler({ event, resolve, auth, building })
}

const authGuard: Handle = ({ event, resolve }) => {
	const pathname = event.url.pathname
	const session = event.locals.session
	const user = event.locals.user

	const hasSession = session !== undefined
	const hasUser = user !== undefined

	if (!hasSession && pathname.startsWith('/app')) {
		throw redirect(303, '/auth')
	}

	if (hasSession && pathname.startsWith('/auth') && event.url.search !== '?/signOut') {
		throw redirect(303, '/app')
	}

	if (hasSession && (pathname === '' || pathname === '/')) {
		throw redirect(303, '/app')
	}

	if (hasUser && user.role !== 'admin' && pathname.startsWith('/app/admin')) {
		console.log('access denied to', user.username)
		throw redirect(303, '/app')
	}

	return resolve(event)
}

export const handle: Handle = sequence(betterAuth, authGuard)
