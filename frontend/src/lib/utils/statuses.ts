// import type { Code, ConnectError } from '@connectrpc/connect';

type HttpErrorCode = 400 | 401 | 403 | 404 | 409 | 422 | 429 | 499 | 500 | 501 | 502 | 503 | 504

/*export function normalizeConnectErrorCode(c: Code): HttpErrorCode {
	const statusMap: Record<Code, HttpErrorCode> = {
		1: 499, // Canceled
		2: 500, // Unknown
		3: 400, // Invalid argument
		4: 504, // Deadline exceeded
		5: 404, // Not found
		6: 409, // Already exists
		7: 403, // Permission denied
		8: 429, // Resource exhausted
		9: 400, // Failed precondition
		10: 409, // Aborted
		11: 400, // Out of range
		12: 501, // Unimplemented
		13: 500, // Internal
		14: 503, // Unavailable
		15: 500, // Data loss
		16: 401 // Unauthenticated
	};

	if (c in statusMap) {
		return statusMap[c];
	}

	return 500;
}*/

export function normalizeErrorStatus(status: number): HttpErrorCode {
	const statusMap: Record<number, HttpErrorCode> = {
		400: 400,
		401: 401,
		403: 403,
		404: 404,
		405: 400,
		406: 400,
		408: 400,
		409: 409,
		410: 404,
		411: 400,
		412: 400,
		413: 400,
		414: 400,
		415: 400,
		416: 400,
		417: 400,
		418: 400,
		421: 400,
		422: 422,
		423: 400,
		424: 400,
		425: 400,
		426: 400,
		428: 400,
		429: 429,
		431: 400,
		451: 403,

		500: 500,
		501: 500,
		502: 502,
		503: 503,
		504: 502,
		505: 500,
		506: 500,
		507: 500,
		508: 500,
		510: 500,
		511: 401
	}

	if (status in statusMap) {
		return statusMap[status]
	}

	if (status >= 400 && status < 500) {
		return 400
	}

	if (status >= 500) {
		return 500
	}

	return 400
}
