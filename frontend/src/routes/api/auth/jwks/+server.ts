import { auth } from "$lib/auth";
import { toSvelteKitHandler } from "better-auth/svelte-kit";
import type { RequestHandler } from './$types';

export const GET: RequestHandler = toSvelteKitHandler(auth);