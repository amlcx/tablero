import * as v from 'valibot'

export const MB = 1024 * 1024

export const signUpSchema = v.pipe(
	v.object({
		email: v.pipe(v.string(), v.trim(), v.email(), v.maxLength(64)),
		username: v.pipe(v.string(), v.trim(), v.minLength(3), v.maxLength(32)),
		password: v.pipe(v.string(), v.minLength(8)),
		confirmPassword: v.string(),
		avatar: v.optional(
			v.pipe(v.file(), v.mimeType(['image/jpeg', 'image/png', 'image/webp']), v.maxSize(1 * MB))
		)
	}),
	v.forward(
		v.partialCheck(
			[['password'], ['confirmPassword']],
			(i) => i.password === i.confirmPassword,
			'Passwords do not match'
		),
		['confirmPassword']
	)
)

export type SignUpSchema = typeof signUpSchema
