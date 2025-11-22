import { BACKEND_URL } from '$env/static/private'
import { GreetService } from '$lib/api/gen/api/v1/greet_pb'
import { createClient } from '@connectrpc/connect'
import { createConnectTransport } from '@connectrpc/connect-node'

const transport = createConnectTransport({
	baseUrl: BACKEND_URL,
	httpVersion: '2'
})

export function createApiClients() {
	console.debug(`Creating client with transport: HTTP 2 and baseUrl: ${BACKEND_URL}`)
	return {
		greetApi: createClient(GreetService, transport)
	}
}
