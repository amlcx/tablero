import { UPLOAD_DIR } from '$env/static/private'
import path from 'path'

const baseDir = UPLOAD_DIR

export const GET = async ({ params }) => {
	const { slug } = params

	const fullpath = path.join(baseDir, slug)

	const resp = new Response(Bun.file(fullpath))

	// TODO: considering that each file has a UUID, we can assume their content will never change.
	// I should add headers to tell clients to cache the response for as long as possible.

	return resp
}
