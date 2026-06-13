'use client'

import { useEffect } from 'react'
import { useRouter } from 'next/router'
import { ROUTES } from '@/lib/routes'

export default function WgIndex() {
	const router = useRouter()
	useEffect(() => {
		router.replace(ROUTES.wg.networks)
	}, [router])
	return null
}


