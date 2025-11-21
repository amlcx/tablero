import { betterAuth } from 'better-auth'
import { db } from './server/database/db'
import { admin, jwt, username } from 'better-auth/plugins'
import { sveltekitCookies } from 'better-auth/svelte-kit'
import { getRequestEvent } from '$app/server'

export const auth = betterAuth({
	database: { db: db, type: 'postgres' },

	emailAndPassword: { enabled: true, minPasswordLength: 8 },

	advanced: {
		database: {
			generateId: () => Bun.randomUUIDv7()
		}
	},

	plugins: [username(), admin(), jwt(), sveltekitCookies(getRequestEvent)]
})
