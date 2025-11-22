import { UPLOAD_DIR } from '$env/static/private'
import { mkdir } from 'node:fs/promises'
import path from 'path'
import sharp from 'sharp'

const baseDir = UPLOAD_DIR

export const saveAvatarToDisk = async (avatar: File, username: string) => {
	await mkdir(baseDir, { recursive: true })

	const filename = `${username}-${Bun.randomUUIDv7()}.webp`
	const filepath = path.join(baseDir, filename)

	const arrayBuf = await avatar.arrayBuffer()
	const buf = Buffer.from(arrayBuf)

	await sharp(buf).resize(300, 300, { fit: 'cover' }).webp({ quality: 85 }).toFile(filepath)

	return `/avatars/${filename}`
}
