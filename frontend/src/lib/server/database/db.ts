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

const pgDb = new SQL({
	hostname: POSTGRES_HOSTNAME,
	port: POSTGRES_PORT,
	database: POSTGRES_DB,
	username: POSTGRES_USER,
	password: POSTGRES_PASSWORD,

	max: 30,
	idleTimeout: 30,
	maxLifetime: 0,
	connectionTimeout: 10
})

export const db = new Kysely({
	dialect: new PostgresJSDialect({
		postgres: pgDb
	})
})
