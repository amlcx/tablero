import { SQL } from 'bun'
import { Kysely } from 'kysely'
import { PostgresJSDialect } from 'kysely-postgres-js'

const pgDb = new SQL({
	url: '',

	max: 30,
	idleTimeout: 30,
	maxLifetime: 0,
	connectionTimeout: 10,

	onconnect: (client) => {
		console.log('connected to database')
	},

	onclose: (client) => {
		console.log(`Connection to database closed`)
	}
})

export const db = new Kysely({
	dialect: new PostgresJSDialect({
		postgres: pgDb
	})
})
