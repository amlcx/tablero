import type { PageServerLoad } from './$types'

const pageTitle = 'Authenticate'

export const load: PageServerLoad = async () => {
	return {
		pageTitle
	}
}
