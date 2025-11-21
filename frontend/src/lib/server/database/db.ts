import {
	POSTGRES_DB,
	POSTGRES_HOSTNAME,
	POSTGRES_PASSWORD,
	POSTGRES_PORT,
	POSTGRES_USER
} from '$env/static/private'
import { SQL } from 'bun'
import { Kysely } from 'kysely'
import { PostgresJSDialect } from 'kysely-postgres-js'

const url = `postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOSTNAME}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable`

const pgDb = new SQL({
	url: url,

	max: 30,
	idleTimeout: 30,
	maxLifetime: 0,
	connectionTimeout: 10,

	onconnect: (client) => {
		console.log('connected to database')
	},

	onclose: (client) => {
		console.log(`Connection to database closed (connection string: ${url})`)
	}
})

export const db = new Kysely({
	dialect: new PostgresJSDialect({
		postgres: pgDb
	})
})
