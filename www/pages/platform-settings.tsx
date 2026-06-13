import { RootLayout } from '@/components/layout'
import { Header } from '@/components/header'
import { PlatformSettingsForm } from '@/components/platform/settings'

export default function PlatformSettingsPage() {
  return (
      <RootLayout mainHeader={<Header />}>
        <div className="w-full py-4">
          <PlatformSettingsForm />
        </div>
      </RootLayout>
  )
}
