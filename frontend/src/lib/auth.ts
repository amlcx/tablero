import { betterAuth } from 'better-auth'
import { admin, jwt, username } from 'better-auth/plugins'
import { sveltekitCookies } from 'better-auth/svelte-kit'
import { getRequestEvent } from '$app/server'
import { drizzleAdapter } from 'better-auth/adapters/drizzle'
import { db } from './server/db'

export const auth = betterAuth({
	database: drizzleAdapter(db, {
		provider: 'pg',
		usePlural: true
	}),

	emailAndPassword: { enabled: true, minPasswordLength: 8 },

	advanced: {
		database: {
			generateId: () => Bun.randomUUIDv7()
		}
	},

	plugins: [username(), admin(), jwt(), sveltekitCookies(getRequestEvent)]
})
