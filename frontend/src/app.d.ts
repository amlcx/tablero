// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			user: {
				id: string
				createdAt: Date
				updatedAt: Date
				email: string
				emailVerified: boolean
				name: string
				image?: string | null | undefined
				banned: boolean | null | undefined
				role?: string | null | undefined
				banReason?: string | null | undefined
				banExpires?: Date | null | undefined
				username?: string | null | undefined
				displayUsername?: string | null | undefined
			}
			session: {
				id: string
				userId: string
				expiresAt: Date
				createdAt: Date
				updatedAt: Date
				token: string
				ipAddress?: string | null | undefined
				userAgent?: string | null | undefined
			}
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
		namespace Superforms {
			type Message = {
				type: 'error' | 'success'
				text: string
			}
		}
	}
}

export {}
