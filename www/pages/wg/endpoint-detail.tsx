'use client'

import { RootLayout } from '@/components/layout'
import { Header } from '@/components/header'
import EndpointDetail from '@/components/wg/endpoint-detail'

export default function EndpointDetailPage() {
	return (
			<RootLayout mainHeader={<Header />}>
				<div className="w-full flex flex-col gap-4">
					<EndpointDetail />
				</div>
			</RootLayout>
	)
}


